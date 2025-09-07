#Go’s fmt package provides three related families of functions:

1. fmt.Scan

- Reads space-separated values from standard input.
- Keeps scanning until all arguments are filled.
- Stops at any whitespace (spaces, tabs, or newlines).

👉 Example:
```go
package main

import "fmt"

func main() {
    var a, b string
    fmt.Print("Enter two words: ")
    fmt.Scan(&a, &b) // space/tab/newline separated
    fmt.Println("a:", a, "b:", b)
}
```

Input:
hello world


Output:
a: hello b: world


Note: Even if you press Enter between them, it still works:

hello
world


2. fmt.Scanln

- Similar to Scan, but stops scanning at a newline.
- If you provide extra input after arguments on the same line, it causes an error (expected newline).

👉 Example:

```go
package main

import "fmt"

func main() {
    var a, b string
    fmt.Print("Enter two words: ")
    fmt.Scanln(&a, &b)
    fmt.Println("a:", a, "b:", b)
}
```

Input:
hello world

Output:
a: hello b: world


Input:
hello world extra

Output:
a: hello b: world
// error: expected newline


3. fmt.Scanf

- Lets you specify a format string (like Printf but for input).
- Very useful when you expect input in a specific pattern.

👉 Example:

```go
package main

import "fmt"

func main() {
    var a int
    var b string
    fmt.Print("Enter: ")
    fmt.Scanf("%d-%s", &a, &b) // expects input like 42-hello
    fmt.Println("a:", a, "b:", b)
}
```

Input:
42-hello

Output:
a: 42 b: hello

Note: If input doesn’t match the format, scanning will fail.


🔑 Summary Table
| Function | Stops At       | Format Control   | Extra Input Handling         |
| -------- | -------------- | ---------------- | ---------------------------- |
| `Scan`   | Whitespace     | No               | Reads next tokens regardless |
| `Scanln` | Newline        | No               | Fails if extra input on line |
| `Scanf`  | Depends on fmt | Yes (`%d`, `%s`) | Strictly follows format      |
