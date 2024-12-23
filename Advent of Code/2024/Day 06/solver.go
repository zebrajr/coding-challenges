package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var puzzle_file_name = "puzzle_input.txt"
var level [][] string
var guard_pos struct {
    col int
    row int
    direction int
}

//              { -1, 0 }
//  { 0, -1}                { 0, +1 }
//              { +1, 0 }
var directions = [][]int {
    {-1, 0},
    {0, 1},
    {1, 0},
    {0, -1},
}

func load_puzzle_from_file(target_file_name string){
    file, err := os.Open(target_file_name)
    if err != nil{
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

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
                guard_pos.col = idx
                guard_pos.row = line_counter
                guard_pos.direction = 0
                line_array = append(line_array, stringified)
            }

            if stringified == "#" {
                line_array = append(line_array, stringified)
            }
        }
        level = append(level, line_array)
        line_counter++ 
    }
}


func solve_movement(){
    fmt.Println("Map Size: ", len(level) -1, len(level[0]) -1)
    fmt.Println("Guard Starting Position: ", guard_pos.row, guard_pos.col, guard_pos.direction)

    // create an indetic sized map to track places that were visited
    var visited [][]int 
    for range level{
        var row []int
        for range level[0] {
            row = append(row, 0)
        }
        visited = append(visited, row)
    }

    // mark starting position as visited
    visited[guard_pos.row][guard_pos.col] = 1
    left_map := false

    for {
        if left_map {
            break
        }

        next_row := guard_pos.row + directions[guard_pos.direction][0]
        next_col := guard_pos.col + directions[guard_pos.direction][1]
        // check for out of bounds
        is_row_negative := next_row < 0
        is_col_negative := next_col < 0
        is_row_oob := next_row > (len(level) -1)
        is_col_oob := next_col > (len(level[0]) -1)
        if is_col_negative || is_row_negative || is_row_oob || is_col_oob {
            left_map = true
            continue
        }

        if level[next_row][next_col] == "#"{
            guard_pos.direction++
            if guard_pos.direction > len(directions) -1 {
                guard_pos.direction = 0
            }
        }

        if level[next_row][next_col] != "#" {
            guard_pos.row = next_row
            guard_pos.col = next_col
            visited[next_row][next_col] = 1
        }

        
    }

    places_visited := 0 
    for row := range visited{
        for col := range visited[row]{
            if visited[row][col] == 1 {
                places_visited++
            }
        }
    }
    fmt.Println("Visited Places: ", places_visited)
}

func time_track(start time.Time, name string){
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}


func main(){
    defer time_track(time.Now(), "main")
    load_puzzle_from_file(puzzle_file_name)
    solve_movement()
}
