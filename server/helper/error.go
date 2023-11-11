package helper

import (
	"fmt"
)

// jika ada error, maka kita akan membuat dia menjadi sebuah `panic`
func HandlePanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleError() {
	// dan akan di handle disini
	err := recover()
	if err != nil {
		fmt.Println("Ada error nih,", err)
	}
}
