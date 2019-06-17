// Draws a decimal spiral of size defined by the user. O(n^2)
// By Sabine Lim
// 17/06/2019

// The implementation consists of 2 parts:
// 1. Loop left to right, up to down, calling an isDigit function to
// know whether to print a digit or a dash for each coordinate
// 2. If a digit should be printed, call a getDigit function to
// calculate the digit to be printed

// The isDigit function is the same the code for the regular spiral, but
// instead of printing an asterisk it returns 1 to whatever code called it

// getDigit works by splitting the spiral into triangular quadrants
//
//   *******
// *  *****  *
// **  ***  **
// ***  *  ***
// **  ***  **
// *  *****  *
//   *******
//
// For the top quadrant, observe the following digits
//
// 6**************
// --------------*
// **0**********-*
// *-----------*-*
// *-**0******-*-*
// *-*-------*-*-*
// *-*-**6**-*-*-*
// *-*-*---*-*-*-*
// *-*-*-***-*-*-*
// *-*-*-----*-*-*
// *-*-*******-*-*
// *-*---------*-*
// *-***********-*
// *-------------*
// ***************
//
// The actual numbers at the location of these digits form a quadratic sequence
// 6  30  70  126
//  24  40  56
//    16  16
// Where the differences between the differences is 16.
// Using this, you can make a quadratic equation for the top left corners
// of each box, and subtract the current column to find the digits for
// coordinates to the right
//
// You'll need to come up with a different quadratic equation for each quadrant
// Since there are 2 types of spiral (last digit in the centre, last digit
// off-centre), that's a total of 8 different quadratic equations
// Use the current column or current row accordingly to determine how much
// to add or subtract from the corner values

package main

import "fmt"

func main() {
    size := 0
    fmt.Printf("Enter size: ")
    fmt.Scanf("%d", &size)
    row := 0
    for row < size {
        col := 0
        for col < size {
            if isDigit(size, row, col) {
                if (row <= size / 2 && col < size / 2 && row == col + 1) {
                    // Special handling for box segments modified to be spirals
                    fmt.Printf("%d", getDigit(size, row - 1, col - 1) % 10)
                } else {
                    fmt.Printf("%d", getDigit(size, row, col) % 10)
                }
            } else {
                fmt.Printf("-")
            }
            col++
        }
        fmt.Printf("\n")
        row++
    }
}

// Returns absolute value of an integer
func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

// Returns false if current coordinate is a dash, true if it's a digit
func isDigit(size int, row int, col int) bool {
    // Absolute row distance from midpoint
    rowDist := abs(row - size / 2)
    // Absolute column distance from midpoint
    colDist := abs(col - size / 2)
    isDigit := false
    if size % 4 == 1 {
        // Type 1 spiral (digit in centre)
        if row <= size / 2 && col < size / 2 && row == col + 1 {
            // Special handling to turn boxes into spirals
            if rowDist % 2 == 0 {
                isDigit = true
            }
        } else if colDist >= rowDist && colDist % 2 == 0 { 
            isDigit = true
        } else if colDist < rowDist && rowDist % 2 == 0 {
            isDigit = true
        }
    } else {
        // Type 2 spiral (no digit in centre)
        if row <= size / 2 && col < size / 2 && row == col + 1 {
            // Special handling to turn boxes into spirals
            if rowDist % 2 == 1 {
                isDigit = true
            }
        } else if colDist >= rowDist && colDist % 2 == 1 {
            isDigit = true
        } else if colDist < rowDist && rowDist % 2 == 1 {
            isDigit = true
        }
    }
    return isDigit
}

// Returns the integer at a specific coordinate on a box for a given size
func getDigit(size int, row int, col int) int {
    // Row displacement from midpoint
    rowDist := row - size / 2
    // Absolute row distance from midpoint
    absRowDist := abs(rowDist)
    // Column displacement from midpoint
    colDist := col - size / 2
    // Absolute column distance from midpoint
    absColDist := abs(colDist)
    // Size of box current coordinate is on
    subSize := 0
    if absRowDist >= absColDist {
        subSize = 2 * absRowDist + 1
    } else {
        subSize = 2 * absColDist + 1
    }
    row = row - (size - subSize) / 2
    col = col - (size - subSize) / 2
    // Layer of current box. 0 is centre
    layer := (subSize + 1) / 4
    if rowDist <= 0 && absRowDist >= absColDist {
        // Top quadrant
        if subSize % 4 == 1 {
            // Type 1 boxes
            return 8 * layer * layer + 8 * layer - col
        } else {
            // Type 2 boxes
            return 8 * layer * layer - 2 - col
        }
    } else if colDist > 0 && absColDist > absRowDist {
        // Right quadrant
        if subSize % 4 == 1 {
            // Type 1 boxes
            return 8 * layer * layer + 4 * layer - row
        } else {
            // Type 2 boxes
            return 8 * layer * layer - 4 * layer - row
        }
    } else if rowDist > 0 && absRowDist >= absColDist {
        // Bottom quadrant
        if subSize % 4 == 1 {
            // Type 1 boxes
            return 8 * layer * layer - 4 * layer + col
        } else {
            // Type 2 boxes
            return 8 * layer * layer - 12 * layer + 4 + col
        }
    } else {
        // Left quadrant
        if subSize % 4 == 1 {
            // Type 1 boxes
            return 8 * layer * layer - 8 * layer + row
        } else {
            // Type 2 boxes
            return 8 * layer * layer - 16 * layer + 6 + row
        }
    }
}