#!/bin/bash
echo "Run with 'source export_datastore_env_vars.sh'"
source <(gcloud beta emulators datastore env-init)
