package main

import (
	"bufio"
	"fmt"
	"os"
)

var file_puzzle_input = "puzzle_input.txt"
var word_to_find_xmas = "XMAS"
var total_words_found = 0
var total_x_mas_found = 0
var directions_for_xmas = [8][2]int {
        {-1, -1}, // Top-left
        {-1, 0},  // Top
        {-1, 1},  // Top-right
        {0, 1},   // Right
        {1, 1},   // Bottom-right
        {1, 0},   // Bottom
        {1, -1},  // Bottom-left
        {0, -1},  // Left
}

var directions_for_x_mas = [4][2]int {
        {-1, -1}, // Top-left
        {-1, 1},  // Top-right
        {1, 1},   // Bottom-right
        {1, -1},  // Bottom-left
}

func read_file_into_matrix(file_puzzle_input string) [][]string {
    puzzle_file, err := os.Open(file_puzzle_input)
    if err != nil {
        // this should be properly handled
    }
    defer puzzle_file.Close()

    var lines []string
    line_count := 0
    line_length := 0
    file_scanner := bufio.NewScanner(puzzle_file)

    for file_scanner.Scan(){
        line := file_scanner.Text()
        lines = append(lines, line)
        line_length = len(line)
        line_count++
    }

    matrix := make([][]string, line_count)
    for idx := 0; idx < line_count; idx++ {
        matrix[idx] = make([]string, line_length)
    }

    for line_idx, line := range lines{
        for char_idx, char := range line{
            matrix[line_idx][char_idx] = string(char) 
        }
    }

    return matrix 
}

func is_word_available(matrix [][]string, starting_row int, starting_col int, word_array []string){
    // directions 
    //  {-1, -1}    {-1, 0}     {-1, 1}
    //  {0, -1}                 {0, 1}
    //  {1, -1}     {1, 0}      {1, 1}
    directions := directions_for_xmas
    
    for _, direction := range directions{
        delta_row := direction[0]
        delta_col := direction[1]

        for step, letter_to_find := range word_array{
            next_row := starting_row + (step * delta_row)
            next_col := starting_col + (step * delta_col)

            // are we within bounds
            if next_row < 0{
                break
            }
            if next_col < 0{
                break
            }
            if next_row > (len(matrix) -1){
                break
            }
            if next_col > (len(matrix[next_row]) -1){
                break
            }

            is_letter_correct := matrix[next_row][next_col] == letter_to_find
            if !is_letter_correct{
                break
            }

            if step >= (len(word_array)-1){
                total_words_found++
            }
        }
    }
}

func check_matrix_for_word(matrix [][]string, word_to_find_xmas string){
    var word_array = make([]string, len(word_to_find_xmas))

    for idx, letter := range word_to_find_xmas{
        word_array[idx] = string(letter)
    }

    for row := range matrix {
        for col, value := range matrix[row]{
            if value != word_array[0]{
                continue
            }

            is_word_available(matrix, row, col, word_array)
        }
    }
}

func check_matrix_for_x_mas(matrix [][]string){
    trigger_letter := string('A')

    for idx_row := range matrix {
        for idx_col, letter := range matrix[idx_row]{
            if letter != trigger_letter{
                continue
            }
            is_x_mas_possible(matrix, idx_row, idx_col)
        }
    }
}

func is_x_mas_possible(matrix [][]string, starting_row int, starting_col int){
    letter_pairs := [2]string {string('M'), string('S')}

    // check for out of bounds
    if starting_row -1 < 0{ return }
    if starting_col -1 < 0{ return }
    if starting_row +1 > len(matrix)-1 { return }
    if starting_col +1 > len(matrix[starting_row])-1 { return }

    top_left := matrix[starting_row -1][starting_col -1]
    top_right:= matrix[starting_row -1][starting_col +1]
    bottom_right:= matrix[starting_row +1][starting_col +1]
    bottom_left := matrix[starting_row +1][starting_col -1]

    passed_left_to_right := false
    passed_right_to_left := false

    if (top_left == letter_pairs[0] && bottom_right == letter_pairs[1]) || (top_left == letter_pairs[1] && bottom_right == letter_pairs[0]){
        passed_left_to_right = true
    }

    if (top_right == letter_pairs[0] && bottom_left == letter_pairs[1]) || (top_right == letter_pairs[1] && bottom_left == letter_pairs[0]){
        passed_right_to_left = true
    }

    if passed_left_to_right && passed_right_to_left {
        total_x_mas_found++
    }
}

func print_matrix(matrix [][]string){
    rows := len(matrix) -1 
    cols := len(matrix[rows]) -1

    fmt.Printf("  ")
    for col := 0; col <= cols; col++{
        fmt.Printf(" %d", col)
    }
    fmt.Println()
    for row := 0; row <= rows; row++{
        fmt.Printf("%d ", row)
        for col := 0; col <= cols; col++{
            fmt.Printf(" %s", matrix[row][col])
        }
        fmt.Println()
    }
    fmt.Println()
}

func main(){
    new_matrix := read_file_into_matrix(file_puzzle_input)
    check_matrix_for_word(new_matrix, word_to_find_xmas)
    fmt.Println("Total Words Found: ", total_words_found)

    check_matrix_for_x_mas(new_matrix)
    fmt.Println("Total X-MAS Found: ", total_x_mas_found)
 }

