/*
@Time : 2019/9/17 10:01
@Author : nancy
@File : reflect_test
*/
package reflect

import (
	"fmt"
	"sync"
	"testing"
)

type TestInterface struct {
	TestValue string
	TestKey  string
}

func convertInterfaceToObject1() *TestInterface{
	m := new(sync.Map)
	m.Store("testKey", TestInterface{TestKey:"key", TestValue:"value"})
	value,err := m.Load("testKey")
	if err{
		switch objectType := value.(type){
		case TestInterface:
			return &objectType
		default:
			return nil
		}

	}
	return nil
}

// go interface 的assert; interface 转成对应的struct
func convertInterfaceToObject2() *TestInterface{
	m := new(sync.Map)
	m.Store("testKey", TestInterface{TestKey:"key", TestValue:"value"})
	value,err := m.Load("testKey")
	if err{
		 p, ok := value.(TestInterface)
		 if ok {
		 	fmt.Println(p.TestValue)
		 	return &p
		 } else {
		 	fmt.Println("Failed to convert interface to TestInterface")
		 	return nil
		 }

	}
	return nil
}


func TestConcurrentRange(t *testing.T) {


	//convertInterfaceToObject1()
	convertInterfaceToObject2()


}