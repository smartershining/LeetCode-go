package main

import (
	"fmt"
	"strings"
)

func main() {

	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
func numUniqueEmails(emails []string) int {

	set := make(map[string]bool)

	for i := 0; i < len(emails); i++ {
		if _, ok := set[parse(emails[i])]; !ok {
			set[parse(emails[i])] = true
		}
	}
	return len(set)
}
func parse(email string) string {
	var result strings.Builder
	for i := 0; i < len(email); {
		switch email[i] {
		case '.':
			i++
		case '+':
			i = strings.IndexByte(email, '@')
		case '@':
			result.WriteString(email[i:])
			i = len(email)
		default:
			result.WriteByte(email[i])
			i++
		}
	}
	return result.String()
}
