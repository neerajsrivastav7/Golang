package main

import (
    "errors"
    "fmt"
)

func openFile(filename string) error {
    return errors.New("file not found")
}

func readFile(filename string) error {
    err := openFile(filename)
    if err != nil {
        return fmt.Errorf("readFile: %w", err)
    }
    return nil
}

func main() {
    err := readFile("example.txt")
    if err != nil {
        fmt.Println(err)
    }
}
