/*
        a-window.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubblegame

/* This file will deal with all Window featuers in general and 
 * will also grant the Lua scripts some functionality as far as windows
 * are concerned.
 */
 
 import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/Shopify/go-lua"
	"trickyunits/mkl"
)

// init
func startWindow(){
	// Make sure all defined windows are destroyed
	BDEFER( func() {
		if window!=nil { window.Destroy() }
	})
	var err error
	flags := uint32(0) // more possibilities will be there soon!
	window,err = sdl.CreateWindow(win_title, 0, 0, int32(win_w), int32(win_h), flags)
	if err!=nil { pi_error("I couldn't create the window:\n\n"+err.Error()) }
}




// Lua api
func lapi_GetWinSize(l *lua.State) int{
	w,h := window.GetSize()
	l.PushNumber(float64(w))
	l.PushNumber(float64(h))
	return 2
}



func init(){
mkl.Version("Bubble Game Engine - Imports - a-window.go","17.12.24")
mkl.Lic    ("Bubble Game Engine - Imports - a-window.go","Mozilla Public License 2.0")
	LuaReg("Win_GetSize",lapi_GetWinSize)
}


