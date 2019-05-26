package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    data := map[string]int{}
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        l := scanner.Text()
        _, e := data[l]
        if !e {
            data[l] = 1
        } else {
            data[l]++
        }
    }

    for k, v := range data { 
        fmt.Printf("%7d %s\n", v, k)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
