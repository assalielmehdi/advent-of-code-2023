package util_io

import (
	"errors"
)

type Tokenizer struct {
	content []byte
	cursor  int
}

func NewTokenizer(content []byte) *Tokenizer {
	return &Tokenizer{
		content: content,
		cursor:  0,
	}
}

func (tkzr *Tokenizer) moveToNext() {
	for tkzr.cursor < len(tkzr.content) && tkzr.content[tkzr.cursor] == ' ' {
		tkzr.cursor++
	}
}

func (tkzr *Tokenizer) HasNext() bool {
	tkzr.moveToNext()
	return tkzr.cursor < len(tkzr.content)
}

func (tkzr *Tokenizer) Next() ([]byte, error) {
	tkzr.moveToNext()

	if !tkzr.HasNext() {
		return nil, errors.New("no more token")
	}

	token := make([]byte, 0)

	for tkzr.cursor < len(tkzr.content) && tkzr.content[tkzr.cursor] != ' ' {
		token = append(token, tkzr.content[tkzr.cursor])
		tkzr.cursor++
	}

	return token, nil
}
