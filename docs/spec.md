# aim-google spec

## Goal

Build a single, clean, modern Go CLI that talks to:

- Gmail API
- Google Calendar API
- Google Chat API
- Google Classroom API
- Google Drive API
- Google Docs API
- Google Sheets API
- Google Forms API
- Apps Script API
- Google Tasks API
- Cloud Identity API (Groups)
- Google People API (Contacts + directory)
- Google Keep API (Workspace-only, service account)

This replaces the existing separate CLIs (`gmcli`, `gccli`, `gdcli`) and the Python contacts server conceptually, but:

- no backwards compatibility
- no migration tooling

## Non-goals

- Preserving legacy command names/flags/output formats
- Importing existing `~/.gmcli`, `~/.gccli`, `~/.gdcli` state
- Running an MCP server (this is a CLI)

## Language/runtime

- Go `1.26` (see `go.mod`)

## CLI framework

- `github.com/alecthomas/kong`
- Root command: `aim-google`
- Global flag:
  - `--color=auto|always|never` (default `auto`)
  - `--json` (JSON output to stdout)
  - `--plain` (TSV output to stdout; stable/parseable; disables colors)
  - `--force` (skip confirmations for destructive commands)
  - `--no-input` (never prompt; fail instead)
  - `--version` (print version)

Notes:

- We run `SilenceUsage: true` and print errors ourselves (colored when possible).
- `NO_COLOR` is respected.

Environment:

- `AIM_GOOGLE_COLOR=auto|always|never` (default `auto`, overridden by `--color`)
- `AIM_GOOGLE_JSON=1` (default JSON output; overridden by flags)
- `AIM_GOOGLE_PLAIN=1` (default plain output; overridden by flags)

## Output (TTY-aware colors)

- `github.com/muesli/termenv` is used to detect rich TTY capabilities and render colored output.
- Colors are enabled when:
  - output is a rich terminal and `--color=auto`, and `NO_COLOR` is not set; or
  - `--color=always`
- Colors are disabled when:
  - `--color=never`; or
  - `NO_COLOR` is set

Implementation: `internal/ui/ui.go`.

## Auth + secret storage

### OAuth client credentials (non-secret-ish)

- Stored on disk in the per-user config directory:
  - `$(os.UserConfigDir())/aim-google/credentials.json` (default client)
  - `$(os.UserConfigDir())/aim-google/credentials-<client>.json` (named clients)
- Written with mode `0600`.
- Command:
  - `aim-google auth credentials <credentials.json>`
  - `aim-google --client <name> auth credentials <credentials.json>`
  - `aim-google auth credentials list`
  - `aim-google auth credentials remove [<client>|all]`
- Supports Google’s downloaded JSON format:
  - `installed.client_id/client_secret` or `web.client_id/client_secret`

Implementation: `internal/config/*`.

### Refresh tokens (secrets)

- Stored in OS credential store via `github.com/99designs/keyring`.
- Key namespace is `aim-google` by default (keyring `ServiceName`); override with `AIM_GOOGLE_KEYRING_SERVICE_NAME`.
- Key format: `token:<client>:<email>` (default client uses `token:default:<email>`)
- Legacy key format: `token:<email>` (migrated on first read)
- Stored payload is JSON (refresh token + metadata like selected services/scopes).
- Fallback: if no OS credential store is available, keyring may use its encrypted "file" backend:
  - Directory: `$(os.UserConfigDir())/aim-google/keyring/` (one file per key)
  - Password: prompts on TTY; for non-interactive runs set `AIM_GOOGLE_KEYRING_PASSWORD`

Current minimal management commands (implemented):

- `aim-google auth tokens list` (keys only)
- `aim-google auth tokens delete <email>`

Implementation: `internal/secrets/store.go`.

### OAuth flow

- Desktop OAuth 2.0 flow using local HTTP redirect on an ephemeral port.
- Supports a browserless/manual flow (paste redirect URL) for headless environments.
- Supports a remote/server-friendly 2-step manual flow:
  - Step 1 prints an auth URL (`aim-google auth add ... --remote --step 1`)
  - Step 2 exchanges the pasted redirect URL and requires `state` validation (`--remote --step 2 --auth-url ...`)
- Refresh token issuance:
  - requests `access_type=offline`
  - supports `--force-consent` to force the consent prompt when Google doesn't return a refresh token
  - uses `include_granted_scopes=true` to support incremental auth re-runs

