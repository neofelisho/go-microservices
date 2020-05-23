#!/bin/bash

# Usage:
# For bash: `source env/env.sh dev` or `source <(env/env.sh dev)`
# For fish: `env/env.sh dev | source`

ENV=$1
if [ -z "$ENV" ]; then
  echo "Please specify one of the available environment:"
  find env/ -maxdepth 1 -name '*.env' | sed -n 's/^env\/\(.*\)\.env$/\1/p'
else
  while read -r i; do
    if [ -n "$i" ]; then
      echo "export $i"
    fi
  done <env/"$ENV".env
fi
