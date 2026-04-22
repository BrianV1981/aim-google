#!/usr/bin/env bash
set -euo pipefail

ACCOUNT=""
ALLOW_NONTEST=false
SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd -- "$SCRIPT_DIR/.." && pwd)"

usage() {
  cat <<'USAGE'
Usage: scripts/live-chat-test.sh [options]

Options:
  --account <email>   Account to use (defaults to AIM_GOOGLE_IT_ACCOUNT or first auth)
  --allow-nontest     Allow running against non-test accounts
  -h, --help          Show this help

Env:
  AIM_GOOGLE_LIVE_CHAT_SPACE=spaces/<id>        Existing space to use for list/send
  AIM_GOOGLE_LIVE_CHAT_THREAD=<id|resource>    Thread id or resource for sends
  AIM_GOOGLE_LIVE_CHAT_DM=user@domain          DM target (workspace user)
  AIM_GOOGLE_LIVE_CHAT_DM_THREAD=<id|resource> Thread id for DM send
  AIM_GOOGLE_LIVE_CHAT_CREATE=1                Create a new space (no cleanup)
  AIM_GOOGLE_LIVE_CHAT_MEMBER=user@domain      Member to add when creating a space
  AIM_GOOGLE_LIVE_ALLOW_NONTEST=1              Allow non-test accounts
USAGE
}

while [ $# -gt 0 ]; do
  case "$1" in
    --account)
      ACCOUNT="$2"
      shift
      ;;
    --allow-nontest)
      ALLOW_NONTEST=true
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown arg: $1" >&2
      usage >&2
      exit 2
      ;;
  esac
  shift
done

BIN="${AIM_GOOGLE_BIN:-$ROOT_DIR/bin/aim-google}"
if [ ! -x "$BIN" ]; then
  make -C "$ROOT_DIR" build >/dev/null
fi

PY="${PYTHON:-python3}"
if ! command -v "$PY" >/dev/null 2>&1; then
  PY="python"
fi

if [ -z "$ACCOUNT" ]; then
  ACCOUNT="${AIM_GOOGLE_IT_ACCOUNT:-}"
fi
if [ -z "$ACCOUNT" ]; then
  acct_json=$($BIN auth list --json)
  ACCOUNT=$($PY -c 'import json,sys; obj=json.load(sys.stdin); print(obj.get("accounts", [{}])[0].get("email", ""))' <<<"$acct_json")
fi
if [ -z "$ACCOUNT" ]; then
  echo "No account available for live tests." >&2
  exit 1
fi

is_test_account() {
  local a
  a=$(echo "$1" | tr 'A-Z' 'a-z')
  case "$a" in
    *test*|*bot*|*sandbox*|*qa*|*staging*|*dev*|*@example.com)
      return 0
      ;;
  esac
  case "$a" in
    *+*)
      return 0
      ;;
  esac
  return 1
}

is_consumer_account() {
  local a domain
  a=$(echo "$1" | tr 'A-Z' 'a-z')
  domain="${a##*@}"
  case "$domain" in
    gmail.com|googlemail.com)
      return 0
      ;;
  esac
  return 1
}

if [ "${ALLOW_NONTEST:-false}" = false ] && [ -z "${AIM_GOOGLE_LIVE_ALLOW_NONTEST:-}" ]; then
  if ! is_test_account "$ACCOUNT"; then
    echo "Refusing to run live tests against non-test account: $ACCOUNT" >&2
    echo "Pass --allow-nontest or set AIM_GOOGLE_LIVE_ALLOW_NONTEST=1 to override." >&2
    exit 2
  fi
fi

if is_consumer_account "$ACCOUNT"; then
  echo "==> chat (skipped; Workspace only)"
  exit 0
fi

aim-google() {
  "$BIN" --account "$ACCOUNT" "$@"
}

TS=$(date +%Y%m%d%H%M%S)

echo "Using account: $ACCOUNT"
echo "==> chat spaces list"
aim-google chat spaces list --json --max 1 >/dev/null

if [ -n "${AIM_GOOGLE_LIVE_CHAT_SPACE:-}" ]; then
  echo "==> chat messages list"
  aim-google chat messages list "$AIM_GOOGLE_LIVE_CHAT_SPACE" --json --max 1 >/dev/null
  echo "==> chat threads list"
  aim-google chat threads list "$AIM_GOOGLE_LIVE_CHAT_SPACE" --json --max 1 >/dev/null
  echo "==> chat messages send"
  if [ -n "${AIM_GOOGLE_LIVE_CHAT_THREAD:-}" ]; then
    aim-google chat messages send "$AIM_GOOGLE_LIVE_CHAT_SPACE" --text "aim-google smoke $TS" --thread "$AIM_GOOGLE_LIVE_CHAT_THREAD" --json >/dev/null
  else
    aim-google chat messages send "$AIM_GOOGLE_LIVE_CHAT_SPACE" --text "aim-google smoke $TS" --json >/dev/null
  fi
else
  echo "==> chat messages/threads (skipped; set AIM_GOOGLE_LIVE_CHAT_SPACE)"
fi

if [ -n "${AIM_GOOGLE_LIVE_CHAT_CREATE:-}" ]; then
  if [ -z "${AIM_GOOGLE_LIVE_CHAT_MEMBER:-}" ]; then
    echo "==> chat spaces create (skipped; set AIM_GOOGLE_LIVE_CHAT_MEMBER)"
  else
    echo "==> chat spaces create"
    aim-google chat spaces create "aim-google-smoke-$TS" --member "$AIM_GOOGLE_LIVE_CHAT_MEMBER" --json >/dev/null
  fi
fi

if [ -n "${AIM_GOOGLE_LIVE_CHAT_DM:-}" ]; then
  echo "==> chat dm space"
  aim-google chat dm space "$AIM_GOOGLE_LIVE_CHAT_DM" --json >/dev/null
  echo "==> chat dm send"
  if [ -n "${AIM_GOOGLE_LIVE_CHAT_DM_THREAD:-}" ]; then
    aim-google chat dm send "$AIM_GOOGLE_LIVE_CHAT_DM" --text "aim-google dm $TS" --thread "$AIM_GOOGLE_LIVE_CHAT_DM_THREAD" --json >/dev/null
  else
    aim-google chat dm send "$AIM_GOOGLE_LIVE_CHAT_DM" --text "aim-google dm $TS" --json >/dev/null
  fi
else
  echo "==> chat dm (skipped; set AIM_GOOGLE_LIVE_CHAT_DM)"
fi

echo "Chat live tests complete."
