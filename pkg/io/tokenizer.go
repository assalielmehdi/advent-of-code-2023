package util_io

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

func NewEmptyTokenizer() *Tokenizer {
	return &Tokenizer{
		content: nil,
		cursor:  0,
	}
}

func (tkzr *Tokenizer) SetContent(content []byte) {
	tkzr.content = content
	tkzr.cursor = 0
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

func (tkzr *Tokenizer) Next() []byte {
	tkzr.moveToNext()

	if !tkzr.HasNext() {
		return nil
	}

	token := make([]byte, 0)

	for tkzr.cursor < len(tkzr.content) && tkzr.content[tkzr.cursor] != ' ' {
		token = append(token, tkzr.content[tkzr.cursor])
		tkzr.cursor++
	}

	return token
}
