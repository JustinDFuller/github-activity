echo "Running pre-commit hook. It worked if it prints OK."

echo "Stashing changes."
# pre-commit.sh
STASH_NAME="pre-commit-$(date +%s)"
git stash save -q --keep-index $STASH_NAME

# Test prospective commit
echo "Running go fmt and test."
go fmt ./... && go test -race ./...
RESULT=$?

# Restore stashed items
echo "Restoring stash."
git stash pop -q

# Exit based on the result of the tests.
[ $RESULT -ne 0 ] && exit 1

echo "OK"
exit 0
