#!/bin/bash

# Step 1: Run docker-compose
echo "Starting docker-compose up --build..."
docker-compose up --build -d

# Step 2: Check if Go is installed
echo "Checking if Go is installed..."
if ! command -v go &> /dev/null; then
  echo "Go is not installed. Installing Go 1.23.3 for Linux/amd64..."
  
  # Set variables for Go installation
  GO_VERSION="1.23.3"
  OS="linux"
  ARCH="amd64"
  GO_TAR="go${GO_VERSION}.${OS}-${ARCH}.tar.gz"
  GO_URL="https://go.dev/dl/${GO_TAR}"
  INSTALL_DIR="/usr/local"

  # Download and install Go
  wget "${GO_URL}" -O "/tmp/${GO_TAR}"
  sudo tar -C "${INSTALL_DIR}" -xzf "/tmp/${GO_TAR}"
  
  # Set up Go environment variables
  export PATH="${INSTALL_DIR}/go/bin:$PATH"
  echo 'export PATH="/usr/local/go/bin:$PATH"' >> ~/.bashrc
  source ~/.bashrc

  # Clean up
  rm "/tmp/${GO_TAR}"
else
  echo "Go is already installed: $(go version)"
fi

# Step 3: Run the Go service
SERVICE_DIR="./services"
echo "Navigating to ${SERVICE_DIR} and starting Go service..."
cd "${SERVICE_DIR}" || { echo "Failed to navigate to ${SERVICE_DIR}"; exit 1; }
go run run.go
