package main

import (
    "gameoflife/render"
)

func main() {
    // Update temperature

    render.Init()

    // Run at some time interval until the program exits
    render.Update()
    render.Render()
}
