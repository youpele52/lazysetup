# Publishing Process

This document details the complete process for publishing lazysetup releases.

## Overview

The publishing workflow is automated via GitHub Actions. When a version tag is pushed, the workflow:
1. Builds binaries for all platforms
2. Generates checksums
3. Creates a GitHub Release with all artifacts

## Prerequisites

- Git repository with remote `origin` pointing to GitHub
- GitHub Actions enabled on the repository
- Release workflow file at `.github/workflows/release.yml`
- Installation scripts: `install.sh`, `uninstall.sh`, `verify.sh`

## Step-by-Step Publishing Process

### 1. Prepare Release

Ensure all changes are committed and the code is ready for release:

```bash
# Check git status
git status

# Ensure all files are committed
git add .
git commit -m "Release v0.1.0"
```

### 2. Create Version Tag

Create a semantic version tag following the pattern `vX.Y.Z`:

```bash
# Create tag
git tag v0.1.0

# Verify tag was created
git tag -l
```

### 3. Push Tag to GitHub

Push the tag to trigger GitHub Actions:

```bash
# Push specific tag
git push origin v0.1.0

# Or push all tags
git push origin --tags
```

### 4. Monitor GitHub Actions

The release workflow automatically starts when a tag is pushed:

1. Go to: `https://github.com/youpele52/lazysetup/actions`
2. Look for "Build and Release" workflow
3. Monitor the build progress

**Workflow steps:**
- Build binaries for each platform (macOS amd64/arm64, Linux amd64/arm64)
- Generate SHA256SUMS file
- Create GitHub Release with all artifacts

**Expected duration:** 2-5 minutes

### 5. Verify Release

Once the workflow completes:

1. Check the release: `https://github.com/youpele52/lazysetup/releases/tag/vX.Y.Z`
2. Verify all binaries are present:
   - `lazysetup-vX.Y.Z-darwin-amd64`
   - `lazysetup-vX.Y.Z-darwin-arm64`
   - `lazysetup-vX.Y.Z-linux-amd64`
   - `lazysetup-vX.Y.Z-linux-arm64`
   - `SHA256SUMS`
3. Verify checksums match binaries

### 6. Test Installation

Test the curl installation with the new release:

```bash
# Test latest release
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/install.sh | bash

# Or specific version
curl -fsSL https://github.com/youpele52/lazysetup/releases/download/vX.Y.Z/install.sh | bash
```

### 7. Announce Release

Update documentation and announce the release:
- Update README with new version info
- Create release notes on GitHub
- Announce on social media/forums if applicable

## Workflow File Details

Location: `.github/workflows/release.yml`

**Triggers:** Any push of a tag matching `v*` pattern

**Jobs:**
1. **build**: Compiles binaries for all platforms
2. **create-release**: Generates checksums and creates GitHub Release

**Platforms built:**
- macOS (darwin): amd64, arm64
- Linux: amd64, arm64

## Troubleshooting

### Workflow Failed to Start

**Issue:** Tag pushed but workflow didn't start

**Solution:**
- Verify workflow file exists: `.github/workflows/release.yml`
- Check GitHub Actions is enabled in repository settings
- Verify tag matches pattern `v*` (e.g., `v1.0.0`)

### Build Failed

**Issue:** Workflow started but build step failed

**Solution:**
- Check workflow logs on GitHub Actions page
- Verify Go version in workflow matches `go.mod`
- Ensure all dependencies are available

### Release Not Created

**Issue:** Build succeeded but release wasn't created

**Solution:**
- Check `create-release` job logs
- Verify GitHub token has permissions
- Ensure SHA256SUMS file was generated

### Binaries Missing from Release

**Issue:** Release created but some binaries are missing

**Solution:**
- Re-run the workflow
- Check individual build job logs
- Verify platform matrix in workflow file

### Release Contains Wrong/Duplicate Binaries

**Issue:** Release contains binaries from previous versions (e.g., both `lazysetup-0.1.2-*` and `lazysetup-v0.3.2-*`)

**Symptoms:**
- Auto-update installs wrong version (downgrades to old version)
- Users see only 5 tools instead of 37
- Binary names inconsistent (some with "v" prefix, some without)

**Root Cause:**
Workflow didn't clean old binaries before organizing artifacts, causing artifacts from previous workflow runs to be included.

**Status:** ✅ **FIXED** as of commit 935c796 (v0.3.2)

The workflow now automatically cleans old binaries using version-based filtering. This ensures only binaries matching the current version tag are included in releases.

**Current Workflow Behavior:**
The `.github/workflows/release.yml` "Organize artifacts" step now includes:

