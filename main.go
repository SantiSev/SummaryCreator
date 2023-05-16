package main

import (
	"fmt"
	gpt "gpt_summary_creator/openaiApi"
)

func main() {
	res := gpt.FetchGPT("how to make a summary of a text")
	fmt.Println(res)
}
