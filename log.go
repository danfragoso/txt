package t

import (
	"fmt"
	"io/ioutil"
)

func log(l interface{}) {
	str := fmt.Sprintln(l)
	ioutil.WriteFile("./log.txt", []byte(str), 0644)
}

func Log(l interface{}) {
	str := fmt.Sprintln(l)
	ioutil.WriteFile("./log.txt", []byte(str), 0644)
}
