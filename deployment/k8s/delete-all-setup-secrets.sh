#!/bin/bash

# Base directory containing the .env files
BASE_DIR="../configs"

# Check if the base directory exists
if [ ! -d "$BASE_DIR" ]; then
  echo "Error: Directory $BASE_DIR does not exist."
  exit 1
fi

# Find all .env files under the base directory
find "$BASE_DIR" -type f -name "*.env" | while read -r env_file; do
  # Get the relative path from the base directory
  relative_path=${env_file#"$BASE_DIR"/}

  # Replace slashes with hyphens, underscores with dashes, and remove the .env extension
  secret_name=$(echo "$relative_path" | sed 's/\//-/g' | sed 's/_/-/g' | sed 's/\.env$//')

  echo "Deleting secret $secret_name"

  # Delete the Kubernetes secret
  kubectl delete secret "$secret_name"

  # Check if the command was successful
  if [ $? -ne 0 ]; then
    echo "Failed to delete secret $secret_name"
  else
    echo "Secret $secret_name deleted successfully"
  fi
done
