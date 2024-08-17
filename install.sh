#!/bin/bash

make build
current_shell=$(basename "$SHELL")

case "$current_shell" in
  "bash")
    rc_file="$HOME/.bashrc"
    ;;
  "zsh")
    rc_file="$HOME/.zshrc"
    ;;
  "fish")
    rc_file="$HOME/.config/fish/config.fish"
    ;;
  *)
    echo "Unsupported shell: $current_shell"
    exit 1
    ;;
esac

curr_dir=$(pwd)

echo "export PATH=\"${curr_dir}/bin/:\$PATH\"" >> ${rc_file}
echo "Please restart your shell or run 'source $rc_file' to apply the changes."
