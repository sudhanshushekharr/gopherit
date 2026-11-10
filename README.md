# gopherit 🐹

*Go for it.* I'm learning Go (Golang) from zero and publishing it here one chapter at a time, as I go.

I'm a working software engineer who mostly knows Java, Python and JS. Each chapter is a set of small runnable programs I wrote while practising what is **different or idiomatic in Go**.

## Progress

| # | Chapter | Status | Finished |
|---|---------|--------|----------|
| 01 | [Toolchain, types and errors](ch01-toolchain-types-errors/) | ✅ done | 2026-11-10 |
| 02 | Control flow, slices and maps | ⬜ | |
| 03 | Structs, methods and embedding | ⬜ | |
| 04 | Interfaces and errors in depth | ⬜ | |
| 05 | Generics | ⬜ | |
| 06 | Concurrency and context | ⬜ | |
| 07 | Testing and tooling | ⬜ | |
| 08 | Building an HTTP service | ⬜ | |

## Running the code

Everything lives in one Go module. From the repo root:

```sh
go run ./ch01-toolchain-types-errors/practice
go vet ./...
go test ./...
```

## Each chapter has

- one small program per idea, each in its own folder with a `main.go` (all written by me)
- a short `README.md` index listing each program and the concepts it practises

## Resources I'm using

- [A Tour of Go](https://go.dev/tour/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Exercism Go track](https://exercism.org/tracks/go)
