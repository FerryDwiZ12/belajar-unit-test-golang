package helper

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Eko")
	if result != "HelloEko" {
		panic(`result is not 'Hello Eko'`)
	}
}