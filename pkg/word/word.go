package word

import (
	"mygame/pkg/constants"
	"mygame/pkg/letter"
)

type Word struct {
	Letters []*letter.Letter
}

func (w *Word) Equals(other string) bool {
	for i, letter := range w.Letters {
		if letter.Char != []rune(other)[i] {
			return false
		}
	}
	return true
}

func (w *Word) ChangeColor(index int, color string) {
	w.Letters[index].Color = color
}

func NewWord(word string) *Word {
	w := &Word{}
	for _, char := range word {
		w.Letters = append(w.Letters, letter.NewLetter(char, constants.Gray))
	}
	return w
}
