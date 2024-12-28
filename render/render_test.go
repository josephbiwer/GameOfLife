package render

import (
    "testing"
    "mugrender/render"
)

func TestInitialize(t *testing.T) {
    render.Init()
    steam := render.GetSteam()

    if len(steam) != 196 {
        t.Error("size of steam array is not correct")
    }
    
    for c := range steam {
        if steam[c] != ' ' {
            t.Error("Steam array not initialized to ' '")
            break
        }
    }

}
