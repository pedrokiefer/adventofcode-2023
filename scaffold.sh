#!/bin/bash

echo "Scaffolding..."

if [ -z "$1" ]
then
  echo "Please provide a day number"
  exit 1
fi

dirName="day$1"

# Create directories
mkdir -p $dirName

echo -e "package main\n\nfunc main() {\n\n}\n" > "$dirName/main.go"
echo -e "package main\n\nimport \"testing\"\n\nfunc Test(t *testing.T) {\n\n}\n" > "$dirName/${dirName}_test.go"
