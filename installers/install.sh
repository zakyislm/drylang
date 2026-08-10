#!/bin/bash
set -e

REPO="zakyislm/drylang"
BIN_NAME="dry"

echo -e "\033[1;36mInstalling dryLang...\033[0m"

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
    linux|darwin) ;;
    *) echo -e "\033[1;31mUnsupported OS: $OS\033[0m"; exit 1 ;;
esac

# Detect Architecture
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo -e "\033[1;31mUnsupported architecture: $ARCH\033[0m"; exit 1 ;;
esac

ASSET_NAME="drylang-${OS}-${ARCH}"
if [ "$OS" = "windows" ]; then
    ASSET_NAME="${ASSET_NAME}.exe"
fi

# Fetch latest release URL
API_URL="https://api.github.com/repos/${REPO}/releases/latest"
echo "Fetching latest release from GitHub..."
DOWNLOAD_URL=$(curl -s $API_URL | grep "browser_download_url" | grep "$ASSET_NAME" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo -e "\033[1;33mWarning: Could not find pre-compiled binary for ${OS}-${ARCH} in the latest release.\033[0m"
    echo -e "You may need to build from source: \033[1;32mgo build -o dry .\033[0m"
    exit 1
fi

echo "Downloading $DOWNLOAD_URL ..."
TMP_BIN="/tmp/$BIN_NAME"
curl -sSL -o "$TMP_BIN" "$DOWNLOAD_URL"
chmod +x "$TMP_BIN"

# Install to /usr/local/bin
INSTALL_DIR="/usr/local/bin"
echo "Installing to $INSTALL_DIR (may require sudo)..."

if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_BIN" "$INSTALL_DIR/$BIN_NAME"
else
    sudo mv "$TMP_BIN" "$INSTALL_DIR/$BIN_NAME"
fi

# Aliasing 'y' -> 'dry'
if [ -w "$INSTALL_DIR" ]; then
    ln -sf "$INSTALL_DIR/$BIN_NAME" "$INSTALL_DIR/y"
else
    sudo ln -sf "$INSTALL_DIR/$BIN_NAME" "$INSTALL_DIR/y"
fi

echo -e "\n\033[1;32mInstallation Complete!\033[0m"
echo "dryLang is installed at: $INSTALL_DIR/$BIN_NAME (and aliased to 'y')"
echo "Type 'dry --version' to verify."
