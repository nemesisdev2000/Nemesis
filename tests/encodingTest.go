package main

import (
	"fmt"
	"os/exec"
)

func main() {
	/*
		var data string
		fmt.Println("Enter string : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			data = scanner.Text()
		}
		fmt.Println("Before encoding : ", data)
		sEnc := b64.StdEncoding.EncodeToString([]byte(data))
		fmt.Println("After encoding ", sEnc)
	*/
	test := "Get-ChildItem"
	cmd := exec.Command("powershell", "-c", test)
	stdout, _ := cmd.Output()

	fmt.Println("-> ", string(stdout))
}
