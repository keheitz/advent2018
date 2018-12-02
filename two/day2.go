package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    var boxIDs []string

    file, err := os.Open("input2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        box := scanner.Text()
        boxIDs = append(boxIDs, box)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    part1(boxIDs)   
    part2(boxIDs)
}

func part1(boxIDs []string) {
    repeat_char := 0
    triple_char := 0
    for _, box := range boxIDs {
        repeat, triple := false, false
        checked_letters := ""
        box_char:= strings.Split(box, "") 
        for _, letter := range box_char {
            if (strings.Count(checked_letters, letter) == 0) {
                checked_letters = checked_letters + string(letter)
                letter_count := strings.Count(box, letter)
                if(letter_count == 2 && repeat == false) {
                    repeat_char++
                    repeat = true
                } else if (letter_count == 3 && triple == false) {
                    triple_char++
                    triple = true
                }
            }
        } 
    }
    checksum := (repeat_char * triple_char)
    fmt.Println("part 1: ", checksum)
}

func part2(boxIDs []string) {
    //ugh, pretty sure this is not the most efficient way to do this...
    //I miss linq :(
    same_letters := ""
    for _, box := range boxIDs {
        box_letters := strings.Split(box, "")
        for _, other_box := range boxIDs {
            difference_count, difference_idx, idx := 0, 0, 0
            other_box_letters := strings.Split(other_box, "")
            for _, letter := range box_letters {
                same := (letter == other_box_letters[idx])
                if (same == false) {
                    difference_count++
                    difference_idx = idx
                }
                if(idx < len(box_letters) -1) { idx++ }
                if (difference_count > 1) {
                    break
                }
                if ((idx == len(box_letters) - 1) && (difference_count == 1)) {
                    box_letters[difference_idx] = ""
                    for _, remaining := range box_letters {
                        same_letters += string(remaining)
                    }
                    break
                }
            }
            if(same_letters != "") {
                break
            }
        }
        if(same_letters != "") {
            break
        }
    }
    fmt.Println("part 2: ", same_letters)
}
