package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Create a scanner to read from standard input
    scanner := bufio.NewScanner(os.Stdin)

    // A boolean variable to track whether we have skipped the headers
    headersSkipped := false

    // Read the input line by line
    for scanner.Scan() {
        line := scanner.Text()

        // Normalize line endings by replacing \r\n or \r with \n temporarily
        line = strings.ReplaceAll(line, "\r", "")

        // If we encounter an empty line, it marks the end of the headers
        if !headersSkipped && line == "" {
            headersSkipped = true
            continue // Skip the empty line itself
        }

        // Once headers are skipped, print the remaining lines (message body)
        if headersSkipped {
            // Normalize output to CRLF (\r\n)
            fmt.Print(line + "\r\n")
        }
    }

    // Error handling for the scanner
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
}