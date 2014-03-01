package main

import(
		"fmt"
		"moof/moof"
		"moof/data"
		"os"
		"io/ioutil"
)

func main(){
	/*
	args:=os.Args
	file,err := Open(args[1])
	if(err!=nil){
		fmt.Println("Unable to load file")
	}else{
	fullFile, err := ioutil.ReadFile(file)
	reader := data.NewReader(fullFile)
	moof := new(Moof)
	fmt.Println(moof.Read(reader))	
	*/
	args:=os.Args
	fullFile, err := ioutil.ReadFile(args[1])
	fmt.Printf("%s",fullFile)
	if err!=nil{
		fmt.Println("Error reading file")
	}else{
		reader := data.NewReader(fullFile)
		moof := new(moof.Moof)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		moof.Read(reader)
		fmt.Println(moof)
	}
	
	
}