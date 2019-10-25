package lexer

type Lexer struct {
	input        string
	position     int  // Current position in input (points to the current character)
	readPosition int  // Position to read from now (next to current character)
	ch           byte // Character currently being inspected
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
