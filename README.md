# 🧭 aim-google — Google in your terminal.

![GitHub Repo Banner](https://ghrb.waren.build/banner?header=aim-google%F0%9F%A7%AD&subheader=Google+in+your+terminal&bg=f3f4f6&color=1f2937&support=true)
<!-- Created with GitHub Repo Banner by Waren Gonzaga: https://ghrb.waren.build -->

Fast, script-friendly CLI for Gmail, Calendar, Chat, Classroom, Drive, Docs, Slides, Sheets, Forms, Apps Script, Contacts, Tasks, People, Admin, Groups (Workspace), and Keep (Workspace-only). JSON-first output, multiple accounts, and flexible auth built in.

## Features

- **Gmail** - search threads/messages, send mail, view attachments, manage labels/drafts/filters/delegation/vacation settings, auto-reply once to matching mail, modify single messages, export filters, inspect history, and run Pub/Sub watch webhooks
- **Email tracking** - track opens for `aim-google gmail send --track` with a small Cloudflare Worker backend
- **Calendar** - list/create/update/delete events, manage invitations, aliases, subscriptions, team calendars, free/busy/conflicts, propose new times, focus/OOO/working-location events, recurrence, and reminders
- **Classroom** - manage courses, roster, coursework/materials, submissions, announcements, topics, invitations, guardians, profiles
- **Chat** - list/find/create spaces, list messages/threads, send messages and DMs, and manage emoji reactions (Workspace-only)
- **Drive** - list/search/upload/download files, replace uploads in-place, convert uploads (including Markdown to Google Doc), manage permissions/comments, organize folders, and list shared drives
- **Contacts** - search/create/update contacts, including addresses, relations, org/title metadata, custom fields, Workspace directory, and other contacts
- **Tasks** - manage tasklists and tasks: get/create/add/update/done/undo/delete/clear, plus repeat schedule materialization with RRULE aliases
- **Sheets** - read/write/update spreadsheets, insert rows/cols, manage tabs and named ranges, format/merge/freeze/resize cells, read/write notes, inspect formats, find/replace text, list links, and create/export sheets
- **Forms** - create/update forms, manage questions, inspect responses, and manage watches
- **Apps Script** - create/get/bind projects, inspect content, and run functions
- **Docs/Slides** - create/copy/export docs/slides, edit Docs by tab, import Markdown, do richer find-replace, export Docs as Markdown/HTML, and generate Slides from Markdown or templates
- **People** - profile lookup and directory search helpers
- **Keep (Workspace only)** - list/get/search/create/delete notes and download attachments (service account + domain-wide delegation)
- **Admin (Workspace only)** - Workspace Admin users/groups commands for common directory operations
- **Groups** - list groups you belong to, view group members (Google Workspace)
- **Local time** - quick local/UTC time display for scripts and agents
- **Multiple accounts** - manage multiple Google accounts simultaneously, with account aliases and per-client OAuth buckets
- **Command allowlist** - restrict top-level commands for sandboxed/agent runs
- **Secure credential storage** using OS keyring or encrypted on-disk keyring (configurable)
- **Auto-refreshing tokens** - authenticate once, use indefinitely
- **Flexible auth** - OAuth refresh tokens, ADC, direct access tokens, service accounts, manual/remote flows, `--extra-scopes`, and proxy-safe callbacks
- **Least-privilege auth** - `--readonly`, `--drive-scope`, and `--gmail-scope` to request fewer scopes
- **Workspace service accounts** - domain-wide delegation auth (preferred when configured)
- **Parseable output** - JSON mode for scripting and automation (Calendar adds day-of-week fields)

## Installation

### Homebrew

```bash
brew install aim-google
```
### Arch User Repository

```bash
yay -S aim-google
```

### Build from Source

```bash
git clone https://github.com/BrianV1981/aim-google.git
cd aim-google
make
```

Run:

```bash
./bin/aim-google --help
```

Help:

- `aim-google --help` shows top-level command groups.
- Drill down with `aim-google <group> --help` (and deeper subcommands).
- For the full expanded command list: `AIM_GOOGLE_HELP=full aim-google --help`.
- Make shortcut: `make aim-google -- --help` (or `make aim-google -- gmail --help`).
- `make aim-google-help` shows CLI help (note: `make aim-google --help` is Make’s own help; use `--`).
- Version: `aim-google --version` or `aim-google version`.

## Quick Start

### 1. Get OAuth2 Credentials

Before adding an account, create OAuth2 credentials from Google Cloud Console:

1. Open the Google Cloud Console credentials page: https://console.cloud.google.com/apis/credentials
1. Create a project: https://console.cloud.google.com/projectcreate
2. Enable the APIs you need:
   - Admin SDK API: https://console.cloud.google.com/apis/api/admin.googleapis.com
   - Apps Script API: https://console.cloud.google.com/apis/api/script.googleapis.com
   - Cloud Identity API (Groups): https://console.cloud.google.com/apis/api/cloudidentity.googleapis.com
   - Gmail API: https://console.cloud.google.com/apis/api/gmail.googleapis.com
   - Google Calendar API: https://console.cloud.google.com/apis/api/calendar-json.googleapis.com
   - Google Chat API: https://console.cloud.google.com/apis/api/chat.googleapis.com
   - Google Docs API: https://console.cloud.google.com/apis/api/docs.googleapis.com
   - Google Drive API: https://console.cloud.google.com/apis/api/drive.googleapis.com
   - Google Classroom API: https://console.cloud.google.com/apis/api/classroom.googleapis.com
   - Google Keep API: https://console.cloud.google.com/apis/api/keep.googleapis.com
   - People API (Contacts): https://console.cloud.google.com/apis/api/people.googleapis.com
   - Google Tasks API: https://console.cloud.google.com/apis/api/tasks.googleapis.com
   - Google Sheets API: https://console.cloud.google.com/apis/api/sheets.googleapis.com
   - Google Forms API: https://console.cloud.google.com/apis/api/forms.googleapis.com
   - Google Slides API: https://console.cloud.google.com/apis/api/slides.googleapis.com
3. Configure OAuth consent screen: https://console.cloud.google.com/auth/branding
4. If your app is in "Testing", add test users: https://console.cloud.google.com/auth/audience
5. Create OAuth client:
   - Go to https://console.cloud.google.com/auth/clients
   - Click "Create Client"
   - Application type: "Desktop app"
   - Download the JSON file (usually named `client_secret_....apps.googleusercontent.com.json`)

### 2. Store Credentials

```bash
aim-google auth credentials ~/Downloads/client_secret_....json
```

For multiple OAuth clients/projects:

```bash
aim-google --client work auth credentials ~/Downloads/work-client.json
aim-google auth credentials list
```

### 3. Provision Sovereign Credentials

**⚠️ CRITICAL A.I.M. DIRECTIVE:** 
This fork (`aim-google`) is hardcoded to remain in Google's "Testing" publishing status. We do NOT provide a global, verified OAuth Client ID for the public. You **must** provision your own private credentials. 

Why? Because true data sovereignty means owning your own gateway. Relying on a third-party OAuth client (like the original `gogcli`) means the creator has the power to revoke your access at any time, or access your data if scopes are mishandled. 

To use this Swarm tool, you will need to spend 5 minutes clicking through the Google Cloud Console to generate your own `credentials.json` file.

**How to create your own aim-google credentials:**
1. Go to the [Google Cloud Console](https://console.cloud.google.com/projectcreate) and create a new project (e.g., `aim-swarm-gateway`).
2. Go to **APIs & Services -> Enabled APIs and services**. Enable the APIs for the services you want the agent to use (e.g., Gmail API, Google Drive API, Google Calendar API, Google Sheets API).
3. Go to **APIs & Services -> OAuth consent screen**.
   - Select **External** (or Internal if you are a Workspace admin).
   - Fill in the required app information.
   - For **Test users**, add the specific Gmail addresses you intend to authorize. (Since the app is in testing mode, only these users can log in).
4. Go to **APIs & Services -> Credentials**.
   - Click **Create Credentials -> OAuth client ID**.
   - Choose **Desktop app**. Name it `aim-google-cli`.
   - Download the JSON file.

Import your sovereign credentials into the CLI:
```bash
aim-google auth credentials /path/to/your/client_secret.json
```

### 4. Authorize Your Account for Autonomous Swarm Agents

To enable headless, background A.I.M. Swarm agents to autonomously use your Google Workspace without hanging on interactive password prompts, you must configure the CLI to use an encrypted file-based keyring and provide the decryption password via an environment variable.

1. **Tell aim-google to use the headless file keyring:**
   ```bash
   aim-google auth keyring file
   ```
2. **Set your encryption password:** Add this to your `~/.bashrc` (or agent environment profile) so the Swarm can autonomously decrypt the OAuth token. (Replace with your actual chosen password):
   ```bash
   export AIM_GOOGLE_KEYRING_PASSWORD="your_secure_password"
   ```
3. **Import your sovereign credentials into this new backend:**
   ```bash
   aim-google auth credentials /path/to/your/client_secret.json
   ```
4. **Authorize the account** (This will open a browser window or provide a remote link):
   ```bash
   aim-google auth add you@gmail.com
   ```

Once authorized, your Swarm agents can silently execute commands in the background without human intervention!

### 5. Install the Bundled A.I.M. Swarm Skill

To transform your autonomous Gemini CLI agent into a Google Workspace expert, install the bundled `.skill` package included in this repository. This provides your agent with explicit instructions on how to natively execute the Go binary, respect the exponential backoff, and use the `--agent` flag for extreme token efficiency.

```bash
gemini skills install agent-skill/aim-google.skill --scope workspace
```
*(After installing, remember to run `/skills reload` in any active agent chats).*

Headless / remote server flows (no browser on the server):

Manual interactive flow (recommended):

```bash
aim-google auth add you@gmail.com --services user --manual
```

- The CLI prints an auth URL. Open it in a local browser.
- After approval, copy the full loopback redirect URL from the browser address bar.
- Paste that URL back into the terminal when prompted.

Split remote flow (`--remote`, useful for two-step/scripted handoff):

```bash
# Step 1: print auth URL (open it locally in a browser)
aim-google auth add you@gmail.com --services user --remote --step 1

# Step 2: paste the full redirect URL from your browser address bar
aim-google auth add you@gmail.com --services user --remote --step 2 --auth-url 'http://127.0.0.1:<port>/oauth2/callback?code=...&state=...'
```

- The `state` is cached on disk for a short time (about 10 minutes). If it expires, rerun step 1.
- Remote step 2 requires a redirect URL that includes `state` (state check mandatory).

Browser OAuth behind proxies / remote tunnels:

```bash
aim-google auth add you@gmail.com --listen-addr 0.0.0.0:8080 --redirect-host aim-google.example.com
aim-google auth manage --listen-addr 0.0.0.0:8080 --redirect-host aim-google.example.com
```

- `--listen-addr` changes where the local callback server binds.
- `--redirect-host` builds `https://<host>/oauth2/callback` for the OAuth redirect URI.
- The redirect URI must also be registered in your OAuth client settings.

Direct access token flow (headless/CI, no stored refresh token):

```bash
aim-google --access-token "$(gcloud auth print-access-token)" gmail labels list
```

- Also available as `AIM_GOOGLE_ACCESS_TOKEN`
- Bypasses stored refresh tokens and keyring lookup
- Token expires in about 1 hour; no auto-refresh

### 4. Test Authentication

```bash
export AIM_GOOGLE_ACCOUNT=you@gmail.com
aim-google gmail labels list
```

## Authentication & Secrets

### Accounts and tokens

`aim-google` stores your OAuth refresh tokens in a “keyring” backend. Default is `auto` (best available backend for your OS/environment).

Before you can run `aim-google auth add`, you must store OAuth client credentials once via `aim-google auth credentials <credentials.json>` (download a Desktop app OAuth client JSON from the Cloud Console). For multiple clients, use `aim-google --client <name> auth credentials ...`; tokens are isolated per client.

List accounts:

```bash
aim-google auth list
```

Verify tokens are usable (helps spot revoked/expired tokens):

```bash
aim-google auth list --check
```

Accounts can be authorized either via OAuth refresh tokens or Workspace service accounts (domain-wide delegation). If a service account key is configured for an account, it takes precedence over OAuth refresh tokens (see `aim-google auth list`).

Show current auth state/services for the active account:

```bash
aim-google auth status
```

### Multiple OAuth clients

Use `--client` (or `AIM_GOOGLE_CLIENT`) to select a named OAuth client:

```bash
aim-google --client work auth credentials ~/Downloads/work.json
aim-google --client work auth add you@company.com
```

Optional domain mapping for auto-selection:

```bash
aim-google --client work auth credentials ~/Downloads/work.json --domain example.com
```

How it works:

- Default client is `default` (stored in `credentials.json`).
- Named clients are stored as `credentials-<client>.json`.
- Tokens are isolated per client (`token:<client>:<email>`); defaults are per client too.

Client selection order (when `--client` is not set):

1) `--client` / `AIM_GOOGLE_CLIENT`
2) `account_clients` config (email -> client)
3) `client_domains` config (domain -> client)
4) Credentials file named after the email domain (`credentials-example.com.json`)
5) `default`