Scope selection note:

- The consent screen shows the scopes the CLI requested.
- Users cannot selectively un-check individual requested scopes in the consent screen; they either approve all requested scopes or cancel.
- To request fewer scopes, choose fewer services via `aim-google auth add --services ...` or use `aim-google auth add --readonly` where applicable.

## Config layout

- Base config dir: `$(os.UserConfigDir())/aim-google/`
- Files:
  - `config.json` (JSON5; comments and trailing commas allowed)
  - `credentials.json` (OAuth client id/secret; default client)
  - `credentials-<client>.json` (OAuth client id/secret; named clients)
- State:
  - `state/gmail-watch/<account>.json` (Gmail watch state)
  - `oauth-manual-state-<state>.json` (temporary manual OAuth state cache; expires quickly; no tokens)
- Secrets:
  - refresh tokens in keyring

We intentionally avoid storing refresh tokens in plain JSON on disk.

Environment:

- `AIM_GOOGLE_ACCOUNT=you@gmail.com` (email or alias; used when `--account` is not set; otherwise uses keyring default or a single stored token)
- `AIM_GOOGLE_CLIENT=work` (select OAuth client bucket; see `--client`)
- `AIM_GOOGLE_KEYRING_PASSWORD=...` (used when keyring falls back to encrypted file backend in non-interactive environments)
- `AIM_GOOGLE_KEYRING_BACKEND={auto|keychain|file}` (force backend; use `file` to avoid Keychain prompts and pair with `AIM_GOOGLE_KEYRING_PASSWORD` for non-interactive)
- `AIM_GOOGLE_KEYRING_SERVICE_NAME=...` (override keyring namespace/service name; default `aim-google`)
- `AIM_GOOGLE_TIMEZONE=America/New_York` (default output timezone; IANA name or `UTC`; `local` forces local timezone)
- `AIM_GOOGLE_ENABLE_COMMANDS=calendar,tasks,gmail.search` (optional allowlist; dot paths allowed)
- `AIM_GOOGLE_DISABLE_COMMANDS=gmail.send,gmail.drafts.send` (optional denylist; dot paths allowed)
- `AIM_GOOGLE_GMAIL_NO_SEND=1` (block Gmail send operations)
- `config.json` can also set `keyring_backend` (JSON5; env vars take precedence)
- `config.json` can also set `default_timezone` (IANA name or `UTC`)
- `config.json` can also set `account_aliases` for `aim-google auth alias` (JSON5)
- `config.json` can also set `account_clients` (email -> client) and `client_domains` (domain -> client)
- `config.json` can also set `gmail_no_send` and `no_send_accounts` for send guards

Flag aliases:
- `--out` also accepts `--output`.
- `--out-dir` also accepts `--output-dir` (Gmail thread attachment downloads).

## Commands (current + planned)

### Implemented