```yaml
- name: Organize artifacts
  run: |
    VERSION=${{ github.ref_name }}
    echo "=== Cleaning old binaries (before) ==="
    rm -f lazysetup-* SHA256SUMS

    echo "=== Binaries directory structure ==="
    ls -laR binaries/

    echo "=== Moving files ==="
    mv binaries/* . 2>/dev/null || true

    echo "=== Cleaning old binaries (after move) ==="
    # Remove any binaries that don't match current version
    ls -la lazysetup-* 2>/dev/null || echo "No binaries found"
    # Keep only binaries matching current version tag
    find . -maxdepth 1 -name "lazysetup-*" ! -name "lazysetup-${VERSION}-*" -type f -delete

    echo "=== Final binaries (version ${VERSION} only) ==="
    ls -la lazysetup-* 2>/dev/null || echo "No binaries in root"
```

**How It Works:**
1. Cleans all old binaries before organizing
2. Moves new binaries from artifacts directory
3. Uses `find` with version filtering to remove any binaries not matching current tag
4. Ensures only `lazysetup-${VERSION}-*` files remain

**If Issue Persists:**
1. **Manual cleanup** - Delete old binaries from release:
```bash
# List current release assets
gh release view v0.3.2 --json assets

# Delete old binaries (replace version numbers as needed)
gh release delete-asset v0.3.2 lazysetup-0.1.2-darwin-amd64 --yes
gh release delete-asset v0.3.2 lazysetup-0.1.2-darwin-arm64 --yes
gh release delete-asset v0.3.2 lazysetup-0.1.2-linux-amd64 --yes
gh release delete-asset v0.3.2 lazysetup-0.1.2-linux-arm64 --yes
```

2. **Verify Fix:**
```bash
# Check only current version binaries remain
gh release view v0.3.2 --json assets | grep '"name":'
# Should only show: lazysetup-v0.3.2-* files
```

## Manual Release (If Needed)

If GitHub Actions fails, you can manually create a release:

```bash
# Build binaries locally (static builds with CGO_ENABLED=0)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o lazysetup-v0.1.0-darwin-amd64 .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o lazysetup-v0.1.0-darwin-arm64 .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lazysetup-v0.1.0-linux-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o lazysetup-v0.1.0-linux-arm64 .

# Generate checksums
sha256sum lazysetup-v0.1.0-* > SHA256SUMS

# Create release on GitHub (requires GitHub CLI)
gh release create v0.1.0 lazysetup-v0.1.0-* SHA256SUMS --generate-notes
```

## Version Numbering

Follow Semantic Versioning (SemVer):

- **MAJOR** (v1.0.0): Breaking changes
- **MINOR** (v0.1.0): New features, backward compatible
- **PATCH** (v0.0.1): Bug fixes, backward compatible

Examples:
- `v0.0.1` - Initial release
- `v0.1.0` - Add new feature
- `v1.0.0` - Major version
- `v1.0.1` - Patch release

## Release Checklist

Before publishing:

- [ ] All code changes committed
- [ ] Tests passing locally
- [ ] README updated with new version
- [ ] CHANGELOG.md updated
- [ ] Version tag created correctly (`vX.Y.Z`)
- [ ] Tag pushed to GitHub

After publishing:

- [ ] GitHub Actions workflow completed successfully
- [ ] All binaries present in release
- [ ] SHA256SUMS file generated
- [ ] Checksums verified
- [ ] Installation tested with curl
- [ ] Release notes published

## Automation Details

### GitHub Actions Workflow

**File:** `.github/workflows/release.yml`

**Triggers on:** `push` with tags matching `v*`

**Build Matrix:**
```yaml
- os: darwin, arch: amd64
- os: darwin, arch: arm64
- os: linux, arch: amd64
- os: linux, arch: arm64
```

**Output:**
- Binary files: `lazysetup-vX.Y.Z-{os}-{arch}`
- Checksums: `SHA256SUMS`
- GitHub Release with all artifacts

### Installation Scripts

**install.sh:**
- Downloads from: `releases/latest/download/install.sh`
- Queries GitHub API for actual version
- Downloads binary and checksums
- Verifies checksum before installation

**uninstall.sh:**
- Finds lazysetup installation
- Prompts for confirmation
- Removes binary safely

**verify.sh:**
- Confirms installation
- Verifies binary integrity
- Shows version information

## Future Enhancements

- [ ] Automated changelog generation
- [ ] GPG signing of releases
- [ ] Docker image publishing
- [ ] Package manager submissions (Homebrew, APT, etc.)
- [ ] Release notes generation from commits
- [ ] Automated version bumping
