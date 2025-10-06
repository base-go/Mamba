#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if version argument is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Version number required${NC}"
    echo "Usage: ./release.sh <version>"
    echo "Example: ./release.sh 1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format (basic semver check)
if ! [[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo -e "${RED}Error: Invalid version format${NC}"
    echo "Version must be in format: X.Y.Z (e.g., 1.0.0)"
    exit 1
fi

echo -e "${YELLOW}Preparing release v${VERSION}${NC}"

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo -e "${RED}Error: Working directory is not clean${NC}"
    echo "Please commit or stash your changes before creating a release"
    exit 1
fi

# Check if on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo -e "${YELLOW}Warning: Not on main branch (current: $CURRENT_BRANCH)${NC}"
    read -p "Continue anyway? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Pull latest changes
echo -e "${YELLOW}Pulling latest changes...${NC}"
git pull origin $CURRENT_BRANCH

# Run tests
echo -e "${YELLOW}Running tests...${NC}"
if ! go test ./...; then
    echo -e "${RED}Error: Tests failed${NC}"
    exit 1
fi

# Run go vet
echo -e "${YELLOW}Running go vet...${NC}"
if ! go vet ./...; then
    echo -e "${RED}Error: go vet failed${NC}"
    exit 1
fi

# Build the project
echo -e "${YELLOW}Building project...${NC}"
if ! go build ./...; then
    echo -e "${RED}Error: Build failed${NC}"
    exit 1
fi

# Delete existing tag if it exists (local and remote)
if git rev-parse "v${VERSION}" >/dev/null 2>&1; then
    echo -e "${YELLOW}Tag v${VERSION} already exists locally, deleting...${NC}"
    git tag -d "v${VERSION}"
fi

if git ls-remote --tags origin | grep -q "refs/tags/v${VERSION}"; then
    echo -e "${YELLOW}Tag v${VERSION} already exists on remote, deleting...${NC}"
    git push origin ":refs/tags/v${VERSION}"
fi

# Create git tag
echo -e "${YELLOW}Creating git tag v${VERSION}...${NC}"
git tag -a "v${VERSION}" -m "Release v${VERSION}"

# Push tag to remote
echo -e "${YELLOW}Pushing tag to remote...${NC}"
git push origin "v${VERSION}"

# Create GitHub release using gh CLI
echo -e "${YELLOW}Creating GitHub release...${NC}"
if command -v gh &> /dev/null; then
    # Extract release notes from CHANGELOG.md for this version
    RELEASE_NOTES=$(awk "/## \[${VERSION}\]/,/## \[/" CHANGELOG.md | grep -v "^## \[" | sed -e '/^$/d' -e '$d')

    if [ -z "$RELEASE_NOTES" ]; then
        RELEASE_NOTES="Release v${VERSION}"
    fi

    echo "$RELEASE_NOTES" | gh release create "v${VERSION}" \
        --title "v${VERSION}" \
        --notes-file -

    echo -e "${GREEN}Successfully created GitHub release v${VERSION}${NC}"
    echo -e "${YELLOW}View release at: https://github.com/base-go/mamba/releases/tag/v${VERSION}${NC}"
else
    echo -e "${YELLOW}Warning: gh CLI not found${NC}"
    echo -e "${YELLOW}Tag v${VERSION} pushed successfully, but GitHub release not created${NC}"
    echo -e "${YELLOW}Install gh CLI or create release manually at:${NC}"
    echo -e "${YELLOW}https://github.com/base-go/mamba/releases/new?tag=v${VERSION}${NC}"
fi
