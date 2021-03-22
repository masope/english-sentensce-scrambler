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

func main() {
	var TargetSentence string
	var TranslatedSentence string

	fmt.Print("영어 문장을 입력해 주세요 >> ")
	StrStdInput(&TargetSentence)

	fmt.Print("한글 문장을 입력해 주세요 >> ")
	StrStdInput(&TranslatedSentence)

	TargetWordsArr := strings.Fields(TargetSentence)
	StrArrShuffle(TargetWordsArr)

	ScrambledTargetSentence := WordsToStr(TargetWordsArr) 
	fmt.Printf("%s\n%s", TranslatedSentence, ScrambledTargetSentence)
}
