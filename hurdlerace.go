// Link: https://www.hackerrank.com/challenges/the-hurdle-race/problem

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

// Complete the hurdleRace function below.
func hurdleRace(k int32, height []int32) int32 {
    max := height[0]
    n := len(height)
    for i := 1; i < n-1; i++ {
        if float64(max) < math.Max( float64(max), float64(height[i])) {
            max = height[i]
        }
    }
    if max > k {
        return max - k
    }
    return 0
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    heightTemp := strings.Split(readLine(reader), " ")

    var height []int32

    for i := 0; i < int(n); i++ {
        heightItemTemp, err := strconv.ParseInt(heightTemp[i], 10, 64)
        checkError(err)
        heightItem := int32(heightItemTemp)
        height = append(height, heightItem)
    }

    result := hurdleRace(k, height)

    fmt.Fprintf(writer, "%d\n", result)

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
