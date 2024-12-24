package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// UnpackageOfString unpacks the given string
func UnpackageOfString(str string) (string, error) {
	errorString := errors.New("incorrect string")
	answ := ""
	strRunes := []rune(str)
	// the letter which we want to unpack
	var letterToMultiply rune
	// counter of repetitions of unpacked letter
	numberOfRepetitions := -1
	var numberOfBackslashs int
	for i := 0; i < len(strRunes); i++ {
		letter := strRunes[i]
		// "\" case
		if letter == '\\' {
			if numberOfBackslashs+1 == 3 {
				answ += "\\"
			}
			numberOfBackslashs++
		} else if unicode.IsDigit(letter) {
			// "\number" case
			flagCountBackslash := false
			if numberOfBackslashs == 1 {
				numberOfRepetitionsBackSlash := 0
				flagCountBackslash = true
				i++
				// count number of repetitions
				for i < len(strRunes) {
					if unicode.IsDigit(strRunes[i]) {
						numberOfRepetitionsBackSlash *= 10
						val, err := strconv.Atoi(string(strRunes[i]))
						if err != nil {
							return "", nil
						}
						numberOfRepetitionsBackSlash += val
						i++
					} else {
						break
					}
				}
				if flagCountBackslash {
					i--
					flagCountBackslash = false
				}
				if numberOfRepetitionsBackSlash > 0 {
					for j := 0; j < numberOfRepetitionsBackSlash; j++ {
						answ += string(letter)
					}
					numberOfRepetitionsBackSlash = 0
					numberOfRepetitions = 1
				} else {
					answ += string(letter)
				}
				numberOfBackslashs = 0
				continue
				// "\\number" case
			} else if numberOfBackslashs == 2 {
				val, err := strconv.Atoi(string(letter))
				if err != nil {
					return "", err
				}
				for i := 0; i < val; i++ {
					answ += "\\"
				}
				numberOfBackslashs = 0
				continue
			}
			numberOfBackslashs = 0
			// begin of string
			if i == 0 {
				return "", errorString
			} else {
				if numberOfRepetitions == -1 {
					numberOfRepetitions = 0
				}
				numberOfRepetitions *= 10
				val, err := strconv.Atoi(string(letter))
				if err != nil {
					return "", err
				}
				numberOfRepetitions += val
			}
		} else {
			if numberOfBackslashs == 1 {
				answ += string(letter)
				numberOfBackslashs = 0
				continue
			}
			if numberOfRepetitions == 0 {
				answ = answ[:len(answ)-1]
			} else if numberOfRepetitions > 1 {
				for i := 0; i < numberOfRepetitions-1; i++ {
					answ += string(letterToMultiply)
				}
			}
			numberOfRepetitions = -1
			letterToMultiply = letter
			answ += string(letter)
		}
	}
	if numberOfRepetitions == 0 {
		answ = answ[:len(answ)-1]
	} else if numberOfRepetitions > 1 {
		for i := 0; i < numberOfRepetitions-1; i++ {
			answ += string(letterToMultiply)
		}
	}
	return answ, nil
}

func main() {
	val, err := UnpackageOfString("1q10p0a5g6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