Config example (JSON5):

```json5
{
  account_clients: { "you@company.com": "work" },
  client_domains: { "example.com": "work" },
}
```

List stored credentials:

```bash
aim-google auth credentials list
```

See `docs/auth-clients.md` for the full client selection and mapping rules.

### Keyring backend: Keychain vs encrypted file

Backends:

- `auto` (default): picks the best backend for the platform.
- `keychain`: macOS Keychain (recommended on macOS; avoids password management).
- `file`: encrypted on-disk keyring (requires a password).

Set backend via command (writes `keyring_backend` into `config.json`):

```bash
aim-google auth keyring file
aim-google auth keyring keychain
aim-google auth keyring auto
```

Show current backend + source (env/config/default) and config path:

```bash
aim-google auth keyring
```

Non-interactive runs (CI/ssh): file backend requires `AIM_GOOGLE_KEYRING_PASSWORD`.

```bash
export AIM_GOOGLE_KEYRING_PASSWORD='...'
aim-google --no-input auth status
```

Force backend via env (overrides config):

```bash
export AIM_GOOGLE_KEYRING_BACKEND=file
```

Precedence: `AIM_GOOGLE_KEYRING_BACKEND` env var overrides `config.json`.

## Configuration

### Account Selection

Specify the account using either a flag or environment variable:

```bash
# Via flag
aim-google gmail search 'newer_than:7d' --account you@gmail.com

# Via alias
aim-google auth alias set work work@company.com
aim-google gmail search 'newer_than:7d' --account work

# Via environment
export AIM_GOOGLE_ACCOUNT=you@gmail.com
aim-google gmail search 'newer_than:7d'

# Auto-select (default account or the single stored token)
aim-google gmail labels list --account auto
```

List configured accounts:

```bash
aim-google auth list
```

### Output

- Default: human-friendly tables on stdout.
- `--plain`: stable TSV on stdout (tabs preserved; best for piping to tools that expect `\t`).
- `--json`: JSON on stdout (best for scripting).
- Human-facing hints/progress go to stderr.
- Colors are enabled only in rich TTY output and are disabled automatically for `--json` and `--plain`.

### Service Scopes

By default, `aim-google auth add` requests access to the **user** services (see `aim-google auth services` for the current list and scopes).

To request fewer scopes:

```bash
aim-google auth add you@gmail.com --services drive,calendar
```

To request read-only scopes (write operations will fail with 403 insufficient scopes):

```bash
aim-google auth add you@gmail.com --services drive,calendar --readonly
```

To control Drive’s scope (default: `full`):

```bash
aim-google auth add you@gmail.com --services drive --drive-scope full
aim-google auth add you@gmail.com --services drive --drive-scope readonly
aim-google auth add you@gmail.com --services drive --drive-scope file
```

To control Gmail’s scope (default: `full`):

```bash
aim-google auth add you@gmail.com --services gmail --gmail-scope full
aim-google auth add you@gmail.com --services gmail --gmail-scope readonly
# Example: readonly on both Gmail and Drive
aim-google auth add you@gmail.com --services gmail,drive --gmail-scope readonly --drive-scope readonly
# Example: append one custom scope beyond the built-in Gmail scope set
aim-google auth add you@gmail.com --services gmail --extra-scopes https://www.googleapis.com/auth/gmail.labels
```

Notes:

- `--drive-scope readonly` is enough for listing/downloading/exporting via Drive (write operations will 403).
- `--drive-scope file` is write-capable (limited to files created/opened by this app) and can’t be combined with `--readonly`.
- `--gmail-scope readonly` requests `gmail.readonly` (no modify/settings write scopes).
- `--extra-scopes` appends additional OAuth scope URIs after the built-in service scope set; remote step-1 guidance replays it so step 2 requests the same token.
- For `--readonly`, `--drive-scope readonly|file`, or `--gmail-scope readonly`, auth disables Google `include_granted_scopes` to prevent old broader grants from silently accumulating.

If you need to add services later and Google doesn't return a refresh token, re-run with `--force-consent`:

```bash
aim-google auth add you@gmail.com --services user --force-consent
# Or add just Sheets
aim-google auth add you@gmail.com --services sheets --force-consent
```

`--services all` is accepted as an alias for `user` for backwards compatibility.

Docs commands are implemented via the Drive API, and `docs` requests both Drive and Docs API scopes.

Service scope matrix (auto-generated; run `go run scripts/gen-auth-services-md.go`):

<!-- auth-services:start -->
| Service | User | APIs | Scopes | Notes |
| --- | --- | --- | --- | --- |
| gmail | yes | Gmail API | `https://www.googleapis.com/auth/gmail.modify`<br>`https://www.googleapis.com/auth/gmail.settings.basic`<br>`https://www.googleapis.com/auth/gmail.settings.sharing` |  |
| calendar | yes | Calendar API | `https://www.googleapis.com/auth/calendar` |  |
| chat | yes | Chat API | `https://www.googleapis.com/auth/chat.spaces`<br>`https://www.googleapis.com/auth/chat.messages`<br>`https://www.googleapis.com/auth/chat.memberships`<br>`https://www.googleapis.com/auth/chat.users.readstate.readonly` |  |
| classroom | yes | Classroom API | `https://www.googleapis.com/auth/classroom.courses`<br>`https://www.googleapis.com/auth/classroom.rosters`<br>`https://www.googleapis.com/auth/classroom.coursework.students`<br>`https://www.googleapis.com/auth/classroom.coursework.me`<br>`https://www.googleapis.com/auth/classroom.courseworkmaterials`<br>`https://www.googleapis.com/auth/classroom.announcements`<br>`https://www.googleapis.com/auth/classroom.topics`<br>`https://www.googleapis.com/auth/classroom.guardianlinks.students`<br>`https://www.googleapis.com/auth/classroom.profile.emails`<br>`https://www.googleapis.com/auth/classroom.profile.photos` |  |
| drive | yes | Drive API | `https://www.googleapis.com/auth/drive` |  |
| docs | yes | Docs API, Drive API | `https://www.googleapis.com/auth/drive`<br>`https://www.googleapis.com/auth/documents` | Export/copy/create via Drive |
| slides | yes | Slides API, Drive API | `https://www.googleapis.com/auth/drive`<br>`https://www.googleapis.com/auth/presentations` | Create/edit presentations |
| contacts | yes | People API | `https://www.googleapis.com/auth/contacts`<br>`https://www.googleapis.com/auth/contacts.other.readonly`<br>`https://www.googleapis.com/auth/directory.readonly` | Contacts + other contacts + directory |
| tasks | yes | Tasks API | `https://www.googleapis.com/auth/tasks` |  |
| sheets | yes | Sheets API, Drive API | `https://www.googleapis.com/auth/drive`<br>`https://www.googleapis.com/auth/spreadsheets` | Export via Drive |
| people | yes | People API | `profile` | OIDC profile scope |
| forms | yes | Forms API | `https://www.googleapis.com/auth/forms.body`<br>`https://www.googleapis.com/auth/forms.responses.readonly` |  |
| appscript | yes | Apps Script API | `https://www.googleapis.com/auth/script.projects`<br>`https://www.googleapis.com/auth/script.deployments`<br>`https://www.googleapis.com/auth/script.processes` |  |
| ads | yes | Google Ads API | `https://www.googleapis.com/auth/adwords` | OAuth scope only |
| groups | no | Cloud Identity API | `https://www.googleapis.com/auth/cloud-identity.groups.readonly` | Workspace only |
| keep | no | Keep API | `https://www.googleapis.com/auth/keep` | Workspace only; service account (domain-wide delegation) |
| admin | no | Admin SDK Directory API | `https://www.googleapis.com/auth/admin.directory.user`<br>`https://www.googleapis.com/auth/admin.directory.group`<br>`https://www.googleapis.com/auth/admin.directory.group.member` | Workspace only; service account with domain-wide delegation required |
<!-- auth-services:end -->

