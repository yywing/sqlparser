package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	spark "github.com/yywing/sqlparser/parser/spark"
)

func IsSpark(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := spark.NewSqlBaseLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := spark.NewSqlBaseParser(stream)
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

func IsMeaningfulSpark(sql string) bool {
	ok := IsHplsql(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := spark.NewSqlBaseLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case spark.SqlBaseLexerSELECT,
			spark.SqlBaseLexerINSERT,
			spark.SqlBaseLexerUPDATE,
			spark.SqlBaseLexerDELETE,
			spark.SqlBaseLexerCREATE,
			spark.SqlBaseLexerALTER,
			spark.SqlBaseLexerDROP:
			return true
		}

	}
	return false
}
