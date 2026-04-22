#!/usr/bin/env bash

set -euo pipefail

run_contacts_tests() {
  if skip "contacts"; then
    echo "==> contacts (skipped)"
    return 0
  fi

  run_required "contacts" "contacts list" aim-google contacts list --json --max 1 >/dev/null

  local contact_json contact_id
  contact_json=$(aim-google contacts create --given "aim-google" --family "smoke-$TS" --email "aim-google-smoke-$TS@example.com" --phone "+1555555$TS" --json)
  contact_id=$(extract_field "$contact_json" resourceName)
  [ -n "$contact_id" ] || { echo "Failed to parse contact resourceName" >&2; exit 1; }

  run_required "contacts" "contacts get" aim-google contacts get "$contact_id" --json >/dev/null
  run_required "contacts" "contacts update" aim-google contacts update "$contact_id" --given "aim-google" --family "smoke-updated-$TS" --email "aim-google-smoke-$TS@example.com" --json >/dev/null
  run_required "contacts" "contacts search" aim-google contacts search "aim-google-smoke-$TS@example.com" --json --max 1 >/dev/null
  run_required "contacts" "contacts delete" aim-google contacts delete "$contact_id" --force >/dev/null

  if is_consumer_account "$ACCOUNT"; then
    echo "==> contacts directory (skipped; Workspace only)"
    echo "==> contacts other (skipped; Workspace only)"
  else
    run_optional "contacts-directory" "contacts directory list" aim-google contacts directory list --json --max 1 >/dev/null
    run_optional "contacts-directory" "contacts directory search" aim-google contacts directory search "aim-google" --json --max 1 >/dev/null
    run_optional "contacts-other" "contacts other list" aim-google contacts other list --json --max 1 >/dev/null
    run_optional "contacts-other" "contacts other search" aim-google contacts other search "aim-google" --json --max 1 >/dev/null
  fi
}
