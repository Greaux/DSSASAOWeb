package main

import (
	"fmt"
	"os"
)

func MainQuestions(url string) {

	var answer int

	fmt.Print("Необходимо пройти опрос\n")
	fmt.Print("К какому типу относится ваш сайт? \n")
	fmt.Print("Форум - 1, Облачное хранилище данных - 2, Онлайн-Банк - 3, Социальная сеть - 4. \n")
	fmt.Fscan(os.Stdin, &answer)

	if answer == 1 {
		Forum(url)
	}
	if answer == 2 {
		Cloud(url)
	}
	if answer == 3 {
		Bank(url)
	}
	if answer == 4 {
		SocialWeb(url)
	} else {
		fmt.Print("Ошибочный тип \n")
	}

}

func Forum(url string)  {
	var answer string
	WriteAnswerToData(url, "Type=Forum")

	fmt.Print("Ваш форум имеет фильтрацию контента? (От ссылок/по ключевым словам)\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Filter="+answer)

	fmt.Print("Ваш форум имеет ограничения для новых пользователей?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Limitations="+answer)

	fmt.Print("Ваш форум имеет достаточно компетентных модераторов?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Moderators="+answer)

	DefaultQuestList(url)
}

func Cloud(url string)  {
	var answer string
	WriteAnswerToData(url, "Type=Cloud")

	fmt.Print("Ваша платформа имеет достаточно компетентных модераторов?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Moderators="+answer)

	fmt.Print("Ваша платформа имеет ограничения для новых пользователей?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Limitations="+answer)

	DefaultQuestList(url)

}

func Bank(url string) {
	var answer string
	WriteAnswerToData(url, "Type=Bank")

	fmt.Print("Ваша платформа имеет достаточно компетентных сотрудников?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Moderators="+answer)

	fmt.Print("Ваша платформа имеет поведенческую биометрию?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Biometric="+answer)

	DefaultQuestList(url)
}

func SocialWeb(url string)  {
	var answer string
	WriteAnswerToData(url, "Type=SW")

	fmt.Print("Ваша платформа имеет достаточно компетентных сотрудников?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Moderators="+answer)

	fmt.Print("Ваша платформа имеет ограничения для новых пользователей?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Limitations="+answer)

	DefaultQuestList(url)
}

func DefaultQuestList(url string)  {
	var answer string

	fmt.Print("Ваша платформа имеет Captcha?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Captcha="+answer)

	fmt.Print("Ваш форум имеет правила на качество пароля?\n")
	fmt.Print("Нет/Не знаю - 0, Да - 1. \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "Password="+answer)

	fmt.Print("Ваша платформа имеет 2FA (Двухфакторную авторизацию)? Вход по SMS/E-Mail/GoogleAuth/Token) \n")
	fmt.Print("Нет/Не знаю - 0, SMS - 1, E-Mail - 2, GoogleAuth - 3, Token - 4 \n")
	fmt.Fscan(os.Stdin, &answer)
	WriteAnswerToData(url, "2FA=" + answer)
}