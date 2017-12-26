/*
        flow-static.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.26
*/

//
package bubblegame

import (
	"trickyunits/mkl"
	"github.com/jpbubble/Base"
)

func init(){
	startvmscript["Static"] = "function bubble_init() end\n\nfunction bubble_main() Crash(\"There is no bubble_main function in this script\")\nend"
mkl.Version("Bubble Game Engine - Imports - flow-static.go","17.12.26")
mkl.Lic    ("Bubble Game Engine - Imports - flow-static.go","Mozilla Public License 2.0")
}

func flowStatic(){
	l := bubble.GetBubble("MAIN")
	bubble.QCall(l,"bubble_init",0)
	bubble.QCall(l,"bubble_main",0)
	Crash()
}
