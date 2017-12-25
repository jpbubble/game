/*
        flow.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.25
*/

//
package bubblegame

import(
	"trickyunits/mkl"
	"github.com/jpbubble/Base"
)

func init(){
mkl.Version("Bubble Game Engine - Imports - flow.go","17.12.25")
mkl.Lic    ("Bubble Game Engine - Imports - flow.go","Mozilla Public License 2.0")
}

func Run(){
	switch flowmode {
		case "Static":
			flowStatic()
		default:
			bubble.Fatal("Internal error! Unknown flow mode: "+flowmode) // As the flowmode was already checked before this should not be possible to happen, that's why it's deemed an internal error. Proof that there is a bug!
	}
}
