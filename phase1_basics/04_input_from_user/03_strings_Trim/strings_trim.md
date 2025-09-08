# strings.TrimSpace() removes:
- Leading spaces (before text)
- Trailing spaces (after text)
- Newline \n and carriage return \r
- But it does not remove spaces inside the text.

1️⃣ strings.TrimSpace(s string) string

- Removes all leading & trailing whitespace, including spaces, tabs, newlines.
- Does not remove internal spaces.

```go
	s := "\t  Hello World  \n"
	fmt.Printf("[%s]\n", strings.TrimSpace(s))
```

2️⃣ strings.Trim(s string, cutset string) string
- Removes all leading & trailing characters that appear in cutset.

Example: remove specific characters instead of whitespace.

```go
s := "!!Hello World!!"
fmt.Printf("[%s]\n", strings.Trim(s, "!"))
```

3️⃣ strings.TrimLeft(s string, cutset string) string
- Removes characters in cutset only from the beginning.

```go
s := "!!Hello World!!"
fmt.Printf("[%s]\n", strings.TrimLeft(s, "!"))
```

4️⃣ strings.TrimRight(s string, cutset string) string
- Removes characters in cutset only from the end.

```go
s := "!!Hello World!!"
fmt.Printf("[%s]\n", strings.TrimRight(s, "!"))
```
