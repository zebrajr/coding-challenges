package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var puzzle_file_name = "puzzle_input.txt"
type level [][] string
type guard_info struct {
    col int
    row int
    direction int
}
var infinite_loop_count = 0

//              { -1, 0 }
//  { 0, -1}                { 0, +1 }
//              { +1, 0 }
var directions = [][]int {
    {-1, 0},
    {0, 1},
    {1, 0},
    {0, -1},
}

func load_puzzle_from_file(target_file_name string) (level, guard_info){
    file, _ := os.Open(target_file_name)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var new_level level
    var new_guard guard_info

    line_counter := 0
    for scanner.Scan(){
        line := scanner.Text()
        var line_array []string
        for idx, individual_char := range(line){
            stringified := string(individual_char)

            if stringified == "."{
                line_array = append(line_array, stringified)
            }

            if stringified == "^"{
                new_guard.col = idx
                new_guard.row = line_counter
                new_guard.direction = 0
                line_array = append(line_array, stringified)
            }

            if stringified == "#" {
                line_array = append(line_array, stringified)
            }
        }
        new_level = append(new_level, line_array)
        line_counter++ 
    }
    return new_level, new_guard
}


func solve_movement(test_level level, test_guard guard_info) (int){
    // create an indetic sized map to track places that were visited
    var visited [][]int 
    for range test_level{
        var row []int
        for range test_level[0] {
            row = append(row, 0)
        }
        visited = append(visited, row)
    }

    // mark starting position as visited
    visited[test_guard.row][test_guard.col] = 1
    left_map := false
    max_visits_assume_infinite_loop := 10
    max_visits_seen := 0

    for {
        if left_map {
            break
        }

        next_row := test_guard.row + directions[test_guard.direction][0]
        next_col := test_guard.col + directions[test_guard.direction][1]
        // check for out of bounds
        is_row_negative := next_row < 0
        is_col_negative := next_col < 0
        is_row_oob := next_row > (len(test_level) -1)
        is_col_oob := next_col > (len(test_level[0]) -1)
        if is_col_negative || is_row_negative || is_row_oob || is_col_oob {
            left_map = true
            continue
        }

        if test_level[next_row][next_col] == "#"{
            test_guard.direction++
            if test_guard.direction > len(directions) -1 {
                test_guard.direction = 0
            }
        }

        if test_level[next_row][next_col] != "#" {
            test_guard.row = next_row
            test_guard.col = next_col
            visited[next_row][next_col]++
            if visited[next_row][next_col] > max_visits_seen{
                max_visits_seen++
            }
        }

        if max_visits_seen > max_visits_assume_infinite_loop {
            infinite_loop_count++
            break
        }
    }

    places_visited := 0 
    for row := range visited{
        for col := range visited[row]{
            if visited[row][col] >= 1 {
                places_visited++
            }
        }
    }

    return places_visited
}

func time_track(start time.Time, name string){
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func clone_level(origin level)(level){
    new_level := make(level, len(origin))
    for i := range origin{
        new_level[i] = make([]string, len(origin[i]))
        copy(new_level[i], origin[i])
    }
    return new_level
}

func test_for_infinite_loop(base_level level, test_guard guard_info){
    for row := range base_level {
        for col := range base_level[row] {
            if test_guard.row == row && test_guard.col == col {
                continue
            }

            if base_level[row][col] == "#"{
                continue
            }

            isolated_level := clone_level(base_level)
            isolated_level[row][col] = "#"

            solve_movement(isolated_level, test_guard)
        }
    }
    fmt.Println("Total Infinite Loops Found: ", infinite_loop_count)
}

func main(){
    defer time_track(time.Now(), "main")
    level, guard := load_puzzle_from_file(puzzle_file_name)
    places_visited := solve_movement(level, guard)
    fmt.Println("Total Places Visited: ", places_visited)
    test_for_infinite_loop(level, guard)
}
