package main

import (
	"fmt"
	"strconv"
)

//func main() {
//	//fmt.Println("do we have a fun?")
//	printMe(nazwaInnegoTypu(10))
//	fmt.Println("/n")
//	printMe(nazwaInnegoTypu(100))
//
//}
//
//type nazwaTypu interface {
//	nazwaFuncji() string
//}
//
//func printMe(s nazwaTypu) {
//	fmt.Print(s.nazwaFuncji())
//}
//
//type nazwaInnegoTypu int
//
//func (s nazwaInnegoTypu) nazwaFuncji() string {
//	zmienna := "moja zmienna to: "
//	var t = strconv.Itoa(int(s))
//	return  zmienna + t
//}

/////////////////////////


func main() {
	funkcjaBioracaTypInterface(typInteger(999))
}

type typInterface interface {
	metodaWykonywalnaNaTypInteger() string
}
type typInteger int


func funkcjaBioracaTypInterface(newtypInterface typInterface) {
	fmt.Print(newtypInterface.metodaWykonywalnaNaTypInteger())
}

func (newTypInteger typInteger) metodaWykonywalnaNaTypInteger() string {
	zmienna := "moja zmienna to: "
	var t = strconv.Itoa(int(newTypInteger))
	return  zmienna + t
}