package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    renderer.test()

    // Clear the screen
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

    fmt.Println("\033[20;40HThis is a test again")
}
