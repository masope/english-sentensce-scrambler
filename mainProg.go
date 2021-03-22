package main

import (
	"fmt"
	"bufio"
	"os"
)

func str_stdInput(valMem *string){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	*valMem = sc.Text()
	return
}

func main(){
	var TargetSentence string
	var TranslatedSentence string
	
	fmt.Print("영어 문장을 입력해 주세요 >> ")
	str_stdInput(&TargetSentence)

	fmt.Print("한글 문장을 입력해 주세요 >> ")
	str_stdInput(&TranslatedSentence)


	fmt.Printf("%s and %s\n", TargetSentence, TranslatedSentence)
}