- `aim-google auth credentials <credentials.json|->`
- `aim-google auth credentials list`
- `aim-google auth credentials remove [<client>|all]`
- `aim-google --client <name> auth credentials <credentials.json|->`
- `aim-google auth add <email> [--services user|all|gmail,calendar,chat,classroom,drive,docs,slides,contacts,tasks,sheets,people,forms,appscript,ads,groups,keep,admin] [--readonly] [--drive-scope full|readonly|file] [--gmail-scope full|readonly] [--extra-scopes CSV] [--manual] [--remote] [--step 1|2] [--auth-url URL] [--listen-addr HOST[:PORT]] [--redirect-host HOST] [--timeout DURATION] [--force-consent]`
- `aim-google auth services [--markdown]`
- `aim-google auth manage [--services ...] [--listen-addr HOST[:PORT]] [--redirect-host HOST]`
- `aim-google auth keep <email> --key <service-account.json>` (Google Keep; Workspace only)
- `aim-google auth list`
- `aim-google auth alias list`
- `aim-google auth alias set <alias> <email>`
- `aim-google auth alias unset <alias>`
- `aim-google auth status`
- `aim-google auth remove <email>`
- `aim-google auth tokens list`
- `aim-google auth tokens delete <email>`
- `aim-google config get <key>`
- `aim-google config keys`
- `aim-google config list`
- `aim-google config path`
- `aim-google config set <key> <value>`
- `aim-google config unset <key>`
- `aim-google version`
- `aim-google drive ls [--all] [--parent ID] [--max N] [--page TOKEN] [--query Q] [--[no-]all-drives]` (`--all` and `--parent` are mutually exclusive)
- `aim-google drive search <text> [--raw-query] [--max N] [--page TOKEN] [--[no-]all-drives]`
- `aim-google drive get <fileId>`
- `aim-google drive download <fileId> [--out PATH] [--format F]` (`--format` only applies to Google Workspace files; `--format md` exports a Google Doc as Markdown)
- `aim-google drive upload <localPath> [--name N] [--parent ID] [--convert] [--convert-to doc|sheet|slides] [--keep-frontmatter]` (Markdown → Google Doc with `--convert` or `--convert-to doc`: leading `---`/`---` frontmatter is stripped before upload unless `--keep-frontmatter`; delimiter-based, not a full YAML parse)
- `aim-google drive mkdir <name> [--parent ID]`
- `aim-google drive delete <fileId> [--permanent]`
- `aim-google drive move <fileId> --parent ID`
- `aim-google drive rename <fileId> <newName>`
- `aim-google drive share <fileId> --to anyone|user|domain [--email addr] [--domain example.com] [--role reader|writer|commenter] [--discoverable]`
- `aim-google drive permissions <fileId> [--max N] [--page TOKEN]`
- `aim-google drive unshare <fileId> <permissionId>`
- `aim-google drive url <fileIds...>`
- `aim-google drive drives [--max N] [--page TOKEN] [--query Q]`
- `aim-google slides thumbnail <presentationId> <slideId> [--size small|medium|large] [--format png|jpeg] [--out PATH]`
- `aim-google calendar calendars`
- `aim-google calendar create-calendar <summary> [--description D] [--timezone TZ] [--location L]`
- `aim-google calendar acl <calendarId>`
- `aim-google calendar events <calendarId> [--cal ID_OR_NAME] [--calendars CSV] [--all] [--from RFC3339] [--to RFC3339] [--max N] [--page TOKEN] [--query Q] [--weekday]`
- `aim-google calendar event|get <calendarId> <eventId>`
- `AIM_GOOGLE_CALENDAR_WEEKDAY=1` defaults `--weekday` for `aim-google calendar events`
- `aim-google calendar create <calendarId> --summary S --from DT --to DT [--description D] [--location L] [--attendees a@b.com,c@d.com] [--all-day] [--event-type TYPE]`
- `aim-google calendar update <calendarId> <eventId> [--summary S] [--from DT] [--to DT] [--description D] [--location L] [--attendees ...] [--add-attendee ...] [--all-day] [--event-type TYPE]`
- `aim-google calendar delete <calendarId> <eventId>`
- `aim-google calendar freebusy [calendarIds] [--cal ID_OR_NAME] [--calendars CSV] [--all] --from RFC3339 --to RFC3339`
- `aim-google calendar conflicts [--cal ID_OR_NAME] [--calendars CSV] [--all] [--from RFC3339|date|relative] [--to RFC3339|date|relative] [--today|--week|--days N]`
- `aim-google calendar respond <calendarId> <eventId> --status accepted|declined|tentative [--send-updates all|none|externalOnly]`
- `aim-google time now [--timezone TZ]`
- `aim-google classroom courses [--state ...] [--max N] [--page TOKEN]`
- `aim-google classroom courses get <courseId>`
- `aim-google classroom courses create --name NAME [--owner me] [--state ACTIVE|...]`
- `aim-google classroom courses update <courseId> [--name ...] [--state ...]`
- `aim-google classroom courses delete <courseId>`
- `aim-google classroom courses archive <courseId>`
- `aim-google classroom courses unarchive <courseId>`
- `aim-google classroom courses join <courseId> [--role student|teacher] [--user me]`
- `aim-google classroom courses leave <courseId> [--role student|teacher] [--user me]`
- `aim-google classroom courses url <courseId...>`
- `aim-google classroom students <courseId> [--max N] [--page TOKEN]`
- `aim-google classroom students get <courseId> <userId>`
- `aim-google classroom students add <courseId> <userId> [--enrollment-code CODE]`
- `aim-google classroom students remove <courseId> <userId>`
- `aim-google classroom teachers <courseId> [--max N] [--page TOKEN]`
- `aim-google classroom teachers get <courseId> <userId>`
- `aim-google classroom teachers add <courseId> <userId>`
- `aim-google classroom teachers remove <courseId> <userId>`
- `aim-google classroom roster <courseId> [--students] [--teachers]`
- `aim-google classroom coursework <courseId> [--state ...] [--topic TOPIC_ID] [--scan-pages N] [--max N] [--page TOKEN]`
- `aim-google classroom coursework get <courseId> <courseworkId>`
- `aim-google classroom coursework create <courseId> --title TITLE [--type ASSIGNMENT|...]`
- `aim-google classroom coursework update <courseId> <courseworkId> [--title ...]`
- `aim-google classroom coursework delete <courseId> <courseworkId>`
- `aim-google classroom coursework assignees <courseId> <courseworkId> [--mode ...] [--add-student ...]`
- `aim-google classroom materials <courseId> [--state ...] [--topic TOPIC_ID] [--scan-pages N] [--max N] [--page TOKEN]`
- `aim-google classroom materials get <courseId> <materialId>`
- `aim-google classroom materials create <courseId> --title TITLE`
- `aim-google classroom materials update <courseId> <materialId> [--title ...]`
- `aim-google classroom materials delete <courseId> <materialId>`
- `aim-google classroom submissions <courseId> <courseworkId> [--state ...] [--max N] [--page TOKEN]`
- `aim-google classroom submissions get <courseId> <courseworkId> <submissionId>`
- `aim-google classroom submissions turn-in <courseId> <courseworkId> <submissionId>`
- `aim-google classroom submissions reclaim <courseId> <courseworkId> <submissionId>`
- `aim-google classroom submissions return <courseId> <courseworkId> <submissionId>`
- `aim-google classroom submissions grade <courseId> <courseworkId> <submissionId> [--draft N] [--assigned N]`
- `aim-google classroom announcements <courseId> [--state ...] [--max N] [--page TOKEN]`
- `aim-google classroom announcements get <courseId> <announcementId>`
- `aim-google classroom announcements create <courseId> --text TEXT`
- `aim-google classroom announcements update <courseId> <announcementId> [--text ...]`
- `aim-google classroom announcements delete <courseId> <announcementId>`
- `aim-google classroom announcements assignees <courseId> <announcementId> [--mode ...]`
- `aim-google classroom topics <courseId> [--max N] [--page TOKEN]`
- `aim-google classroom topics get <courseId> <topicId>`
- `aim-google classroom topics create <courseId> --name NAME`
- `aim-google classroom topics update <courseId> <topicId> --name NAME`
- `aim-google classroom topics delete <courseId> <topicId>`
- `aim-google classroom invitations [--course ID] [--user ID]`
- `aim-google classroom invitations get <invitationId>`
- `aim-google classroom invitations create <courseId> <userId> --role STUDENT|TEACHER|OWNER`
- `aim-google classroom invitations accept <invitationId>`
- `aim-google classroom invitations delete <invitationId>`
- `aim-google classroom guardians <studentId> [--max N] [--page TOKEN]`
- `aim-google classroom guardians get <studentId> <guardianId>`
- `aim-google classroom guardians delete <studentId> <guardianId>`
- `aim-google classroom guardian-invitations <studentId> [--state ...] [--max N] [--page TOKEN]`
- `aim-google classroom guardian-invitations get <studentId> <invitationId>`
- `aim-google classroom guardian-invitations create <studentId> --email EMAIL`
- `aim-google classroom profile [userId]`
- `aim-google gmail search <query> [--max N] [--page TOKEN]`
- `aim-google gmail messages search <query> [--max N] [--page TOKEN] [--include-body] [--full]`
- `aim-google gmail autoreply <query> [--max N] [--subject S] [--body B|--body-file PATH|--body-html HTML] [--from addr] [--reply-to addr] [--label L] [--archive] [--mark-read] [--skip-bulk] [--allow-self]`
- `aim-google gmail thread get <threadId> [--download]`
- `aim-google gmail thread modify <threadId> [--add ...] [--remove ...]`
- `aim-google gmail get <messageId> [--format full|metadata|raw] [--headers ...]`
- `aim-google gmail attachment <messageId> <attachmentId> [--out PATH] [--name NAME]`
- `aim-google gmail url <threadIds...>`
- `aim-google gmail forward <messageId> --to a@b.com [--cc ...] [--bcc ...] [--note TEXT|--note-file PATH] [--from addr] [--skip-attachments]`
- `aim-google gmail labels list`
- `aim-google gmail labels get <labelIdOrName>`
- `aim-google gmail labels create <name>`
- `aim-google gmail labels rename <labelIdOrName> <newName>`
- `aim-google gmail labels modify <threadIds...> [--add ...] [--remove ...]`
- `aim-google gmail send --to a@b.com --subject S [--body B] [--body-html H] [--cc ...] [--bcc ...] [--reply-to-message-id <messageId>] [--reply-to addr] [--attach <file>...]`
- `aim-google gmail drafts list [--max N] [--page TOKEN]`
- `aim-google gmail drafts get <draftId> [--download]`
- `aim-google gmail drafts create --subject S [--to a@b.com] [--body B] [--body-html H] [--cc ...] [--bcc ...] [--reply-to-message-id <messageId>] [--reply-to addr] [--attach <file>...]`
- `aim-google gmail drafts update <draftId> --subject S [--to a@b.com] [--body B] [--body-html H] [--cc ...] [--bcc ...] [--reply-to-message-id <messageId>] [--reply-to addr] [--attach <file>...]`
- `aim-google gmail drafts send <draftId>`
- `aim-google gmail drafts delete <draftId>`
- `aim-google gmail watch start|status|renew|stop|serve`
- `aim-google gmail history --since <historyId>`
- `aim-google chat spaces list [--max N] [--page TOKEN]`
- `aim-google chat spaces find <displayName> [--max N] [--exact]`
- `aim-google chat spaces create <displayName> [--member email,...]`
- `aim-google chat messages list <space> [--max N] [--page TOKEN] [--order ORDER] [--thread THREAD] [--unread]`
- `aim-google chat messages send <space> --text TEXT [--thread THREAD]`
- `aim-google chat threads list <space> [--max N] [--page TOKEN]`
- `aim-google chat dm space <email>`
- `aim-google chat dm send <email> --text TEXT [--thread THREAD]`
- `aim-google tasks lists [--max N] [--page TOKEN]`
- `aim-google tasks lists create <title>`
- `aim-google tasks list <tasklistId> [--max N] [--page TOKEN]`
- `aim-google tasks get <tasklistId> <taskId>`
- `aim-google tasks add <tasklistId> --title T [--notes N] [--due RFC3339|YYYY-MM-DD] [--repeat daily|weekly|monthly|yearly] [--repeat-count N] [--repeat-until DT] [--parent ID] [--previous ID]`
- `aim-google tasks update <tasklistId> <taskId> [--title T] [--notes N] [--due RFC3339|YYYY-MM-DD] [--status needsAction|completed]`
- `aim-google tasks done <tasklistId> <taskId>`
- `aim-google tasks undo <tasklistId> <taskId>`
- `aim-google tasks delete <tasklistId> <taskId>`
- `aim-google tasks clear <tasklistId>`
- `aim-google contacts search <query> [--max N]`
- `aim-google contacts list [--max N] [--page TOKEN]`
- `aim-google contacts get <people/...|email>`
- `aim-google contacts create --given NAME [--family NAME] [--email addr] [--phone num] [--relation type=person]`
- `aim-google contacts update <people/...> [--given NAME] [--family NAME] [--email addr] [--phone num] [--birthday YYYY-MM-DD] [--notes TEXT] [--relation type=person] [--from-file PATH|-] [--ignore-etag]`
- `aim-google contacts delete <people/...>`
- `aim-google contacts directory list [--max N] [--page TOKEN]`
- `aim-google contacts directory search <query> [--max N] [--page TOKEN]`
- `aim-google contacts other list [--max N] [--page TOKEN]`
- `aim-google contacts other search <query> [--max N]`
- `aim-google people me`
- `aim-google people get <people/...|userId>`
- `aim-google people search <query> [--max N] [--page TOKEN]`
- `aim-google people relations [<people/...|userId>] [--type TYPE]`

