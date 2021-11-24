package main 

import "fmt"

func main(){

	var numbers [10]int;

	for i:=0 ; i < 10; i++{
		numbers[i] = i;
	}

	for i:=0; i < 10; i++{
		fmt.Println("Numbers : ", i)
	}

}