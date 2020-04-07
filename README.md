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

## Database Migrations:
We utilize https://github.com/golang-migrate/migrate for handling database migrations. At a simple level,

To migrate to most recent version, run `migrate -source file://driver/migrations -database postgres://user:password@host:port/database up`

To migrate to a specific version, run `migrate -source file://driver/migrations -database postgres://user:password@host:port/database up {version number}`

To migrate down to a specific version, run `migrate -source file://driver/migrations -database postgres://user:password@host:port/database down {version number}`

*Be sure to let other team members know when migrating so that there are no schema conflicts/unexpected updates. It's best to let a single person maintain this to avoid conflicts.
