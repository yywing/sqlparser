package parser

import "testing"

func TestIsSpark(t *testing.T) {
	sqls := map[string]bool{
		"CREATE TEMPORARY VIEW `test_tag` AS\nSELECT\n  *\nFROM `moss`.`test_tag`; SELECT a from A;;;;;": true,
		"create table if not exists t (i int, d decimal) stored as orc;":                                 true,
		"create view v1 as select c from t":                                                              true,
		"drop table if exists t":                                                                         true,
		"select a from table where date = '20250806'":                                                    true,

		"I am a good select body gg where": false,
		"select":                           false,

		"":      false,
		";":     true,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsSpark(sql)
		if result != expect {
			t.Errorf("IsSpark(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulSpark(t *testing.T) {
	sqls := map[string]bool{
		"":       false,
		";":      false,
		"-- aa":  false,
		"# aaa":  false,
		"select": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulSpark(sql)
		if result != expect {
			t.Errorf("IsMeaningfulSpark(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
