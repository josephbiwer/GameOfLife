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

func TestSteamPropagation(t *testing.T) {

    for c := 0; c < 5; c++ {
        
        render.Init()
        steam := render.GetSteam()[5]
        render.Update()
        newSteam := render.GetSteam()[31]

        if newSteam != steam {
            t.Errorf("Steam was not properly copied. Itteration: %d, Risen steam: %c, old steam: %c", c, steam, newSteam)
        }

    }

}
