package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    var changes []int

    file, err := os.Open("input1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        change, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }
        changes = append(changes, change)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    part1(changes)
    part2(changes)
}

func part1(changes []int) {
    freq := 0
    for _, change := range changes {
        freq += change
    }
    fmt.Println("part 1: ", freq)
}

func part2(changes []int) {
    frequencies := []int{0}
    change_index := 0
    freq := 0
    found := false

    for found == false {
        freq += changes[change_index]
        found = contains(frequencies, freq)
        if (found == false) {
            if (change_index < len(changes) - 1)  {
            change_index++
            } else {
                change_index = 0
            }
        }
        frequencies = append(frequencies, freq)
    }
    fmt.Println("part 2: ", freq)
}

// Contains tells whether a contains x.
func contains(a []int, x int) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}