package main

import "fmt"
import "regexp"

func isValidFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func main(){
	email:="example@gmail.com"
	valid := isValidFormat(email)
	if valid {
		fmt.Println("✅ Email is valid")
	} else {
		fmt.Println("❌ Email is invalid")
	}
	fmt.Println(email)
}