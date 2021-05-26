package main

import (
	"fmt"
	"gotest/utils"
)

func retInter() utils.Interface_{
	return &utils.Struct_{}
}

func main() {
	ret := retInter()
	fmt.Println(ret.F())
}

func client(junjie utils.Interface_)  {
	fmt.Println(junjie.F())
}