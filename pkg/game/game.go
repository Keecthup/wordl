package game

import (
	"fmt"
	"mygame/pkg/constants"
	"mygame/pkg/persistence"
	"mygame/pkg/word"
	"net/http"
	"text/template"
)
/// попытки и загаданное слово
var (
	attempts  []*word.Word
	rightWord string
)
/// сообщение пользователя и попытки
type PageData struct {
	Message  string
	Attempts []*word.Word
}
/// хэндлеры
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := PageData{Message: "", Attempts: GetAttempts()}
		tmpl.Execute(w, data)
		rightWord = persistence.GetRandomWord()
	} else if r.Method == "POST" {
		r.ParseForm()
		word2 := r.FormValue("word")

		message := ""
		if word2 == "" {
			message = "Пожалуйста, введите слово!"
		} else if len([]rune(word2)) != constants.WordLength {
			message = fmt.Sprintf("Слово должно быть длиной %d букв!", constants.WordLength)
		} else if !persistence.IsWordValid(word2) {
			message = "Слово не найдено в словаре!"
		} else {
			current := word.NewWord(word2)
			Compare(current, word.NewWord(rightWord))
			AddAttempt(current)

			if current.Equals(rightWord) {
				message = fmt.Sprintf("Поздравляем, вы отгадали слово: %s", rightWord)
				ResetGame()
			} else if len(GetAttempts()) >= constants.AttemptsNumber {
				message = fmt.Sprintf("Слово не угадано. Загаданное слово было: %s", rightWord)
				ResetGame()
			}
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		data := PageData{Message: message, Attempts: GetAttempts()}
		tmpl.Execute(w, data)
	}
}
/// получение попыток
func GetAttempts() []*word.Word {
	return attempts
}
/// добавление попыток
func AddAttempt(attempt *word.Word) {
	attempts = append(attempts, attempt)
}
/// рестарт игры
func ResetGame() {

	attempts = []*word.Word{}
	rightWord = persistence.GetRandomWord()
}
/// сравнение слов
func Compare(current, correct *word.Word) {
	usedIndices := make(map[int]bool)

	for i, letter := range current.Letters {
		if letter.Char == correct.Letters[i].Char {
			current.ChangeColor(i, "green")
			usedIndices[i] = true
		}
	}

	for i, letter := range current.Letters {
		if current.Letters[i].Color == constants.Gray {
			for j, correctLetter := range correct.Letters {
				if letter.Char == correctLetter.Char && !usedIndices[j] {
					current.ChangeColor(i, "yellow")
					usedIndices[j] = true
					break
				}
			}
		}
	}
}
