#!/bin/bash

if [[ -z "$1" ]];  then
  echo "Usage:
  gen.sh <output path> [types file]
  "
  exit 1
fi

targetDir="$1"
sourceFile="${2:-types.csv}"
goTargets=(
  'basic-getter'
  'basic-setter'
  'err-getter'
  'err-setter'
  'slice-getter'
  'slice-setter'
  'slice-err-getter'
  'slice-err-setter'
)

while read t; do
  split=(`echo $t | sed 's/,/ /g'`)
  type="${split[0]:?}"
  name="${split[1]:?}"

  if [[ -z "$name" ]] || [[ -z "$type" ]]; then
    echo "Name and type must be provided to run generator"
    echo "$name $type"
    exit 1
  fi

  echo "Process: $type"

  for tgt in "${goTargets[@]}"; do
    echo "  -> $tgt"
    go run gen/$tgt.go \
      -name "$name" \
      -type "$type" > "$targetDir/$tgt-$type.go"
  done

 done < $sourceFile
