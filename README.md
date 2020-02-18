# ready-set-golang

Just a simple demo application to demonstrate some of the Go programming language features. It uses the Marvel API to retrieve a random character.
Only Go's standard libraries are used, no external dependencies.

This is part of my Ready, Set, Golang! talk. You can find the [slides here](https://docs.google.com/presentation/d/1B-mAV0r42zPXk9AIYmE2I8XEQdDuCWsrNLxESl6CLTM/edit?usp=sharing)

## How to run

1. Install the Go tools in your machine. [How to here](https://golang.org/doc/install)
1. Get API keys from developers.marvel.com [Slides](https://docs.google.com/presentation/d/1B-mAV0r42zPXk9AIYmE2I8XEQdDuCWsrNLxESl6CLTM/edit?usp=sharing)
2. Set your public and private API keys as environment variables with the following names, e.g.

```bash
export MARVEL_PUBLIC_KEY=your_public_key
export MARVEL_PRIVATE_KEY=your_private_key
```

3. Build and run `go run main.go`