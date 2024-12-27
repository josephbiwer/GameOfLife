package render

import (
    "fmt"
    "strconv"
)

// Steam architecture
/*
2
6
10
14
18
22
26
*/

var steam [96]byte

func Init() {
    for c := range steam {
        steam[c] = ' '
    }
}

func Update() {

}

func getCharacter(index int) byte {
    
    return ' '
}

func Render() {
    xOff := 103 
    yOff := 20

    fmt.Println("\033[2J") // Clear screen

    // Draw mug steam
    steamXOff := 0
    // steamYOff := 6 
    for y := yOff - 1; y > (yOff - 8); y-- {
        xOffset := steamXOff * 2
        fmt.Print("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(xOff + xOffset - 1) + "H|")
        for i := 0; i < (26 - 2 * xOffset); i++ {
            fmt.Print("x")
        }
        fmt.Print("|")
        steamXOff++;
    }


    // Draw Mug
    fmt.Println("\033[" + strconv.Itoa(yOff) + ";103H|------------------------|")
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
