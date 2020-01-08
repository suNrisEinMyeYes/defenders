package main
import (
    "fmt"
    "os"
    "net"
    "io"
    "os/exec"
    "strings"
    "io/ioutil"
    "encoding/json"
)
type jsfile struct{
	Stage string 
	String string 
	Key string 
	Init_key string 
}
func main() {
    jsonFile, err := os.Open("clients.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data jsfile
	json.Unmarshal(byteValue, &data)
    fmt.Println(data.Stage)
    s:="python3 Client_prot.py"
    execute(s)
                                                                                                                                                                            key:= data.Key

    conn, err := net.Dial("tcp", "127.0.0.1:4545")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    io.Copy(os.Stdout, conn)
    fmt.Println("\nDone")
}

func execute(s string){
    args:= strings.Split(s," ")
    cmd:=exec.Command(args[0], args[1:]...)
    cmd.CombinedOutput()
  }
  
