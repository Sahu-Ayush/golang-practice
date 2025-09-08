# Golang I/O (Input/Output) from User using both

1. fmt.Scanln (basic, easy, but limited)
2. bufio (powerful, flexible, recommended for more complex inputs).
    2.1. bufio.Reader
    2.2. bufio.Scanner

## 1. Using fmt.Scanln:
- fmt.Scanln is part of Go’s fmt package, used to read input from standard input (stdin).
- Reads user input until "first space" or "newline" i.e, ("Enter" key).
- Automatically splits input into variables.
- Limited: not great for spaces in strings.

✅ Basics

```go
fmt.Scanln(&a, &b, &c)
```
- Reads space-separated values into the provided variables.
- Stops reading at newline (\n).
- Returns the number of items successfully scanned and an error.

⚡ Key Behaviors

1. Reading stops at whitespace (space or newline)

    ```go
    var name string
    fmt.Scanln(&name)
    // Input: "Ayush Sahu"
    // name = "Ayush", "Sahu" remains in buffer
    ```

2. Leaves leftovers in the buffer
 - If the next Scanln reads a different type (like int), it may fail because "Sahu" isn’t an integer.

3. Newline consumption
 - After hitting Enter, \n marks the end of input.
 - If variables are fewer than words typed, extra words remain in the buffer.

4. Multiple variables

    ```go
    var first, last string
    fmt.Scanln(&first, &last)
    // Input: "Ayush Sahu"
    // first = "Ayush", last = "Sahu"
    ```
5. Error handling

    ```go
    n, err := fmt.Scanln(&age)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Scanned items:", n)
    }
    ```

🧨 Edge Cases

1. Extra tokens in input

    ```go
    var x int
    fmt.Scanln(&x)
    // Input: "10 20"
    // x = 10, "20" stays in buffer → may cause issues later
    ```

2. Invalid type

    ```go
    var age int
    fmt.Scanln(&age)
    // Input: "Ayush"
    // error: expected int, got string → age = 0
    ```


3. Empty input

    ```go
    var name string
    fmt.Scanln(&name)
    // Just pressing Enter → no scan, error: unexpected newline
    ```

4. Trailing newline issue with multiple Scanln calls

    ```go
    var name string
    var age int
    fmt.Scanln(&name) // Input: "Ayush Sahu"
    fmt.Scanln(&age)  // age = 0, leftover "Sahu" prevents reading
    ```

💡 Tips & Workarounds

- Single word input only → Scanln is fine.
- Multi-word strings (like full names) → use:
    - fmt.Scanf("%[^\n]s", &name)
    - or bufio.Reader with ReadString('\n').
- Always check errors from Scanln to detect invalid input.
- To clear leftovers, read the rest of the line manually if needed.

✅ Rule of thumb

- Use Scanln when you expect fixed, space-separated single tokens.
- Use Scanf or bufio.Reader when you need more control (multi-word strings, strict formats).


## 2.1. Using bufio.Reader

- Reads the entire line (with spaces), means Reads input until the specified delimiter (often \n for Enter key).
- Returns a string including the delimiter.
- Can read very long lines, because buffer size is internally managed.
- Works with os.Stdin.
- Usually paired with strings.TrimSpace to remove \n.


## 2.2. Using bufio.Scanner

- Scanner reads input token by token (default: line by line).
- More control than Scanln, less than Reader.
- scanner.Scan() advances to next token.
- scanner.Text() returns the current token as string.
- Automatically ignores the newline in Text().
- Simpler when reading line-by-line input interactively.
- Maximum line length: Default 64 KB

Notes: Quick Rule of Thumb
- Use Scanner for simple interactive line input.
- Use Reader.ReadString('\n') when reading large input or files and need exact control.