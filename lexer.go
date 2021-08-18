package clanglex

import (
	"fmt"
	"strings"
)

type Lexer struct {
	input string
	pos   int
}

type Token struct {
	TokenType int
	Literal   string
}

const (
	Eof = iota
	Word
	Integer
	Float
	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash
	Percent
	Lt
	Gt
	Eq
	Ne
	Gteq
	Lteq
	Semicolon
	Lparen
	Rparen
	Comma
	Lbrace
	Rbrace
	Lbracket
	Rbracket
	Ampersand
	Tilde
	Caret
	Vertical
	Colon
	Question
	Period
	Backslash
	Str
	Letter
	Arrow
	LeftShift
	RightShift
	Increment
	Decrement
	And
	Or
	PlusAssigne
	MinusAssigne
	AsteriskAssigne
	SlashAssigne
	VerticalAssigne
	AmpersandAssigne
	LeftShiftAssigne
	RightShiftAssigne
	TildeAssigne
	CaretAssigne
	PercentAssigne
	KeyReturn
	KeyIf
	KeyElse
	KeyWhile
	KeyDo
	KeyGoto
	KeyFor
	KeyBreak
	KeyContinue
	KeySwitch
	KeyCase
	KeyDefault
	KeyExtern
	KeyVolatile
	KeyConst
	KeyTypedef
	KeyUnion
	KeyStruct
	KeyEnum
	KeyAttribute
	KeyVoid
	KeyAsm
	KeySizeof
	Comment
	Illegal
)

func (t *Token) String() string {
	return fmt.Sprintf("tokenType:%v, literal:%s", t.TokenType, t.Literal)
}

func NewLexer(src string) *Lexer {
	return &Lexer{input: src, pos: 0}
}

