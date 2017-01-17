package base

import (
	"log"
	"regexp"
)

func CheckUsername(name string) bool {
	pattern := `^[a-zA-Z][0-9a-zA-Z_]{3,19}$`
	matched, err := regexp.MatchString(pattern, name)
	if err != nil {
		log.Println(err)
		return false
	}
	return matched
}

func CheckEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		log.Println(err)
		return false
	}
	return matched
}

// TODO
func CheckPhone(phone string) bool {
	return true
}
