package bubblegame

import "os"
import "path"
import "trickyunits/mkl"
import "trickyunits/qstr"
import "log"

const darwininitdebugchat = true

func init(){
  mkl.Version("","")
  mkl.Lic("","")
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
