package main

import (
	"golog/configure"
	"golog/gologger"
	_ "golog/gologger"
)

func main() {

	configure.Timer()
	configure.CallingFunction()
	gologger.D("Yashish")
	gologger.E("Dua")

}

