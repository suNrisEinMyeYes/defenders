package main

import(
    "os/exec"
    "strings"
)

func main(){
  s:="python script.py"
  args:= strings.Split(s," ")
  cmd:=exec.Command(args[0], args[1:]...)
  cmd.CombinedOutput()
  
}
