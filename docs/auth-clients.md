# OAuth Clients

Use multiple OAuth client credentials (for different Google Cloud projects or brands) without mixing refresh tokens.

## How it works

- Default client name: `default`
- Default credentials file: `$(os.UserConfigDir())/aim-google/credentials.json`
- Named credentials files: `$(os.UserConfigDir())/aim-google/credentials-<client>.json`
- Tokens are stored per client (`token:<client>:<email>`). Default client also writes legacy keys for backwards compatibility.
- Default account is stored per client, with a legacy global fallback for the default client.

## Selecting a client

Use `--client` (or `AIM_GOOGLE_CLIENT`) to pick which credentials + token bucket to use:

```
aim-google --client work auth credentials ~/Downloads/work-client.json
aim-google --client work auth add you@company.com
aim-google --client work gmail search "is:unread"
```

When `--client` is not set, `aim-google` resolves the client in this order:

1) `--client` / `AIM_GOOGLE_CLIENT` override
2) `account_clients` map in config
3) `client_domains` map in config
4) Credentials file named after the email domain (e.g. `credentials-example.com.json`)
5) `default`

## Domain auto-map

To auto-select a client for a domain:

```
aim-google --client work auth credentials ~/Downloads/work.json --domain example.com
```

This writes `client_domains` into `config.json` so any `@example.com` account selects the `work` client.

## Listing stored credentials

```
aim-google auth credentials list
```

Shows stored credential files plus any configured domain mappings.

## Config example

```
{
  keyring_backend: "auto",
  account_clients: {
    "you@company.com": "work",
  },
  client_domains: {
    "example.com": "work",
  },
}
```

## Migration notes

- Legacy `token:<email>` entries are copied to `token:default:<email>` the first time they are read.
- Legacy `default_account` is still respected for the default client.
