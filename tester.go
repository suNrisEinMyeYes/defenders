package main
import(
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)
type js struct{
	Stage string 
	String string 
	Key string 
	Init_key string 
}

func main(){
	jsonFile, err := os.Open("clients.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data js
	json.Unmarshal(byteValue, &data)
	fmt.Println(data.Stage)
	data.Stage = "golangggg"
	//data.String = ""
	//data.key = ""
	//data.init_key = ""
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("info_for_golang.json", file, 0644)
}