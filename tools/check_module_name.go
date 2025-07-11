package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const expectedModule = "github.com/ayushsahu/go-practice" // 🔁 Change as needed

func main() {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("❌ Error opening go.mod:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			actualModule := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			if actualModule == expectedModule {
				fmt.Println("✅ Module name is correct:", actualModule)
			} else {
				fmt.Printf("⚠️  Mismatch!\nExpected: %s\nFound:    %s\n", expectedModule, actualModule)
			}
			return
		}
	}

	fmt.Println("❌ No module declaration found in go.mod")
}
