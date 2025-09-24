package parser

import "testing"

func TestIsClickhouse(t *testing.T) {
	sqls := map[string]bool{
		"SELECT COLUMNS('a'), COLUMNS('c'), toTypeName(COLUMNS('c')) FROM col_names":                                            true,
		"CREATE TABLE IF NOT EXISTS all_hits ON CLUSTER cluster (p Date, i Int32) ENGINE = Distributed(cluster, default, hits)": true,
		"INSERT INTO t VALUES (1, 'Hello, world'), (2, 'abc'), (3, 'def')":                                                      true,

		"I am a good select body gg where": false,
		"select":                           false,

		"":      true,
		";":     false,
		"-- aa": true,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsClickhouse(sql)
		if result != expect {
			t.Errorf("IsClickhouse(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulClickhouse(t *testing.T) {
	sqls := map[string]bool{
		"":      false,
		";":     false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulClickhouse(sql)
		if result != expect {
			t.Errorf("IsMeaningfulClickhouse(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
