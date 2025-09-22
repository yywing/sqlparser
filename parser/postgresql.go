package parser

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	postgresql "github.com/yywing/sqlparser/parser/postgresql"
)

type PostgresqlParser struct {
	*postgresql.BasePostgreSQLParserListener
	stmts []string
}

func NewPostgresqlParser() *PostgresqlParser {
	return &PostgresqlParser{
		BasePostgreSQLParserListener: &postgresql.BasePostgreSQLParserListener{},
	}
}

// EnterStmt is called when production stmt is entered.
func (s *PostgresqlParser) EnterStmt(ctx *postgresql.StmtContext) {
	s.stmts = append(s.stmts, ctx.GetText())
}

func IsPostgresql(statement string) bool {
	trimmedStatement := strings.TrimRightFunc(statement, unicode.IsSpace)
	if len(trimmedStatement) > 0 && !strings.HasSuffix(trimmedStatement, ";") {
		// Add a semicolon to the end of the statement to allow users to omit the semicolon
		// for the last statement in the script.
		statement += ";"
	}

	inputStream := antlr.NewInputStream(statement)
	lexer := postgresql.NewPostgreSQLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := postgresql.NewPostgreSQLParser(stream)
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
	_ = p.Root()

	if lexerErrorListener.Err != nil {
		return false
	}

	if parserErrorListener.Err != nil {
		return false
	}
	return true
}

func IsMeaningfulPostgresql(sql string) bool {
	ok := IsPostgresql(sql)
	if !ok {
		return false
	}
	inputStream := antlr.NewInputStream(sql)
	lexer := postgresql.NewPostgreSQLLexer(inputStream)

	var t int
Outer:
	for _, token := range lexer.GetAllTokens() {
		switch token.GetTokenType() {
		case postgresql.PostgreSQLLexerINSERT,
			postgresql.PostgreSQLLexerUPDATE,
			postgresql.PostgreSQLLexerDELETE_P,
			postgresql.PostgreSQLLexerCREATE,
			postgresql.PostgreSQLLexerALTER,
			postgresql.PostgreSQLLexerDROP:
			return true
		case postgresql.PostgreSQLLexerSELECT:
			t = token.GetTokenType()
			break Outer
		}

	}

	// 特殊处理一下 select 语句，里面可能有非 select 语句
	if t == postgresql.PostgreSQLLexerSELECT {
		lexer.Reset()
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := postgresql.NewPostgreSQLParser(stream)
		p.RemoveErrorListeners()
		tree := p.Root()

		l := NewPostgresqlParser()
		antlr.ParseTreeWalkerDefault.Walk(l, tree)

		if len(l.stmts) > 0 {
			for _, stmt := range l.stmts {
				if strings.TrimSpace(strings.ToUpper(stmt)) != "SELECT" {
					return true
				}
			}
		}
		return false
	}

	return false
}
