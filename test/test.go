package test

import (
	"fmt"
	"github.com/pkg/errors"
)

func GoTest(){
	i:=0
	for j:=0;j<1000000000;j++{
		i++

	}

}


func GoTestFailed(){
	fmt.Println("Test Hello World")
	panic(errors.New("test"))
}
