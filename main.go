package main

import (
	"fmt"
	"os"
)

func main() {
	//show directory of current file
	dir, err := os.Getwd()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)

	isErr := fileCreate("test.txt", "test 2: hi hello..")
	fmt.Println(isErr)


	fi, err := os.Stat("ami.txt")
	if err!=nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fi.IsDir())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())




}

func fileCreate(fineName,content string) bool {
	//create directory
	posf,err := os.Create(fineName)
	if err!=nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	//write in directory
	_ ,err = posf.Write([]byte(content))
	// fmt.Println(n)
	if err!=nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}