package lexer

import (
	"errors"
	"log"
)

const (
	JSON_QUOTE = '"'
)

func lexString(str string) (string, string, error) {
	jStr := ""

	if str[0] == JSON_QUOTE {
		str = str[1:]
	} else {
		return "", str, nil
	}

	for _, char := range str {
		if char == JSON_QUOTE {
			return jStr, str[len(str)+1:], nil
		} else {
			jStr = jStr + string(char)
		}
	}
	return "", "", errors.New("Expected end of quote")
}

func lex(str string) {
	tokens := []interface{}{}

	for len(str) != 0 {
		jString, rStr, errLexStr := lexString(str)
		if errLexStr != nil {
			log.Fatal(errLexStr)
		}
		str = rStr
		tokens = append(tokens, jString)
	}

}
