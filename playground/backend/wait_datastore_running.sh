#!/bin/bash

current_dir="$(dirname "$0")"
source "$current_dir/envs_and_functions.sh"

waitport "$DATASTORE_PORT"