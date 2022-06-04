wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Watch for .go file changes
CompileDeamon --build="go build -o main main.go" --command=./main