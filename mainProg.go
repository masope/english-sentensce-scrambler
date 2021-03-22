package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func StrStdInput(valMem *string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	*valMem = sc.Text()
	return
}

func StrArrShuffle(arr []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func WordsToStr(wordsArr []string) string {
	var sentence string
	sentence = ""
	for indx, word := range wordsArr {
		if indx == 0 {
			sentence = "( "
			sentence += word
		}else {
			sentence += " , " + word
		}
	}; sentence += " )"
	return sentence
}

func EvaluateExit(value string) bool {
	if value == "exit" || value == "0" {
		return true
	} else {
		return false
	}
} //Do I Really Need This Function??

func main() {
	fmt.Println("Remember! : Enter \"exit\" or integer 0 to terminate the program.")
	for ;; {
		var TargetSentence string
		var TranslatedSentence string
	
		fmt.Print("영어 문장을 입력해 주세요 >> ")
		StrStdInput(&TargetSentence)
		if EvaluateExit(TargetSentence) { break }
	
		fmt.Print("한글 문장을 입력해 주세요 >> ")
		StrStdInput(&TranslatedSentence)
		if EvaluateExit(TranslatedSentence) { break }
	
		TargetWordsArr := strings.Fields(TargetSentence)
		StrArrShuffle(TargetWordsArr)
	
		ScrambledTargetSentence := WordsToStr(TargetWordsArr) 
		fmt.Printf("%s\n%s", TranslatedSentence, ScrambledTargetSentence)
	
		
		fmt.Printf("\n\n\n")
	}
}
