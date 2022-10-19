package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)


var tops_a []string = []string{"1", "Jorge", "181.79457213348914", "10-17-2022 07:29:00", "910067180706627594"}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

func main() {
    lines, err := readLines("foo.in.txt")
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    for i, line := range lines {
        fmt.Println(i, line)
    }

    if err := writeLines(tops_a, "foo.out.txt"); err != nil {
        log.Fatalf("writeLines: %s", err)
    }
}