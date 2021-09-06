package clanglex

import (
	_ "reflect"
	"testing"
)

func TestLexicalize(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  []*Token
	}{
		{
			"test0",
			``,
			[]*Token{
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test1",
			`   char   `,
			[]*Token{
				{
					Word,
					"char",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test2",
			`   char		hoge   `,
			[]*Token{
				{
					Word,
					"char",
				},
				{
					Word,
					"hoge",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test3",
			`=`,
			[]*Token{
				{
					Assign,
					"=",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test4",
			`=+-!*/<>;(),{}[]`,
			[]*Token{
				{
					Assign,
					"=",
				},
				{
					Plus,
					"+",
				},
				{
					Minus,
					"-",
				},
				{
					Bang,
					"!",
				},
				{
					Asterisk,
					"*",
				},
				{
					Slash,
					"/",
				},
				{
					Lt,
					"<",
				},
				{
					Gt,
					">",
				},
				{
					Semicolon,
					";",
				},
				{
					Lparen,
					"(",
				},
				{
					Rparen,
					")",
				},
				{
					Comma,
					",",
				},
				{
					Lbrace,
					"{",
				},
				{
					Rbrace,
					"}",
				},
				{
					Lbracket,
					"[",
				},
				{
					Rbracket,
					"]",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test5",
			` = + - ! * / % < > ; ( ) , { } [ ] `,
			[]*Token{
				{
					Assign,
					"=",
				},
				{
					Plus,
					"+",
				},
				{
					Minus,
					"-",
				},
				{
					Bang,
					"!",
				},
				{
					Asterisk,
					"*",
				},
				{
					Slash,
					"/",
				},
				{
					Percent,
					"%",
				},
				{
					Lt,
					"<",
				},
				{
					Gt,
					">",
				},
				{
					Semicolon,
					";",
				},
				{
					Lparen,
					"(",
				},
				{
					Rparen,
					")",
				},
				{
					Comma,
					",",
				},
				{
					Lbrace,
					"{",
				},
				{
					Rbrace,
					"}",
				},
				{
					Lbracket,
					"[",
				},
				{
					Rbracket,
					"]",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test6",
			`&~^|:?.\-><<>>++--||&&==!=>=<=`,
			[]*Token{
				{
					Ampersand,
					"&",
				},
				{
					Tilde,
					"~",
				},
				{
					Caret,
					"^",
				},
				{
					Vertical,
					"|",
				},
				{
					Colon,
					":",
				},
				{
					Question,
					"?",
				},
				{
					Period,
					".",
				},
				{
					Backslash,
					"\\",
				},
				{
					Arrow,
					"->",
				},
				{
					LeftShift,
					`<<`,
				},
				{
					RightShift,
					`>>`,
				},
				{
					Increment,
					`++`,
				},
				{
					Decrement,
					`--`,
				},
				{
					Or,
					`||`,
				},
				{
					And,
					`&&`,
				},
				{
					Eq,
					`==`,
				},
				{
					Ne,
					`!=`,
				},
				{
					Gteq,
					`>=`,
				},
				{
					Lteq,
					`<=`,
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test7",
			`+= -= *= /= |= &= <<= >>= ~= ^= %=`,
			[]*Token{
				{
					PlusAssigne,
					"+=",
				},
				{
					MinusAssigne,
					"-=",
				},
				{
					AsteriskAssigne,
					"*=",
				},
				{
					SlashAssigne,
					"/=",
				},
				{
					VerticalAssigne,
					"|=",
				},
				{
					AmpersandAssigne,
					"&=",
				},
				{
					LeftShiftAssigne,
					"<<=",
				},
				{
					RightShiftAssigne,
					">>=",
				},
				{
					TildeAssigne,
					"~=",
				},
				{
					CaretAssigne,
					"^=",
				},
				{
					PercentAssigne,
					"%=",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test8",
			`   ident00+123;   `,
			[]*Token{
				{
					Word,
					"ident00",
				},
				{
					Plus,
					"+",
				},
				{
					Integer,
					"123",
				},
				{
					Semicolon,
					";",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test9",
			`# 1 "hoge.c"`,
			[]*Token{
				{
					Comment,
					" 1 \"hoge.c\"",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test10",
			`# 1 "hoge.c"
# 1 "<built-in>" 1`,
			[]*Token{
				{
					Comment,
					" 1 \"hoge.c\"",
				},
				{
					Comment,
					" 1 \"<built-in>\" 1",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test11",
			`0 0U 123 0xA1c 0765 0b0110 567u 567U 567l 567L 567lu 567UL`,
			[]*Token{
				{
					Integer,
					"0",
				},
				{
					Integer,
					"0U",
				},
				{
					Integer,
					"123",
				},
				{
					Integer,
					"0xA1c",
				},
				{
					Integer,
					"0765",
				},
				{
					Integer,
					"0b0110",
				},
				{
					Integer,
					"567u",
				},
				{
					Integer,
					"567U",
				},
				{
					Integer,
					"567l",
				},
				{
					Integer,
					"567L",
				},
				{
					Integer,
					"567lu",
				},
				{
					Integer,
					"567UL",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test12",
			`0.123 987.123 123.`,
			[]*Token{
				{
					Float,
					"0.123",
				},
				{
					Float,
					"987.123",
				},
				{
					Float,
					"123.",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test13",
			`return if else while do goto for break continue switch case default
 extern volatile const typedef union struct enum __attribute__ void`,
			[]*Token{
				{
					KeyReturn,
					"return",
				},
				{
					KeyIf,
					"if",
				},
				{
					KeyElse,
					"else",
				},
				{
					KeyWhile,
					"while",
				},
				{
					KeyDo,
					"do",
				},
				{
					KeyGoto,
					"goto",
				},
				{
					KeyFor,
					"for",
				},
				{
					KeyBreak,
					"break",
				},
				{
					KeyContinue,
					"continue",
				},
				{
					KeySwitch,
					"switch",
				},
				{
					KeyCase,
					"case",
				},
				{
					KeyDefault,
					"default",
				},
				{
					KeyExtern,
					"extern",
				},
				{
					KeyVolatile,
					"volatile",
				},
				{
					KeyConst,
					"const",
				},
				{
					KeyTypedef,
					"typedef",
				},
				{
					KeyUnion,
					"union",
				},
				{
					KeyStruct,
					"struct",
				},
				{
					KeyEnum,
					"enum",
				},
				{
					KeyAttribute,
					"__attribute__",
				},
				{
					KeyVoid,
					"void",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test13",
			`char hoge[] = "hello";`,
			[]*Token{
				{
					Word,
					"char",
				},
				{
					Word,
					"hoge",
				},
				{
					Lbracket,
					"[",
				},
				{
					Rbracket,
					"]",
				},
				{
					Assign,
					"=",
				},
				{
					Str,
					"\"hello\"",
				},
				{
					Semicolon,
					";",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test14",
			`int hoge = 0;`,
			[]*Token{
				{
					Word,
					"int",
				},
				{
					Word,
					"hoge",
				},
				{
					Assign,
					"=",
				},
				{
					Integer,
					"0",
				},
				{
					Semicolon,
					";",
				},
				{
					Eof,
					"eof",
				},
			},
		},

		{
			"test15",
			`
# 1 "hoge.c"

int func(int a) {
      a = a + (10);
        return a;
    }
`,
			[]*Token{
				{
					Comment,
					` 1 "hoge.c"`,
				},
				{
					Word,
					`int`,
				},
				{
					Word,
					`func`,
				},
				{
					Lparen,
					`(`,
				},
				{
					Word,
					`int`,
				},
				{
					Word,
					`a`,
				},
				{
					Rparen,
					`)`,
				},
				{
					Lbrace,
					`{`,
				},
				{
					Word,
					`a`,
				},
				{
					Assign,
					`=`,
				},
				{
					Word,
					`a`,
				},
				{
					Plus,
					`+`,
				},
				{
					Lparen,
					`(`,
				},
				{
					Integer,
					`10`,
				},
				{
					Rparen,
					`)`,
				},
				{
					Semicolon,
					`;`,
				},
				{
					KeyReturn,
					`return`,
				},
				{
					Word,
					`a`,
				},
				{
					Semicolon,
					`;`,
				},
				{
					Rbrace,
					`}`,
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test16",
			`'A' '\n'`,
			[]*Token{
				{
					Letter,
					"A",
				},
				{
					Letter,
					"\\n",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test17",
			`
            '"' '\\' '\b' '\f' '\n' '\r' '\t' '\033' '\'' '\0'
`,
			[]*Token{
				{
					Letter,
					`"`,
				},
				{
					Letter,
					`\\`,
				},
				{
					Letter,
					`\b`,
				},
				{
					Letter,
					`\f`,
				},
				{
					Letter,
					`\n`,
				},
				{
					Letter,
					`\r`,
				},
				{
					Letter,
					`\t`,
				},
				{
					Letter,
					`\033`,
				},
				{
					Letter,
					`\'`,
				},
				{
					Letter,
					`\0`,
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test18",
			`
             "\\\\"
             "\\b"
             "\\f"
             "\\n"
             "\\r"
             "\\t"
            `,
			[]*Token{
				{
					Str,
					`"\\\\"`,
				},
				{
					Str,
					`"\\b"`,
				},
				{
					Str,
					`"\\f"`,
				},
				{
					Str,
					`"\\n"`,
				},
				{
					Str,
					`"\\r"`,
				},
				{
					Str,
					`"\\t"`,
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test19",
			`
             "\""
             "\" ***\\"
`,
			[]*Token{
				{
					Str,
					`"\""`,
				},
				{
					Str,
					`"\" ***\\"`,
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test20",
			`
            sizeof
`,
			[]*Token{
				{
					KeySizeof,
					"sizeof",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test gcc 1",
			`__asm`,
			[]*Token{
				{
					KeyAsm,
					"__asm",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test static 1",
			`static`,
			[]*Token{
				{
					KeyStatic,
					"static",
				},
				{
					Eof,
					"eof",
				},
			},
		},
	}

	for _, tt := range testTbl {
		t.Logf("%s", tt.comment)
		l := NewLexer(tt.src)
		got, err := l.lexicalize()
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != len(tt.expect) {
			t.Fatalf("got len=%v, expect len=%v", len(got), len(tt.expect))
		}
		for i, v := range got {
			e := tt.expect[i]
			if v.TokenType != e.TokenType {
				t.Errorf("got type=%v, expect type=%v", v.TokenType, tt.expect[i].TokenType)
			}
			if v.Literal != e.Literal {
				t.Errorf("got literal=%v, expect literal=%v", v.Literal, tt.expect[i].Literal)
			}
		}
	}
}

func TestLexComment(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  []*Token
	}{
		{
			"test comment 1",
			`/**/`,
			[]*Token{
				{
					Comment,
					"",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test comment 2",
			`/* hoge */`,
			[]*Token{
				{
					Comment,
					" hoge ",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test comment 3",
			`/** */`,
			[]*Token{
				{
					Comment,
					"* ",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test comment 4",
			`/* **/`,
			[]*Token{
				{
					Comment,
					" *",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test comment 5",
			`/* /* */`,
			[]*Token{
				{
					Comment,
					" /* ",
				},
				{
					Eof,
					"eof",
				},
			},
		},
		{
			"test comment 6",
			`/* * / */`,
			[]*Token{
				{
					Comment,
					" * / ",
				},
				{
					Eof,
					"eof",
				},
			},
		},
	}

	for _, tt := range testTbl {
		t.Logf("%s", tt.comment)
		l := NewLexer(tt.src)
		got, err := l.lexicalize()
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != len(tt.expect) {
			t.Fatalf("got len=%v, expect len=%v", len(got), len(tt.expect))
		}
		for i, v := range got {
			e := tt.expect[i]
			if v.TokenType != e.TokenType {
				t.Errorf("got type=%v, expect type=%v", v.TokenType, tt.expect[i].TokenType)
			}
			if v.Literal != e.Literal {
				t.Errorf("got literal=%v, expect literal=%v", v.Literal, tt.expect[i].Literal)
			}
		}
	}
}

func TestBin(t *testing.T) {
	testTbl := []struct {
		comment string
		bin     []byte
		expect  []*Token
	}{
		{
			"test0",
			[]byte{
				0x3f, 0x3f, 0x3f, 0x3f, 0x07, 0x00, 0x00, 0x01, 0x03, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
				0x0b, 0x00, 0x00, 0x00, 0x30, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x19, 0x00, 0x00, 0x00, 0x48, 0x00, 0x00, 0x00, 0x5f, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5a, 0x45,
				0x52, 0x4f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00, 0x78, 0x02, 0x00, 0x00,
				0x5f, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x07, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5f, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x10, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0xd5, 0x84, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x10, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x04, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5f, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x73, 0x74, 0x75, 0x62, 0x31, 0x00, 0x00,
				0x5f, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x3f, 0x3f, 0x0a, 0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x3f, 0x3f, 0x0a, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x08, 0x04, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5f, 0x5f, 0x72, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			[]*Token{
				{
					Eof,
					"eof",
				},
			},
		},
	}

	for _, tt := range testTbl {
		t.Logf("%s", tt.comment)
		l := NewLexer(string(tt.bin))
		_, err := l.lexicalize()
		if err == nil {
			t.Fatal(err)
		}
	}
}
