# ready-set-golang

Just a simple demo application to demonstrate some of the Go programming language features. It uses the Marvel API to retrieve a random character.
Only Go's SDK is used, no external dependencies.

## How to run

1. Get API keys from developers.marvel.com
2. Set your public and private API keys as environment variables with the following names, e.g.

```bash
export MARVEL_PUBLIC_KEY=your_public_key
export MARVEL_PRIVATE_KEY=your_private_key
```

3. Build and run `go run main.go`