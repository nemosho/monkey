package lexer

type Lexer struct {
	input        string
	position     int  // Current position in input (points to the current character)
	readPosition int  // Position to read from now (next to current character)
	ch           byte // Character currently being inspected
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
