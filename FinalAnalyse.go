package main

import (
	"bufio"
	"log"
	"os"
)

func FinalAnalyse(url string) {
	file, err := os.Open(url + ".data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	WriteResultData(url, "Рекомендации:")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch str := scanner.Text(); str {
		case "HTTPS=0":
			WriteResultData(url, "Настоятельно рекомендуется включить принудительное перенаправление на https версию сайта, это поможет пользователям в защите их данных при использовании публичных точек доступа а также так называемый зелёный замочек будет показывать пользователям что скорее всего они находятся на достоверном сайте")
		case "ExpiresDate error":
			WriteResultData(url, "Пожалуйста, продлевайте ваш домен хотя бы за 2 месяца до окончания срока его регистрации, это поможет вам избежать разных форс-мажоров а так же кражи домена злоумышленниками")
		case "DomainName error":
			WriteResultData(url, "Вполне возможно что ваш домен имеет перенаправление на другой домен, если это так то рекомендуется либо использовать тот домен на который происходит перенаправление, либо проигнорируйте данное уведомление.")
		case "Type=Forum":
			WriteResultData(url, "Тип - форум")
			scanner.Scan()
			filter := scanner.Text()
			if filter == "Filter=0" {
				WriteResultData(url, "Наличие фильтрации контента для форума является обязательным, рекомендуется добавить на платформу хотя бы простой спам фильтр")
			}

			scanner.Scan()
			limits := scanner.Text()
			if limits == "Limitations=0" {
				WriteResultData(url, "Ограничения для новых пользователей рекомендуется добавлять на случай если ваша платформа имеет достаточно крупную базу пользователей.")
			}

			scanner.Scan()
			mods := scanner.Text()
			if mods == "Moderators=0" {
				WriteResultData(url, "Модераторы являются важной частью защиты платформы, потому что только человек способен проанализировать то что написано другим человеком достаточно достоверно при наличии критического мышления и достаточной информационной культуры")
			}

			scanner.Scan()
			captcha := scanner.Text()
			if captcha == "Captcha=0" {
				WriteResultData(url, "Добавление капчи на плафтрому позволяет уменьшить уровень угрозы от ботов. Очень рекомендуется добавить её на платформу.")
			}

			scanner.Scan()
			pass := scanner.Text()
			if pass == "Password=0" {
				WriteResultData(url, "Ограничения на количество символов и качество пороля очень важно для защиты пользоватей. Рекомендуется принуждать пользователей использовать спец символы, цифры и разный регистр букв в одном пароле, а так же длинну пароля от 10 символов")
			}

			scanner.Scan()
			twofa := scanner.Text()
			if twofa == "2FA=0" {
				WriteResultData(url, "2FA не является обязательным, однако если ваша платформа является масштабной, то желательно добавить возможность 2FA авторизации. Рекомендуется выбирать Email или Google Auth.")
			}
			break

		case "Type=Cloud":
			WriteResultData(url, "Тип - облако")
			break

		case "Type=Bank":
			WriteResultData(url, "Тип - банк")
			break

		case "Type=SW":
			WriteResultData(url, "Тип - социальная сеть")
			break

		}


	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}