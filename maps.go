package main

import (
	"fmt"
	"strings"
)

type Books struct {
	author string
	title  string
	price  float64
}

var testMap map[string]Books

func main() {
	testMap = make(map[string]Books)
	testMap["White Guard"] = Books{
		"Bulgakov", "White Guard", 20.22,
	}

	testMap["War and Peace"] = Books{
		"Leo Tolstoy", "War and Peace", 1488.228,
	}

	fmt.Println(testMap["White Guard"])

	//_________________________________________//

	productQuantity := make(map[string]int)
	//insert into map
	productQuantity["Avocado"] = 12
	productQuantity["VODKA"] = 0
	productQuantity["Pen"] = 819

	//retrieve from map
	avocadoQuantity := productQuantity["Avocado"]
	fmt.Println(avocadoQuantity)

	// delete an element
	delete(productQuantity, "Avocado")
	fmt.Println(productQuantity["Avocado"])

	// test that a key is present
	value, result := productQuantity["Pen"]
	fmt.Println("value is ", value, "present: ", result)

	value, result = productQuantity["ASDASDAS"]
	fmt.Println("value is ", value, "present: ", result)

	stringToCount := "God save our glorious King and make them GOIDA"
	fmt.Println(charCount(stringToCount))

	stringToWordCount := "I have to say that implemented method is not convenient for society method I have to say"
	fmt.Println(wordCount(stringToWordCount))
}

func charCount(s string) map[string]int {
	alphabet := make(map[string]int)
	for i := 0; i < len(s); i++ {
		alphabet[string(s[i])] += 1
	}
	return alphabet
}

func wordCount(s string) map[string]int {
	wordsArray := strings.Fields(s)
	wordsMap := make(map[string]int)
	for i := 0; i < len(wordsArray); i++ {
		wordsMap[string(wordsArray[i])] += 1
	}
	return wordsMap
}
