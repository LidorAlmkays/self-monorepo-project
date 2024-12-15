#!/bin/bash
# apply.sh - Applies all Kubernetes manifests in the specified folder

# Folder containing the Kubernetes deployment files
MANIFESTS_FOLDER="./manifests"

# Check if the folder exists
if [ ! -d "$MANIFESTS_FOLDER" ]; then
  echo "Error: Directory '$MANIFESTS_FOLDER' does not exist."
  exit 1
fi

# Apply all manifest files in the folder
kubectl apply -f "$MANIFESTS_FOLDER"

# Check the exit status
if [ $? -eq 0 ]; then
  echo "All Kubernetes manifests applied successfully."
else
  echo "Failed to apply Kubernetes manifests."
  exit 1
fi