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
	"trickyunits/gini"
	"trickyunits/jcr6/jcr6main"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strconv"
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
	var err error
	// Init the SDL routines
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err!=nil { pi_error("Could not init SDL\n\n"+err.Error()) }
	BDEFER(sdl.Quit)
	err = ttf.Init()
	if err!=nil { pi_error("Could not init SDL font library\n\n"+err.Error()) }
	BDEFER(ttf.Quit)
	// Check for resource
	if !qff.Exists(resfile) { pi_error("I could not find "+resfile) }
	r:=jcr6main.Recognize(resfile)
	switch r {
		case "NONE":
			pi_error("Unrecognized resource: "+resfile)
		case "WAD":
			pi_error("Are you crazy? Using a WAD file for this game? Well if you want... I don't advice you to, due to the 8 char entry name restriction. It makes looking for the required ID files already impossible :P")
		}
	// Read resource
	jcr=jcr6main.Dir(resfile)
	if jcr6main.JCR6Error!="" { pi_error(jcr6main.JCR6Error) }
	// Check id and load it
	if !jcr6main.HasEntry(jcr,"ID/Identify.gini") { pi_error("Resource has no identify data") }
	giniload := jcr6main.JCR_B(jcr,"ID/Identify.gini")
	ini = gini.ParseBytes(giniload)
	// Let's check the data for engine correctness and version requirments
	if ini.C("ENGINE")=="" { pi_error("No engine data present") }
	if EngineName=="" { pi_error("Hey! I need to have the name of the engine itself, but the programmer didn't give me that.... STOOOOPID!") }
	if ini.C("ENGINE")!=EngineName { pi_error("Sorry, but this resource was written for the "+ini.C("ENGINE")+" engine, and this is the "+EngineName+" engine, so this resource is useless for me!") }
	// Parse Window 
	if ini.C("TITLE")=="" { win_title=EngineName+" project" } else { win_title=ini.C("TITLE") }
	win_w,err = strconv.ParseInt(ini.C("Width") ,0,32); if err!=nil  { pi_error("Error getting desired window width\n\n" +err.Error()) }
	win_h,err = strconv.ParseInt(ini.C("Height"),0,32); if err!=nil { pi_error("Error getting desired window height\n\n"+err.Error()) }
	startWindow()
	// Parse Debug Console
	setupConsole()
}
