#!/bin/bash
# uninstall.sh - Remove lazysetup from your system

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}lazysetup Uninstaller${NC}"
echo ""

# Find lazysetup installation
LAZYSETUP_PATH=$(which lazysetup 2>/dev/null || echo "")

if [ -z "$LAZYSETUP_PATH" ]; then
  echo -e "${RED}✗ lazysetup is not installed${NC}"
  echo "  Nothing to uninstall"
  exit 0
fi

echo -e "${YELLOW}→${NC} Found lazysetup at: $LAZYSETUP_PATH"
echo ""

# Confirm uninstall
read -p "Are you sure you want to uninstall lazysetup? (y/N) " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
  echo -e "${YELLOW}→${NC} Uninstall cancelled"
  exit 0
fi

echo ""
echo -e "${YELLOW}→${NC} Removing lazysetup..."

# Check if we need sudo
if [ ! -w "$(dirname "$LAZYSETUP_PATH")" ]; then
  echo -e "${YELLOW}→${NC} Requesting sudo access..."
  if ! sudo rm -f "$LAZYSETUP_PATH"; then
    echo -e "${RED}✗ Failed to remove lazysetup${NC}"
    exit 1
  fi
else
  if ! rm -f "$LAZYSETUP_PATH"; then
    echo -e "${RED}✗ Failed to remove lazysetup${NC}"
    exit 1
  fi
fi

echo -e "${GREEN}✓${NC} lazysetup removed successfully"

# Verify removal
if command -v lazysetup &> /dev/null; then
  echo -e "${YELLOW}⚠${NC} Warning: lazysetup still found in PATH"
  echo "  Run: which lazysetup"
else
  echo -e "${GREEN}✓${NC} Verified: lazysetup is no longer installed"
fi

echo ""
echo -e "${GREEN}✓ Uninstall complete!${NC}"
