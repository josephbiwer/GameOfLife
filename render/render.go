package render

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)


var steamRows = [14]int{26, 26, 22, 22, 18, 18, 14, 14, 10, 10, 6, 6, 2, 2}
var steam [196]byte
var steamChars = []byte {'/', '|', '\\', ' '}

func Init() {
    rand.Seed(time.Now().UnixNano())
    charLen := len(steamChars)

    // Initialize steam to random character
    for c := range steam {
        steam[c] = steamChars[rand.Intn(charLen)]
    }
}

func Update() {

    for i := 0; i < len(steamRows) - 1; i++ {

    }

}

func getCharacter(index int) byte {
    return ' '
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
