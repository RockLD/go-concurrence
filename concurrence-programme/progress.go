package concurrence_programme

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"io"
)

func GetPid() {
	fmt.Printf("the pid is %d\n",os.Getpid())
	fmt.Printf("the ppid is %d\n",os.Getppid())
}

func Exec() {
	cmd0 := exec.Command("echo","-n","test cmd")

	stdout0,err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	if err:=cmd0.Start();err != nil{
		fmt.Println("error:",err)
		return
	}

	var outputBuf0 bytes.Buffer
	for {
		tempOutput := make([]byte,5)
		n,err := stdout0.Read(tempOutput)
		if err != nil {
			if err == io.EOF{
				break
			}else{
				fmt.Printf("err:",err)
				return
			}
		}
		if n>0 {
			outputBuf0.Write(tempOutput)
		}
	}

	fmt.Printf("%s\n",outputBuf0.String())

}