### Service Accounts (Workspace only)

A service account is a non-human Google identity that belongs to a Google Cloud project. In Google Workspace, a service account can impersonate a user via **domain-wide delegation** (admin-controlled) and access APIs like Gmail/Calendar/Drive as that user.

In `aim-google`, service accounts are an **optional auth method** that can be configured per account email. If a service account key is configured for an account, it takes precedence over OAuth refresh tokens (see `aim-google auth list`).

#### 1) Create a Service Account (Google Cloud)

1. Create (or pick) a Google Cloud project.
2. Enable the APIs you’ll use (e.g. Gmail, Calendar, Drive, Sheets, Docs, People, Tasks, Cloud Identity).
3. Go to **IAM & Admin → Service Accounts** and create a service account.
4. In the service account details, enable **Domain-wide delegation**.
5. Create a key (**Keys → Add key → Create new key → JSON**) and download the JSON key file.

#### 2) Allowlist scopes (Google Workspace Admin Console)

Domain-wide delegation is enforced by Workspace admin settings.

1. Open **Admin console → Security → API controls → Domain-wide delegation**.
2. Add a new API client:
   - Client ID: use the service account’s “Client ID” from Google Cloud.
   - OAuth scopes: comma-separated list of scopes you want to allow (copy from `aim-google auth services` and/or your `aim-google auth add --services ...` usage).

If a scope is missing from the allowlist, service-account token minting can fail (or API calls will 403 with insufficient permissions).

#### 3) Configure `aim-google` to use the service account

Store the key for the user you want to impersonate:

```bash
aim-google auth service-account set you@yourdomain.com --key ~/Downloads/service-account.json
```

Verify `aim-google` is preferring the service account for that account:

```bash
aim-google --account you@yourdomain.com auth status
aim-google auth list
```

### Google Keep (Workspace only)

Keep requires Workspace + domain-wide delegation. You can configure it via the generic service-account command above (recommended), or the legacy Keep helper:

```bash
aim-google auth service-account set you@yourdomain.com --key ~/Downloads/service-account.json
aim-google keep list --account you@yourdomain.com
aim-google keep get <noteId> --account you@yourdomain.com
aim-google keep create --title "Todo" --item "Milk" --item "Eggs" --account you@yourdomain.com
aim-google keep delete <noteId> --account you@yourdomain.com --force
```

### Environment Variables

- `AIM_GOOGLE_ACCOUNT` - Default account email or alias to use (avoids repeating `--account`; otherwise uses keyring default or a single stored token)
- `AIM_GOOGLE_ACCESS_TOKEN` - Use a provided access token directly (headless/CI; no auto-refresh)
- `AIM_GOOGLE_CLIENT` - OAuth client name (selects stored credentials + token bucket)
- `AIM_GOOGLE_JSON` - Default JSON output
- `AIM_GOOGLE_PLAIN` - Default plain output
- `AIM_GOOGLE_COLOR` - Color mode: `auto` (default), `always`, or `never`
- `AIM_GOOGLE_TIMEZONE` - Default output timezone for Calendar/Gmail (IANA name, `UTC`, or `local`)
- `AIM_GOOGLE_ENABLE_COMMANDS` - Comma-separated allowlist of commands; dot paths allowed (e.g., `calendar,tasks,gmail.search`)
- `AIM_GOOGLE_DISABLE_COMMANDS` - Comma-separated denylist of commands; dot paths allowed (e.g., `gmail.send,gmail.drafts.send`)
- `AIM_GOOGLE_GMAIL_NO_SEND` - Block Gmail send operations
- `AIM_GOOGLE_KEYRING_SERVICE_NAME` - Override the keyring namespace/service name (default: `aim-google`)

### Config File (JSON5)

Find the actual config path in `aim-google --help` or `aim-google auth keyring`.

Typical paths:

- macOS: `~/Library/Application Support/aim-google/config.json`
- Linux: `~/.config/aim-google/config.json` (or `$XDG_CONFIG_HOME/aim-google/config.json`)
- Windows: `%AppData%\\aim-google\\config.json`

Example (JSON5 supports comments and trailing commas):

```json5
{
  // Avoid macOS Keychain prompts
  keyring_backend: "file",
  // Default output timezone for Calendar/Gmail (IANA, UTC, or local)
  default_timezone: "UTC",
  // Optional account aliases
  account_aliases: {
    work: "work@company.com",
    personal: "me@gmail.com",
  },
  // Optional per-account OAuth client selection
  account_clients: {
    "work@company.com": "work",
  },
  // Optional domain -> client mapping
  client_domains: {
    "example.com": "work",
  },
  // Optional safety guard: block Gmail send operations
  gmail_no_send: true,
  no_send_accounts: {
    "agent@example.com": true,
  },
}
```

### Config Commands

```bash
aim-google config path
aim-google config list
aim-google config keys
aim-google config get timezone
aim-google config set timezone UTC
aim-google config unset timezone
```

### Account Aliases

```bash
aim-google auth alias set work work@company.com
aim-google auth alias list
aim-google auth alias unset work
```

Aliases work anywhere you pass `--account` or `AIM_GOOGLE_ACCOUNT` (reserved: `auto`, `default`).

### Command Guards (Sandboxing)

```bash
# Only allow calendar + tasks commands for an agent
aim-google --enable-commands calendar,tasks calendar events --today

# Allow one Gmail read path, but block Gmail writes
aim-google --enable-commands gmail.search --disable-commands gmail.send gmail search from:me

# Same via env
export AIM_GOOGLE_ENABLE_COMMANDS=calendar,tasks
export AIM_GOOGLE_DISABLE_COMMANDS=gmail.send,gmail.drafts.send
aim-google tasks list <tasklistId>

# Extra Gmail send guard
aim-google --gmail-no-send gmail send --to someone@example.com --subject Test --body Test
aim-google config no-send set agent@example.com
```
 
## Security

### Credential Storage

OAuth credentials are stored securely in your system's keychain:
- **macOS**: Keychain Access
- **Linux**: Secret Service (GNOME Keyring, KWallet)
- **Windows**: Credential Manager

