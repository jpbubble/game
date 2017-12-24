/*
        console.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubblegame

import (
	"unsafe"
	"trickyunits/mkl"
	"trickyunits/tricon"
	"trickyunits/jcr6/jcr6main"
	"github.com/jpbubble/Base"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/sdl"
	)
	
func bubbleconsole_bc_write(color,txt string) {
	r,g,b:=Col(color)
	tricon.Write(txt,r,g,b)
}


func bubbleconsole_bc_writeln(color,txt string) {
	r,g,b:=Col(color)
	tricon.WriteLn(txt,r,g,b)
}

func bubbleconsole_bc_error(txt string,fatal ...bool) {
	bubbleconsole_bc_writeln("Red","ERROR!")
	bubbleconsole_bc_writeln("Yellow",txt)
	if len(fatal)>0 {
		if fatal[0] { 
			bubbleconsole_bc_writeln("Magenta","This is a fatal error! Terminating!")
			bubbleconsole_bc_writeln("Purple","Press any key to end this program!")
			tricon.Show()
			tricon.Pause()
			Crash(20)
		}
	}
}

func bubbleconsole_bc_warn(txt string) {
	bubbleconsole_bc_writeln("Amber","WARNING!")
	bubbleconsole_bc_writeln("Yellow",txt)
}

var cfont *ttf.Font

func killConfont() { cfont.Close() }


func setupConsole() {
	bubble.SetConsole(bubbleconsole_bc_write,bubbleconsole_bc_writeln,bubbleconsole_bc_error,bubbleconsole_bc_warn)
	if !jcr6main.HasEntry(jcr,"Console/Font.ttf") { pi_error("I don't have Console/Font.ttf") }
	bfont:=jcr6main.JCR_B(jcr,"Console/Font.ttf") 
	rfont:=sdl.RWFromMem(unsafe.Pointer(&bfont[0]), len(bfont))
	tfont,err:=ttf.OpenFontRW(rfont, 0, 14) 
	if err!=nil { pi_error("Error loading console font\n\n"+err.Error()) }
	cfont=tfont
	tricon.Setup(window,cfont)
	BDEFER(killConfont)
	BDEFER(tricon.Kill)
	bubble.WriteLn("Yellow",EngineName+" -- version: "+mkl.Newest())
	bubble.WriteLn("Cyan",EngineCopyright)
}


func init(){
mkl.Version("Bubble Game Engine - Imports - console.go","17.12.24")
mkl.Lic    ("Bubble Game Engine - Imports - console.go","Mozilla Public License 2.0")

}
