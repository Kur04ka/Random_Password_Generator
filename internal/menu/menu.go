package menu

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/Kur04ka/RandomPassword/internal/utils"
	"github.com/Kur04ka/RandomPassword/internal/variables"
)

func MainMenu() {
	variables.All_chars = ""
	fmt.Println("1) Сгенерировать рандомный пароль\n2) Проверить энтропию собственного пароля.\n3) Выход")
	fmt.Fscan(os.Stdin, &variables.UserInput)
	switch variables.UserInput {
	case "1":
		all_chars, passwordLength := DialogWindow()
		// fmt.Printf("Array given: %s;\nPassword length: %d;\n", all_chars, passwordLength)

		generatedPassword := utils.GeneratePassword([]rune(all_chars), passwordLength)
		fmt.Printf("Random password: %s\n", generatedPassword)

		entropy, attemptsToHack := utils.CheckPasswordEntropy(generatedPassword)
		fmt.Printf("Entropy: %v, it means that you have to make %v passwords in order to hack the password\n",
			math.Floor(entropy), math.Floor(attemptsToHack))

		MainMenu()
	case "2":
		fmt.Println("Введите ваш пароль...")
		fmt.Fscan(os.Stdin, &variables.UserInput)

		entropy, attemptsToHack := utils.CheckPasswordEntropy(variables.UserInput)
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
	variables.PasswordLength = 0
	for str, chars := range variables.Dialog_map {
		fmt.Printf("Требуются ли в пароле %s? (Y/N)  ", str)
		fmt.Fscan(os.Stdin, &variables.YesOrNo)
		utils.IsNeeded(variables.YesOrNo, chars)
	}

	fmt.Print("Сколько символов требуется в пароле?  ")
	fmt.Fscan(os.Stdin, &variables.PasswordLength)
	if variables.All_chars == "" {
		log.Println("There is no chars to make password from")
		DialogWindow()
	}
	if variables.PasswordLength <= 0 {
		log.Println("Password length should be more than 0 symbols")
		DialogWindow()
	}

	return variables.All_chars, variables.PasswordLength
}
