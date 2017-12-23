/*
        resource_darwin.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.23
*/
package bubblegame

import "os"
import "path"
import "trickyunits/mkl"
import "trickyunits/qstr"
import "log"
import "strings"

const darwininitdebugchat = true

func init(){
mkl.Version("Bubble Game Engine - Imports - resource_darwin.go","17.12.23")
mkl.Lic    ("Bubble Game Engine - Imports - resource_darwin.go","Mozilla Public License 2.0")
  trueapp,err:=os.Executable()
  if err!=nil { panic("Error when getting exe") }
  dirsplit:=strings.Split(trueapp,"/")
  appexe=""
  for i:=0;i<len(dirsplit);i++ {
    if appexe!="" { appexe+="/" }
    appexe+=dirsplit[i]
    }
  appdir=path.Dir(appexe)
  resdir=appexe+"/Contents/Resources"
  resfile=resdir+qstr.StripAll(appexe)+".jcr"
  if darwininitdebugchat {
     log.Print("Appdir:"+appdir)
     log.Print("AppExe:"+appexe)
     log.Print("Resdir:"+resdir)
     log.Print("Resfil:"+resfile)
   }
}
