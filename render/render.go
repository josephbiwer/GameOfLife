package render

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)


var steamRows = [14]int{26, 26, 22, 22, 18, 18, 14, 14, 10, 10, 6, 6, 2, 2}
var steam [196]byte
var steamChars = []byte {'/', '|', '\\', ' ', ' '} // Add duplicate characters if you want to increase the probability that they appear

func Init() {
    rand.Seed(time.Now().UnixNano())
    charLen := len(steamChars)

    // Initialize steam array to be empty characters
    for c := range steam {
        steam[c] = ' '
    }

    // Initialize steam to random character for the first row
    for c := 0; c < steamRows[0]; c++ {
        steam[c] = steamChars[rand.Intn(charLen)]
    }
}

func Update() {
    // Calculate index offset
    steamOff := 0

    for i := 0; i < len(steamRows) - 1; i++ {
    
        // Don't need to start at the very beginning of the array if the next row is not the same length of characters
        if(steamRows[i] == steamRows[i + 1]) {
            steamOff = 1
        }
            
        for j := steamOff; j < steamRows[i] - steamOff; j++ {
            
            indexOff := 0
            for k := 0; k < i; k++ {
                indexOff += steamRows[k]
            }

            // indexRowAbove := indexOff + steamRows[i]
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

                case '/':
                    // If not in right range
                    if(!(j == (steamRows[i] - 1) || (steamOff == 1 && j == (steamRows[i] - 2)))) {
                        steam[indexAbove + 1] = '/'
                    }
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
    for x := 1; x < 11; x++ {
        str := "\033["
        str += strconv.Itoa(yOff + x)
        str += ";" + strconv.Itoa(xOff) + "H|                        |"
        fmt.Println(str)
    }

    fmt.Println("\033[" + strconv.Itoa(yOff + 11) + ";" + strconv.Itoa(xOff)     + "H\\                        /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 12) + ";" + strconv.Itoa(xOff + 1) + "H\\                      /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 13) + ";" + strconv.Itoa(xOff + 2) + "H\\                    /")
    fmt.Println("\033[" + strconv.Itoa(yOff + 14) + ";" + strconv.Itoa(xOff + 3) + "H\\------------------/")
}
