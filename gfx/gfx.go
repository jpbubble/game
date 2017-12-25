/*
        gfx.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.25
*/
// ---- //

// This is just a lib meant to simplefy the usage of graphics in SDL.
// This was basically only set up to allow me to set up the API used
// in the Game import more easily, but this way you can also use it as
// an independent lib... if you like
package gfx

import(
			"fmt"
			"path"
			"unsafe"
			"errors"
			"strings"
			"strconv"
			"trickyunits/mkl"
			"trickyunits/qstr"
			"trickyunits/gini"
	jcr6	"trickyunits/jcr6/jcr6main"
			"github.com/veandco/go-sdl2/sdl"
			"github.com/veandco/go-sdl2/img"
_			"github.com/veandco/go-sdl2/ttf"
)



type myIndPic struct{
	surf *sdl.Surface
	text *sdl.Texture
	lrnd *sdl.Renderer
}

type MyPic struct{
	indpic []myIndPic
	hotx int32
	hoty int32
	//scale int
}

var drawr,drawg,drawb uint8 = 255,255,255


// Define the current renderer on here.
// When not set we cannot work.
var GREND *sdl.Renderer

func loadindpic(j jcr6.TJCR6Dir,entry string) (myIndPic,error){
	if GREND==nil { return myIndPic{},errors.New("No features in this library may be used prior to setting a renderer!") }
	if !jcr6.HasEntry(j,entry) { return myIndPic{},errors.New("Entry "+entry+" not found in requested JCR6 resource") }
	bpic:=jcr6.JCR_B(j,entry)
	if jcr6.JCR6Error!="" { return myIndPic{},errors.New("JCR6 Error> "+jcr6.JCR6Error) }
	rpic:=sdl.RWFromMem(unsafe.Pointer(&bpic[0]), len(bpic))
	defer rpic.FreeRW()
	var lpic *sdl.Surface
	var e error
	if img.IsPNG(rpic) {
		lpic,e=img.LoadPNGRW(rpic)
	} else if img.IsBMP(rpic) {
		lpic,e=img.LoadBMPRW(rpic)
	} else if img.IsJPG(rpic) {
		lpic,e=img.LoadJPGRW(rpic)
	} else if img.IsICO(rpic) {
		lpic,e=img.LoadICORW(rpic)
	} else if img.IsLBM(rpic) {
		lpic,e=img.LoadLBMRW(rpic)
	} else if img.IsPCX(rpic) {
		lpic,e=img.LoadPCXRW(rpic)
	} else if img.IsTIF(rpic) {
		lpic,e=img.LoadTIFRW(rpic)
	} else {
		e = errors.New("Sorry, what kind of picture format is image: "+entry+"?")
	}
	ret:=myIndPic{ lpic,nil,nil }
	return ret,e
}


func loadsingle(j jcr6.TJCR6Dir,entry string) (*MyPic,error){
	ret:=MyPic{}
	mIP,err:=loadindpic(j,entry)
	ret.indpic=append(ret.indpic,mIP)
	ret.hotx=0
	ret.hoty=0
	//ret.scale=100
	return &ret,err
}

func loadbundle(j jcr6.TJCR6Dir,entry string) (*MyPic,error){
	ret:=MyPic{}
	e:=strings.ToUpper(entry)
	el:=len(e)
	for _,ent := range jcr6.EntryList(j){
		if qstr.Left(ent,el)==e && strings.ToUpper(path.Base(ent))!="DATA.GINI"{
			mIP,err:=loadindpic(j,ent)
			if err!=nil { return nil,err }
			ret.indpic = append(ret.indpic,mIP)
		}
	}
	var per error
	if jcr6.HasEntry(j,entry+"/Data.gini") {
		bgini:=gini.ParseBytes(jcr6.JCR_B(j,entry+"/Data.gini"))
		i64:=int64(0)
		if bgini.C("HOT.X")!="" { i64,per =strconv.ParseInt(bgini.C("HOT.X"),0,32); ret.hotx=int32(i64) }
		if bgini.C("HOT.Y")!="" { i64,per =strconv.ParseInt(bgini.C("HOT.Y"),0,32); ret.hoty=int32(i64) }
	}
	return &ret,per
}

// Load image from a JCR resource.
// if suffixed with ".BPB" (Bubble picture bundle) it will mean the file is a directory with that name (suffix included) with in it all picture files to load, which will load in the order in which they are found inside the JCR file.
// If a data.gini file is found inside this bundle it can be use to easily hotspot data, and perhaps it can get more functionality in the future.
func LoadImage(j jcr6.TJCR6Dir,entry string) (*MyPic,error){
	ce:=strings.ToUpper(entry)
	if qstr.Right(ce,4)==".BPB" { 
		return loadbundle(j,entry)
	} else {
		return loadsingle(j,entry)
	}
}

// Draws a loaded picture
func (pic *MyPic) Draw(x,y int32, frame int) error {
	var err error
	if GREND==nil { return errors.New("I don't have a renderer!") }
	if frame>=len(pic.indpic) { return errors.New(fmt.Sprintf("You tried to access frame %d, but this image only has %d frame(s)",frame,len(pic.indpic))) }
	if len(pic.indpic)==0 { return errors.New("Empty picture resource!") }
	ip:=pic.indpic[frame]
	if GREND!=ip.lrnd {
		if ip.text!=nil { ip.text.Destroy() }
		ip.text=nil
	}
	if ip.text==nil {
		ip.text,err = GREND.CreateTextureFromSurface(ip.surf)
		ip.lrnd = GREND
		if err!= nil { return err }
	}
	ip.text.SetColorMod(drawr,drawg,drawb)
	ctsr:=sdl.Rect{0, 0, ip.surf.W, ip.surf.H}
	cttr:=sdl.Rect{x,y,ip.surf.W,ip.surf.H}
	GREND.Copy(ip.text,&ctsr,&cttr)
	return nil
}

func (pic *MyPic) Frames() int {
	return len(pic.indpic)
}

// Fress the memory from everything taken up by this image.
func (pic *MyPic) Kill() {
	for _,idp:=range pic.indpic {
		if idp.surf!=nil { idp.surf.Free() }
		if idp.text!=nil { idp.text.Destroy() }
	}
}

func Cls() { GREND.Clear() }

func Color(r,g,b uint8) {
	drawr,drawg,drawb=r,g,b
}

func init(){
mkl.Version("Bubble Game Engine - Imports - gfx.go","17.12.25")
mkl.Lic    ("Bubble Game Engine - Imports - gfx.go","Mozilla Public License 2.0")
}
