---
summary: "Release checklist for aim-google (GitHub release + Homebrew tap)"
---

# Releasing `aim-google`

This playbook mirrors the Homebrew + GitHub flow used in `../camsnap`.

Always do **all** steps below (CI + changelog + tag + GitHub release artifacts + tap update + Homebrew sanity install). No partial releases.

Shortcut scripts (preferred, keep notes non-empty):
```sh
scripts/release.sh X.Y.Z
scripts/verify-release.sh X.Y.Z
```

Assumptions:
- Repo: `BrianV1981/aim-google`
- Tap repo: `../homebrew-tap` (tap: `BrianV1981/tap`)
- Homebrew formula name: `aim-google` (installs the `aim-google` binary)

## 0) Prereqs
- Clean working tree on `main`.
- Go toolchain installed (Go version comes from `go.mod`).
- `make` works locally.
- Access to the tap repo (e.g. `BrianV1981/homebrew-tap`).

## 1) Verify build is green
```sh
make ci
```

Confirm GitHub Actions `ci` is green for the commit you’re tagging:
```sh
gh run list -L 5 --branch main
```

## 2) Update changelog
- Update `CHANGELOG.md` for the version you’re releasing.

Example heading:
- `## 0.1.0 - 2025-12-12`

## 3) Commit, tag & push
```sh
git checkout main
git pull

# commit changelog + any release tweaks
git commit -am "release: vX.Y.Z"

git tag -a vX.Y.Z -m "Release X.Y.Z"
git push origin main --tags
```

## 4) Verify GitHub release artifacts
The tag push triggers `.github/workflows/release.yml` (GoReleaser). Ensure it completes successfully and the release has assets.

```sh
gh run list -L 5 --workflow release.yml
gh release view vX.Y.Z
```

Ensure GitHub release notes are not empty (mirror the changelog section).

If the workflow needs a rerun:
```sh
gh workflow run release.yml -f tag=vX.Y.Z
```

## 5) Update (or add) the Homebrew formula
In the tap repo (assumed sibling at `../homebrew-tap`), create/update `Formula/aim-google.rb`.

Recommended formula shape (build-from-source, no binary assets needed):
- `version "X.Y.Z"`
- `url "https://github.com/BrianV1981/aim-google/archive/refs/tags/vX.Y.Z.tar.gz"`
- `sha256 "<sha256>"`
- `depends_on "go" => :build`
- Build:
  - `system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/aim-google"`

Compute the SHA256 for the tag tarball:
```sh
curl -L -o /tmp/aim-google.tar.gz https://github.com/BrianV1981/aim-google/archive/refs/tags/vX.Y.Z.tar.gz
shasum -a 256 /tmp/aim-google.tar.gz
```

Commit + push in the tap repo:
```sh
cd ../homebrew-tap
git add Formula/aim-google.rb
git commit -m "aim-google vX.Y.Z"
git push origin main
```

## 6) Sanity-check install from tap
```sh
brew update
brew uninstall aim-google || true
brew untap BrianV1981/tap || true
brew tap BrianV1981/tap
brew install BrianV1981/tap/aim-google
brew test BrianV1981/tap/aim-google

aim-google --help
```

## Notes
- `aim-google --version` / `aim-google version` should report the release version post-install.
