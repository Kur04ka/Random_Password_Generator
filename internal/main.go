package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func main() {
	MainMenu()
}

// #region dialog

func MainMenu() {
	all_chars = ""
	fmt.Println("1) Сгенерировать рандомный пароль\n2) Проверить энтропию собственного пароля.\n3) Выход")
	fmt.Fscan(os.Stdin, &userInput)
	switch userInput {
	case "1":
		all_chars, passwordLength := DialogWindow()
		// fmt.Printf("Array given: %s;\nPassword length: %d;\n", all_chars, passwordLength)

		generatedPassword := generatePassword([]rune(all_chars), passwordLength)
		fmt.Printf("Random password: %s\n", generatedPassword)

		entropy, attemptsToHack := checkPasswordEntropy(generatedPassword)
		fmt.Printf("Entropy: %v, it means that you have to make %v passwords in order to hack the password\n",
			math.Floor(entropy), math.Floor(attemptsToHack))

		MainMenu()
	case "2":
		fmt.Println("Введите ваш пароль...")
		fmt.Fscan(os.Stdin, &userInput)

		entropy, attemptsToHack := checkPasswordEntropy(userInput)
		fmt.Printf("Entropy: %v, it means that you have to make %v passwords in order to hack the password\n",
			math.Floor(entropy), math.Floor(attemptsToHack))

		MainMenu()
	case "3":
		os.Exit(2)
	default:
		fmt.Println("Неверный ввод!")
		MainMenu()
	}
}

func DialogWindow() (string, int) {
	passwordLength = 0
	for str, chars := range dialog_map {
		fmt.Printf("Требуются ли в пароле %s? (Y/N)  ", str)
		fmt.Fscan(os.Stdin, &yesOrNo)
		isNeeded(yesOrNo, chars)
	}

	fmt.Print("Сколько символов требуется в пароле?  ")
	fmt.Fscan(os.Stdin, &passwordLength)
	if all_chars == "" {
		log.Println("There is no chars to make password from")
		DialogWindow()
	}
	if passwordLength <= 0 {
		log.Println("Password length should be more than 0 symbols")
		DialogWindow()
	}

	return all_chars, passwordLength
}

// #endregion

// #region auxiliary functions

func checkPasswordEntropy(userInput string) (entropy float64, attemptsToHack float64) {
	if all_chars == "" {
		for key, value := range regexp_map {
			if consists, _ := regexp.MatchString(key, userInput); consists {
				all_chars += value
			}
		}
	}
	entropy = checkEntropy(len(userInput))
	attemptsToHack = math.Pow(2, entropy)
	return entropy, attemptsToHack
}

func convertToBool(str string) (bool, error) {
	if str == "Y" || str == "y" {
		return true, nil
	} else if str == "N" || str == "n" {
		return false, nil
	}
	return false, fmt.Errorf("failed to convert client answer to boolean")
}

func isNeeded(yesOrNo, chars string) {
	if isNeeded, err := convertToBool(yesOrNo); isNeeded {
		all_chars += chars
	} else if err != nil {
		log.Println(err)
	}
}

func checkEntropy(passLength int) float64 {
	N := float64(len(all_chars))
	L := float64(passLength)
	return L * math.Log2(N)
}

func generatePassword(all_chars []rune, passwordsLength int) string {
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

// #endregion

// #region variables

var chars_EN, chars_en string = "ABCDEFGHIJKLMNOPQRRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"
var chars_specials, chars_digits string = "~=+%^*/()[]{}/!@#$?|", "0123456789"
var chars_RU, chars_ru string = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЭЮЯ", "абвгдеёжзийклмнопрстуфхцчшщэюя"
var yesOrNo, userInput, all_chars string
var passwordLength int
var regexp_map = map[string]string{
	`[A-Z]`: chars_EN, `[a-z]`: chars_en, `[А-Я]`: chars_RU, `[а-я]`: chars_ru,
	`\d`: chars_digits, `[~=\+\%\^\*/\(\)\[\]\{\}/!@#\$?\|]`: chars_specials,
}
var dialog_map = map[string]string{
	"A-Z": chars_EN, "a-z": chars_en, "А-Я": chars_RU, "а-я": chars_ru,
	"0-9": chars_digits, "~=+%^*/()[]{}/!@#$?|": chars_specials,
}

// #endregion
