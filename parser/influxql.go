package parser

import (
	"strings"

	"github.com/influxdata/influxql"
)

func IsInfluxql(statement string) bool {
	scanner := influxql.NewScanner(strings.NewReader(statement))

	// 收集所有 BOUNDPARAM
	boundParams := make(map[string]interface{})
	for {
		tok, _, lit := scanner.Scan()
		if tok == influxql.EOF {
			break
		}
		if tok == influxql.BOUNDPARAM {
			boundParams[strings.TrimPrefix(lit, "$")] = true
		}
	}

	parser := influxql.NewParser(strings.NewReader(statement))

	parser.SetParams(boundParams)

	_, err := parser.ParseStatement()
	return err == nil
}
