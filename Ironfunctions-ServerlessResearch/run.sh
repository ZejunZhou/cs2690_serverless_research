#!/bin/bash

# Check if fn CLI is available
if ! command -v fn &> /dev/null; then
  echo "Error: fn CLI is not installed. Please install fn CLI first."
  exit 1
fi

# Create the benchmark application
echo "Creating Fn application 'benchmark'..."
fn apps create benchmark 2>/dev/null
if [ $? -eq 0 ]; then
  echo "Fn application 'benchmark' created successfully."
else
  echo "Warning: Fn application 'benchmark' already exists. Continuing..."
fi

# Define hotel services and their corresponding routes
declare -A hotel_services
hotel_services=(
  ["hotel_dates"]="/dates"
  ["hotel_recommendation"]="/recommendation"
  ["hotel_reservation"]="/reservation"
  ["hotel_user"]="/user"
)

# Iterate through each service directory
for service_dir in "${!hotel_services[@]}"; do
  if [ -d "$service_dir" ]; then
    echo "Processing directory: $service_dir"
    cd "$service_dir" || { echo "Error: Unable to enter directory $service_dir"; exit 1; }

    # Run fn build
    echo "Running 'fn build' in $service_dir..."
    fn build
    if [ $? -ne 0 ]; then
      echo "Error: 'fn build' failed in $service_dir."
      exit 1
    fi
    echo "'fn build' completed successfully in $service_dir."

    # Create the route
    route=${hotel_services["$service_dir"]}
    echo "Creating route: $route for service in $service_dir..."
    fn routes create benchmark "$route"
    if [ $? -ne 0 ]; then
      echo "Warning: Route $route already exists or failed to create. Continuing..."
    else
      echo "Route $route created successfully for $service_dir."
    fi

    # Return to the parent directory
    cd - > /dev/null
  else
    echo "Warning: Directory $service_dir does not exist. Skipping."
  fi
done

echo "All operations completed successfully."
