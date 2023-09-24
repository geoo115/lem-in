	#!/bin/bash

# Compile the Go code
go build -o main main.go

# Input file names
input_files=("example00.txt" "example01.txt" "example02.txt" "example03.txt" "example04.txt" "example05.txt" "example06.txt" "example07.txt" "example08.txt" "badexample00.txt" "badexample01.txt")

# Run the code with each input file
for input_file in "${input_files[@]}"; do
  example_name="${input_file%.txt}"  # Extract example name without extension
  example_name="${example_name^}"     # Capitalize first letter

  echo "================= $example_name ====================="
  go run . "$input_file"
  # echo "=========================="
done

# Clean up the executable
rm main