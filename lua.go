/*
        lua.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubblegame

import(
	"github.com/Shopify/go-lua"
	"github.com/jpbubble/Base"
	"trickyunits/mkl"
	)
	
	
func LuaReg(name string, f func(l *lua.State) int ) {
	bubbleapi = append(bubbleapi,lua.RegistryFunction{ name, f } )
}


func init(){
bubble.BREG(func (l *lua.State) { lua.SetFunctions(l,bubbleapi,0) })
mkl.Lic    ("Bubble Game Engine - Imports - lua.go","Mozilla Public License 2.0")
mkl.Version("Bubble Game Engine - Imports - lua.go","17.12.24")
}
