#!/bin/bash
# unapply.sh - Stops and deletes all Kubernetes resources defined in the specified folder

# Folder containing the Kubernetes deployment files
MANIFESTS_FOLDER="./manifests"

# Check if the folder exists
if [ ! -d "$MANIFESTS_FOLDER" ]; then
  echo "Error: Directory '$MANIFESTS_FOLDER' does not exist."
  exit 1
fi

# Gracefully scale down all Deployments to zero replicas
for deployment in $(kubectl get deployments -o jsonpath='{.items[*].metadata.name}' -n default); do
  echo "Scaling down Deployment: $deployment"
  kubectl scale deployment "$deployment" --replicas=0 -n default
  if [ $? -ne 0 ]; then
    echo "Failed to scale down Deployment: $deployment"
    exit 1
  fi
  echo "Deployment $deployment scaled down successfully."
  sleep 5 # Wait to ensure pods are terminated
done

# Delete all Services in the namespace
for service in $(kubectl get services -o jsonpath='{.items[*].metadata.name}' -n default); do
  echo "Deleting Service: $service"
  kubectl delete service "$service" -n default
  if [ $? -ne 0 ]; then
    echo "Failed to delete Service: $service"
    exit 1
  fi
  echo "Service $service deleted successfully."
  sleep 2 # Wait to ensure service is removed
done

# Finally, delete all manifest files in the folder
kubectl delete -f "$MANIFESTS_FOLDER"

# Check the exit status
if [ $? -eq 0 ]; then
  echo "All Kubernetes manifests deleted successfully."
else
  echo "Failed to delete Kubernetes manifests."
  exit 1
fi
