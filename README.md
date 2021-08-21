# clanglex

Go Library which lexicalizes C language

## Description

プリプロセス処理済みの C ソースを字句解析するための Go 言語用ライブラリ. 

## Usage

sample.go
``` go
package main

import (
        "fmt"

        lex "github.com/kita127/clanglex"
)

func main() {

        src := `
int g_var;
static unsigned long s_var = (long)(100U);

int main(void){
    char local_var;
    g_var = s_var;
    g_var++
    return g_var;
}
`

        tokens := lex.Lexicalize(src)

        fmt.Println(tokens)
}
```
    $ go run ./sample.go 
    [TokenType:Word, Literal:int TokenType:Word, Literal:g_var TokenType:Semicolon, Literal:; TokenType:KeyStatic, Literal:static TokenType:Word, Literal:unsigned
     TokenType:Word, Literal:long TokenType:Word, Literal:s_var TokenType:Assign, Literal:= TokenType:Lparen, Literal:( TokenType:Word, Literal:long TokenType:Rpa
    ren, Literal:) TokenType:Lparen, Literal:( TokenType:Integer, Literal:100U TokenType:Rparen, Literal:) TokenType:Semicolon, Literal:; TokenType:Word, Literal:
    int TokenType:Word, Literal:main TokenType:Lparen, Literal:( TokenType:KeyVoid, Literal:void TokenType:Rparen, Literal:) TokenType:Lbrace, Literal:{ TokenType
    :Word, Literal:char TokenType:Word, Literal:local_var TokenType:Semicolon, Literal:; TokenType:Word, Literal:g_var TokenType:Assign, Literal:= TokenType:Word,
     Literal:s_var TokenType:Semicolon, Literal:; TokenType:Word, Literal:g_var TokenType:Increment, Literal:++ TokenType:KeyReturn, Literal:return TokenType:Word
    , Literal:g_var TokenType:Semicolon, Literal:; TokenType:Rbrace, Literal:} TokenType:Eof, Literal:eof]


### トークンの種類

| トークンタイプ       | トークンの内容                 | 例                                |
| -------------------- | -----------------------------  | --------------------------------- |
| Eof                  | EOF                            | -                                 |
| Word                 | 単語                           | `int`, `var_name`, `AnyType`      |
| Integer              | 整数リテラル                   | `100`, `0x00`, `20U`, `02`        |
| Float                | 浮動小数点リテラル             | `0.1230`, `.3211`, `0.2345f`      |
| Assign               | `=`                            | -                                 |
| Plus                 | `+`                            | -                                 |
| Minus                | `-`                            | -                                 |
| Bang                 | `!`                            | -                                 |
| Asterisk             | `*`                            | -                                 |
| Slash                | `/`                            | -                                 |
| Percent              | `%`                            | -                                 |
| Lt                   | `<`                            | -                                 |
| Gt                   | `>`                            | -                                 |
| Eq                   | `==`                           | -                                 |
| Ne                   | `!=`                           | -                                 |
| Gteq                 | `>=`                           | -                                 |
| Lteq                 | `<=`                           | -                                 |
| Semicolon            | `;`                            | -                                 |
| Lparen               | `(`                            | -                                 |
| Rparen               | `)`                            | -                                 |
| Comma                | `,`                            | -                                 |
| Lbrace               | `{`                            | -                                 |
| Rbrace               | `}`                            | -                                 |
| Lbracket             | `[`                            | -                                 |
| Rbracket             | `]`                            | -                                 |
| Ampersand            | `&`                            | -                                 |
| Tilde                | `~`                            | -                                 |
| Caret                | `^`                            | -                                 |
| Vertical             | `|`                            | -                                 |
| Colon                | `:`                            | -                                 |
| Question             | `?`                            | -                                 |
| Period               | `.`                            | -                                 |
| Backslash            | `/`                            | -                                 |
| Str                  | 文字列リテラル                 | "moji"                            |
| Letter               | 文字リテラル                   | 'c', 'h'                          |
| Arrow                | `->`                           | -                                 |
| LeftShift            | `<<`                           | -                                 |
| RightShift           | `>>`                           | -                                 |
| Increment            | `++`                           | -                                 |
| Decrement            | `--`                           | -                                 |
| And                  | `&&`                           | -                                 |
| Or                   | `||`                           | -                                 |
| PlusAssigne          | `+=`                           | -                                 |
| MinusAssigne         | `-=`                           | -                                 |
| AsteriskAssigne      | `*=`                           | -                                 |
| SlashAssigne         | `/=`                           | -                                 |
| VerticalAssigne      | `|=`                           | -                                 |
| AmpersandAssigne     | `&=`                           | -                                 |
| LeftShiftAssigne     | `<<=`                          | -                                 |
| RightShiftAssigne    | `>>=`                          | -                                 |
| TildeAssigne         | `~=`                           | -                                 |
| CaretAssigne         | `^=`                           | -                                 |
| PercentAssigne       | `%=`                           | -                                 |
| KeyReturn            | `return`                       | -                                 |
| KeyIf                | `if`                           | -                                 |
| KeyElse              | `else`                         | -                                 |
| KeyWhile             | `while`                        | -                                 |
| KeyDo                | `do`                           | -                                 |
| KeyGoto              | `goto`                         | -                                 |
| KeyFor               | `for`                          | -                                 |
| KeyBreak             | `break`                        | -                                 |
| KeyContinue          | `continue`                     | -                                 |
| KeySwitch            | `switch`                       | -                                 |
| KeyCase              | `case`                         | -                                 |
| KeyDefault           | `default`                      | -                                 |
| KeyExtern            | `extern`                       | -                                 |
| KeyVolatile          | `volatile`                     | -                                 |
| KeyConst             | `const`                        | -                                 |
| KeyTypedef           | `typedef`                      | -                                 |
| KeyUnion             | `union`                        | -                                 |
| KeyStruct            | `struct`                       | -                                 |
| KeyEnum              | `enum`                         | -                                 |
| KeyAttribute         | `__attribute__`                | -                                 |
| KeyVoid              | `void`                         | -                                 |
| KeyAsm               | `__asm`                        | -                                 |
| KeySizeof            | `sizeof`                       | -                                 |
| KeyStatic            | `static`                       | -                                 |
| Comment              | コメント                       | `/* comment */`                   |
| Illegal              | 字句解析できなかったトークン   | -                                 |

### メソッド・関数

#### IsToken

トークンの判別に使用する. 

``` go

tokens := lex.Lexicalize(src)

for _, t := range tokens

    if t.IsToken(KeyExtern){
        // extern の時の処理
    }

```

## License

This software is released under the MIT License, see LICENSE.
