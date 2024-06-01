#!/bin/sh -l

# Exit on any command failure
set -e

# Run Trivy scan
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasecurity/trivy:latest image --severity $1 --ignore-unfixed=$2 myapp:latest > trivy-output.txt || true

# Check if the scan found any issues
if grep -q "VULNERABILITY ID" trivy-output.txt; then
  echo "Trivy scan failed"
  exit 1
else
  echo "Trivy scan passed"
  exit 0
fi
