package bubblegame

import(
	"trickyunits/mkl"
	"github.com/jpbubble/Base"
	"github.com/jpbubble/game/gfx"
	"github.com/Shopify/go-lua"
	"fmt"
)

var images = map[string] *gfx.MyPic {}


func pic_assign(img *gfx.MyPic,entry string){
}

func pichave(entry string) bool{
	_,ret:=images[entry]
	return ret
}

func picfree(entry string) {
	pic,ok:=images[entry]
	if !ok { return } // Trying to delete a non-existent picture will be a no-op.
	shared:=false
	for k,v:=range images {
		shared = shared || (k!=entry && v==pic)
	}
	if !shared { pic.Kill() } // If there are multiple pointers, then DON'T free the data yet!
	delete(images,entry)
}

func lapi_Cls(l *lua.State) int{
	gfx.Cls()
	return 0
}

func lapi_FreeImage(l *lua.State) int {
	img:=lua.CheckString(l,1)
	picfree(img)
	return 0
}

func lapi_AssignImage(l *lua.State) int{
	ori:=lua.CheckString(l,1)
	tgt:=lua.CheckString(l,2)
	freeoriginal:=lua.OptInteger(l,3,1)>0
	if !pichave(ori) { bubble.Fatal("No picture has been registered at id '"+ori+"'") }
	if  pichave(tgt) { picfree(tgt) }
	images[tgt]=images[ori]
	if freeoriginal {delete(images,ori)}
	return 0
}

func lapi_LoadImage(l *lua.State) int {
	file:=lua.CheckString(l,1)
	entry:=lua.OptString(l,2,"")
	img,err:=gfx.LoadImage(jcr,file)
	if err!=nil { bubble.Fatal("LoadImage(\""+file+"\",\""+entry+"\"): "+err.Error()) }
	if entry=="" {
		hi:=1
		hx:=fmt.Sprintf("%X",hi)
		for pichave(hx){
			hi++
			hx=fmt.Sprintf("%X",hi)
		}
		entry=hx
	} else {
		if pichave(entry) { picfree(entry) }
	}
	images[entry]=img
	l.PushString(entry)
	return 1
}

func lapi_DrawImage(l *lua.State) int {
	imgkey:=lua.CheckString(l,1)
	x:=int32(lua.CheckInteger(l,2))
	y:=int32(lua.CheckInteger(l,3))
	frame:=lua.OptInteger(l,4,0)
	pic,ok:=images[imgkey]
	if !ok { bubble.Fatal(fmt.Sprintf("LoadImage(\"%s\",%d,%d,%d): Image doesn't exist!",imgkey,x,y,frame)) }
	err:=pic.Draw(x,y,frame)
	if err!=nil { bubble.Fatal(fmt.Sprintf("LoadImage(\"%s\",%d,%d,%d): %s",imgkey,x,y,frame,err.Error())) }
	return 0
}

func lapi_ImageFrames(l *lua.State) int {
	imgkey:=lua.CheckString(l,1)
	pic,ok:=images[imgkey]
	if !ok { bubble.Fatal(fmt.Sprintf("ImageFrames(\"%s\"): Image doesn't exist")) }
	l.PushNumber(float64(pic.Frames()))
	return 1
}

func lapi_Flip(l *lua.State) int {
	gfx.GREND.Present()
	return 0
}




func gfxAPIinit() { // This is needed as there are some very specific settings here!
	LuaReg("Cls",lapi_Cls)
	LuaReg("LoadImage",lapi_LoadImage)
	LuaReg("DrawImage",lapi_DrawImage)
	LuaReg("ImageFrames",lapi_ImageFrames)
	LuaReg("AssignImage",lapi_AssignImage)
	LuaReg("FreeImage",lapi_FreeImage)
	if flowmode!="CallBack" {
		LuaReg("Flip",lapi_Flip)
	}
}



func init(){
mkl.Version("","")
mkl.Lic("","")
}
