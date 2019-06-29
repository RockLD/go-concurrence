package grammar

import "fmt"

func TypeGrammar(){
	//array
	var ipv4 [4]uint8 = [...]uint8{192,168,0,1}
	fmt.Println("array ipv4 is : ",ipv4)
	var ipv [4]uint8 = [...]uint8{192,168,0,1}
	fmt.Println("array ipv is :",ipv)

	//slice
	var ips = []string{"192","168","0","1"}
	fmt.Println("slice ips is : ",ips)
	cap_ips := cap(ips)
	fmt.Println("slice ips cap is :",cap_ips)
	ips = append(ips,"appended")
	fmt.Println("slice ips is :",ips)

	ips_make := make([]string,100)
	fmt.Println("slice ips_make is :",ips_make)
	fmt.Println("slice ips_make len is :",len(ips_make))

	//map
	var ipSwitchs = map[string]bool{}
	ipSwitchs["192.168.0.1"] = true
	ipSwitchs["192.168.0.3"] = true
	ipSwitchs["192.168.0.2"] = true
	fmt.Println("ipSwitchs is :",ipSwitchs)
	delete(ipSwitchs,"192.168.0.2")
	fmt.Println("ipSwitchs after del is :",ipSwitchs)

	//function
	/**
		func divide(divid int,divisor int)(int,error){

		}
	 */

	var talk Talk = new(myTalk)
	talk.Hello("li")
	_,ok := talk.(*myTalk)
	fmt.Println("talk is the *myTalk type",ok)

	var chatbot Chatbot = new(myTalk)
	_,ok = chatbot.(*myTalk)
	fmt.Println("chatbot is the *myTalk type",ok)
	//struct
	var simpleCN simpleCN
	simpleCN.name = "ld"
	fmt.Println(simpleCN)
}

type myInt int

func (i *myInt) add(another int) myInt {
	*i = *i + myInt(another)
	return *i
}

type Talk interface {
	Hello(userName string)string
	Talk(head string)(saying string,end bool,err error)
}

type myTalk string

func (my *myTalk)Hello(userName string)string {
	return ""
}

func (my *myTalk)Talk(head string)(saying string,end bool,err error){
	return "",true,nil
}

func (my *myTalk)Name()string {
	return ""
}

func (my *myTalk)Begin()(string, error) {
	return "",nil
}

func (my *myTalk)ReportError(err error)string {
	return ""
}

func (my *myTalk)End()error {
	return nil
}

type Chatbot interface {
	Name() string
	Begin()(string ,error)
	Talk
	ReportError(err error)string
	End() error
}

type simpleCN struct {
	name string
	talk Talk
}