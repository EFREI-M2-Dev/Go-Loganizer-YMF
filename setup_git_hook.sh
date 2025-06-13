#!/bin/bash

GIT_DIR=".git"

if [ ! -d "$GIT_DIR" ]; then
    echo "Erreur : Ce n'est pas un dépôt Git."
    exit 1
fi

mkdir -p "$GIT_DIR/hooks"

PRE_COMMIT_HOOK="$GIT_DIR/hooks/pre-commit"

cat << 'EOF' > "$PRE_COMMIT_HOOK"
#!/bin/sh

FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

if [ -n "$FILES" ]; then
  gofmt -w $FILES
  git add $FILES
fi
EOF

chmod +x "$PRE_COMMIT_HOOK"

echo "Le pre-commit hook pour go fmt a été installé avec succès."
