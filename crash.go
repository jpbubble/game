package bubblegame

import "trickyunits/mkl"
import "os"

func init(){
mkl.Version("","")
mkl.Lic("","")
}

type tCrashFunc struct{
	f func()
}

var crashFuncs []tCrashFunc

// Similar to Go's "defer" keyword, but it only takes functions without parameters, returning nothing and they are used whenever bubble is closed dow in the way it should.
func BDEFER(f func()){
	crashFuncs = append(crashFuncs,tCrashFunc{f})
}


// Executes all BDEFER registered funcions in reversed order as they were given and closes
// Optionally the first parameter may contain the exit code
func Crash(ec ...int){
	for i:=len(crashFuncs)-1;i>=0;i-- {
		crashFuncs[i].f()
	}
	ecc:=0
	if len(ec)>=1 { ecc=ec[0] }
	os.Exit(ecc)
}
