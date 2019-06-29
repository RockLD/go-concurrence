package grammar

import "fmt"

var v = "1,2,3"

func Redeclare(){
	v := []int{1,2,3}
	if v != nil {
		var v = 123
		fmt.Printf("%v\n",v)
	}
}

func Defer(){
	defer fmt.Println("deffffffer")
	defer fmt.Println("deffffffer2")
	defer fmt.Println("deffffffer3")
	defer fmt.Println("deffffffer4")
	defer fmt.Println("deffffffer5")
	fmt.Println("first")
}

func PanicRecover() {
	//defer func() {
	//	if p:=recover();p!=nil{
	//		fmt.Println("Recover panic :",p)
	//	}
	//}()
	panic("test panice1")
}