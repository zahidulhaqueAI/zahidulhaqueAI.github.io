// Link https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

 func binarySearch(ranked []int32, low int32, high int32, score int32) int32 {
     mid := low + (high - low)/2
     if ranked[mid] == score {
         return mid
     } else if score > ranked[mid] && score < ranked[mid-1] {
         return mid
     } else if score > ranked[mid] && score < ranked[mid-1] {
         return mid-1
     }
     if score > ranked[mid]{
         return binarySearch(ranked,low, mid-1,score)
     } else {
         return binarySearch(ranked,mid+1,high,score)
     }
  //   return -1
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    // no of games
    n := len(player)
    m := len(ranked)
    // result of rank array
    result := make([]int32,n)
    // populate the rank based on the ranked array
    rank := make([]int32,m)
    // ranked array in ascending order,index 0 will always be 1
    rank[0] = 1
    for i := 1; i < m; i++{
        if ranked[i] == ranked[i-1]{
            // means previous and current index same rank.
            // assign same rank value
            rank[i] = rank[i-1]
        } else {
            // assign next rank lower to the previous rank
            rank[i] = rank[i-1]+1
        }
    }
    // our rank array is ready based on the already given score i.e. ranked array
    // start fetching the rank of the palyer now
    for i := range player{
        // check for the boundry condition
        score := player[i]
        if score >= ranked[0] {
            result[i] = 1
        } else if score < ranked[m-1]{
            result[i] = rank[m-1]+1
        }else {
            // look for other index
            index := binarySearch(ranked, 0, int32(m-1), score)
            result[i] = rank[index]
        }
    }
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ranked []int32

    for i := 0; i < int(rankedCount); i++ {
        rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
        checkError(err)
        rankedItem := int32(rankedItemTemp)
        ranked = append(ranked, rankedItem)
    }

    playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var player []int32

    for i := 0; i < int(playerCount); i++ {
        playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
        checkError(err)
        playerItem := int32(playerItemTemp)
        player = append(player, playerItem)
    }

    result := climbingLeaderboard(ranked, player)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
