#!/usr/bin/env bash

set -euo pipefail

run_people_tests() {
  run_required "people" "people me" aim-google people me --json >/dev/null
}
