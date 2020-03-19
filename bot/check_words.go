package bot

import (
	"log"
	"regexp"
	s "strings"
)

// checkWords find in text a key words
func checkWords(text string) bool {

	var keyWords = "вест lada lаda ladа vest vеst bест beст becт"

	text = s.ToLower(text)

	reg, err := regexp.Compile("[^а-яА-Яa-zA-Z]+")
	if err != nil {
		log.Fatalln(err)
	}

	clearText := reg.ReplaceAllString(text, " ")

	words := s.Fields(keyWords)
	arr := s.Fields(clearText)

	for _, keyWord := range words {
		re := regexp.MustCompile(keyWord + `.{1,6}`)
		for _, target := range arr {
			if len(re.FindString(s.ToLower(target))) > 0 {
				log.Printf("Found template %v string %v \n", re, re.FindString(target))
				return true
			}
		}
	}
	return false
}
