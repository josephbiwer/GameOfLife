package main

import (
    "gameoflife/render"
    "time"
    "flag"
)

func main() {

    // Handling command line arguments
    interval := flag.Int("interval", 333, "Provide an interval in milliseconds to update the animation")
    flag.Parse()

    render.Init()

    for {
        // Run at some time interval until the program exits
        render.Update()
        render.Render()

        time.Sleep(time.Duration(*interval) * time.Millisecond)
    }
}
