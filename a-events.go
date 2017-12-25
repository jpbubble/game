/*
        a-events.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.26
*/
package bubblegame

import (
	"github.com/Shopify/go-lua"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/jpbubble/Base"
	"trickyunits/mkl"
	"time"
)


func nano() int64{
	t:=time.Now()
	return t.UnixNano()
}

var vapiquit = false
var keypressed = map[string] bool {}
var keydown = map[string] bool {}
var mousex,mousey int32
var mousedown = [16]bool{}
var mousehit = [16]bool{}
var lastnano = nano()

// Yeah this feature is not part of the API itself but the API can call to it.
// This is because all mode types need to be able to handle stuff on their own way.
func eventhandler() {
	l:=bubble.GetBubble(currentflow)
	vapiquit=false
	for k := range keypressed { delete(keypressed, k) }
	for i := 0; i<16; i++ { mousehit[i]=false }
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e:=event.(type) {
			case *sdl.QuitEvent:
				vapiquit=true
				if flowmode=="CallBack" { bubble.QCall(l,"bubble_quit",0) }
			//case *sdl.TextInputEvent:
				//tricon.Write(string(e.Text),180,0,255)
			case *sdl.KeyboardEvent:
				//tricon.WriteLn(fmt.Sprint(e.State,":",e.Keysym.Scancode,">",sdl.GetKeyName(e.Keysym.Sym)),180,255,0)
				key:=sdl.GetKeyName(e.Keysym.Sym)
				keydown[key] = e.State==1
				if e.State==1 { 
					keypressed[key]=true
					if flowmode=="CallBack" { bubble.QCall(l,"bubble_keydown",0,key) }
				} else {					
					if flowmode=="CallBack" { bubble.QCall(l,"bubble_keyup",0,key) }
				}
			case *sdl.MouseMotionEvent:
				mousex,mousey=e.X,e.Y
				if flowmode=="CallBack" { 
					l.Global("bubble_mousemove")
					l.PushNumber(float64(mousex))
					l.PushNumber(float64(mousey))
					bubble.Call(l,2,0)
				}
			case *sdl.MouseButtonEvent:
				mousex,mousey=e.X,e.Y
				but:=e.Button
				mousedown[but]=e.State==sdl.PRESSED
				if mousedown[but] {
					mousehit[but] = true
					l.Global("bubble_mousedown")
					l.PushNumber(float64(mousex))
					l.PushNumber(float64(mousey))
					l.PushNumber(float64(but))
					bubble.Call(l,3,0)
				} else {
					l.Global("bubble_mousedown")
					l.PushNumber(float64(mousex))
					l.PushNumber(float64(mousey))
					l.PushNumber(float64(but))
					bubble.Call(l,3,0)
				}
		}
	}
}

func checktime() bool {
	nu:=nano()
	r:=nu-lastnano
	if r<0 { r = r * (-1) }
	ret:=r>=timetick
	if ret { 
		lastnano = nu 
		if flowmode=="CallBack" { 
			l:=bubble.GetBubble(currentflow)
			bubble.QCall(l,"bubble_tick",0)
		}
	}
	return ret
}

func waittick() {
	for !checktime() {}
}

func lapi_SetTick(l *lua.State) int{
	timetick = int64(lua.CheckInteger(l,1))
	return 0
}

func lapi_WaitTick(l *lua.State) int{
	waittick()
	return 0
}

func lapi_KeyPressed(l *lua.State) int{
	key:=lua.CheckString(l,1)
	v,ok:=keypressed[key]
	l.PushBoolean(v && ok)
	return 1
}

func lapi_KeyDown(l *lua.State) int{
	key:=lua.CheckString(l,1)
	v,ok:=keydown[key]
	l.PushBoolean(v && ok)
	return 1
}

func lapi_MouseHit(l *lua.State) int{
	key:=lua.CheckInteger(l,1)
	v:=mousehit[key]
	l.PushBoolean(v)
	return 1
}

func lapi_MouseDown(l *lua.State) int{
	key:=lua.CheckInteger(l,1)
	v:=mousedown[key]
	l.PushBoolean(v)
	return 1
}

func lapi_MouseCoords(l *lua.State) int{
	l.PushNumber(float64(mousex))
	l.PushNumber(float64(mousey))
	return 2
}

func lapi_ClosureRequested(l *lua.State) int{
	l.PushBoolean(vapiquit)
	return 1
}

func initEventAPI(){
	// all modes
	LuaReg("SetTick",lapi_SetTick)
	LuaReg("KeyPressed",lapi_KeyPressed)
	LuaReg("KeyDown",lapi_KeyDown)
	LuaReg("MouseHit",lapi_MouseHit)
	LuaReg("MouseDown",lapi_MouseDown)
	LuaReg("MouseCoords",lapi_MouseCoords)
	LuaReg("ClosureRequested",lapi_ClosureRequested)
	// mode specific
	switch flowmode{
		case "Static":
			LuaReg("WaitTick",lapi_WaitTick)
		case "Cyclic":
			LuaReg("WaitTick",lapi_WaitTick)
		case "CallBack":
	}
}

func init(){
mkl.Version("Bubble Game Engine - Imports - a-events.go","17.12.26")
mkl.Lic    ("Bubble Game Engine - Imports - a-events.go","Mozilla Public License 2.0")
}
