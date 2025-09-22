package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	hive "github.com/yywing/sqlparser/parser/hive"
)

func IsHive(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && strings.HasSuffix(trimmedStatement, ";") {
		// 去掉结尾的 ;
		statement = statement[:len(statement)-1]
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := hive.NewHiveLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := hive.NewHiveParser(stream)
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

	_ = p.Statement()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulHive(sql string) bool {
	ok := IsHive(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := hive.NewHiveLexer(inputStream)

	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case hive.HiveLexerKW_SELECT,
			hive.HiveLexerKW_INSERT,
			hive.HiveLexerKW_UPDATE,
			hive.HiveLexerKW_DELETE,
			hive.HiveLexerKW_CREATE,
			hive.HiveLexerKW_ALTER,
			hive.HiveLexerKW_DROP:
			return true
		}

	}
	return false
}
