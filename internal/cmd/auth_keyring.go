package cmd

import (
	"context"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/BrianV1981/aim-google/internal/config"
	"github.com/BrianV1981/aim-google/internal/outfmt"
	"github.com/BrianV1981/aim-google/internal/secrets"
	"github.com/BrianV1981/aim-google/internal/ui"
)

type AuthKeyringCmd struct {
	Backend  string `arg:"" optional:"" name:"backend" help:"Keyring backend: auto|keychain|file"`
	Backend2 string `arg:"" optional:"" name:"backend2" help:"(compat) Use: aim-google auth keyring set <backend>"`
}

func (c *AuthKeyringCmd) Run(ctx context.Context, flags *RootFlags) error {
	u := ui.FromContext(ctx)

	const keyringPasswordEnv = "AIM_GOOGLE_KEYRING_PASSWORD" //nolint:gosec // env var name, not a credential

	backend := strings.ToLower(strings.TrimSpace(c.Backend))
	backend2 := strings.ToLower(strings.TrimSpace(c.Backend2))

	// Backwards compat for earlier suggestion: `aim-google auth keyring set <backend>`.
	if backend == "set" {
		backend = backend2
		backend2 = ""
	}

	// No args: show current config.
	if backend == "" {
		path, _ := config.ConfigPath()
		info, err := secrets.ResolveKeyringBackendInfo()
		if err != nil {
			return err
		}

		if outfmt.IsJSON(ctx) {
			return outfmt.WriteJSON(ctx, os.Stdout, map[string]any{
				"keyring_backend": info.Value,
				"source":          info.Source,
				"path":            path,
			})
		}

		if u == nil {
			return nil
		}
		u.Out().Printf("path\t%s", path)
		u.Out().Printf("keyring_backend\t%s", info.Value)
		u.Out().Printf("source\t%s", info.Source)
		u.Err().Println("Hint: aim-google auth keyring <auto|keychain|file>")
		return nil
	}

	if backend2 != "" {
		return usagef("too many args: %q %q", c.Backend, c.Backend2)
	}

	if backend == "default" {
		backend = "auto"
	}

	allowed := map[string]struct{}{
		"auto":     {},
		"keychain": {},
		strFile:    {},
	}
	if _, ok := allowed[backend]; !ok {
		return usagef("invalid backend: %q (expected auto, keychain, or file)", c.Backend)
	}

	path, _ := config.ConfigPath()
	if err := dryRunExit(ctx, flags, "auth.keyring.set", map[string]any{
		"path":            path,
		"keyring_backend": backend,
	}); err != nil {
		return err
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}
	cfg.KeyringBackend = backend
	if err := config.WriteConfig(cfg); err != nil {
		return err
	}

	// Env var wins; warn so it doesn't look "broken".
	if v := strings.TrimSpace(os.Getenv("AIM_GOOGLE_KEYRING_BACKEND")); v != "" &&
		u != nil &&
		!outfmt.IsJSON(ctx) &&
		!outfmt.IsPlain(ctx) {
		u.Err().Printf("NOTE: AIM_GOOGLE_KEYRING_BACKEND=%s overrides config.json", v)
	}

	if backend == strFile &&
		u != nil &&
		!outfmt.IsJSON(ctx) &&
		!outfmt.IsPlain(ctx) {
		if v := strings.TrimSpace(os.Getenv(keyringPasswordEnv)); v != "" {
			u.Err().Println("AIM_GOOGLE_KEYRING_PASSWORD found in environment.")
		} else if !term.IsTerminal(int(os.Stdin.Fd())) { //nolint:gosec // os file descriptor fits int on supported targets
			u.Err().Printf("NOTE: file keyring backend in non-interactive context requires %s", keyringPasswordEnv)
		} else {
			u.Err().Printf("Hint: set %s for non-interactive use (CI/ssh)", keyringPasswordEnv)
		}
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(ctx, os.Stdout, map[string]any{
			"written":         true,
			"path":            path,
			"keyring_backend": backend,
		})
	}

	if u == nil {
		return nil
	}

	u.Out().Printf("written\ttrue")
	u.Out().Printf("path\t%s", path)
	u.Out().Printf("keyring_backend\t%s", backend)
	return nil
}
