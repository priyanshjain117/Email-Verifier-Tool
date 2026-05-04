package main

import "fmt"

func main(){
	email:="example@gmail.com"
	valid := VerifyEmail(email)
	fmt.Println(valid)
	fmt.Println(email)
}