Date/time input conventions (shared parser):

- Date-only: `YYYY-MM-DD`
- Datetime: `RFC3339` / `RFC3339Nano` / `YYYY-MM-DDTHH:MM[:SS]` / `YYYY-MM-DD HH:MM[:SS]`
- Numeric timezone offset accepted: `YYYY-MM-DDTHH:MM:SS-0800`
- Calendar range flags also accept relatives: `now`, `today`, `tomorrow`, `yesterday`, weekday names (`monday`, `next friday`)
- Tracking `--since` also accepts durations like `24h`

### Planned high-level command tree

- `aim-google auth …`
  - `aim-google auth credentials <credentials.json>`
  - `aim-google auth credentials list`
  - `aim-google --client <name> auth credentials <credentials.json>`
- `aim-google gmail …`
- `aim-google chat …`
- `aim-google calendar …`
- `aim-google drive …`
- `aim-google contacts …`
- `aim-google tasks …`
- `aim-google people …`

Planned service identifiers (canonical):

- `gmail`
- `calendar`
- `chat`
- `drive`
- `contacts`
- `tasks`
- `people`

## Google API dependencies (planned)

- `golang.org/x/oauth2`
- `golang.org/x/oauth2/google`
- `google.golang.org/api/option`
- `google.golang.org/api/gmail/v1`
- `google.golang.org/api/calendar/v3`
- `google.golang.org/api/chat/v1`
- `google.golang.org/api/drive/v3`
- `google.golang.org/api/people/v1`
- `google.golang.org/api/tasks/v1`

