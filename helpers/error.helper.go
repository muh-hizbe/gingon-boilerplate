package helpers

import "log"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogIfError(label string, err error) {
	if err != nil {
		log.Println(label, " ===>> ", err)
	}
}
