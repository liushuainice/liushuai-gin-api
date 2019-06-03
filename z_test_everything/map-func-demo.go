package main

import "fmt"

type Commendx struct {
	Commend string
	Msg     string
}

//func getCommandxList() map[string]func(*Commendx) bool {
func getCommandxList(comm string) func(*Commendx) bool {
	maps := map[string]func(*Commendx) bool{
		"help": gmHelp,
	}
	//return maps
	return maps[comm]
}

// 显示所有命令
func gmHelp(command *Commendx) bool {
	command.Msg = `test--commdend`
	return true
}
func main() {
	comm := Commendx{
		Commend: "help111",
		Msg:     "fail",
	}
	/*	commList := getCommandxList()
		commFunc,ok := commList[comm.Commend]
		if ok {
			commFunc(&comm)
			fmt.Println(comm)
		}else {
			fmt.Println(666)
		}*/
	commFunc := getCommandxList(comm.Commend)
	if commFunc != nil { //查不到方法会返回nil
		commFunc(&comm)
	} else {
		fmt.Println(666)
	}
	fmt.Println(comm)
}
