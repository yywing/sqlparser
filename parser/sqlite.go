package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	postgresql "github.com/yywing/sqlparser/parser/postgresql"
	sqlite "github.com/yywing/sqlparser/parser/sqlite"
)

func IsSQLite(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := sqlite.NewSQLiteLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := sqlite.NewSQLiteParser(stream)
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

	_ = p.Parse()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulSQLite(sql string) bool {
	ok := IsSQLite(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := postgresql.NewPostgreSQLLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case sqlite.SQLiteLexerSELECT_,
			sqlite.SQLiteLexerINSERT_,
			sqlite.SQLiteLexerUPDATE_,
			sqlite.SQLiteLexerDELETE_,
			sqlite.SQLiteLexerCREATE_,
			sqlite.SQLiteLexerALTER_,
			sqlite.SQLiteLexerDROP_:
			return true
		}

	}
	return false
}
