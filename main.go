package main

import (
	"fmt"

	"os"
)
func main(){
	// Read the file the source file
	data, err := os.ReadFile("myfile.txt") 
	if err != nil {
		fmt.Println("Error reading the file: ",err)
		return
	}
	// Prints the source file
	fmt.Println(string(data))
	// Open the destination file
	data2,err := os.OpenFile("myfile2.txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
	if err != nil {
		fmt.Println("Error opening the file:",err)
		return
	}
	defer data2.Close()
	// appends the source file to the destination file
	_, err = data2.Write(data)
	if err != nil {
		fmt.Println("error appending the file; ",err)
		return
	}
	// read the new version of the modified file
	newData , err := os.ReadFile("myfile2.txt")
	if err != nil {
		fmt.Println("error reading the file: ", err)
		return
	}
	fmt.Println(string(newData))

}