#!/bin/bash

# Installs the latest stable version of Python via pyenv for the current user.
# Dependencies: bash, curl, grep, tail, tr, apt-get, sudo

set -e

echo "=== Installing latest Python ==="

echo "Installing Python build dependencies..."
sudo apt-get update -qq
sudo apt-get install -y -q git build-essential liblzma-dev libssl-dev zlib1g-dev \
    libbz2-dev libreadline-dev libsqlite3-dev libncurses5-dev libncursesw5-dev \
    xz-utils tk-dev libffi-dev

echo "Installing pyenv..."
if [ ! -d "$HOME/.pyenv" ]; then
    curl -s https://pyenv.run | bash
fi

export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)"

PY_LATEST=$(pyenv install --list | grep -E '^[[:space:]]*3\.[0-9]+\.[0-9]+$' | tail -1 | tr -d '[:space:]')
echo "Installing Python ${PY_LATEST} (this may take a few minutes)..."
pyenv install -s "$PY_LATEST"
pyenv global "$PY_LATEST"

if ! grep -q "PYENV_ROOT" "$HOME/.bashrc"; then
    {
        echo 'export PYENV_ROOT="$HOME/.pyenv"'
        echo 'export PATH="$PYENV_ROOT/bin:$PATH"'
        echo 'eval "$(pyenv init -)"'
    } >> "$HOME/.bashrc"
fi
echo "Python version: $(python --version)"

echo "=== Python installation complete ==="
