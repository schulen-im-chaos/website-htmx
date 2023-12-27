#!/bin/bash

# Function to build the project
build_project() {
    echo "Building the project..."

    # Change directory and build CSS
    cd ./static/content && bunx tailwindcss -i ./../../web/input.css -o ./dist/output.css --minify

    # Return to the project root directory
    cd ./../../



    # Generate templates and build the Go project
    templ generate
    go build .
}

# Function to bundle the project
bundle_project() {
    echo "Bundling the project..."

    # Remove the existing bundle directory if it exists
    if [ -d "bundle" ]; then
        echo "Removing existing bundle directory..."
        rm -rf bundle
    fi

    # Create the necessary directories for the bundle
    mkdir -p bundle/files
    mkdir -p bundle/static/content/dist
    mkdir -p bundle/static/content/node_modules/htmx.org/dist
    mkdir -p bundle/static/img

    # Copy the necessary files and directories
    cp -r ./files/* ./bundle/files/
    cp ./static/content/dist/output.css ./bundle/static/content/dist/
    cp ./static/content/node_modules/htmx.org/dist/htmx.min.js ./bundle/static/content/node_modules/htmx.org/dist/
    cp ./static/favicon.png ./bundle/static/
    cp -r ./static/img/* ./bundle/static/img/

    # Copy the built Go binary and Deleting it
    cp ./website ./bundle/
    rm ./website
}

# Execute the build and bundle functions
build_project
bundle_project

echo "Export script completed."
