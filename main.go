package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func getRow(indicator string) int {
    min := 0
    max := 127

    for _, chr := range indicator {
        switch chr {
        case 'F':
            max = (min + max) / 2

        case 'B':
            min = ((min + max) / 2) + 1
        }
    }

    return min
}

func getCol(indicator string) int {
    min := 0
    max := 7

    for _, chr := range indicator {
        switch chr {
        case 'L':
            max = (min + max) / 2

        case 'R':
            min = ((min + max) / 2) + 1
        }
    }

    return min
}

func calcId(row int, col int) int {
    return row*8 + col
}

func main() {
    fc, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Println(err)
        return
    }

    passes := strings.Split(string(fc), "\n")
    ids := make(map[int]int)
    highest := -1

    seats := [128][8]int{}
    for i := 0; i < len(seats); i++ {
        for j := 0; j < len(seats[i]); j++ {
            seats[i][j] = 1
        }
    }

    for _, pass := range passes {
        ri := pass[:7]
        ci := pass[7:]

        r := getRow(ri)
        c := getCol(ci)
        id := calcId(r, c)

        // Indicate the seat is taken
        seats[r][c] = 0

        ids[id] = 1

        if id > highest {
            highest = id
        }
    }

    fmt.Println(highest)

    // Part 2
    for r := 1; r < len(seats)-1; r++ {
        for c := 0; c < len(seats[r]); c++ {
            if seats[r][c] == 1 {
                id := calcId(r, c)

                _, ok := ids[id-1]
                if ok {
                    _, ok = ids[id+1]

                    if ok {
                        fmt.Println(id)
                        return
                    }
                }
            }
        }
    }
}
