#!/bin/sh
# sky-log Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/anzellai/sky-log/main/install.sh | sh
#
# Environment variables:
#   SKY_LOG_VERSION      - specific version to install (default: latest)
#   SKY_LOG_INSTALL_DIR  - installation directory (default: /usr/local/bin)
set -e

REPO="anzellai/sky-log"
BINARY_NAME="sky-log"
INSTALL_DIR="${SKY_LOG_INSTALL_DIR:-/usr/local/bin}"

RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

info() { printf "${CYAN}==>${NC} %s\n" "$1"; }
success() { printf "${GREEN}==>${NC} %s\n" "$1"; }
error() { printf "${RED}error:${NC} %s\n" "$1" >&2; exit 1; }

detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux)  PLATFORM="linux" ;;
        Darwin) PLATFORM="darwin" ;;
        MINGW*|MSYS*|CYGWIN*) PLATFORM="windows" ;;
        *) error "Unsupported OS: $OS" ;;
    esac

    case "$ARCH" in
        x86_64|amd64)  ARCH="x64" ;;
        arm64|aarch64) ARCH="arm64" ;;
        *) error "Unsupported architecture: $ARCH" ;;
    esac
}

get_latest_version() {
    if ! command -v curl >/dev/null 2>&1; then
        error "curl is required but not installed"
    fi
    # Use GITHUB_TOKEN if available (avoids API rate limits in CI)
    AUTH_HEADER=""
    if [ -n "$GITHUB_TOKEN" ]; then
        AUTH_HEADER="-H \"Authorization: token $GITHUB_TOKEN\""
    fi
    VERSION=$(eval curl -fsSL $AUTH_HEADER "https://api.github.com/repos/$REPO/releases/latest" 2>/dev/null | grep '"tag_name"' | sed 's/.*"v\(.*\)".*/\1/')
    if [ -z "$VERSION" ]; then
        error "Could not determine latest version. Check https://github.com/$REPO/releases"
    fi
}

main() {
    printf "\n${BOLD}sky-log Installer${NC}\n\n"

    detect_platform

    if [ -n "$SKY_LOG_VERSION" ]; then
        VERSION="$SKY_LOG_VERSION"
    else
        get_latest_version
    fi

    info "Platform: ${PLATFORM}/${ARCH}"
    info "Version:  v${VERSION}"
    echo ""

    EXT=""
    if [ "$PLATFORM" = "windows" ]; then
        EXT=".exe"
    fi

    ARTIFACT="${BINARY_NAME}-${PLATFORM}-${ARCH}${EXT}"
    URL="https://github.com/$REPO/releases/download/v${VERSION}/${ARTIFACT}"

    info "Downloading ${BINARY_NAME} v${VERSION}..."

    TMPFILE=$(mktemp)
    trap 'rm -f "$TMPFILE"' EXIT

    if ! curl -fsSL "$URL" -o "$TMPFILE"; then
        error "Failed to download $URL\nCheck that v${VERSION} exists at https://github.com/$REPO/releases"
    fi

    chmod +x "$TMPFILE"

    if [ -w "$INSTALL_DIR" ]; then
        mv "$TMPFILE" "$INSTALL_DIR/${BINARY_NAME}${EXT}"
    else
        info "Requires sudo to install to $INSTALL_DIR"
        sudo mv "$TMPFILE" "$INSTALL_DIR/${BINARY_NAME}${EXT}"
    fi

    success "Installed ${BINARY_NAME} -> ${INSTALL_DIR}/${BINARY_NAME}${EXT}"

    printf "\n${GREEN}${BOLD}sky-log v${VERSION} installed successfully!${NC}\n\n"
    echo "  Usage:"
    echo "    sky-log                    # start the Sky.Live log viewer"
    echo "    sky-log --help             # show options"
    echo ""
}

main "$@"
