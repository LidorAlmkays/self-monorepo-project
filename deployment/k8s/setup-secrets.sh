#!/bin/bash

# Stop the script if any command fails
set -e

# Base directory containing the .env files
BASE_DIR="../configs"

# Directory to store generated YAML files
OUTPUT_DIR="./secrets"

# Function to check if a directory exists
check_directory_exists() {
  local dir=$1
  if [ ! -d "$dir" ]; then
    echo "Error: Directory $dir does not exist."
    exit 1
  fi
}

# Function to create the output directory and setup .gitignore
setup_output_directory() {
  local output_dir=$1
  mkdir -p "$output_dir"
  echo "*" > "$output_dir/.gitignore"
  echo "!README.md" >> "$output_dir/.gitignore"

  cat <<EOF > "$output_dir/README.md"
# Secrets YAML Folder

This folder contains generated YAML files for Kubernetes secrets. These files are ignored by Git for security reasons.
EOF
}

# Function to process the .env file and generate a temporary file with cleaned-up content
process_env_file() {
  local env_file=$1
  local tmp_env_file
  tmp_env_file=$(mktemp)

  awk -F'=' '
  /^[^#]/ { # Skip commented lines
    key=tolower($1);               # Convert the key to lowercase
    gsub(/_/, "-", key);           # Replace underscores with dashes
    if (key ~ /^[a-z][-._a-z0-9]*$/) { # Validate the key name
      gsub(/^[ \t]+|[ \t]+$/, "", $2); # Trim spaces around the value
      print key "=" $2;
    } else {
      printf "Invalid key skipped: %s\n", $1 > "/dev/stderr";
    }
  }' "$env_file" > "$tmp_env_file"

  echo "$tmp_env_file"
}

# Function to generate a Kubernetes secret YAML file from a .env file
generate_yaml_from_env() {
  local env_file=$1
  local tmp_env_file=$2
  local secret_name=$3
  local yaml_file=$4

  kubectl create secret generic "$secret_name" \
    --from-env-file="$tmp_env_file" \
    --dry-run=client \
    -o yaml > "$yaml_file"
}

# Function to apply the secret YAML to the Kubernetes cluster
apply_secret_to_cluster() {
  local yaml_file=$1
  echo "Applying secret from $yaml_file to the cluster"
  kubectl apply -f "$yaml_file"
}

# Main function to orchestrate the process
main() {
  # Check if the base directory exists
  check_directory_exists "$BASE_DIR"

  # Create the output directory and set up .gitignore
  setup_output_directory "$OUTPUT_DIR"

  # Trap to ensure temporary files are cleaned up
  trap 'rm -f "$tmp_env_file"' EXIT

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

    # Process the .env file
    tmp_env_file=$(process_env_file "$env_file")

    # Generate the YAML for the secret
    generate_yaml_from_env "$env_file" "$tmp_env_file" "$secret_name" "$yaml_file"

    echo "YAML for secret $secret_name created successfully at $yaml_file"

    # Apply the secret to the cluster
    apply_secret_to_cluster "$yaml_file"

    echo "Secret $secret_name applied successfully"
  done
}

# Call the main function
main
