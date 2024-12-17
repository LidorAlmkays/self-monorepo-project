#!/bin/bash

# Base directory containing the .env files
BASE_DIR="../configs"

# Directory to store generated YAML files
OUTPUT_DIR="./secrets"

# Function to sanitize an .env file by removing comments and empty lines
sanitize_env() {
  local input_file=$1
  local output_file=$2

  awk ' 
  /^[ \t]*#/ { next }  # Skip comment lines
  /^[ \t]*$/ { next }  # Skip blank lines
  {
    gsub(/\\n/, "", $0); # Remove newline characters
    print $0;
  }' "$input_file" > "$output_file"
}

# Function to preprocess a sanitized .env file
preprocess_env() {
  local input_file=$1
  local output_file=$2

  awk -F'=' '
  {
    # Trim single quotes, spaces, and invalid characters
    gsub(/'\''/, "", $2); # Remove single quotes
    gsub(/ /, "", $2);    # Remove spaces
    key=tolower($1);      # Convert key to lowercase
    gsub(/_/, "-", key);  # Replace underscores with dashes
    print key "=" $2;
  }' "$input_file" > "$output_file"
}

# Function to create YAML from an .env file
create_yaml() {
  local secret_name=$1
  local env_file=$2
  local yaml_file=$3

  kubectl create secret generic "$secret_name" \
    --from-env-file="$env_file" \
    --dry-run=client \
    -o yaml > "$yaml_file"

  if [ $? -ne 0 ]; then
    echo "Failed to create YAML for secret $secret_name"
    return 1
  else
    echo "YAML for secret $secret_name created successfully at $yaml_file"
    return 0
  fi
}

# Function to apply a YAML file to the Kubernetes cluster
apply_yaml() {
  local yaml_file=$1
  local secret_name=$2

  kubectl apply -f "$yaml_file"

  if [ $? -ne 0 ]; then
    echo "Failed to apply secret $secret_name to the cluster"
    return 1
  else
    echo "Secret $secret_name applied successfully"
    return 0
  fi
}

# Function to handle ME_CONFIG_MONGODB_URL encoding
encode_mongodb_url() {
  local env_file=$1
  local yaml_file=$2

  if grep -q 'ME_CONFIG_MONGODB_URL' "$env_file"; then
    url_value=$(grep 'ME_CONFIG_MONGODB_URL' "$env_file" | cut -d '=' -f2 | tr -d '"')
    encoded_url=$(echo -n "$url_value" | base64)
    sed -i "s|me-config-mongodb-url:.*|me-config-mongodb-url: $encoded_url|" "$yaml_file"
  fi
}

# Function to create a dated output directory for YAML files
create_dated_output_dir() {
  local base_dir=$1
  local date_dir=$(date +%Y-%m-%d)
  local full_path="$base_dir/$date_dir"

  mkdir -p "$full_path"

  # Create or overwrite a .gitignore file to ignore all files in this directory
  echo "*" > "$OUTPUT_DIR/.gitignore"
  echo "!README.md" >> "$OUTPUT_DIR/.gitignore"

  # Optional: Create a README.md to explain why this folder is ignored
  cat <<EOF > "$OUTPUT_DIR/README.md"
# Secrets YAML Folder

This folder contains generated YAML files for Kubernetes secrets. These files are ignored by Git for security reasons.
EOF

  echo "$full_path"
}

# Main function
main() {
  # Check if the base directory exists
  if [ ! -d "$BASE_DIR" ]; then
    echo "Error: Directory $BASE_DIR does not exist."
    exit 1
  fi

  # Create the dated output directory
  OUTPUT_DIR=$(create_dated_output_dir "$OUTPUT_DIR")

  # Find all .env files under the base directory
  find "$BASE_DIR" -type f -name "*.env" | while read -r env_file; do
    # Get the relative path from the base directory
    relative_path=${env_file#"$BASE_DIR"/}

    # Replace slashes with hyphens, underscores with dashes, and remove the .env extension
    base_secret_name=$(echo "$relative_path" | sed 's/\//-/g' | sed 's/_/-/g' | sed 's/\.env$//')
    secret_name="${base_secret_name}-secret"

    # Define the YAML output path
    yaml_file="$OUTPUT_DIR/${secret_name}.yaml"

    echo "Creating YAML for secret $secret_name from $env_file"

    # Sanitize the .env file
    tmp_env_file=$(mktemp)
    sanitize_env "$env_file" "$tmp_env_file"

    # Preprocess the sanitized .env file
    tmp_processed_file=$(mktemp)
    preprocess_env "$tmp_env_file" "$tmp_processed_file"

    # Create YAML for the secret
    create_yaml "$secret_name" "$tmp_processed_file" "$yaml_file"
    if [ $? -ne 0 ]; then
      rm "$tmp_env_file" "$tmp_processed_file"
      continue
    fi

    # Ensure MongoDB URL is encoded correctly
    encode_mongodb_url "$env_file" "$yaml_file"

    # Apply the YAML to the cluster
    apply_yaml "$yaml_file" "$secret_name"

    # Clean up the temporary files
    rm "$tmp_env_file" "$tmp_processed_file"
  done
}

# Run the main function
main
