package test

import (
	"testing"
	)

//func TestGoTest(t *testing.T) {
//	GoTest()
//}

func TestGoTestFailed(t *testing.T) {
	GoTest()
}


func BenchmarkGoTest(b *testing.B) {
	for i:=0;i<10;i++{
		GoTest()

	}
}