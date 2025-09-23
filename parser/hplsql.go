package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	hplsql "github.com/yywing/sqlparser/parser/hplsql"
)

func IsHplsql(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := hplsql.NewHplsqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := hplsql.NewHplsqlParser(stream)
	lexerErrorListener := &ParseErrorListener{
		Statement: statement,
		BaseLine:  0,
	}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)

	parserErrorListener := &ParseErrorListener{
		Statement: statement,
		BaseLine:  0,
	}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrorListener)

	p.BuildParseTrees = false

	_ = p.Program()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulHplsql(sql string) bool {
	ok := IsHplsql(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := hplsql.NewHplsqlLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case hplsql.HplsqlLexerT_SELECT,
			hplsql.HplsqlLexerT_INSERT,
			hplsql.HplsqlLexerT_UPDATE,
			hplsql.HplsqlLexerT_DELETE,
			hplsql.HplsqlLexerT_CREATE,
			hplsql.HplsqlLexerT_ALTER,
			hplsql.HplsqlLexerT_DROP:
			return true
		}

	}
	return false
}
