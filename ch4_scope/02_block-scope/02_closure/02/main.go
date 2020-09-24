package main

import "fmt"

var x = 0

func increment() int  {
	x ++
	return x
}

func main(){
	fmt.Println(increment())
	fmt.Println(increment())
}
/*
闭包帮助我们显著多个函数使用变量的范围，如果多个函数访问需要是package scope
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
 */