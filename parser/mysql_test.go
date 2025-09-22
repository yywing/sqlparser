package parser

import "testing"

func TestIsMysql(t *testing.T) {
	sqls := map[string]bool{
		"SELECT name from aaa":                                             true,
		"select /* back-quote num */ `2a` from t":                          true,
		"select * from t1 where k like 'Müller' collate latin1_german2_ci": true,

		"INSERT INTO aaa (name, age) VALUES ('a', 1)": true,
		"update /* bool in update */ a set b = true":  true,
		"delete /* where */ from a where a = b":       true,

		"alter table a reorganize partition b into (partition c values less than (?), partition d values less than (maxvalue))": true,
		"create table a (\n\t`a` int\n)":    true,
		"drop view a":                       true,
		"SELECT name from aaa; drop view a": true,

		"I am a good select body gg where": false,
		"select":                           false,

		"":      true,
		"-- aa": true,
		"# aaa": true,
	}
	for sql, expect := range sqls {
		result := IsMysql(sql)
		if result != expect {
			t.Errorf("IsMysql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulMysql(t *testing.T) {
	sqls := map[string]bool{
		"":      false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulMysql(sql)
		if result != expect {
			t.Errorf("IsMeaningfulMysql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
