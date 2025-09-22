package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	mysql "github.com/yywing/sqlparser/parser/mysql"
)

func IsMysql(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := mysql.NewMySQLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := mysql.NewMySQLParser(stream)
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

	_ = p.Queries()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulMysql(sql string) bool {
	ok := IsMysql(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := mysql.NewMySQLLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case mysql.MySQLLexerSELECT_SYMBOL,
			mysql.MySQLLexerINSERT_SYMBOL,
			mysql.MySQLLexerUPDATE_SYMBOL,
			mysql.MySQLLexerDELETE_SYMBOL,
			mysql.MySQLLexerCREATE_SYMBOL,
			mysql.MySQLLexerALTER_SYMBOL,
			mysql.MySQLLexerDROP_SYMBOL:
			return true
		}

	}
	return false
}
