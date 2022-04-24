package main

import (
	"flag"
	"fmt"
)

func transformationRune(s *string, myChannel chan []rune) {
	output := []rune(*s)
	myChannel <- output
}

func searching(input *[]rune, search *[]rune) bool {
	tmp := make([]rune, 0, len(*search)*2)
	for i := 0; i < len(*input) && len(tmp) != len(*search); i++ {
		for j := 0; j < len(*search); j++ {
			if (*input)[i] == (*search)[j] {
				tmp = append(tmp, (*search)[j])
				if i < len(*input)-1 {
					i++
				}
			} else {
				tmp = nil
				break
			}
		}
	}
	if len(tmp) != len(*search) {
		return false
	} else {
		return true
	}
}

func main() {
	var input, search string
	flag.StringVar(&input, "str", "", "set input")
	flag.StringVar(&search, "substr", "", "set search")
	flag.Parse()
	fmt.Printf("Предложение для поиска: %s.\nПодстрока которую ищем: %s\n", input, search)
	myChanelOne := make(chan []rune)
	myChanelTwo := make(chan []rune)
	go transformationRune(&input, myChanelOne)
	go transformationRune(&search, myChanelTwo)
	inputRune, searchRune := <-myChanelOne, <-myChanelTwo
	fmt.Println(searching(&inputRune, &searchRune))
}
