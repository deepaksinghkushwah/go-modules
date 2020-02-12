package helpers

import "fmt"

func init() {
	fmt.Println("From helpers.go init func in helper package")
}

func Show() {
	fmt.Println("this is from show func in helpers main packge")
}
