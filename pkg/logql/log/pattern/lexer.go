package pattern

type lexer struct {
	data        []byte
	p, pe, cs   int
	ts, te, act int

	lastnewline int
	curline     int

	errs []parseError
	expr []node
}

func newLexer() *lexer {
	lex := &lexer{}
	lex.init()
	return lex
}

func (lex *lexer) setData(data []byte) {
	lex.data = data
	lex.pe = len(data)
	lex.lastnewline = -1
	lex.curline = 1
}

// Error implements exprLexer interface generated by yacc (yyLexer)
func (lex *lexer) Error(e string) {
	lex.errs = append(lex.errs, newParseError(e, lex.curline, lex.curcol()))
}

// curcol calculates the current token's start column based on the last newline position
// returns a 1-indexed value
func (lex *lexer) curcol() int {
	return (lex.ts + 1 /* 1-indexed columns */) - (lex.lastnewline + 1 /* next after newline */)
}

func (lex *lexer) handle(token int, err error) int {
	if err != nil {
		lex.Error(err.Error())
		return LEXER_ERROR
	}
	return token
}

func (lex *lexer) token() string {
	return string(lex.data[lex.ts:lex.te])
}

// nolint
func (lex *lexer) identifier(out *exprSymType) (int, error) {
	t := lex.token()
	out.str = t[1 : len(t)-1]
	return IDENTIFIER, nil
}

// nolint
func (lex *lexer) literal(out *exprSymType) (int, error) {
	out.literal = rune(lex.data[lex.ts])
	return LITERAL, nil
}