package render

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)


var steamRows = [14]int{26, 26, 22, 22, 18, 18, 14, 14, 10, 10, 6, 6, 2, 2}
var steam [196]byte
var steamChars = []byte {'(', ')', ' ', ' ', ' ', ' ', ' '} // Add duplicate characters if you want to increase the probability that they appear

func Init() {
    rand.Seed(time.Now().UnixNano())

    // Initialize steam array to be empty characters
    for c := range steam {
        steam[c] = ' '
    }

}

func Update() {

    for c := 52; c < len(steam); c++ {
        steam[c] = ' '
    }

    copy(steam[26:52], steam[0:26])

    charLen := len(steamChars)
    // Initialize steam to random character for the first row
    for c := 0; c < steamRows[0]; c++ {
        steam[c] = steamChars[rand.Intn(charLen)]
    }

    for i := 1; i < len(steamRows) - 1; i++ {
        steamOff := 0
    
        // Don't need to start at the very beginning of the array if the next row is not the same length of characters
        if(steamRows[i] == steamRows[i + 1]) {
            steamOff = 1
        }
            
        for j := steamOff; j < steamRows[i] - steamOff; j++ {
            
            indexOff := 0
            for k := 0; k < i; k++ {
                indexOff += steamRows[k]
            }

            // Replace current character with something else before evaluating character above
            if rand.Intn(4) == 0 {
                steam[indexOff + j] = steamChars[rand.Intn(charLen)]
            }

            indexAbove := indexOff + j + steamRows[i]
            // Get index of cell above current value being checked
            if(steamRows[i] != steamRows[i + 1]) {
                // left range, middle range, right range
                // if j > left range and j < right range, there is an index right above
                indexAbove -= ((steamRows[i] - steamRows[i + 1]) / 2)
            }
            
            switch steam[indexOff + j] {
                case '|':
                    // If in middle range and not on the left/right range when the row above is decreased in width
                    if !(j == 1 && indexOff == 1) ||  
                        !((j == steamRows[i] - 1) && indexOff == 1) || 
                        (j >= 2 && j <= (steamRows[i] - 2) && indexOff == 0) {

                        steam[indexAbove] = '|'
                    }

                case '\\':
                    // If not in left range
                    if(!(j == steamOff || (steamOff == 1 && j == 2))) {
                        steam[indexAbove - 1] = '\\'
                    }

                case ')':
                    // If not in left range
                    if(!(j == steamOff || (steamOff == 1 && j == 2))) {
                        // steam[indexAbove - 1] = '('
                        steam[indexAbove] = '('
                    }

                case '/':
                    // If not in right range
                    if(!(j == (steamRows[i] - 1) || (steamOff == 1 && j == (steamRows[i] - 2)))) {
                        steam[indexAbove + 1] = '/'
                    }
                    
                case '(':
                    // If not in right range
                    if(!(j == (steamRows[i] - 1) || (steamOff == 1 && j == (steamRows[i] - 2)))) {
                        // steam[indexAbove + 1] = ')'
                        steam[indexAbove] = ')'
                    }
                
                default:
                    steam[indexAbove] = ' '

            }
        }
    }
}

func Render() {
    xOff := 103 
    yOff := 20

    fmt.Println("\033[2J") // Clear screen

    // Draw mug steam
    for k := 0; k < len(steamRows); k++ {

        y := yOff - 2 - k

        // Calculate index offset
        indexOff := 0
        for j := 0; j < k; j++ {
            indexOff += steamRows[j]
        }

        fmt.Print("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(xOff + (26 - steamRows[k]) / 2) + "H")
        for i := 0; i < 26 - (26 - steamRows[k] ); i++ {
            fmt.Printf("%c", steam[indexOff + i])
        }

    }

    // Draw Mug
    fmt.Println("\033[" + strconv.Itoa(yOff) + ";103H/------------------------\\")
    fmt.Println("\033[" + strconv.Itoa(yOff + 1) + ";102H|                          |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 2) + ";103H\\------------------------/")
    fmt.Println("\033[" + strconv.Itoa(yOff + 3) + ";103H|                        |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 4) + ";103H|                        | ______")
    fmt.Println("\033[" + strconv.Itoa(yOff + 5) + ";103H|                        |/      \\")
    fmt.Println("\033[" + strconv.Itoa(yOff + 6) + ";103H|                        |  (--)  \\")
    fmt.Println("\033[" + strconv.Itoa(yOff + 7) + ";103H|                        |  |  |  |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 8) + ";103H|                        |  )  (  |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 9) + ";103H|                        |  |  |  |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 10) + ";103H|                        |  (--)  |")
    fmt.Println("\033[" + strconv.Itoa(yOff + 11) + ";103H|                        |        /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 12) + ";103H|                        |\\______/")


    fmt.Println("\033[" + strconv.Itoa(yOff + 13) + ";" + strconv.Itoa(xOff)     + "H\\                        /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 14) + ";" + strconv.Itoa(xOff + 1) + "H\\                      /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 15) + ";" + strconv.Itoa(xOff + 2) + "H\\                    /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 16) + ";" + strconv.Itoa(xOff + 3) + "H\\------------------/")
}
