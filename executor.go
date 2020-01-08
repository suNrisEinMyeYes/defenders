package executor

import(
    "os/exec"
    "strings"
)

func execute(s string){
  args:= strings.Split(s," ")
  cmd:=exec.Command(args[0], args[1:]...)
  cmd.CombinedOutput()
}
