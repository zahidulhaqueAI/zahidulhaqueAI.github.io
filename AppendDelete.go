// link: https://www.hackerrank.com/challenges/append-and-delete/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
    lens := len(s)
    lent := len(t)
    var count int
    count = 0
    // loop to find the common letters
    for i := float64(0); i<math.Min(float64(lens),float64(lent));i++{
        if s[int(i)] == t[int(i)]{
            count++
        } else {
            break
        }
    }
    diff_s := lens - count
    diff_t := lent - count

    if int32(diff_s + diff_t) > k {
        return "No"
    }
    return "Yes"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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
