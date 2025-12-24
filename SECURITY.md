# Security Policy

## Installation Security

lazysetup is designed with security as a top priority. Here's how we keep your installation safe:

### Checksum Verification

Every binary release is accompanied by a `SHA256SUMS` file containing cryptographic checksums. The installer automatically:

1. **Downloads the binary** from GitHub Releases
2. **Downloads the checksums** from the same release
3. **Verifies the checksum** before installation
4. **Aborts if verification fails** - protecting you from corrupted or tampered files

```bash
# The installer does this automatically:
sha256sum lazysetup-v0.0.1-linux-amd64
# Expected: matches SHA256SUMS entry
```

### Secure Download

- **HTTPS only**: All downloads use encrypted HTTPS connections
- **Retry logic**: Automatic retries with exponential backoff for failed downloads
- **Timeout protection**: 10-second connection timeout, 5-minute download timeout
- **Temp directory**: Files are downloaded to a temporary directory and cleaned up after installation

### Installation Safety

- **Minimal permissions**: Only requests `sudo` when writing to `/usr/local/bin`
- **Atomic installation**: Binary is moved (not copied) to prevent partial installations
- **Executable verification**: Binary is marked executable before installation
- **Error handling**: Installation aborts with clear error messages on any failure

## Verification

After installation, verify the integrity of your binary:

```bash
# Verify installation
curl -fsSL https://raw.githubusercontent.com/youpele52/lazysetup/main/verify.sh | bash

# Manual verification
sha256sum $(which lazysetup)
```

## Release Process

### How Binaries Are Built

1. **GitHub Actions** builds binaries automatically on every release tag
2. **Cross-platform compilation** ensures binaries for:
   - macOS (amd64, arm64)
   - Linux (amd64, arm64)
3. **Checksums generated** automatically by GitHub Actions
4. **Uploaded to GitHub Releases** with both binaries and checksums

### Checksum Generation

```bash
# GitHub Actions generates checksums:
sha256sum lazysetup-* > SHA256SUMS

# Example output:
# a1b2c3d4... lazysetup-v0.0.1-linux-amd64
# e5f6g7h8... lazysetup-v0.0.1-darwin-amd64
```

## Reporting Security Issues

If you discover a security vulnerability, please email security@example.com instead of using the issue tracker.

Please include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

We will acknowledge receipt within 48 hours and provide updates on our progress.

## Security Best Practices

### For Users

1. **Always verify checksums** - Don't skip checksum verification
2. **Use HTTPS** - Always download from https://github.com/youpele52/lazysetup
3. **Keep updated** - Install the latest version for security patches
4. **Review scripts** - Read `install.sh` before running it
5. **Check permissions** - Verify `/usr/local/bin/lazysetup` is owned by root or your user

### For Developers

1. **Code review** - All changes reviewed before merge
2. **Dependency management** - Keep Go dependencies updated
3. **Testing** - Automated tests run on every commit
4. **Signed commits** - Use GPG signatures for releases
5. **Audit logs** - GitHub Actions logs all build activities

## Transparency

- All source code is open source and publicly available
- Build process is automated and transparent
- Release checksums are publicly available
- No telemetry or tracking in the application

## Supported Versions

| Version | Status | Security Updates |
|---------|--------|------------------|
| v0.0.1+ | Active | Yes |

## Dependencies

lazysetup uses minimal dependencies to reduce attack surface:

```go
require (
    github.com/jesseduffield/gocui v0.5.0
)
```

All dependencies are:
- Regularly updated
- Reviewed for security
- Pinned to specific versions in `go.mod`

## Changelog

Security updates are documented in [CHANGELOG.md](CHANGELOG.md) with the `[SECURITY]` tag.
