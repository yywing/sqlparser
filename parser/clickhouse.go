package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	clickhouse "github.com/yywing/sqlparser/parser/clickhouse"
)

func IsClickhouse(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := clickhouse.NewClickHouseLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := clickhouse.NewClickHouseParser(stream)
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

	_ = p.ClickhouseFile()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulClickhouse(sql string) bool {
	ok := IsClickhouse(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := clickhouse.NewClickHouseLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case clickhouse.ClickHouseLexerSELECT,
			clickhouse.ClickHouseLexerINSERT,
			clickhouse.ClickHouseLexerUPDATE,
			clickhouse.ClickHouseLexerDELETE,
			clickhouse.ClickHouseLexerCREATE,
			clickhouse.ClickHouseLexerALTER,
			clickhouse.ClickHouseLexerDROP:
			return true
		}

	}
	return false
}
