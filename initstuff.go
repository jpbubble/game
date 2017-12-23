/*
        initstuff.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubblegame

import (
	"trickyunits/mkl"
	"trickyunits/qff"
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	)
	
	
	
func init(){
mkl.Version("Bubble Game Engine - Imports - initstuff.go","17.12.24")
mkl.Lic    ("Bubble Game Engine - Imports - initstuff.go","Mozilla Public License 2.0")
}

func pi_error(errormsg string){
        buttons := []sdl.MessageBoxButtonData{
                {0, 0, "Really?"},
                {sdl.MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "Oh crap!"},
                {sdl.MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, 2, "Goodbye!"},
        }

        colorScheme := sdl.MessageBoxColorScheme{
                Colors: [5]sdl.MessageBoxColor{
                        sdl.MessageBoxColor{255, 0, 0},
                        sdl.MessageBoxColor{0, 255, 0},
                        sdl.MessageBoxColor{255, 255, 0},
                        sdl.MessageBoxColor{0, 0, 255},
                        sdl.MessageBoxColor{255, 0, 255},
                },
        }

        messageboxdata := sdl.MessageBoxData{
                sdl.MESSAGEBOX_ERROR,
                nil,
                "FATAL ERROR",
                errormsg,
                int32(len(buttons)),
                buttons,
                &colorScheme,
        }

        var buttonid int32
        var err error
        if buttonid, err = sdl.ShowMessageBox(&messageboxdata); err != nil {
                fmt.Println("error displaying message box")
                return
        }
	/*
        if buttonid == -1 {
                fmt.Println("no selection")
        } else {
                fmt.Println("selection was", buttons[buttonid].Text)
        }
        */
        fmt.Printf("Use clicked button number %d\nNot that it really matters, though!\n\n",buttonid) // Go will otherwise refuse to compile. Silly huh?
        Crash(1)
}



// Reads the ID/Identify.gini file from the JCR resource.
// The JCR file can basically be any file JCR6 supports (except for WAD, as it has a limitation of 8 characters per entry name)
// but it must be named the same as the executable file running it with it's regular extension replaced with .jcr
// Please note that on Mac this file must be located inside the application bundle's Resources folder (mac users expect this).
func InitBubble(){
	if !qff.Exists(resfile) { pi_error("I could not find "+resfile) }
}
