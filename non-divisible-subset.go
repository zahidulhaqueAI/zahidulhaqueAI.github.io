
package main
// Link: https://www.hackerrank.com/challenges/non-divisible-subset/problem

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
    len := int32(len(s))
    // frequency array
    freq := make([] int32, k)
    for i :=int32(0); i < len; i++ {
        index := s[i]%k
        freq[index] += 1
    } 
    // result count
    res := int32(0)
    res = int32(math.Min(float64(freq[0]),float64(1)))  
    for i := int32(1); i <= k/2; i++ {
        if i != k-i{
            res += int32(math.Max(float64(freq[i]), float64(freq[k-i])))
        } else {
            res += int32(math.Min(float64(freq[i]),float64(1)))
        }
    }
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int32

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int32(sItemTemp)
        s = append(s, sItem)
    }

    result := nonDivisibleSubset(k, s)

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
