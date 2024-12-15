#!/bin/bash

# Base directory containing the .env files
BASE_DIR="../configs"

# Directory to store generated YAML files
OUTPUT_DIR="./secrets_yaml"

# Check if the base directory exists
if [ ! -d "$BASE_DIR" ]; then
  echo "Error: Directory $BASE_DIR does not exist."
  exit 1
fi

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Create or overwrite a .gitignore file to ignore all files in this directory
echo "*" > "$OUTPUT_DIR/.gitignore"
echo "!README.md" >> "$OUTPUT_DIR/.gitignore"

# Optional: Create a README.md to explain why this folder is ignored
cat <<EOF > "$OUTPUT_DIR/README.md"
# Secrets YAML Folder

This folder contains generated YAML files for Kubernetes secrets. These files are ignored by Git for security reasons.
EOF

# Find all .env files under the base directory
find "$BASE_DIR" -type f -name "*.env" | while read -r env_file; do
  # Get the relative path from the base directory
  relative_path=${env_file#"$BASE_DIR"/}

  # Replace slashes with hyphens, underscores with dashes, and remove the .env extension
  secret_name=$(echo "$relative_path" | sed 's/\//-/g' | sed 's/_/-/g' | sed 's/\.env$//')

  # Define the YAML output path
  yaml_file="$OUTPUT_DIR/$secret_name.yaml"

  echo "Creating YAML for secret $secret_name from $env_file"

  # Generate the YAML for the secret
  kubectl create secret generic "$secret_name" \
    --from-env-file="$env_file" \
    --dry-run=client \
    -o yaml > "$yaml_file"

  # Check if the YAML file was successfully created
  if [ $? -ne 0 ]; then
    echo "Failed to create YAML for secret $secret_name"
    continue
  else
    echo "YAML for secret $secret_name created successfully at $yaml_file"
  fi

  # Apply the YAML to the cluster
  echo "Applying secret $secret_name to the cluster"
  kubectl apply -f "$yaml_file"

  # Check if the secret was successfully applied
  if [ $? -ne 0 ]; then
    echo "Failed to apply secret $secret_name to the cluster"
  else
    echo "Secret $secret_name applied successfully"
  fi
done
