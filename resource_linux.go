/*
        resource_linux.go
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
import "strings"
import "trickyunits/mkl"
//import "trickyunits/qstr"


func init(){
mkl.Version("Bubble Game Engine - Imports - resource_linux.go","17.12.23")
mkl.Lic    ("Bubble Game Engine - Imports - resource_linux.go","Mozilla Public License 2.0")
	trueapp,err:=os.Executable()
	if err!=nil { panic("Error when getting exe") }
	appexe =trueapp
	appdir =path.Dir(appexe)
	resdir =appdir
	resfile=appexe+".jcr"
}
