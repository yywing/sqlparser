package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	flink "github.com/yywing/sqlparser/parser/flink"
)

func IsFlink(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := flink.NewFlinkSqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := flink.NewFlinkSqlParser(stream)
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

	_ = p.SqlStatements()

	if lexerErrorListener.Err != nil {
		fmt.Println("aa error: ", parserErrorListener.Err)
		return false
	}

	if parserErrorListener.Err != nil {
		fmt.Println("parser error: ", parserErrorListener.Err)
		return false
	}
	return true
}

func IsMeaningfulFlink(sql string) bool {
	ok := IsFlink(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := flink.NewFlinkSqlLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case flink.FlinkSqlLexerKW_SELECT,
			flink.FlinkSqlLexerKW_INSERT,
			flink.FlinkSqlLexerKW_UPDATE,
			flink.FlinkSqlLexerKW_DELETE,
			flink.FlinkSqlLexerKW_CREATE,
			flink.FlinkSqlLexerKW_ALTER,
			flink.FlinkSqlLexerKW_DROP:
			return true
		}

	}
	return false
}
