#!/bin/bash
# install.sh - Secure installer for lazysetup
# Downloads and installs lazysetup with checksum verification

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
VERSION="${1:-latest}"
GITHUB_REPO="youpele52/lazysetup"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

# Resolve 'latest' to actual version tag
if [ "$VERSION" = "latest" ]; then
  echo -e "${YELLOW}→${NC} Resolving latest version..."
  VERSION=$(curl -fsSL "https://api.github.com/repos/${GITHUB_REPO}/releases/latest" | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": "\([^"]*\)".*/\1/')
  if [ -z "$VERSION" ]; then
    echo -e "${RED}✗ Failed to resolve latest version${NC}"
    exit 1
  fi
  echo -e "${GREEN}✓${NC} Latest version: $VERSION"
fi

# Detect OS and architecture
detect_platform() {
  local OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  local ARCH=$(uname -m)

  # Map architecture names
  case $ARCH in
    x86_64)
      ARCH="amd64"
      ;;
    aarch64)
      ARCH="arm64"
      ;;
    arm64)
      ARCH="arm64"
      ;;
    *)
      echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
      exit 1
      ;;
  esac

  # Map OS names
  case $OS in
    darwin)
      OS="darwin"
      ;;
    linux)
      OS="linux"
      ;;
    *)
      echo -e "${RED}✗ Unsupported OS: $OS${NC}"
      exit 1
      ;;
  esac

  echo "${OS}-${ARCH}"
}

# Download file with retry logic
download_file() {
  local url=$1
  local output=$2
  local max_attempts=3
  local attempt=1

  while [ $attempt -le $max_attempts ]; do
    echo -e "${YELLOW}→${NC} Downloading (attempt $attempt/$max_attempts)..."
    if curl -fsSL --connect-timeout 10 --max-time 300 -o "$output" "$url"; then
      return 0
    fi
    attempt=$((attempt + 1))
    if [ $attempt -le $max_attempts ]; then
      sleep 2
    fi
  done

  return 1
}

# Verify checksum
verify_checksum() {
  local file=$1
  local expected_sha=$2

  echo -e "${YELLOW}→${NC} Verifying checksum..."
  
  local actual_sha
  if command -v sha256sum &> /dev/null; then
    actual_sha=$(sha256sum "$file" | awk '{print $1}')
  elif command -v shasum &> /dev/null; then
    actual_sha=$(shasum -a 256 "$file" | awk '{print $1}')
  else
    echo -e "${YELLOW}⚠${NC} Warning: sha256sum/shasum not found, skipping checksum verification"
    return 0
  fi

  if [ "$actual_sha" != "$expected_sha" ]; then
    echo -e "${RED}✗ Checksum verification failed!${NC}"
    echo "  Expected: $expected_sha"
    echo "  Got:      $actual_sha"
    return 1
  fi

  echo -e "${GREEN}✓${NC} Checksum verified"
  return 0
}

# Main installation
main() {
  echo -e "${GREEN}lazysetup Installer${NC}"
  echo "Version: $VERSION"
  echo ""

  # Detect platform
  PLATFORM=$(detect_platform)
  echo -e "${GREEN}✓${NC} Detected platform: $PLATFORM"

  # Determine binary name
  BINARY_NAME="lazysetup-${VERSION}-${PLATFORM}"
  if [[ $PLATFORM == *"windows"* ]]; then
    BINARY_NAME="${BINARY_NAME}.exe"
  fi

  # Download URLs
  DOWNLOAD_URL="https://github.com/${GITHUB_REPO}/releases/download/${VERSION}/${BINARY_NAME}"
  CHECKSUMS_URL="https://github.com/${GITHUB_REPO}/releases/download/${VERSION}/SHA256SUMS"

  echo ""
  echo -e "${YELLOW}→${NC} Downloading lazysetup..."
  if ! download_file "$DOWNLOAD_URL" "$TEMP_DIR/lazysetup"; then
    echo -e "${RED}✗ Failed to download lazysetup${NC}"
    echo "  URL: $DOWNLOAD_URL"
    exit 1
  fi
  echo -e "${GREEN}✓${NC} Downloaded successfully"

  echo ""
  echo -e "${YELLOW}→${NC} Downloading checksums..."
  if ! download_file "$CHECKSUMS_URL" "$TEMP_DIR/SHA256SUMS"; then
    echo -e "${RED}✗ Failed to download checksums${NC}"
    echo "  URL: $CHECKSUMS_URL"
    exit 1
  fi
  echo -e "${GREEN}✓${NC} Downloaded checksums"

  echo ""
  # Extract expected checksum
  local expected_sha=$(grep "$BINARY_NAME" "$TEMP_DIR/SHA256SUMS" | awk '{print $1}')
  if [ -z "$expected_sha" ]; then
    echo -e "${RED}✗ Checksum not found for $BINARY_NAME${NC}"
    exit 1
  fi

  # Verify checksum
  if ! verify_checksum "$TEMP_DIR/lazysetup" "$expected_sha"; then
    exit 1
  fi

  echo ""
  # Make executable
  chmod +x "$TEMP_DIR/lazysetup"

  # Install
  echo -e "${YELLOW}→${NC} Installing to $INSTALL_DIR..."
  if [ ! -w "$INSTALL_DIR" ]; then
    echo -e "${YELLOW}→${NC} Requesting sudo access..."
    if ! sudo mv "$TEMP_DIR/lazysetup" "$INSTALL_DIR/lazysetup"; then
      echo -e "${RED}✗ Installation failed${NC}"
      exit 1
    fi
  else
    if ! mv "$TEMP_DIR/lazysetup" "$INSTALL_DIR/lazysetup"; then
      echo -e "${RED}✗ Installation failed${NC}"
      exit 1
    fi
  fi
  echo -e "${GREEN}✓${NC} Installed successfully"

  echo ""
  echo -e "${GREEN}✓ Installation complete!${NC}"
  echo ""
  echo "Run lazysetup with:"
  echo -e "  ${GREEN}lazysetup${NC}"
  echo ""
  echo "For help:"
  echo -e "  ${GREEN}lazysetup --help${NC}"
}

main "$@"
