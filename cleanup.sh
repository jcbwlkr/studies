#!/bin/bash

# Find all directories under here that don't have a . somewhere in their name.
for dir in $(find . -type d -not -wholename '*.*'); do

  # Get the name of just the directory.
  f=$(basename $dir);

  # Delete any files in that directory whose name match the directory name.
  rm -f "${dir}/${f}";
done
