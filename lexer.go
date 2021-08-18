package symc

import (
	"fmt"
	"strings"
)

type Lexer struct {
	input string
	pos   int
}

type Token struct {
	tokenType int
	literal   string
}

const (
	eof = iota
	word
	integer
	float
	assign
	plus
	minus
	bang
	asterisk
	slash
	percent
	lt
	gt
	eq
	ne
	gteq
	lteq
	semicolon
	lparen
	rparen
	comma
	lbrace
	rbrace
	lbracket
	rbracket
	ampersand
	tilde
	caret
	vertical
	colon
	question
	period
	backslash
	str
	letter
	arrow
	leftShift
	rightShift
	increment
	decrement
	and
	or
	plusAssigne
	minusAssigne
	asteriskAssigne
	slashAssigne
	verticalAssigne
	ampersandAssigne
	leftShiftAssigne
	rightShiftAssigne
	tildeAssigne
	caretAssigne
	percentAssigne
	keyReturn
	keyIf
	keyElse
	keyWhile
	keyDo
	keyGoto
	keyFor
	keyBreak
	keyContinue
	keySwitch
	keyCase
	keyDefault
	keyExtern
	keyVolatile
	keyConst
	keyTypedef
	keyUnion
	keyStruct
	keyEnum
	keyAttribute
	keyVoid
	keyAsm
	keySizeof
	comment
	illegal
)

func (t *Token) String() string {
	return fmt.Sprintf("tokenType:%v, literal:%s", t.tokenType, t.literal)
}

func NewLexer(src string) *Lexer {
	return &Lexer{input: src, pos: 0}
}

func (l *Lexer) lexicalize() []*Token {
	ts := []*Token{}
	for {
		t := l.nextToken()
		ts = append(ts, t)
		if t.tokenType == eof {
			break
		}
	}
	return ts
}

