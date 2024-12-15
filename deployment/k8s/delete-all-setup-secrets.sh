#!/bin/bash

# Stop the script if any command fails
set -e

# Base directory containing the .env files
BASE_DIR="../configs"

# Function to check if a directory exists
check_directory_exists() {
  local dir=$1
  if [ ! -d "$dir" ]; then
    echo "Error: Directory $dir does not exist."
    exit 1
  fi
}

# Function to generate the secret name based on the .env file path
generate_secret_name() {
  local env_file=$1
  local base_dir=$2

  # Get the relative path from the base directory
  relative_path=${env_file#"$base_dir"/}

  # Replace slashes with hyphens, underscores with dashes, and remove the .env extension
  base_secret_name=$(echo "$relative_path" | sed 's/\//-/g' | sed 's/_/-/g' | sed 's/\.env$//')
  echo "${base_secret_name}-secret"
}

# Function to delete the Kubernetes secret
delete_k8s_secret() {
  local secret_name=$1

  echo "Deleting secret $secret_name"

  kubectl delete secret "$secret_name"

  # Check if the command was successful
  if [ $? -ne 0 ]; then
    echo "Failed to delete secret $secret_name"
  else
    echo "Secret $secret_name deleted successfully"
  fi
}

# Main function to orchestrate the process
main() {
  # Check if the base directory exists
  check_directory_exists "$BASE_DIR"

  # Find all .env files under the base directory
  find "$BASE_DIR" -type f -name "*.env" | while read -r env_file; do
    # Generate the secret name based on the .env file path
    secret_name=$(generate_secret_name "$env_file" "$BASE_DIR")

    # Delete the Kubernetes secret
    delete_k8s_secret "$secret_name"
  done
}

# Call the main function
main
