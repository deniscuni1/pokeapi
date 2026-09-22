package main
import "strings"
func cleanInput(text string) []string {
	lower_case := strings.ToLower(text)
	words:=strings.Fields(lower_case)
    list_to_give := []string{}
	for i:=0;i<len(words);i++ {
		list_to_give = append(list_to_give, words[i])
	}
	return list_to_give
}