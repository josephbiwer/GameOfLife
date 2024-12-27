package main

import (
    "gameoflife/render"
)

func main() {
    // Update temperature

    render.Init()
    render.Update()
    render.Render()
}
