# auth-api

## Local setup:
1. Run make local to start the server
2. Navigate to [https://localhost:8000]
3. You can check [https://localhost:8000/status] to check the status of the server

TODO: Setup docker

## Formatting:
We format the project using the go imports package.

1. Install VS Code's Go package
2. Install go imports `go get golang.org/x/tools/cmd/goimports`
3. If it hasn't been done automatically, create a folder `.vscode` and a file within it `settings.json` and add `"go.formatTool": "goimports"` to modify your formatter.
