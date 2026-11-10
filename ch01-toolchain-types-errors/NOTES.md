# Chapter 01 notes: Toolchain, types and errors

What I learned in chapter 1, coming from Java/Python/JS. The code for this chapter is in [`practice/main.go`](practice/main.go).

> [!TIP]
> **Chapter 1 in six lines**
> - One `go` binary is the build tool, package manager, formatter and test runner.
> - A package is a **folder**. A capitalised name is exported. Unused imports and variables are compile errors.
> - Every variable has a **zero value**. Strings are `""`, never null. A nil map panics when you write to it.
> - There are no implicit numeric conversions.
> - `:=` only works inside functions. Use `var` at package level.
> - Failures are **returned values** (`error`), not exceptions.

---

## 1. Toolchain and modules

| Task | Go | What I already knew |
|---|---|---|
| Start a project | `go mod init example.com/hello` | `npm init`, `pom.xml` |
| Run | `go run .` (the space matters) | `python main.py` |
| Build a binary | `go build` | `mvn package` |
| Add a dependency | `go get github.com/x/y` | `npm install`, `pip install` |
| Clean up deps | `go mod tidy` | nothing really |
| Format | `go fmt ./...` | Prettier / Black, except here it's mandatory |
| Test | `go test ./...` | JUnit / pytest |

- **`go.mod`** is the shopping list (like `package.json`): module path, Go version, `require` lines.
- **`go.sum`** is the sealed receipt: a checksum for every module downloaded. Go refuses to build if one doesn't match. It only appears after you add your first dependency.
- You never edit either file by hand (`go get` and `go mod tidy` do it), and you commit both.
- No venv and no `node_modules`. Dependencies live in one shared module cache.
- The output is a single static binary, so the target machine doesn't need a runtime.

## 2. Packages and `main`

```
hello/
├── go.mod
├── main.go        ← package main
└── greet/         ← package greet (the folder is the package)
    ├── hello.go   ← package greet
    └── bye.go     ← package greet
```

- `package main` plus `func main()` makes an executable. Any other package name makes a library.
- Files in the same folder share one namespace. They can call each other, even lowercase names, without importing anything. To use another folder, import it: `import "example.com/hello/greet"`, then call `greet.Hello()`.
- **A capital letter means public.** `ToUpper` can be called from outside `strings`; `toUpper` couldn't be. Privacy is per **package**, not per class, which is different from Java.
- An unused import fails the build. Go is strict about this on purpose.

## 3. Types and zero values

