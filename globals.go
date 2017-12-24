/*
        globals.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubblegame


import(
	"trickyunits/jcr6/jcr6main"
	"trickyunits/mkl"
	"trickyunits/gini"
)


var appdir,appexe,resdir,resfile string


// This must be set prior it initizing.
// This string will be used to check the data in the ID/Identify.gini file
var EngineName string




var jcr jcr6main.TJCR6Dir

var ini gini.TGINI


func init(){
	mkl.Version("","")
	mkl.Lic("","")
}
