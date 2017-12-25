/*
        globals.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.25
*/
package bubblegame


import(
	"trickyunits/jcr6/jcr6main"
	"trickyunits/mkl"
	"trickyunits/gini"
	"github.com/Shopify/go-lua"
	"github.com/veandco/go-sdl2/sdl"
)


var appdir,appexe,resdir,resfile string


// This must be set prior it initizing.
// This string will be used to check the data in the ID/Identify.gini file
var EngineName string
var EngineCopyright = "(c) Jeroen P. Broks"




var jcr jcr6main.TJCR6Dir

var ini gini.TGINI


func init(){
mkl.Version("Bubble Game Engine - Imports - globals.go","17.12.25")
mkl.Lic    ("Bubble Game Engine - Imports - globals.go","Mozilla Public License 2.0")
bubcol["Red"]        = tbubcol{255,  0,  0}
bubcol["Green"]      = tbubcol{  0,255,  0}
bubcol["Blue"]       = tbubcol{  0,  0,255}
bubcol["Yellow"]     = tbubcol{255,255,  0}
bubcol["Magenta"]    = tbubcol{255,  0,255}
bubcol["Cyan"]       = tbubcol{  0,255,  0}
bubcol["Amber"]      = tbubcol{255,180,  0}
bubcol["LightGreen"] = tbubcol{180,255,  0}
bubcol["Purple"]     = tbubcol{180,  0,255}
bubcol["SkyBlue"]    = tbubcol{  0,180,255}
bubcol["White"]      = tbubcol{255,255,255}
}



// window globals
var window *sdl.Window
var win_title string
var win_w,win_h int64


// Lua
var bubbleapi = []lua.RegistryFunction{}

// colors
type tbubcol struct{
	r,g,b uint8
}
var bubcol = map[string]tbubcol{}

// A quick routine for quickly seeing color values.
// Most of use for the tricon support, but why not make it "public"? :P
func Col(name string) (r,g,b uint8) { 
	if ret,ok:=bubcol[name];ok{
		return ret.r,ret.g,ret.b
	} else {
		return 255,255,255
	}
}


// Flowmode
var flowmode = "Static"


// Start VM script
var startvm = "MAIN"
var startscript "Script/Main.lua"


var startvmscript = map[string]string{}
