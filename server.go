package main
import (
    "fmt"
    "net"
    "os/exec"
    "strings"
)
func main() {
    message := "Hello, I am a server"   // отправляемое сообщение
    listener, err := net.Listen("tcp", ":4545")

    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        
        conn.Write([]byte(message))
        conn.Close()
    }
}
func execute(s string){
    args:= strings.Split(s," ")
    cmd:=exec.Command(args[0], args[1:]...)
    cmd.CombinedOutput()
  }