## Scopes (planned)

We store a single refresh token per Google account email.

- `aim-google auth add` requests a union of scopes based on `--services`.
- Each API client refreshes an access token for the subset of scopes needed for that service.
- If you later want additional services, re-run `aim-google auth add <email> --services ...` (may require `--force-consent` to mint a new refresh token).

- Gmail: `https://mail.google.com/` (or narrower scopes if we decide later)
- Calendar: `https://www.googleapis.com/auth/calendar`
- Chat:
  - `https://www.googleapis.com/auth/chat.spaces`
  - `https://www.googleapis.com/auth/chat.messages`
  - `https://www.googleapis.com/auth/chat.memberships`
  - `https://www.googleapis.com/auth/chat.users.readstate.readonly`
- Drive: `https://www.googleapis.com/auth/drive`
- Contacts/Directory:
  - `https://www.googleapis.com/auth/contacts`
  - `https://www.googleapis.com/auth/contacts.other.readonly`
  - `https://www.googleapis.com/auth/directory.readonly`
- People:
  - `profile` (OIDC)

## Output formats

Default: human-friendly tables (stdlib `text/tabwriter`).

- Parseable stdout:
  - `--json`: JSON objects/arrays suitable for scripting
  - `--plain`: stable TSV (tabs preserved; no alignment; no colors)