The CLI uses [github.com/99designs/keyring](https://github.com/99designs/keyring) for secure storage.

If no OS keychain backend is available (e.g., Linux/WSL/container), keyring can fall back to an encrypted on-disk store and may prompt for a password; for non-interactive runs set `AIM_GOOGLE_KEYRING_PASSWORD`.

### Keychain Prompts (macOS)

macOS Keychain may prompt more than you’d expect when the “app identity” keeps changing (different binary path, `go run` temp builds, rebuilding to new `./bin/aim-google`, multiple copies). Keychain treats those as different apps, so it asks again.

Options:

- **Default (recommended):** keep using Keychain (secure) and run a stable `aim-google` binary path to reduce repeat prompts.
- **Force Keychain:** `AIM_GOOGLE_KEYRING_BACKEND=keychain` (disables any file-backend fallback).
- **Avoid Keychain prompts entirely:** `AIM_GOOGLE_KEYRING_BACKEND=file` (stores encrypted entries on disk under your config dir).
  - To avoid password prompts too (CI/non-interactive): set `AIM_GOOGLE_KEYRING_PASSWORD=...` (tradeoff: secret in env).
- **Use a separate keyring namespace:** `AIM_GOOGLE_KEYRING_SERVICE_NAME=custom-aim-google` (default: `aim-google`).

### Best Practices

- **Never commit OAuth client credentials** to version control
- Store client credentials outside your project directory
- Use different OAuth clients for development and production
- Re-authorize with `--force-consent` if you suspect token compromise
- Remove unused accounts with `aim-google auth remove <email>`

### OAuth Client IDs in Open Source

Some open source Google CLIs ship a pre-configured OAuth client ID/secret copied from other desktop apps to avoid OAuth consent verification, testing-user limits, or quota issues. This makes the consent screen/security emails show the other app’s name and can stop working at any time.

`aim-google` does not do this. Supported auth:

- Your own OAuth Desktop client JSON via `aim-google auth credentials ...` + `aim-google auth add ...`
- Google Workspace service accounts with domain-wide delegation (Workspace only)

## Commands

Flag aliases:
- `--out` also accepts `--output`.
- `--out-dir` also accepts `--output-dir` (Gmail thread attachment downloads).

### Authentication

```bash
aim-google auth credentials <path>           # Store OAuth client credentials
aim-google auth credentials list             # List stored OAuth client credentials
aim-google auth credentials remove work      # Remove one OAuth client plus its tokens/domain mappings
aim-google auth credentials remove all       # Remove all stored OAuth clients plus their tokens/domain mappings
aim-google --client work auth credentials <path>  # Store named OAuth client credentials
aim-google auth add <email>                  # Authorize and store refresh token
aim-google auth add <email> --services gmail --gmail-scope readonly  # Gmail read-only token
aim-google auth add <email> --listen-addr 0.0.0.0:8080 --redirect-host aim-google.example.com
aim-google auth service-account set <email> --key <path>  # Configure service account impersonation (Workspace only)
aim-google auth service-account status <email>            # Show service account status
aim-google auth service-account unset <email>             # Remove service account
aim-google auth keep <email> --key <path>                 # Legacy alias (Keep)
aim-google auth keyring [backend]            # Show/set keyring backend (auto|keychain|file)
aim-google auth status                       # Show current auth state/services
aim-google auth services                     # List available services and OAuth scopes
aim-google auth list                         # List stored accounts
aim-google auth list --check                 # Validate stored refresh tokens
aim-google auth remove <email>               # Remove a stored refresh token
aim-google auth manage                       # Open accounts manager in browser
aim-google auth manage --listen-addr 0.0.0.0:8080 --redirect-host aim-google.example.com
aim-google auth tokens                       # Manage stored refresh tokens
```

### Keep (Workspace only)

```bash
aim-google keep list --account you@yourdomain.com
aim-google keep get <noteId> --account you@yourdomain.com
aim-google keep search <query> --account you@yourdomain.com
aim-google keep create --title "Todo" --item "Milk" --item "Eggs" --account you@yourdomain.com
aim-google keep create --title "Note" --text "Remember this" --account you@yourdomain.com
aim-google keep delete <noteId> --account you@yourdomain.com --force
aim-google keep attachment <attachmentName> --account you@yourdomain.com --out ./attachment.bin
```

### Gmail

```bash
# Search and read
aim-google gmail search 'newer_than:7d' --max 10
aim-google gmail thread get <threadId>
aim-google gmail thread get <threadId> --download              # Download attachments to current dir
aim-google gmail thread get <threadId> --download --out-dir ./attachments
aim-google gmail get <messageId>
aim-google gmail get <messageId> --format metadata
aim-google gmail attachment <messageId> <attachmentId>
aim-google gmail attachment <messageId> <attachmentId> --out ./attachment.bin
aim-google gmail url <threadId>              # Print Gmail web URL
aim-google gmail thread modify <threadId> --add STARRED --remove INBOX

# Send and compose
aim-google gmail send --to a@b.com --subject "Hi" --body "Plain fallback"
aim-google gmail send --to a@b.com --subject "Hi" --body-file ./message.txt
aim-google gmail send --to a@b.com --subject "Hi" --body-file -   # Read body from stdin
aim-google gmail send --to a@b.com --subject "Hi" --body "Plain fallback" --body-html "<p>Hello</p>"
aim-google gmail forward <messageId> --to a@b.com --note "FYI"
aim-google gmail forward <messageId> --to a@b.com --skip-attachments
# Reply + include quoted original message (auto-generates HTML quote unless you pass --body-html)
aim-google gmail send --reply-to-message-id <messageId> --quote --to a@b.com --subject "Re: Hi" --body "My reply"
# Draft reply + quote (create requires explicit reply target)
aim-google gmail drafts create --reply-to-message-id <messageId> --quote --subject "Re: Hi" --body "My reply"
# Draft reply + quote (update accepts explicit target; else falls back to latest non-draft, non-self message in thread)
aim-google gmail drafts update <draftId> --reply-to-message-id <messageId> --quote --subject "Re: Hi" --body "My reply"
aim-google gmail drafts update <draftId> --quote --subject "Re: Hi" --body "My reply"
aim-google gmail drafts list
aim-google gmail drafts create --subject "Draft" --body "Body"
aim-google gmail drafts create --to a@b.com --subject "Draft" --body "Body"
aim-google gmail drafts update <draftId> --subject "Draft" --body "Body"
aim-google gmail drafts update <draftId> --to a@b.com --subject "Draft" --body "Body"
aim-google gmail drafts send <draftId>
aim-google gmail autoreply 'from:alerts@example.com newer_than:7d' --body-file ./reply.txt --label AutoReplied --dry-run

# Labels
aim-google gmail labels list
aim-google gmail labels get INBOX --json  # Includes message counts
aim-google gmail labels create "My Label"
aim-google gmail labels rename "Old Label" "New Label"
aim-google gmail labels style "My Label" --text-color "#ffffff" --background-color "#4285f4"
aim-google gmail labels modify <threadId> --add STARRED --remove INBOX
aim-google gmail labels delete <labelIdOrName>  # Deletes user label (guards system labels; confirm)

# Batch operations
aim-google gmail batch delete <messageId> <messageId>
aim-google gmail batch modify <messageId> <messageId> --add STARRED --remove INBOX

# Filters
aim-google gmail filters list
aim-google gmail filters create --from 'noreply@example.com' --add-label 'Notifications'
aim-google gmail filters delete <filterId>
aim-google gmail filters export --out ./filters.json

# Settings
aim-google gmail autoforward get
aim-google gmail autoforward enable --email forward@example.com
aim-google gmail autoforward disable
aim-google gmail forwarding list
aim-google gmail forwarding add --email forward@example.com
aim-google gmail sendas list
aim-google gmail sendas create --email alias@example.com
aim-google gmail vacation get
aim-google gmail vacation enable --subject "Out of office" --message "..."
aim-google gmail vacation disable

# Delegation (G Suite/Workspace)
aim-google gmail delegates list
aim-google gmail delegates add --email delegate@example.com
aim-google gmail delegates remove --email delegate@example.com

# Watch (Pub/Sub push)
aim-google gmail watch start --topic projects/<p>/topics/<t> --label INBOX
aim-google gmail watch serve --bind 127.0.0.1 --token <shared> --hook-url http://127.0.0.1:18789/hooks/agent
aim-google gmail watch serve --bind 0.0.0.0 --verify-oidc --oidc-email <svc@...> --hook-url <url>
aim-google gmail watch serve --bind 127.0.0.1 --token <shared> --fetch-delay 5 --hook-url http://127.0.0.1:18789/hooks/agent
aim-google gmail watch serve --bind 127.0.0.1 --token <shared> --exclude-labels SPAM,TRASH --hook-url http://127.0.0.1:18789/hooks/agent
aim-google gmail history --since <historyId>
```

Gmail watch (Pub/Sub push):
- Create Pub/Sub topic + push subscription (OIDC preferred; shared token ok for dev).
- Full flow + payload details: `docs/watch.md`.
- `watch serve --fetch-delay` defaults to `3s` and helps avoid Gmail History indexing races after push delivery.
- `watch serve --exclude-labels` defaults to `SPAM,TRASH`; IDs are case-sensitive.

### Email Tracking

Track when recipients open your emails:

```bash
# Set up local tracking config (per-account; generates keys; follow printed deploy steps)
aim-google gmail track setup --worker-url https://aim-google-email-tracker.<acct>.workers.dev

# Send with tracking
aim-google gmail send --to recipient@example.com --subject "Hello" --body-html "<p>Hi!</p>" --track

# Check opens
aim-google gmail track opens <tracking_id>
aim-google gmail track opens --to recipient@example.com

# View status
aim-google gmail track status
```

Docs: `docs/email-tracking.md` (setup/deploy) + `docs/email-tracking-worker.md` (internals).

**Notes:** `--track` requires exactly 1 recipient (no cc/bcc) and an HTML body (`--body-html` or `--quote`). Use `--track-split` to send per-recipient messages with individual tracking ids. The tracking worker stores IP/user-agent + coarse geo by default.

### Calendar

```bash
# Calendars
aim-google calendar calendars
aim-google calendar create-calendar "Team Calendar" --timezone Europe/London
aim-google calendar acl <calendarId>         # List access control rules
aim-google calendar colors                   # List available event/calendar colors
aim-google calendar time --timezone America/New_York
aim-google calendar users                    # List workspace users (use email as calendar ID)

# Events (with timezone-aware time flags)
aim-google calendar events <calendarId> --today                    # Today's events
aim-google calendar events <calendarId> --tomorrow                 # Tomorrow's events
aim-google calendar events <calendarId> --week                     # This week (Mon-Sun by default; use --week-start)
aim-google calendar events <calendarId> --days 3                   # Next 3 days
aim-google calendar events <calendarId> --from today --to friday   # Relative dates
aim-google calendar events <calendarId> --from today --to friday --weekday   # Include weekday columns
aim-google calendar events <calendarId> --from 2025-01-01T00:00:00Z --to 2025-01-08T00:00:00Z
aim-google calendar events --all             # Fetch events from all calendars
aim-google calendar events --calendars 1,3   # Fetch events from calendar indices (see aim-google calendar calendars)
aim-google calendar events --cal Work --cal Personal  # Fetch events from calendars by name/ID
aim-google calendar event <calendarId> <eventId>
aim-google calendar get <calendarId> <eventId>                     # Alias for event
aim-google calendar search "meeting" --today
aim-google calendar search "meeting" --tomorrow
aim-google calendar search "meeting" --days 365
aim-google calendar search "meeting" --from 2025-01-01T00:00:00Z --to 2025-01-31T00:00:00Z --max 50

# Search defaults to 30 days ago through 90 days ahead unless you set --from/--to/--today/--week/--days.
# Tip: set AIM_GOOGLE_CALENDAR_WEEKDAY=1 to default --weekday for calendar events output.

# JSON event output includes timezone and localized times (useful for agents).
aim-google calendar get <calendarId> <eventId> --json
# {
#   "event": {
#     "id": "...",
#     "summary": "...",
#     "startDayOfWeek": "Friday",
#     "endDayOfWeek": "Friday",
#     "timezone": "America/Los_Angeles",
#     "eventTimezone": "America/New_York",
#     "startLocal": "2026-01-23T20:45:00-08:00",
#     "endLocal": "2026-01-23T22:45:00-08:00",
#     "start": { "dateTime": "2026-01-23T23:45:00-05:00" },
#     "end": { "dateTime": "2026-01-24T01:45:00-05:00" }
#   }
# }

# Team calendars (requires Cloud Identity API for Google Workspace)
aim-google calendar team <group-email> --today           # Show team's events for today
aim-google calendar team <group-email> --week            # Show team's events for the week (use --week-start)
aim-google calendar team <group-email> --freebusy        # Show only busy/free blocks (faster)
aim-google calendar team <group-email> --query "standup" # Filter by event title

# Create and update
aim-google calendar create <calendarId> \
  --summary "Meeting" \
  --from 2025-01-15T10:00:00Z \
  --to 2025-01-15T11:00:00Z

aim-google calendar create <calendarId> \
  --summary "Team Sync" \
  --from 2025-01-15T14:00:00Z \
  --to 2025-01-15T15:00:00Z \
  --attendees "alice@example.com,bob@example.com" \
  --location "Zoom"

aim-google calendar update <calendarId> <eventId> \
  --summary "Updated Meeting" \
  --from 2025-01-15T11:00:00Z \
  --to 2025-01-15T12:00:00Z

# Send notifications when creating/updating
aim-google calendar create <calendarId> \
  --summary "Team Sync" \
  --from 2025-01-15T14:00:00Z \
  --to 2025-01-15T15:00:00Z \
  --send-updates all

aim-google calendar update <calendarId> <eventId> \
  --send-updates externalOnly

# Default: no attendee notifications unless you pass --send-updates.
aim-google calendar delete <calendarId> <eventId> \
  --send-updates all --force

# Recurrence + reminders
aim-google calendar create <calendarId> \
  --summary "Payment" \
  --from 2025-02-11T09:00:00-03:00 \
  --to 2025-02-11T09:15:00-03:00 \
  --rrule "RRULE:FREQ=MONTHLY;BYMONTHDAY=11" \
  --reminder "email:3d" \
  --reminder "popup:30m"

# Special event types via --event-type (focus-time/out-of-office/working-location)
aim-google calendar create primary \
  --event-type focus-time \
  --from 2025-01-15T13:00:00Z \
  --to 2025-01-15T14:00:00Z

aim-google calendar create primary \
  --event-type out-of-office \
  --from 2025-01-20 \
  --to 2025-01-21 \
  --all-day

aim-google calendar create primary \
  --event-type working-location \
  --working-location-type office \
  --working-office-label "HQ" \
  --from 2025-01-22 \
  --to 2025-01-23

# Dedicated shortcuts (same event types, more opinionated defaults)
aim-google calendar focus-time --from 2025-01-15T13:00:00Z --to 2025-01-15T14:00:00Z
aim-google calendar out-of-office --from 2025-01-20 --to 2025-01-21 --all-day
aim-google calendar working-location --type office --office-label "HQ" --from 2025-01-22 --to 2025-01-23
# Add attendees without replacing existing attendees/RSVP state
aim-google calendar update <calendarId> <eventId> \
  --add-attendee "alice@example.com,bob@example.com"

aim-google calendar delete <calendarId> <eventId>

# Invitations
aim-google calendar respond <calendarId> <eventId> --status accepted
aim-google calendar respond <calendarId> <eventId> --status declined
aim-google calendar respond <calendarId> <eventId> --status tentative
aim-google calendar respond <calendarId> <eventId> --status declined --send-updates externalOnly

# Propose a new time (browser-only flow; API limitation)
aim-google calendar propose-time <calendarId> <eventId>
aim-google calendar propose-time <calendarId> <eventId> --open
aim-google calendar propose-time <calendarId> <eventId> --decline --comment "Can we do 5pm?"

# Availability
aim-google calendar freebusy --calendars "primary,work@example.com" \
  --from 2025-01-15T00:00:00Z \
  --to 2025-01-16T00:00:00Z
aim-google calendar freebusy --cal Work --from 2025-01-15T00:00:00Z --to 2025-01-16T00:00:00Z

aim-google calendar conflicts --calendars "primary,work@example.com" \
  --today                             # Today's conflicts
aim-google calendar conflicts --all --today # Check conflicts across all calendars
```

### Time

```bash
aim-google time now
aim-google time now --timezone UTC
```

### Drive

When you turn a Markdown file into a Google Doc, use **`--convert`** (extension-based) or **`--convert-to doc`**. Leading YAML frontmatter between **`---`** lines is **removed before upload** unless you pass **`--keep-frontmatter`**. That step only looks for opening and closing delimiter lines—it is **not** a full YAML parse, so odd edge cases may need **`--keep-frontmatter`** or editing the file first.

```bash
# List and search
aim-google drive ls --max 20
aim-google drive ls --parent <folderId> --max 20
aim-google drive ls --all --max 20               # List across all accessible files (cannot combine with --parent)
aim-google drive ls --no-all-drives            # Only list from "My Drive"
aim-google drive search "invoice" --max 20
aim-google drive search "invoice" --no-all-drives
aim-google drive search "mimeType = 'application/pdf'" --raw-query
aim-google drive get <fileId>                # Get file metadata
aim-google drive url <fileId>                # Print Drive web URL
aim-google drive copy <fileId> "Copy Name"

# Upload and download
aim-google drive upload ./path/to/file --parent <folderId>
aim-google drive upload ./path/to/file --replace <fileId>  # Replace file content in-place (preserves shared link)
aim-google drive upload ./report.docx --convert
aim-google drive upload ./chart.png --convert-to sheet
aim-google drive upload ./report.docx --convert --name report.docx
aim-google drive upload ./notes.md --convert                              # Markdown → Google Doc (or use --convert-to doc)
aim-google drive download <fileId> --out ./downloaded.bin
aim-google drive download <fileId> --format pdf --out ./exported.pdf     # Google Workspace files only
aim-google drive download <fileId> --format docx --out ./doc.docx
aim-google drive download <fileId> --format md --out ./note.md            # Google Doc → Markdown
aim-google drive download <fileId> --format pptx --out ./slides.pptx

# Organize
aim-google drive mkdir "New Folder"
aim-google drive mkdir "New Folder" --parent <parentFolderId>
aim-google drive rename <fileId> "New Name"
aim-google drive move <fileId> --parent <destinationFolderId>
aim-google drive delete <fileId>             # Move to trash
aim-google drive delete <fileId> --permanent # Permanently delete

# Permissions
aim-google drive permissions <fileId>
aim-google drive share <fileId> --to user --email user@example.com --role reader
aim-google drive share <fileId> --to user --email user@example.com --role writer
aim-google drive share <fileId> --to user --email reviewer@example.com --role commenter
aim-google drive share <fileId> --to domain --domain example.com --role reader
aim-google drive unshare <fileId> --permission-id <permissionId>

# Shared drives (Team Drives)
aim-google drive drives --max 100
```

### Docs / Slides / Sheets

```bash
# Docs
aim-google docs info <docId>
aim-google docs cat <docId> --max-bytes 10000
aim-google docs create "My Doc"
aim-google docs create "My Doc" --file ./doc.md            # Import markdown
aim-google docs create "My Doc" --pageless
aim-google docs copy <docId> "My Doc Copy"
aim-google docs export <docId> --format pdf --out ./doc.pdf
aim-google docs list-tabs <docId>
aim-google docs cat <docId> --tab "Notes"
aim-google docs cat <docId> --all-tabs
aim-google docs update <docId> --text "Append this later"
aim-google docs update <docId> --text "Only in this tab" --tab-id t.notes
aim-google docs update <docId> --file ./insert.txt --index 25 --pageless
aim-google docs write <docId> --text "Fresh content"
aim-google docs write <docId> --text "Rewrite one tab" --tab-id t.notes
aim-google docs write <docId> --file ./body.txt --append --pageless
aim-google docs write <docId> --file ./body.md --replace --markdown
aim-google docs find-replace <docId> "old" "new"
aim-google docs find-replace <docId> "old" "new" --tab-id t.notes

# Slides
aim-google slides info <presentationId>
aim-google slides create "My Deck"
aim-google slides create-from-markdown "My Deck" --content-file ./slides.md
aim-google slides create-from-template <templateId> "My Deck" --replace "name=John" --replace "date=2026-02-15"
aim-google slides copy <presentationId> "My Deck Copy"
aim-google slides export <presentationId> --format pdf --out ./deck.pdf
aim-google slides list-slides <presentationId>
aim-google slides add-slide <presentationId> ./slide.png --notes "Speaker notes"
aim-google slides update-notes <presentationId> <slideId> --notes "Updated notes"
aim-google slides replace-slide <presentationId> <slideId> ./new-slide.png --notes "New notes"

# Sheets
aim-google sheets copy <spreadsheetId> "My Sheet Copy"
aim-google sheets export <spreadsheetId> --format pdf --out ./sheet.pdf
aim-google sheets format <spreadsheetId> 'Sheet1!A1:B2' --format-json '{"textFormat":{"bold":true}}' --format-fields 'userEnteredFormat.textFormat.bold'
aim-google sheets format <spreadsheetId> 'Sheet1!A1:B2' --format-json '{"borders":{"top":{"style":"SOLID"}}}' --format-fields 'userEnteredFormat.borders.top.style'
aim-google sheets merge <spreadsheetId> 'Sheet1!A1:B2'
aim-google sheets number-format <spreadsheetId> 'Sheet1!C:C' --type CURRENCY --pattern '$#,##0.00'
aim-google sheets freeze <spreadsheetId> --rows 1 --cols 1
aim-google sheets resize-columns <spreadsheetId> 'Sheet1!A:C' --auto
aim-google sheets read-format <spreadsheetId> 'Sheet1!A1:B2'
aim-google sheets insert <spreadsheetId> "Sheet1" rows 2 --count 3
aim-google sheets notes <spreadsheetId> 'Sheet1!A1:B10'
aim-google sheets find-replace <spreadsheetId> "old" "new"
aim-google sheets find-replace <spreadsheetId> "old" "new" --sheet Sheet1 --match-entire
aim-google sheets links <spreadsheetId> 'Sheet1!A1:B10'
aim-google sheets add-tab <spreadsheetId> <tabName> --index 0
aim-google sheets rename-tab <spreadsheetId> <oldName> <newName>
aim-google sheets delete-tab <spreadsheetId> <tabName> --force
```

### Contacts

```bash
# Personal contacts
aim-google contacts list --max 50
aim-google contacts search "Ada" --max 50
aim-google contacts get people/<resourceName>
aim-google contacts get user@example.com     # Get by email

# Other contacts (people you've interacted with)
aim-google contacts other list --max 50
aim-google contacts other search "John" --max 50

# Create and update
aim-google contacts create \
  --given "John" \
  --family "Doe" \
  --email "john@example.com" \
  --phone "+1234567890" \
  --address "12 St James's Square, London" \
  --gender "male" \
  --relation "spouse=Jane Doe"

aim-google contacts update people/<resourceName> \
  --given "Jane" \
  --email "jane@example.com" \
  --address "1 Infinite Loop, Cupertino" \
  --birthday "1990-05-12" \
  --gender "female" \
  --notes "Met at WWDC" \
  --relation "friend=Bob"

# Update via JSON (see docs/contacts-json-update.md)
aim-google contacts get people/<resourceName> --json | \
  jq '(.contact.urls //= []) | (.contact.urls += [{"value":"obsidian://open?vault=notes&file=People/John%20Doe","type":"profile"}])' | \
  aim-google contacts update people/<resourceName> --from-file -

aim-google contacts delete people/<resourceName>

# Workspace directory (requires Google Workspace)
aim-google contacts directory list --max 50
aim-google contacts directory search "Jane" --max 50
```

### Tasks

```bash
# Task lists
aim-google tasks lists --max 50
aim-google tasks lists create <title>

# Tasks in a list
aim-google tasks list <tasklistId> --max 50
aim-google tasks get <tasklistId> <taskId>
aim-google tasks add <tasklistId> --title "Task title"
aim-google tasks add <tasklistId> --title "Weekly sync" --due 2025-02-01 --repeat weekly --repeat-count 4
aim-google tasks add <tasklistId> --title "Daily standup" --due 2025-02-01 --repeat daily --repeat-until 2025-02-05
aim-google tasks add <tasklistId> --title "Bi-weekly review" --due 2025-02-01 --recur-rrule "FREQ=WEEKLY;INTERVAL=2" --repeat-count 3
aim-google tasks update <tasklistId> <taskId> --title "New title"
aim-google tasks done <tasklistId> <taskId>
aim-google tasks undo <tasklistId> <taskId>
aim-google tasks delete <tasklistId> <taskId>
aim-google tasks clear <tasklistId>

# Note: Google Tasks treats due dates as date-only; time components may be ignored.
# Note: Public Google Tasks API does not expose true recurring-task metadata; `--repeat*`/`--recur*` materialize concrete tasks.
# See docs/dates.md for all supported date/time input formats across commands.
```

### Sheets

```bash
# Read
aim-google sheets metadata <spreadsheetId>
aim-google sheets get <spreadsheetId> 'Sheet1!A1:B10'
aim-google sheets get <spreadsheetId> MyNamedRange

# Export (via Drive)
aim-google sheets export <spreadsheetId> --format pdf --out ./sheet.pdf
aim-google sheets export <spreadsheetId> --format xlsx --out ./sheet.xlsx

# Write
aim-google sheets update <spreadsheetId> 'A1' 'val1|val2,val3|val4'
aim-google sheets update <spreadsheetId> 'A1' --values-json '[["a","b"],["c","d"]]'
aim-google sheets update <spreadsheetId> 'Sheet1!A1:C1' 'new|row|data' --copy-validation-from 'Sheet1!A2:C2'
aim-google sheets update <spreadsheetId> MyNamedRange 'new|row|data'
aim-google sheets update <spreadsheetId> 'Sheet1!A1:C1' 'new|row|data' --copy-validation-from MyValidationNamedRange
aim-google sheets append <spreadsheetId> 'Sheet1!A:C' 'new|row|data'
aim-google sheets append <spreadsheetId> 'Sheet1!A:C' 'new|row|data' --copy-validation-from 'Sheet1!A2:C2'
aim-google sheets find-replace <spreadsheetId> "old" "new"
aim-google sheets find-replace <spreadsheetId> "old" "new" --sheet Sheet1 --regex
aim-google sheets update-note <spreadsheetId> 'Sheet1!A1' --note ''
aim-google sheets append <spreadsheetId> MyNamedRange 'new|row|data'
aim-google sheets clear <spreadsheetId> 'Sheet1!A1:B10'
aim-google sheets clear <spreadsheetId> MyNamedRange

# Format
aim-google sheets format <spreadsheetId> 'Sheet1!A1:B2' --format-json '{"textFormat":{"bold":true}}' --format-fields 'userEnteredFormat.textFormat.bold'
aim-google sheets format <spreadsheetId> MyNamedRange --format-json '{"textFormat":{"bold":true}}' --format-fields 'userEnteredFormat.textFormat.bold'
aim-google sheets format <spreadsheetId> 'Sheet1!A1:B2' --format-json '{"borders":{"top":{"style":"SOLID"}}}' --format-fields 'userEnteredFormat.borders.top.style'
aim-google sheets merge <spreadsheetId> 'Sheet1!A1:B2'
aim-google sheets unmerge <spreadsheetId> 'Sheet1!A1:B2'
aim-google sheets number-format <spreadsheetId> 'Sheet1!C:C' --type CURRENCY --pattern '$#,##0.00'
aim-google sheets freeze <spreadsheetId> --rows 1 --cols 1
aim-google sheets resize-columns <spreadsheetId> 'Sheet1!A:C' --auto
aim-google sheets resize-rows <spreadsheetId> 'Sheet1!1:10' --height 36
aim-google sheets read-format <spreadsheetId> 'Sheet1!A1:B2'
aim-google sheets read-format <spreadsheetId> 'Sheet1!A1:B2' --effective

# Named ranges
aim-google sheets named-ranges <spreadsheetId>
aim-google sheets named-ranges get <spreadsheetId> MyNamedRange
aim-google sheets named-ranges add <spreadsheetId> MyNamedRange 'Sheet1!A1:B2'
aim-google sheets named-ranges add <spreadsheetId> MyCols 'Sheet1!A:C'
aim-google sheets named-ranges update <spreadsheetId> MyNamedRange --name MyNamedRange2
aim-google sheets named-ranges delete <spreadsheetId> MyNamedRange2

# Charts
aim-google sheets chart list <spreadsheetId>
aim-google sheets chart get <spreadsheetId> <chartId> --json > chart.json
aim-google sheets chart create <spreadsheetId> --spec-json @chart.json
aim-google sheets chart create <spreadsheetId> --spec-json '{"title":"Revenue","basicChart":{"chartType":"COLUMN"}}' --sheet Sheet1 --anchor E10
aim-google sheets chart update <spreadsheetId> <chartId> --spec-json '{"title":"New Title","basicChart":{"chartType":"PIE"}}'
aim-google sheets chart delete <spreadsheetId> <chartId>

# Insert rows/cols
aim-google sheets insert <spreadsheetId> "Sheet1" rows 2 --count 3
aim-google sheets insert <spreadsheetId> "Sheet1" cols 3 --after

# Notes
aim-google sheets notes <spreadsheetId> 'Sheet1!A1:B10'
aim-google sheets links <spreadsheetId> 'Sheet1!A1:B10'   # Includes rich-text links

# Create
aim-google sheets create "My New Spreadsheet" --sheets "Sheet1,Sheet2"

# Tab management
aim-google sheets add-tab <spreadsheetId> <tabName> --index 0
aim-google sheets rename-tab <spreadsheetId> <oldName> <newName>
aim-google sheets delete-tab <spreadsheetId> <tabName>          # use --force to skip confirmation
```

### Forms

```bash
# Forms
aim-google forms get <formId>
aim-google forms create --title "Weekly Check-in" --description "Friday async update"
aim-google forms update <formId> --title "Weekly Sync" --quiz true
aim-google forms add-question <formId> --title "What shipped?" --type paragraph --required
aim-google forms move-question <formId> 3 1
aim-google forms delete-question <formId> 2 --force

# Responses
aim-google forms responses list <formId> --max 20
aim-google forms responses get <formId> <responseId>

# Watches
aim-google forms watch create <formId> --topic projects/<project>/topics/<topic>
aim-google forms watch list <formId>
aim-google forms watch renew <formId> <watchId>
aim-google forms watch delete <formId> <watchId>
```

### Apps Script

```bash
# Projects
aim-google appscript get <scriptId>
aim-google appscript content <scriptId>
aim-google appscript create --title "Automation Helpers"
aim-google appscript create --title "Bound Script" --parent-id <driveFileId>

# Execute functions
aim-google appscript run <scriptId> myFunction --params '["arg1", 123, true]'
aim-google appscript run <scriptId> myFunction --dev-mode
```

### People

```bash
# Profile
aim-google people me
aim-google people get people/<userId>

# Search the Workspace directory
aim-google people search "Ada Lovelace" --max 5

# Relations (defaults to people/me)
aim-google people relations
aim-google people relations people/<userId> --type manager
```

### Chat

```bash
# Spaces
aim-google chat spaces list
aim-google chat spaces find "Engineering"
aim-google chat spaces find "Engineering" --exact
aim-google chat spaces create "Engineering" --member alice@company.com --member bob@company.com

# Messages
aim-google chat messages list spaces/<spaceId> --max 5
aim-google chat messages list spaces/<spaceId> --thread <threadId>
aim-google chat messages list spaces/<spaceId> --unread
aim-google chat messages send spaces/<spaceId> --text "Build complete!" --thread spaces/<spaceId>/threads/<threadId>
aim-google chat messages reactions list spaces/<spaceId>/messages/<messageId>
aim-google chat messages react spaces/<spaceId>/messages/<messageId> "👍"  # shorthand for reactions create
aim-google chat messages reactions delete spaces/<spaceId>/messages/<messageId>/reactions/<reactionId>

# Threads
aim-google chat threads list spaces/<spaceId>

# Direct messages
aim-google chat dm space user@company.com
aim-google chat dm send user@company.com --text "ping"
```

Note: Chat commands require a Google Workspace account (consumer @gmail.com accounts are not supported).

### Admin

```bash
# Requires a Workspace service account with domain-wide delegation.
aim-google admin users list --domain example.com
aim-google admin users get user@example.com
aim-google admin users create user@example.com --given Ada --family Lovelace --password 'TempPass123!'
aim-google admin users suspend user@example.com --force

aim-google admin groups list --domain example.com
aim-google admin groups members list engineering@example.com
aim-google admin groups members add engineering@example.com user@example.com --role MEMBER
aim-google admin groups members remove engineering@example.com user@example.com --force
```

### Groups (Google Workspace)

```bash
# List groups you belong to
aim-google groups list

# List members of a group
aim-google groups members engineering@company.com
```

Note: Groups commands require the Cloud Identity API and the `cloud-identity.groups.readonly` scope. If you get a permissions error, re-authenticate:

```bash
aim-google auth add your@email.com --services groups --force-consent
```

### Classroom (Google Workspace for Education)

```bash
# Courses
aim-google classroom courses list
aim-google classroom courses list --role teacher
aim-google classroom courses get <courseId>
aim-google classroom courses create --name "Math 101"
aim-google classroom courses update <courseId> --name "Math 102"
aim-google classroom courses archive <courseId>
aim-google classroom courses unarchive <courseId>
aim-google classroom courses url <courseId>

# Roster
aim-google classroom roster <courseId>
aim-google classroom roster <courseId> --students
aim-google classroom students add <courseId> <userId>
aim-google classroom teachers add <courseId> <userId>

# Coursework
aim-google classroom coursework list <courseId>
aim-google classroom coursework get <courseId> <courseworkId>
aim-google classroom coursework create <courseId> --title "Homework 1" --type ASSIGNMENT --state PUBLISHED
aim-google classroom coursework update <courseId> <courseworkId> --title "Updated"
aim-google classroom coursework assignees <courseId> <courseworkId> --mode INDIVIDUAL_STUDENTS --add-student <studentId>

# Materials
aim-google classroom materials list <courseId>
aim-google classroom materials create <courseId> --title "Syllabus" --state PUBLISHED

# Submissions
aim-google classroom submissions list <courseId> <courseworkId>
aim-google classroom submissions get <courseId> <courseworkId> <submissionId>
aim-google classroom submissions grade <courseId> <courseworkId> <submissionId> --grade 85
aim-google classroom submissions return <courseId> <courseworkId> <submissionId>
aim-google classroom submissions turn-in <courseId> <courseworkId> <submissionId>
aim-google classroom submissions reclaim <courseId> <courseworkId> <submissionId>

# Announcements
aim-google classroom announcements list <courseId>
aim-google classroom announcements create <courseId> --text "Welcome!"
aim-google classroom announcements update <courseId> <announcementId> --text "Updated"
aim-google classroom announcements assignees <courseId> <announcementId> --mode INDIVIDUAL_STUDENTS --add-student <studentId>

# Topics
aim-google classroom topics list <courseId>
aim-google classroom topics create <courseId> --name "Unit 1"
aim-google classroom topics update <courseId> <topicId> --name "Unit 2"

# Invitations
aim-google classroom invitations list
aim-google classroom invitations create <courseId> <userId> --role student
aim-google classroom invitations accept <invitationId>

# Guardians
aim-google classroom guardians list <studentId>
aim-google classroom guardians get <studentId> <guardianId>
aim-google classroom guardians delete <studentId> <guardianId>

# Guardian invitations
aim-google classroom guardian-invitations list <studentId>
aim-google classroom guardian-invitations create <studentId> --email parent@example.com

# Profiles
aim-google classroom profile get
aim-google classroom profile get <userId>
```

Note: Classroom commands require a Google Workspace for Education account. Personal Google accounts have limited Classroom functionality.

### Docs

```bash
# Export (via Drive)
aim-google docs export <docId> --format pdf --out ./doc.pdf
aim-google docs export <docId> --format docx --out ./doc.docx
aim-google docs export <docId> --format txt --out ./doc.txt
aim-google docs export <docId> --format md --out ./doc.md
aim-google docs export <docId> --format html --out ./doc.html

# Sed-style regex editing with Markdown formatting (sedmat)
aim-google docs sed <docId> 's/pattern/replacement/g'

# Formatting in replacements
aim-google docs sed <docId> 's/hello/**hello**/'          # bold
aim-google docs sed <docId> 's/hello/*hello*/'             # italic
aim-google docs sed <docId> 's/hello/~~hello~~/'           # strikethrough
aim-google docs sed <docId> 's/hello/`hello`/'             # monospace
aim-google docs sed <docId> 's/hello/__hello__/'           # underline
aim-google docs sed <docId> 's/Google/[Google](https://google.com)/'  # link

# Images
aim-google docs sed <docId> 's/{{LOGO}}/![](https://example.com/logo.png)/'
aim-google docs sed <docId> 's/{{HERO}}/![](https://example.com/hero.jpg){width=600}/'

# Tables — create, populate, modify
aim-google docs sed <docId> 's/{{TABLE}}/|3x4|/'            # create 3-row, 4-col table
aim-google docs sed <docId> 's/|1|[A1]/**Name**/'           # set cell A1 (bold)
aim-google docs sed <docId> 's/|1|[1,*]/**&**/'             # bold entire row 1
aim-google docs sed <docId> 's/|1|[row:+2]//'               # insert row before row 2
aim-google docs sed <docId> 's/|1|[col:$+]//'               # append column at end
```

> See [docs/sedmat.md](docs/sedmat.md) for the full sedmat syntax reference.

### Slides

```bash
# Export (via Drive)
aim-google slides export <presentationId> --format pptx --out ./deck.pptx
aim-google slides export <presentationId> --format pdf --out ./deck.pdf

# Create from template with text replacements
aim-google slides create-from-template <templateId> "Q1 Report" \
  --replace "quarter=Q1 2026" \
  --replace "revenue=$1.2M" \
  --replace "growth=15%"

# Use JSON file for many replacements
cat > replacements.json <<EOF
{
  "name": "John Doe",
  "title": "Sales Manager",
  "date": "2026-02-15",
  "sales": "125",
  "target": "100"
}
EOF

aim-google slides create-from-template <templateId> "Monthly Report" \
  --replacements replacements.json

# Read slide content (text, notes, images)
aim-google slides read-slide <presentationId> <slideId>

# Include grouped elements, word art, and tables
aim-google slides read-slide <presentationId> <slideId> --recursive --json

# Get a rendered slide thumbnail URL
aim-google slides thumbnail <presentationId> <slideId>

# Download a rendered slide thumbnail
aim-google slides thumbnail <presentationId> <slideId> --output ./slide.png

# Control thumbnail size and format
aim-google slides thumbnail <presentationId> <slideId> --size medium --format jpeg --output ./slide.jpg
```

## Output Formats

### Text

Human-readable output with colors (default):

```bash
$ aim-google gmail search 'newer_than:7d' --max 3
THREAD_ID           SUBJECT                           FROM                  DATE
18f1a2b3c4d5e6f7    Meeting notes                     alice@example.com     2025-01-10
17e1d2c3b4a5f6e7    Invoice #12345                    billing@vendor.com    2025-01-09
16d1c2b3a4e5f6d7    Project update                    bob@example.com       2025-01-08
```

Message-level search (one row per email; add `--include-body` to fetch/decode bodies, or `--full` for untruncated text bodies):

```bash
$ aim-google gmail messages search 'newer_than:7d' --max 3
ID                  THREAD             SUBJECT                           FROM                  DATE
18f1a2b3c4d5e6f7    9e8d7c6b5a4f3e2d    Meeting notes                     alice@example.com     2025-01-10
17e1d2c3b4a5f6e7    9e8d7c6b5a4f3e2d    Invoice #12345                    billing@vendor.com    2025-01-09
16d1c2b3a4e5f6d7    7f6e5d4c3b2a1908    Project update                    bob@example.com       2025-01-08
```

### JSON

Machine-readable output for scripting and automation:

```bash
$ aim-google gmail search 'newer_than:7d' --max 3 --json
{
  "threads": [
    {
      "id": "18f1a2b3c4d5e6f7",
      "snippet": "Meeting notes from today...",
      "messages": [...]
    },
    ...
  ]
}
```

```bash
$ aim-google gmail messages search 'newer_than:7d' --max 3 --json
{
  "messages": [
    {
      "id": "18f1a2b3c4d5e6f7",
      "threadId": "9e8d7c6b5a4f3e2d",
      "subject": "Meeting notes",
      "from": "alice@example.com",
      "date": "2025-01-10"
    },
    ...
  ]
}
```

```bash
$ aim-google gmail messages search 'newer_than:7d' --max 1 --full --json
{
  "messages": [
    {
      "id": "18f1a2b3c4d5e6f7",
      "threadId": "9e8d7c6b5a4f3e2d",
      "subject": "Meeting notes",
      "from": "alice@example.com",
      "date": "2025-01-10",
      "body": "Hi team — meeting notes..."
    }
  ]
}
```

Data goes to stdout, errors and progress to stderr for clean piping:

```bash
aim-google --json drive ls --max 5 | jq '.files[] | select(.mimeType=="application/pdf")'
```

Useful pattern:

- `aim-google --json ... | jq .`

Calendar JSON convenience fields:

- `startDayOfWeek` / `endDayOfWeek` on event payloads (derived from start/end).

## Examples

### Search recent emails and download attachments

```bash
# Search for emails from the last week
aim-google gmail search 'newer_than:7d has:attachment' --max 10

# Get thread details and download attachments
aim-google gmail thread get <threadId> --download
```

### Modify labels on a thread

```bash
# Archive and star a thread
aim-google gmail thread modify <threadId> --remove INBOX --add STARRED
```

### Create a calendar event with attendees

```bash
# Find a free time slot
aim-google calendar freebusy --calendars "primary" \
  --from 2025-01-15T00:00:00Z \
  --to 2025-01-16T00:00:00Z

# Create the meeting
aim-google calendar create primary \
  --summary "Team Standup" \
  --from 2025-01-15T10:00:00Z \
  --to 2025-01-15T10:30:00Z \
  --attendees "alice@example.com,bob@example.com"
```

### Find and download files from Drive

```bash
# Search for PDFs
aim-google drive search "invoice filetype:pdf" --max 20 --json | \
  jq -r '.files[] | .id' | \
  while read fileId; do
    aim-google drive download "$fileId"
  done
```

### Manage multiple accounts

```bash
# Check personal Gmail
aim-google gmail search 'is:unread' --account personal@gmail.com

# Check work Gmail
aim-google gmail search 'is:unread' --account work@company.com

# Or set default
export AIM_GOOGLE_ACCOUNT=work@company.com
aim-google gmail search 'is:unread'
```

### Update a Google Sheet from a CSV

```bash
# Convert CSV to pipe-delimited format and update sheet
cat data.csv | tr ',' '|' | \
  aim-google sheets update <spreadsheetId> 'Sheet1!A1'
```

### Export Sheets / Docs / Slides

```bash
# Sheets
aim-google sheets export <spreadsheetId> --format pdf

# Docs
aim-google docs export <docId> --format docx

# Slides
aim-google slides export <presentationId> --format pptx
```

### Batch process Gmail threads

```bash
# Mark all emails from a sender as read
aim-google --json gmail search 'from:noreply@example.com' --max 200 | \
  jq -r '.threads[].id' | \
  xargs -n 50 aim-google gmail labels modify --remove UNREAD

# Archive old emails
aim-google --json gmail search 'older_than:1y' --max 200 | \
  jq -r '.threads[].id' | \
  xargs -n 50 aim-google gmail labels modify --remove INBOX

# Label important emails
aim-google --json gmail search 'from:boss@example.com' --max 200 | \
  jq -r '.threads[].id' | \
  xargs -n 50 aim-google gmail labels modify --add IMPORTANT
```

## Advanced Features

### Verbose Mode

Enable verbose logging for troubleshooting:

```bash
aim-google --verbose gmail search 'newer_than:7d'
# Shows API requests and responses
```

## Global Flags

All commands support these flags:

- `--account <email|alias|auto>` - Account to use (overrides AIM_GOOGLE_ACCOUNT)
- `--enable-commands <csv>` - Allowlist commands; dot paths allowed (e.g., `calendar,tasks,gmail.search`)
- `--disable-commands <csv>` - Denylist commands; dot paths allowed (e.g., `gmail.send,gmail.drafts.send`)
- `--gmail-no-send` - Block Gmail send operations
- `--json` - Output JSON to stdout (best for scripting)
- `--plain` - Output stable, parseable text to stdout (TSV; no colors)
- `--color <mode>` - Color mode: `auto`, `always`, or `never` (default: auto)
- `--force` - Skip confirmations for destructive commands
- `--no-input` - Never prompt; fail instead (useful for CI)
- `--verbose` - Enable verbose logging
- `--help` - Show help for any command

## Shell Completions

Generate shell completions for your preferred shell:

### Bash

```bash
# macOS (with Homebrew)
aim-google completion bash > $(brew --prefix)/etc/bash_completion.d/aim-google

# Linux
aim-google completion bash > /etc/bash_completion.d/aim-google

# Or load directly in your current session
source <(aim-google completion bash)
```

### Zsh

```zsh
# Generate completion file
aim-google completion zsh > "${fpath[1]}/_gog"

# Or add to .zshrc for automatic loading
echo 'eval "$(aim-google completion zsh)"' >> ~/.zshrc

# Enable completions if not already enabled
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

### Fish

```fish
aim-google completion fish > ~/.config/fish/completions/aim-google.fish
```

### PowerShell

```powershell
# Load for current session
aim-google completion powershell | Out-String | Invoke-Expression

# Or add to profile for all sessions
aim-google completion powershell >> $PROFILE
```

After installing completions, start a new shell session for changes to take effect.

## Development

After cloning, install tools:

```bash
make tools
```

Pinned tools (installed into `.tools/`):

- Format: `make fmt` (goimports + gofumpt)
- Lint: `make lint` (golangci-lint)
- Test: `make test`

CI runs format checks, tests, and lint on push/PR.

Regenerate the expanded command reference from the live schema when CLI surface changes:

```bash
make docs-commands
```

### Integration Tests (Live Google APIs)

Opt-in tests that hit real Google APIs using your stored `aim-google` credentials/tokens.

```bash
# Optional: override which account to use
export AIM_GOOGLE_IT_ACCOUNT=you@gmail.com
export AIM_GOOGLE_CLIENT=work
go test -tags=integration ./...
```

Tip: if you want to avoid macOS Keychain prompts during these runs, set `AIM_GOOGLE_KEYRING_BACKEND=file` and `AIM_GOOGLE_KEYRING_PASSWORD=...` (uses encrypted on-disk keyring).

### Live Test Script (CLI)

Fast end-to-end smoke checks against live APIs:

```bash
scripts/live-test.sh --fast
scripts/live-test.sh --account you@gmail.com --skip groups,keep,calendar-enterprise
scripts/live-test.sh --client work --account you@company.com
```

Script toggles:

- `--auth all,groups` to re-auth before running
- `--client <name>` to select OAuth client credentials
- `--strict` to fail on optional features (groups/keep/enterprise)
- `--allow-nontest` to override the test-account guardrail

Go test wrapper (opt-in):

```bash
AIM_GOOGLE_LIVE=1 go test -tags=integration ./internal/integration -run Live
```

Optional env:
- `AIM_GOOGLE_LIVE_FAST=1`
- `AIM_GOOGLE_LIVE_SKIP=groups,keep`
- `AIM_GOOGLE_LIVE_AUTH=all,groups`
- `AIM_GOOGLE_LIVE_ALLOW_NONTEST=1`
- `AIM_GOOGLE_LIVE_EMAIL_TEST=BrianV1981+aimgoogletest@gmail.com`
- `AIM_GOOGLE_LIVE_GROUP_EMAIL=group@domain`
- `AIM_GOOGLE_LIVE_CLASSROOM_COURSE=<courseId>`
- `AIM_GOOGLE_LIVE_CLASSROOM_CREATE=1`
- `AIM_GOOGLE_LIVE_CLASSROOM_ALLOW_STATE=1`
- `AIM_GOOGLE_LIVE_TRACK=1`
- `AIM_GOOGLE_LIVE_GMAIL_BATCH_DELETE=1`
- `AIM_GOOGLE_LIVE_GMAIL_FILTERS=1`
- `AIM_GOOGLE_LIVE_GMAIL_WATCH_TOPIC=projects/.../topics/...`
- `AIM_GOOGLE_LIVE_CALENDAR_RESPOND=1`
- `AIM_GOOGLE_LIVE_CALENDAR_RECURRENCE=1`
- `AIM_GOOGLE_KEEP_SERVICE_ACCOUNT=/path/to/service-account.json`
- `AIM_GOOGLE_KEEP_IMPERSONATE=user@workspace-domain`

### Make Shortcut

Build and run:

```bash
make aim-google auth add you@gmail.com
```

For clean stdout when scripting:

- Use `--` when the first arg is a flag: `make aim-google -- --json gmail search "from:me" | jq .`

## License

MIT

## Links

- [GitHub Repository](https://github.com/BrianV1981/aim-google)
- [Gmail API Documentation](https://developers.google.com/gmail/api)
- [Google Calendar API Documentation](https://developers.google.com/calendar)
- [Google Drive API Documentation](https://developers.google.com/drive)
- [Google People API Documentation](https://developers.google.com/people)
- [Google Tasks API Documentation](https://developers.google.com/tasks)
- [Google Sheets API Documentation](https://developers.google.com/sheets)
- [Cloud Identity API Documentation](https://cloud.google.com/identity/docs/reference/rest)

## Credits

This project is inspired by Mario Zechner's original CLIs:

- [gmcli](https://github.com/badlogic/gmcli)
- [gccli](https://github.com/badlogic/gccli)
- [gdcli](https://github.com/badlogic/gdcli)

<!-- AIM_ECOSYSTEM_START -->
### 🧬 The A.I.M. Ecosystem

Modular A.I.M. (Actual Intelligent Memory) repositories. **Flagship engine: [aim-agy](https://github.com/BrianV1981/aim-agy).**

**Active vessels (CLI hosts):**
- **[aim-agy](https://github.com/BrianV1981/aim-agy)** — Core engine / *soul* (Antigravity CLI). *Flagship.* Shared nested `aim-agy_os/` ships here first.
- **[aim-grok](https://github.com/BrianV1981/aim-grok)** — Grok CLI vessel (hybrid memory, GitOps, wiki, fleet orchestration tooling).
- **[aim-opencode](https://github.com/BrianV1981/aim-opencode)** — OpenCode CLI vessel.
- **[aim-codex](https://github.com/BrianV1981/aim-codex)** — OpenAI Codex CLI vessel (greenfield nested soul + Codex overlays; primary `main`).

**Tools & workspaces:**
- **[aim-connect](https://github.com/BrianV1981/aim-connect)** — Self-hosted remote workspace web UI.
- **[aim-tmux-dashboard](https://github.com/BrianV1981/aim-tmux-dashboard)** — Terminal multi-session monitor.
- **[aim-browser](https://github.com/BrianV1981/aim-browser)** — Headed Chromium CDP engine + browser **skill suite**.
- **[aim-google](https://github.com/BrianV1981/aim-google)** — Google Workspace CLI (Gmail, Drive, Calendar, …).
- **[aim-flight-recorder](https://github.com/BrianV1981/aim-flight-recorder)** — Forensic Markdown session extractor.
- **[aim-boardroom](https://github.com/BrianV1981/aim-boardroom)** — Multi-agent orchestration room (OS multiplexing + artifacts).
- **[aim-skills](https://github.com/BrianV1981/aim-skills)** — **Skills index / multi-CLI install registry** (agy, grok, opencode, codex).

**DNA, comms & lore:**
- **[aim-coagents](https://github.com/BrianV1981/aim-coagents)** — DNA bank for sovereign co-agent blueprints.
- **[aim-knowledge](https://github.com/BrianV1981/aim-knowledge)** — Public Obsidian vault / deep-lore archive.
- **[aim-chalkboard](https://github.com/BrianV1981/aim-chalkboard)** — Optional cross-host async git mailbox (PoC; default same-host comms = **aim-communicate** skill).

**Deprecated / not maintained:**
- **[aim](https://github.com/BrianV1981/aim)** — Original **Gemini CLI** framework. Deprecated after loss of practical subscription access; **Great Migration → aim-agy**.
- **[aim-swarm](https://github.com/BrianV1981/aim-swarm)** — Legacy Python swarm factory → use **aim-coagents** + aim-agy spawn.
- **aim-claude / Anthropic-line vessels** — **Done.** Operator does not develop against Anthropic. Use **aim-agy / aim-grok / aim-opencode / aim-codex**.

Full map: see **aim-skills** `docs/AIM_ECOSYSTEM_MAP.md` or Operator artifact `AIM_ECOSYSTEM_MAP.md`.
<!-- AIM_ECOSYSTEM_END -->

