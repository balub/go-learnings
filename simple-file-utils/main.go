package main

import (
	"fmt"
	"github.com/balub/file-utils/file-funcs"
)

func main(){

	contents, err := fileUtils.ReadFile("test.txt")
	if err !=nil{
		fmt.Println(err)
	}
	 fmt.Println(contents)

	 err = fileUtils.WriteToFile("test.txt","Im being written programmatically")
	 if err != nil{
		fmt.Println(err)
	 }
	
}