The everyday types are `int` (like Java's `long`), `float64` (`double`), `string` (immutable) and `bool`. There's also `byte` (= `uint8`) and `rune` (= `int32`, one Unicode code point).

| Type | Zero value |
|---|---|
| numbers | `0` |
| `bool` | `false` |
| `string` | `""` |
| pointer, slice, map, chan, func, interface | `nil` |
| struct | each field gets its own zero value |

- **A variable always has a value.** A Java `String` field starts as `null` and throws an NPE later. In Go, `len(name)` on a string you never set is just `0`.
- **Why a Go string can never be null:** a Java `String` is an arrow that can point at nothing. A Go `string` is a box that always holds text. If you forget to set it you get `""`, and `var s string = nil` won't compile.
- **`nil` is for handles.** A map variable is like a locker key, and `var m map[string]int` is a key with no locker assigned yet:
  ```go
  var scores map[string]int
  fmt.Println(scores["alice"])   // 0: reading a nil map is safe
  scores["alice"] = 10           // panic: assignment to entry in nil map
  scores = make(map[string]int)  // fix: allocate it first
  ```
  In Java, a null `Map` throws on both get and put. In Go only the write panics.
- **No implicit numeric conversion.** `int + float64` doesn't compile, and neither does `int32 + int64`. You have to write `float64(count) + avg`. Go makes you say it so nothing overflows or loses precision silently.
- Idiom: design types so their zero value is ready to use, the way `sync.Mutex` and `bytes.Buffer` are.

## 4. `var` vs `:=`

```go
var count int          // zero value 0
var name = "gopher"    // type inferred
limit := 10            // short declaration, only inside functions
```

- **`:=` creates a variable and `=` updates one.** Think of the colon as meaning "new".
- `:=` never takes a type. If you want to write the type, use `var timeout int = 30`.
- `:=` doesn't work at package level, because every top-level line has to start with a keyword (`var`, `func`, `const`, `import`).
- **The roll-call rule:** with several names on the left, `:=` needs at least one of them to be new.
  ```go
  a, b := 1, 2  // both new ✅
  c, b := 3, 4  // c is new, b gets updated ✅
  a, b := 5, 6  // nothing new ❌ "no new variables on left side of :="
  a, b = 5, 6   // fix: plain =
  ```
  This is what lets you write `result, err := ...` over and over and keep reusing `err`.

> [!WARNING]
> **Shadowing.** `:=` inside `{ }` creates a *second* variable with the same name, and it only lives inside that block:
> ```go
> status := "pending"
> if true {
>     status := "done"    // a NEW inner status
>     fmt.Println(status) // done
> }
> fmt.Println(status)     // pending: the outer one never changed
> ```
> The fix is `status = "done"` with no colon. Java forbids this kind of redeclaration. Go allows it and compiles it without complaint, which is exactly why it causes bugs.

## 5. Multiple returns and errors as values

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
	fmt.Println("failed:", err)
	return
}
```

- A function can return several values: `q, r := divmod(17, 5)`.
- `error` is just an interface with one method, `Error() string`. `nil` means success.
- **Java vs Go:** Java hands you the goods and may throw the receipt at you later. Go hands you both at once and makes you check the receipt right away.
- **Why Go doesn't "ignore errors":** (a) an unused variable won't compile, so you must either use `err` or throw it away visibly with `_`; (b) `error` is part of the function signature, so every caller can see the failure path.
- **`_` throws a value away.** Writing `n, _ := strconv.Atoi("abc")` is a visible red flag.
- **Wrap errors to add context:** `fmt.Errorf("doubleIt: %w", err)` gives `doubleIt: strconv.Atoi: parsing "oops": invalid syntax`. You get a trail of breadcrumbs instead of a stack trace, and `%w` keeps the original error so it can be inspected later with `errors.Is` / `errors.As`.
- **`panic`** is for bugs and impossible states (like an index out of range). For expected failures, return an `error`. It's roughly Java's `AssertionError` compared with a checked exception.
- Idiom: return early on errors so the happy path stays unindented. Avoid `else` where you can.

---

## Mistakes I made (and the fixes)

| I wrote | Problem | Fix |
|---|---|---|
| `float54(x)` | typo | `float64(x)` |
| `timeout int := 30` | `:=` can't take a type | `var timeout int = 30` or `var timeout = 30` |
| `import { "fmt", "errors" }` | wrong brackets, and commas | `import (` then one package per line `)` |
| `func divide(a, b float64) (result, error)` | the return list needs types | `(float64, error)` |
| `}` / `else` / `{` on separate lines | Go inserts `;` at line ends, so `else` ends up orphaned | `} else {` on one line |
| `return n*2` in a function returning `(int, error)` | missing the second value | `return n * 2, nil` |
| `a, b := 5, 6` after both already existed | no new variables on the left | `a, b = 5, 6` |
| `go run.` | unknown command | `go run .` |

Reading compiler errors: `./main.go:6:20` means file, line, column.

## Quiz yourself

<details><summary>A package maps to a … ?</summary>A directory, not a file.</details>
<details><summary>How do you make an identifier public?</summary>Start it with a capital letter. Privacy is per package.</details>
<details><summary>What happens when you import a package and never use it?</summary>A compile error, not a warning. The same goes for unused local variables.</details>
<details><summary>What's the zero value of a string?</summary><code>""</code>. A string is never nil.</details>
<details><summary>Reading vs writing a nil map?</summary>Reading returns the zero value. Writing panics.</details>
<details><summary>Why doesn't <code>var y float64 = x</code> compile when x is an int?</summary>Go has no implicit numeric conversions. Write <code>float64(x)</code>.</details>
<details><summary>Where is <code>:=</code> not allowed?</summary>At package level. Use <code>var</code> there.</details>
<details><summary>Why does <code>a, err := f()</code> followed by <code>b, err := g()</code> compile?</summary><code>b</code> is new, so <code>err</code> is just reassigned.</details>
<details><summary>How does <code>:=</code> inside an if-block cause a bug?</summary>It shadows the outer variable, so the outer one keeps its old value after the block.</details>
<details><summary>What does <code>%w</code> do?</summary>It wraps the original error so <code>errors.Is</code> / <code>errors.As</code> can still find it.</details>
<details><summary>When should you panic?</summary>Only for programmer bugs or impossible states, never for expected failures.</details>
<details><summary>go.mod vs go.sum?</summary>go.mod lists the module, the Go version and the deps. go.sum holds checksums. Commit both.</details>

## Further reading
- [Tutorial: Get started](https://go.dev/doc/tutorial/getting-started) · [Create a module](https://go.dev/doc/tutorial/create-module)
- [Effective Go: Multiple returns](https://go.dev/doc/effective_go#multiple-returns)
- [Errors are values (Rob Pike)](https://go.dev/blog/errors-are-values) · [Error handling and Go](https://go.dev/blog/error-handling-and-go)
