package main

import (
	"container/list"
	"fmt"
)

//go 的 list 使用方法 -> linkedList
func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在“fist”之后添加”high”
	l.InsertAfter("high", element)
	// 在“fist”之前添加”noon”
	l.InsertBefore("noon", element)
	// 使用
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//l.MoveToBack(element)//移动到最后
	//l.MoveToFront(element)//移动到最前
	//l.Remove(element)
	fmt.Println()
	for {
		v := l.Front().Value
		v = "asdf  " + v.(string) //类型断言,v 不是string会报错
		fmt.Println(v)            //value 是interface类型
		switch e := v.(type) {
		case string:
			fmt.Println(e)
		}
		l.Remove(l.Front()) //POP方法
		if l.Len() <= 0 {
			break
		}
	}
}
