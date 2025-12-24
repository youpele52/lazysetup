#!/bin/bash
set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}lazysetup Uninstaller${NC}"
echo ""

LAZYSETUP_PATH=$(which lazysetup 2>/dev/null || echo "")

if [ -z "$LAZYSETUP_PATH" ]; then
  echo -e "${RED}✗ Not installed${NC}"
  exit 0
fi

echo -e "${YELLOW}→${NC} Found at: $LAZYSETUP_PATH"
echo ""
echo "Uninstall lazysetup? (y/N)"
read -n 1 -r REPLY
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
  echo -e "${YELLOW}→${NC} Cancelled"
  exit 0
fi

echo -e "${YELLOW}→${NC} Removing..."

DIR=$(dirname "$LAZYSETUP_PATH")
if [ ! -w "$DIR" ]; then
  echo -e "${YELLOW}→${NC} Need sudo..."
  sudo rm -f "$LAZYSETUP_PATH" || exit 1
else
  rm -f "$LAZYSETUP_PATH" || exit 1
fi

echo -e "${GREEN}✓${NC} Removed successfully"

if command -v lazysetup &> /dev/null; then
  echo -e "${YELLOW}⚠${NC} Still in PATH"
else
  echo -e "${GREEN}✓${NC} Verified removed"
fi

echo ""
echo -e "${GREEN}✓ Done!${NC}"
