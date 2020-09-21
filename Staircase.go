// link
// https://www.hackerrank.com/challenges/staircase/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func staircase(n int32) {
    for i := int32(1); i <=n; i++{
        for j := int32(1); j <= n-i; j++{
            fmt.Print(" ")
        }
        for k := int32(1); k <=i; k++{
            fmt.Print("#")
        }
        fmt.Println()
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
