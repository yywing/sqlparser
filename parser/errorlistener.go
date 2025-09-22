package parser

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// ParseErrorListener is a custom error listener for PLSQL parser.
type ParseErrorListener struct {
	BaseLine  int
	Err       error
	Statement string
}

// SyntaxError returns the errors.
func (l *ParseErrorListener) SyntaxError(_ antlr.Recognizer, token any, line, column int, message string, _ antlr.RecognitionException) {
	if l.Err != nil {
		return
	}

	errMessage := ""
	if token, ok := token.(*antlr.CommonToken); ok {
		stream := token.GetInputStream()
		start := token.GetStart() - 40
		if start < 0 {
			start = 0
		}
		stop := token.GetStop()
		if stop >= stream.Size() {
			stop = stream.Size() - 1
		}
		errMessage = fmt.Sprintf("related text: %s", stream.GetTextFromInterval(antlr.NewInterval(start, stop)))
	}
	if l.Statement == "" {
		l.Err = errors.New(errMessage)
		return
	}
	l.Err = errors.New(errMessage)
}

// ReportAmbiguity reports an ambiguity.
func (*ParseErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// ReportAttemptingFullContext reports an attempting full context.
func (*ParseErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

// ReportContextSensitivity reports a context sensitivity.
func (*ParseErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}
