#!/bin/bash
# Prompt the user for the location of the .go file
read -p "Enter the path of the .go file to compile: " go_file
# Check if the provided file exists
if [[ ! -f "$go_file" ]]; then
    echo "The file '$go_file' does not exist."
    exit 1
fi

# Extract the filename without extension
filename=$(basename -- "$go_file")
filename_no_ext="${filename%.*}"

# Create Output Directory
mkdir -p "./bin"
compiler_dir=$(pwd)

# Get the directory of the .go file
go_dir=$(dirname "$go_file")

# Compile for Linux
echo "Compiling for Linux..."
(
    cd "$go_dir" || exit
    GOOS=linux GOARCH=amd64 go build -o "${compiler_dir}/bin/${filename_no_ext}-linux" "$filename"
)
if [[ $? -eq 0 ]]; then
    echo "Successfully compiled to '${compiler_dir}/bin/${filename_no_ext}-linux'."
else
    echo "Failed to compile for Linux."
fi

# Compile for Windows
echo "Compiling for Windows..."
(
    cd "$go_dir" || exit
    GOOS=windows GOARCH=amd64 go build -o "${compiler_dir}/bin/${filename_no_ext}-windows.exe" "$filename"
)
if [[ $? -eq 0 ]]; then
    echo "Successfully compiled to '${compiler_dir}/bin/${filename_no_ext}-windows.exe'."
else
    echo "Failed to compile for Windows."
fi