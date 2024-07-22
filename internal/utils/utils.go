package utils

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strings"

	"github.com/Kur04ka/RandomPassword/internal/variables"
)

func CheckPasswordEntropy(userInput string) (entropy float64, attemptsToHack float64) {
	if variables.All_chars == "" {
		for key, value := range variables.Regexp_map {
			if consists, _ := regexp.MatchString(key, userInput); consists {
				variables.All_chars += value
			}
		}
	}
	entropy = CheckEntropy(len(userInput))
	attemptsToHack = math.Pow(2, entropy)
	return entropy, attemptsToHack
}

func ConvertToBool(str string) (bool, error) {
	if str == "Y" || str == "y" {
		return true, nil
	} else if str == "N" || str == "n" {
		return false, nil
	}
	return false, fmt.Errorf("failed to convert client answer to boolean")
}

func IsNeeded(yesOrNo, chars string) {
	if isNeeded, err := ConvertToBool(yesOrNo); isNeeded {
		variables.All_chars += chars
	} else if err != nil {
		log.Println(err)
	}
}

func CheckEntropy(passLength int) float64 {
	N := float64(len(variables.All_chars))
	L := float64(passLength)
	return L * math.Log2(N)
}

func GeneratePassword(all_chars []rune, passwordsLength int) string {
	var b strings.Builder

	rand.Shuffle(len(all_chars), func(i, j int) {
		all_chars[i], all_chars[j] = all_chars[j], all_chars[i]
	})

	// fmt.Printf("Array given: %s;\nPassword length: %d;\n", string(all_chars), passwordLength)

	for i := 0; i < passwordsLength; i++ {
		b.WriteRune(all_chars[rand.Intn(len(all_chars))])
	}
	return b.String()
}
