package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//some fake comment

func main() {
	fmt.Println("Podaj liczbe:")

	reader := bufio.NewReader(os.Stdin)
	imput, dupa := reader.ReadString('\n')

	if dupa == nil {
		fmt.Println("Twoja liczba to: ", imput)
	} else {
		log.Fatal(dupa)
	}
}

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

//func main() {
//	funkcja1(typInterface("I to jest obiektowosc"))
//}
//
//type typInterface interface {
//	metodaWykonywalnaNaTypInteger() string
//}
//type typString string
//
//
//func funkcja1(newtypInterface typInterface) {
//	fmt.Print(newtypInterface.metodaWykonywalnaNaTypInteger())
//}
//
//func (newTypString typString) metodaWykonywalnaNaTypInteger() typString {
//	return  newTypString
//}
