package ts

import(
		"fmt"
		"regifted/moof"
		"regifted/data"
		"os"
		"io/ioutil"
)

func main(){
	args:=os.Args
	fullFile, err := ioutil.ReadFile(args[1])
	if err!=nil{
		fmt.Println("Error reading file")
	}else{
		reader := data.NewReader(fullFile)
		moof := new(moof.Moof)
		moof.Read(reader)
		fmt.Println(moof)
	}
	
	
}