func (l *Lexer) nextToken() *Token {
	// スペースをとばす
	for {
		i := l.pos
		if i >= len(l.input) {
			break
		}
		c := l.input[i]
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			break
		}
		l.pos++
	}

	// ソースの終端
	if l.pos >= len(l.input) {
		return &Token{tokenType: eof, literal: "eof"}
	}

	var tk *Token
	c := l.input[l.pos]
	switch c {
	case '=':
		tk = &Token{tokenType: assign, literal: "="}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: eq, literal: "=="}
			l.pos++
		}
	case '+':
		tk = &Token{tokenType: plus, literal: "+"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '+' {
			tk = &Token{tokenType: increment, literal: "++"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: plusAssigne, literal: "+="}
			l.pos++
		}
	case '-':
		tk = &Token{tokenType: minus, literal: "-"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			// ->
			tk = &Token{tokenType: arrow, literal: "->"}
			l.pos++
		} else if l.input[l.pos] == '-' {
			tk = &Token{tokenType: decrement, literal: "--"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: minusAssigne, literal: "-="}
			l.pos++
		}
	case '!':
		tk = &Token{tokenType: bang, literal: "!"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: ne, literal: "!="}
			l.pos++
		}
	case '*':
		tk = &Token{tokenType: asterisk, literal: "*"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: asteriskAssigne, literal: "*="}
			l.pos++
		}
	case '/':
		tk = &Token{tokenType: slash, literal: "/"}
		l.pos++
		if l.pos >= len(l.input) {
			// 何もしない
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: slashAssigne, literal: "/="}
			l.pos++
		} else if l.input[l.pos] == '*' {
			// comment
			l.pos++
			com := l.readComment()
			tk = &Token{tokenType: comment, literal: com}
		}
	case '<':
		tk = &Token{tokenType: lt, literal: "<"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '<' {
			tk = &Token{tokenType: leftShift, literal: "<<"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{tokenType: leftShiftAssigne, literal: "<<="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: lteq, literal: "<="}
			l.pos++
		}
	case '>':
		tk = &Token{tokenType: gt, literal: ">"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			tk = &Token{tokenType: rightShift, literal: ">>"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{tokenType: rightShiftAssigne, literal: ">>="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: gteq, literal: ">="}
			l.pos++
		}
	case ';':
		tk = &Token{tokenType: semicolon, literal: ";"}
		l.pos++
	case '(':
		tk = &Token{tokenType: lparen, literal: "("}
		l.pos++
	case ')':
		tk = &Token{tokenType: rparen, literal: ")"}
		l.pos++
	case ',':
		tk = &Token{tokenType: comma, literal: ","}
		l.pos++
	case '{':
		tk = &Token{tokenType: lbrace, literal: "{"}
		l.pos++
	case '}':
		tk = &Token{tokenType: rbrace, literal: "}"}
		l.pos++
	case '[':
		tk = &Token{tokenType: lbracket, literal: "["}
		l.pos++
	case ']':
		tk = &Token{tokenType: rbracket, literal: "]"}
		l.pos++
	case '&':
		tk = &Token{tokenType: ampersand, literal: "&"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '&' {
			tk = &Token{tokenType: and, literal: "&&"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: ampersandAssigne, literal: "&="}
			l.pos++
		}
	case '~':
		tk = &Token{tokenType: tilde, literal: "~"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: tildeAssigne, literal: "~="}
			l.pos++
		}
	case '^':
		tk = &Token{tokenType: caret, literal: "^"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: caretAssigne, literal: "^="}
			l.pos++
		}
	case '|':
		tk = &Token{tokenType: vertical, literal: "|"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '|' {
			tk = &Token{tokenType: or, literal: "||"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: verticalAssigne, literal: "|="}
			l.pos++
		}
	case '%':
		tk = &Token{tokenType: percent, literal: "%"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{tokenType: percentAssigne, literal: "%="}
			l.pos++
		}
	case ':':
		tk = &Token{tokenType: colon, literal: ":"}
		l.pos++
	case '?':
		tk = &Token{tokenType: question, literal: "?"}
		l.pos++
	case '.':
		tk = &Token{tokenType: period, literal: "."}
		l.pos++
	case '\\':
		tk = &Token{tokenType: backslash, literal: "\\"}
		l.pos++
	case '\'':
		tk = l.readLetter()
	case '"':
		tk = l.readString()
	case '#':
		tk = l.readHashComment()
		l.pos++
	default:
		if isLetter(c) {
			tk = l.readWord()
		} else if isDec(c) {
			tk = l.readNumber()
		}
	}
	return tk
}

func (l *Lexer) readWord() *Token {
	// ワードの終わりの次まで pos を進める
	var next int
	for next = l.pos; next < len(l.input); next++ {
		c := l.input[next]
		if !isLetter(c) && !isDec(c) {
			break
		}
	}
	w := l.input[l.pos:next]
	tk := l.determineKeyword(w)
	l.pos = next
	return tk
}

func (l *Lexer) readNumber() *Token {
	var next int
	isFloat := false

	next = l.pos
	c := l.input[next]
	if c == '0' {
		next++
		c = l.input[next]
		switch c {
		case 'x':
			// 16進数
			next++
		case 'b':
			// 2進数
			next++
		case '.':
			// 小数
			next++
			isFloat = true
		default:
			if isDec(c) {
				// 8進数
				next++
			} else {
				// ゼロ

			}
		}
	}

	// ワードの終わりの次まで pos を進める
	for ; next < len(l.input); next++ {
		c = l.input[next]
		if c == '.' {
			isFloat = true
		} else if c == 'u' || c == 'U' || c == 'l' || c == 'L' {
			continue
		} else if !isHex(c) {
			break
		}
	}
	w := l.input[l.pos:next]
	l.pos = next

	var tk *Token
	if isFloat {
		tk = &Token{tokenType: float, literal: w}
	} else {
		tk = &Token{tokenType: integer, literal: w}
	}
	return tk
}

func (l *Lexer) readString() *Token {
	var next int

	// 次の " を探す
	for next = l.pos + 1; next < len(l.input); next++ {
		// エスケープシーケンス考慮
		if l.input[next] == '\\' && l.input[next+1] == '\\' {
			next++
		} else if l.input[next] == '\\' && l.input[next+1] == '"' {
			next++
		} else if l.input[next] == '"' {
			break
		}
	}
	// 次の pos に進める
	next++
	w := l.input[l.pos:next]
	l.pos = next
	return &Token{tokenType: str, literal: w}
}

func (l *Lexer) readHashComment() *Token {
	// # の次の文字に移動
	l.pos++
	var next int
	for i := l.pos; i <= len(l.input); i++ {
		next = i
		if next >= len(l.input) {
			break
		}
		c := l.input[next]
		if c == '\n' || c == '\r' {
			break
		}
	}
	tk := &Token{tokenType: comment, literal: l.input[l.pos:next]}
	l.pos = next
	return tk
}

func (l *Lexer) readLetter() *Token {

	l.pos++
	var s []byte
	c := l.input[l.pos]
	if c == '\\' {
		s = l.getEscC()
		l.pos++
	} else {
		s = append(s, c)
		l.pos++
		l.pos++
	}
	return &Token{tokenType: letter, literal: string(s)}
}

func (l *Lexer) getEscC() []byte {
	res := []byte{}
	res = append(res, l.input[l.pos])
	l.pos++
	if l.input[l.pos] >= '0' && l.input[l.pos] <= '9' {
		for l.input[l.pos] >= '0' && l.input[l.pos] <= '9' {
			res = append(res, l.input[l.pos])
			l.pos++
		}
	} else {
		res = append(res, l.input[l.pos])
		l.pos++
	}

	return res
}

func (l *Lexer) newIllegal() *Token {
	tk := &Token{tokenType: illegal, literal: l.input[l.pos:]}
	l.pos = len(l.input)
	return tk
}

func (l *Lexer) determineKeyword(w string) *Token {
	if strings.Compare("return", w) == 0 {
		return &Token{tokenType: keyReturn, literal: w}
	} else if strings.Compare("if", w) == 0 {
		return &Token{tokenType: keyIf, literal: w}
	} else if strings.Compare("else", w) == 0 {
		return &Token{tokenType: keyElse, literal: w}
	} else if strings.Compare("while", w) == 0 {
		return &Token{tokenType: keyWhile, literal: w}
	} else if strings.Compare("do", w) == 0 {
		return &Token{tokenType: keyDo, literal: w}
	} else if strings.Compare("goto", w) == 0 {
		return &Token{tokenType: keyGoto, literal: w}
	} else if strings.Compare("for", w) == 0 {
		return &Token{tokenType: keyFor, literal: w}
	} else if strings.Compare("break", w) == 0 {
		return &Token{tokenType: keyBreak, literal: w}
	} else if strings.Compare("continue", w) == 0 {
		return &Token{tokenType: keyContinue, literal: w}
	} else if strings.Compare("switch", w) == 0 {
		return &Token{tokenType: keySwitch, literal: w}
	} else if strings.Compare("case", w) == 0 {
		return &Token{tokenType: keyCase, literal: w}
	} else if strings.Compare("default", w) == 0 {
		return &Token{tokenType: keyDefault, literal: w}
	} else if strings.Compare("extern", w) == 0 {
		return &Token{tokenType: keyExtern, literal: w}
	} else if strings.Compare("volatile", w) == 0 {
		return &Token{tokenType: keyVolatile, literal: w}
	} else if strings.Compare("const", w) == 0 {
		return &Token{tokenType: keyConst, literal: w}
	} else if strings.Compare("typedef", w) == 0 {
		return &Token{tokenType: keyTypedef, literal: w}
	} else if strings.Compare("union", w) == 0 {
		return &Token{tokenType: keyUnion, literal: w}
	} else if strings.Compare("struct", w) == 0 {
		return &Token{tokenType: keyStruct, literal: w}
	} else if strings.Compare("enum", w) == 0 {
		return &Token{tokenType: keyEnum, literal: w}
	} else if strings.Compare("__attribute__", w) == 0 {
		return &Token{tokenType: keyAttribute, literal: w}
	} else if strings.Compare("void", w) == 0 {
		return &Token{tokenType: keyVoid, literal: w}
	} else if strings.Compare("__asm", w) == 0 {
		return &Token{tokenType: keyAsm, literal: w}
	} else if strings.Compare("sizeof", w) == 0 {
		return &Token{tokenType: keySizeof, literal: w}
	} else {
		return &Token{tokenType: word, literal: w}
	}
}

func (l *Lexer) readComment() string {
	res := []byte{}

	for !l.isCommentEnd() {
		res = append(res, l.input[l.pos])
		l.pos++
	}
	l.pos++
	l.pos++
	// next

	return string(res)
}

func (l *Lexer) isCommentEnd() bool {
	// */ か確認
	return l.input[l.pos] == '*' && l.input[l.pos+1] == '/'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isHex(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

func isDec(c byte) bool {
	return '0' <= c && c <= '9'
}

// isTypeToken
func (t *Token) isTypeToken() bool {

	switch t.tokenType {
	case word:
	case asterisk:
	case keyConst:
	case keyVoid:
	case keyStruct:
	case keyUnion:
	case keyEnum:
	case keyVolatile:
	case caret:
		// clang でコンパイルした場合型の種類に^が含まれる？
	default:
		return false
	}

	return true
}

func (t *Token) isOperator() bool {
	switch t.tokenType {
	case assign:
	case plus:
	case minus:
	case asterisk:
	case slash:
	case lt:
	case gt:
	case eq:
	case gteq:
	case lteq:
	case ne:
	case ampersand:
	case tilde:
	case caret:
	case vertical:
	case question:
	case leftShift:
	case rightShift:
	case increment:
	case decrement:
	case or:
	case and:
	case percent:
	case colon:
	case plusAssigne:
	case minusAssigne:
	case asteriskAssigne:
	case slashAssigne:
	case verticalAssigne:
	case ampersandAssigne:
	case leftShiftAssigne:
	case rightShiftAssigne:
	case tildeAssigne:
	case caretAssigne:
	case percentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) isPrefixExpression() bool {
	switch t.tokenType {
	case minus:
	case plus:
	case increment:
	case decrement:
	case tilde:
	case bang:
	case asterisk:
	case ampersand:
	default:
		return false
	}
	return true
}

func (t *Token) isPostExpression() bool {
	switch t.tokenType {
	case lparen:
	case increment:
	case decrement:
	default:
		return false
	}
	return true
}

func (t *Token) isCompoundOp() bool {
	switch t.tokenType {
	case plusAssigne:
	case minusAssigne:
	case asteriskAssigne:
	case slashAssigne:
	case verticalAssigne:
	case ampersandAssigne:
	case leftShiftAssigne:
	case rightShiftAssigne:
	case tildeAssigne:
	case caretAssigne:
	case percentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) isToken(t2 int) bool {
	return t.tokenType == t2
}
