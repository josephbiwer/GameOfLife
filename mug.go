package main

import (
    "mugrender/render"
    "time"
    "flag"
)

func main() {

    // Handling command line arguments
    interval := flag.Int("interval", 333, "Provide an interval in milliseconds to update the animation")
    debug := flag.Bool("debug", false, "Set to true if you would like to not run the program animation")
    flag.Parse()

    render.Init()
    render.Update()

    if *debug {
        render.Render()
        return
    }

    for {
        // Run at some time interval until the program exits
        render.Update()
        render.Render()

        time.Sleep(time.Duration(*interval) * time.Millisecond)
    }
}
