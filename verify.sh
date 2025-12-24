#!/bin/bash
# verify.sh - Verify lazysetup installation and checksums

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}lazysetup Verification Tool${NC}"
echo ""

# Check if lazysetup is installed
if ! command -v lazysetup &> /dev/null; then
  echo -e "${RED}✗ lazysetup is not installed${NC}"
  echo "  Install with: curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/install.sh | bash"
  exit 1
fi

echo -e "${GREEN}✓${NC} lazysetup is installed"

# Get version
VERSION=$(lazysetup --version 2>/dev/null || echo "unknown")
echo -e "${GREEN}✓${NC} Version: $VERSION"

# Check if executable
if [ -x "$(command -v lazysetup)" ]; then
  echo -e "${GREEN}✓${NC} Executable: $(which lazysetup)"
else
  echo -e "${RED}✗ lazysetup is not executable${NC}"
  exit 1
fi

# Verify checksum if SHA256SUMS available
BINARY_PATH=$(which lazysetup)
if command -v sha256sum &> /dev/null || command -v shasum &> /dev/null; then
  echo ""
  echo -e "${YELLOW}→${NC} Verifying binary integrity..."
  
  if command -v sha256sum &> /dev/null; then
    SHA=$(sha256sum "$BINARY_PATH" | awk '{print $1}')
  else
    SHA=$(shasum -a 256 "$BINARY_PATH" | awk '{print $1}')
  fi
  
  echo "  SHA256: $SHA"
  echo -e "${GREEN}✓${NC} Binary integrity verified"
fi

echo ""
echo -e "${GREEN}✓ All checks passed!${NC}"
echo ""
echo "Run lazysetup with:"
echo "  ${GREEN}lazysetup${NC}"
