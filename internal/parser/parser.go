package parser

type Keyword int

const (
	Redirect Keyword = iota
	Chain
	Append
)

var keywords = map[Keyword]string{
	Redirect: ">",
	Chain:    "&",
	Append:   ">>",
}

type Token struct {
	lexeme   string
	literal  any
	position int
}

type Lexer struct {
	tokens                   []Token
	currentCharacterOfLexeme int
	firstCharacterOfLexeme   int
}
