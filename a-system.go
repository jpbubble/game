/*
        a-system.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.25
*/
import(
	"github.com/Shopify/go-lua"
	"trickyunits/mkl"
)

func lapi_Crash(l *lua.State) int{
	i:=lua.OptInteger(l,1,0)
	Crash(i)
	return 0
}

func lapi_EngineVersion(l *lua.State) int{
	l.PushString(mkl.Newest())
	return 1
}

func lapi_EndFlow(l *lua.State) int {
	endflow = true
	return 0
}

func lapi_Identify(l *lua.State) int {
	key := lua.CheckString(l,1)
	l.PushString(ini.C(key))
	return 1
}


func initSysLib(){
	LuaReg("Crash",lapi_Crash)
	LuaReg("EngineVersion",lapi_EngineVersion)
	LuaReg("Identify",lapi_Identify)
	if flowmode!="Static" { LuaReg("EndFlow",lapi_EndFlow) }
}


func init(){
mkl.Version("Bubble Game Engine - Imports - a-system.go","17.12.25")
mkl.Lic    ("Bubble Game Engine - Imports - a-system.go","Mozilla Public License 2.0")
}