func (l *Lexer) lexicalize() []*Token {
	ts := []*Token{}
	for {
		t := l.nextToken()
		ts = append(ts, t)
		if t.TokenType == Eof {
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
		return &Token{TokenType: Eof, Literal: "eof"}
	}

	var tk *Token
	c := l.input[l.pos]
	switch c {
	case '=':
		tk = &Token{TokenType: Assign, Literal: "="}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: Eq, Literal: "=="}
			l.pos++
		}
	case '+':
		tk = &Token{TokenType: Plus, Literal: "+"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '+' {
			tk = &Token{TokenType: Increment, Literal: "++"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: PlusAssigne, Literal: "+="}
			l.pos++
		}
	case '-':
		tk = &Token{TokenType: Minus, Literal: "-"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			// ->
			tk = &Token{TokenType: Arrow, Literal: "->"}
			l.pos++
		} else if l.input[l.pos] == '-' {
			tk = &Token{TokenType: Decrement, Literal: "--"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: MinusAssigne, Literal: "-="}
			l.pos++
		}
	case '!':
		tk = &Token{TokenType: Bang, Literal: "!"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: Ne, Literal: "!="}
			l.pos++
		}
	case '*':
		tk = &Token{TokenType: Asterisk, Literal: "*"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: AsteriskAssigne, Literal: "*="}
			l.pos++
		}
	case '/':
		tk = &Token{TokenType: Slash, Literal: "/"}
		l.pos++
		if l.pos >= len(l.input) {
			// 何もしない
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: SlashAssigne, Literal: "/="}
			l.pos++
		} else if l.input[l.pos] == '*' {
			// comment
			l.pos++
			com := l.readComment()
			tk = &Token{TokenType: Comment, Literal: com}
		}
	case '<':
		tk = &Token{TokenType: Lt, Literal: "<"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '<' {
			tk = &Token{TokenType: LeftShift, Literal: "<<"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{TokenType: LeftShiftAssigne, Literal: "<<="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: Lteq, Literal: "<="}
			l.pos++
		}
	case '>':
		tk = &Token{TokenType: Gt, Literal: ">"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			tk = &Token{TokenType: RightShift, Literal: ">>"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{TokenType: RightShiftAssigne, Literal: ">>="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: Gteq, Literal: ">="}
			l.pos++
		}
	case ';':
		tk = &Token{TokenType: Semicolon, Literal: ";"}
		l.pos++
	case '(':
		tk = &Token{TokenType: Lparen, Literal: "("}
		l.pos++
	case ')':
		tk = &Token{TokenType: Rparen, Literal: ")"}
		l.pos++
	case ',':
		tk = &Token{TokenType: Comma, Literal: ","}
		l.pos++
	case '{':
		tk = &Token{TokenType: Lbrace, Literal: "{"}
		l.pos++
	case '}':
		tk = &Token{TokenType: Rbrace, Literal: "}"}
		l.pos++
	case '[':
		tk = &Token{TokenType: Lbracket, Literal: "["}
		l.pos++
	case ']':
		tk = &Token{TokenType: Rbracket, Literal: "]"}
		l.pos++
	case '&':
		tk = &Token{TokenType: Ampersand, Literal: "&"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '&' {
			tk = &Token{TokenType: And, Literal: "&&"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: AmpersandAssigne, Literal: "&="}
			l.pos++
		}
	case '~':
		tk = &Token{TokenType: Tilde, Literal: "~"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: TildeAssigne, Literal: "~="}
			l.pos++
		}
	case '^':
		tk = &Token{TokenType: Caret, Literal: "^"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: CaretAssigne, Literal: "^="}
			l.pos++
		}
	case '|':
		tk = &Token{TokenType: Vertical, Literal: "|"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '|' {
			tk = &Token{TokenType: Or, Literal: "||"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: VerticalAssigne, Literal: "|="}
			l.pos++
		}
	case '%':
		tk = &Token{TokenType: Percent, Literal: "%"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: PercentAssigne, Literal: "%="}
			l.pos++
		}
	case ':':
		tk = &Token{TokenType: Colon, Literal: ":"}
		l.pos++
	case '?':
		tk = &Token{TokenType: Question, Literal: "?"}
		l.pos++
	case '.':
		tk = &Token{TokenType: Period, Literal: "."}
		l.pos++
	case '\\':
		tk = &Token{TokenType: Backslash, Literal: "\\"}
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
		tk = &Token{TokenType: Float, Literal: w}
	} else {
		tk = &Token{TokenType: Integer, Literal: w}
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
	return &Token{TokenType: Str, Literal: w}
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
	tk := &Token{TokenType: Comment, Literal: l.input[l.pos:next]}
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
	return &Token{TokenType: Letter, Literal: string(s)}
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
	tk := &Token{TokenType: Illegal, Literal: l.input[l.pos:]}
	l.pos = len(l.input)
	return tk
}

func (l *Lexer) determineKeyword(w string) *Token {
	if strings.Compare("return", w) == 0 {
		return &Token{TokenType: KeyReturn, Literal: w}
	} else if strings.Compare("if", w) == 0 {
		return &Token{TokenType: KeyIf, Literal: w}
	} else if strings.Compare("else", w) == 0 {
		return &Token{TokenType: KeyElse, Literal: w}
	} else if strings.Compare("while", w) == 0 {
		return &Token{TokenType: KeyWhile, Literal: w}
	} else if strings.Compare("do", w) == 0 {
		return &Token{TokenType: KeyDo, Literal: w}
	} else if strings.Compare("goto", w) == 0 {
		return &Token{TokenType: KeyGoto, Literal: w}
	} else if strings.Compare("for", w) == 0 {
		return &Token{TokenType: KeyFor, Literal: w}
	} else if strings.Compare("break", w) == 0 {
		return &Token{TokenType: KeyBreak, Literal: w}
	} else if strings.Compare("continue", w) == 0 {
		return &Token{TokenType: KeyContinue, Literal: w}
	} else if strings.Compare("switch", w) == 0 {
		return &Token{TokenType: KeySwitch, Literal: w}
	} else if strings.Compare("case", w) == 0 {
		return &Token{TokenType: KeyCase, Literal: w}
	} else if strings.Compare("default", w) == 0 {
		return &Token{TokenType: KeyDefault, Literal: w}
	} else if strings.Compare("extern", w) == 0 {
		return &Token{TokenType: KeyExtern, Literal: w}
	} else if strings.Compare("volatile", w) == 0 {
		return &Token{TokenType: KeyVolatile, Literal: w}
	} else if strings.Compare("const", w) == 0 {
		return &Token{TokenType: KeyConst, Literal: w}
	} else if strings.Compare("typedef", w) == 0 {
		return &Token{TokenType: KeyTypedef, Literal: w}
	} else if strings.Compare("union", w) == 0 {
		return &Token{TokenType: KeyUnion, Literal: w}
	} else if strings.Compare("struct", w) == 0 {
		return &Token{TokenType: KeyStruct, Literal: w}
	} else if strings.Compare("enum", w) == 0 {
		return &Token{TokenType: KeyEnum, Literal: w}
	} else if strings.Compare("__attribute__", w) == 0 {
		return &Token{TokenType: KeyAttribute, Literal: w}
	} else if strings.Compare("void", w) == 0 {
		return &Token{TokenType: KeyVoid, Literal: w}
	} else if strings.Compare("__asm", w) == 0 {
		return &Token{TokenType: KeyAsm, Literal: w}
	} else if strings.Compare("sizeof", w) == 0 {
		return &Token{TokenType: KeySizeof, Literal: w}
	} else {
		return &Token{TokenType: Word, Literal: w}
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

// IsTypeToken
func (t *Token) IsTypeToken() bool {

	switch t.TokenType {
	case Word:
	case Asterisk:
	case KeyConst:
	case KeyVoid:
	case KeyStruct:
	case KeyUnion:
	case KeyEnum:
	case KeyVolatile:
	case Caret:
		// clang でコンパイルした場合型の種類に^が含まれる？
	default:
		return false
	}

	return true
}

func (t *Token) IsOperator() bool {
	switch t.TokenType {
	case Assign:
	case Plus:
	case Minus:
	case Asterisk:
	case Slash:
	case Lt:
	case Gt:
	case Eq:
	case Gteq:
	case Lteq:
	case Ne:
	case Ampersand:
	case Tilde:
	case Caret:
	case Vertical:
	case Question:
	case LeftShift:
	case RightShift:
	case Increment:
	case Decrement:
	case Or:
	case And:
	case Percent:
	case Colon:
	case PlusAssigne:
	case MinusAssigne:
	case AsteriskAssigne:
	case SlashAssigne:
	case VerticalAssigne:
	case AmpersandAssigne:
	case LeftShiftAssigne:
	case RightShiftAssigne:
	case TildeAssigne:
	case CaretAssigne:
	case PercentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) IsPrefixExpression() bool {
	switch t.TokenType {
	case Minus:
	case Plus:
	case Increment:
	case Decrement:
	case Tilde:
	case Bang:
	case Asterisk:
	case Ampersand:
	default:
		return false
	}
	return true
}

func (t *Token) IsPostExpression() bool {
	switch t.TokenType {
	case Lparen:
	case Increment:
	case Decrement:
	default:
		return false
	}
	return true
}

func (t *Token) IsCompoundOp() bool {
	switch t.TokenType {
	case PlusAssigne:
	case MinusAssigne:
	case AsteriskAssigne:
	case SlashAssigne:
	case VerticalAssigne:
	case AmpersandAssigne:
	case LeftShiftAssigne:
	case RightShiftAssigne:
	case TildeAssigne:
	case CaretAssigne:
	case PercentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) IsToken(t2 int) bool {
	return t.TokenType == t2
}