- Human-facing hints/progress are written to stderr so stdout can be safely captured.
- Colors are only used for human-facing output and are disabled automatically for `--json` and `--plain`.

We avoid heavy table deps unless we decide we need them.

## Code layout (current)

- `cmd/aim-google/main.go` — binary entrypoint
- `internal/cmd/*` — kong command structs
- `internal/ui/*` — color + printing
- `internal/config/*` — config paths + credential parsing/writing
- `internal/secrets/*` — keyring store

## Formatting, linting, tests

### Formatting

Pinned tools, installed into local `.tools/` via `make tools`:

- `mvdan.cc/gofumpt@v0.7.0`
- `golang.org/x/tools/cmd/goimports@v0.38.0`
- `github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2`

Commands:

- `make fmt` — applies `goimports` + `gofumpt`
- `make fmt-check` — formats and fails if Go files or `go.mod/go.sum` change

### Lint

- `golangci-lint` with config in `.golangci.yml`
- `make lint`

### Tests

- stdlib `testing` (+ `httptest` when we add OAuth/API tests)
- `make test`

### Integration tests (local only)

There is an opt-in integration test suite guarded by build tags (not run in CI).

- Requires:
  - stored `credentials.json` (or `credentials-<client>.json`) via `aim-google auth credentials ...`
  - refresh token in keyring via `aim-google auth add <email>`
- Run:
  - `AIM_GOOGLE_IT_ACCOUNT=you@gmail.com go test -tags=integration ./internal/integration`
  - optional: `AIM_GOOGLE_CLIENT=work` to select a non-default OAuth client

## CI (GitHub Actions)

Workflow: `.github/workflows/ci.yml`

- runs on push + PR
- uses `actions/setup-go` with `go-version-file: go.mod`
- runs:
  - `make tools`
  - `make fmt-check`
  - `go test ./...`
  - `golangci-lint` (pinned `v1.62.2`)

## Next implementation steps

- Expand Gmail further (labels by name everywhere, richer body rendering, compose edge cases).
- Improve People updates (multi-field + richer contact data).
- Harden UX (consistent output formats, retries/backoff on specific transient errors).
