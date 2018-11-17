package __1

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func TestInsertionSort(t *testing.T) {
	testData:=make([]int,0)
	rander:=rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0;i<100;i++{
		testData = append(testData,rander.Intn(100))
	}
	result:=InsertionSort(testData)
	fmt.Println(result)
}
