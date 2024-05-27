package word

import (
	"mygame/pkg/constants"
	"mygame/pkg/letter"
)
/// структура, описывающая слово
type Word struct {
	Letters []*letter.Letter
}
/// функция, отвечающая за сравнение полученного слова с загаданным
func (w *Word) Equals(other string) bool {
	for i, letter := range w.Letters {
		if letter.Char != []rune(other)[i] {
			return false
		}
	}
	return true
}
/// функция, меняющая цвет символов
func (w *Word) ChangeColor(index int, color string) {
	w.Letters[index].Color = color
}
/// функция добавляет новое слово
func NewWord(word string) *Word {
	w := &Word{}
	for _, char := range word {
		w.Letters = append(w.Letters, letter.NewLetter(char, constants.Gray))
	}
	return w
}
