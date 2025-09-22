package parser

import "testing"

func TestIsPostgresql(t *testing.T) {
	sqls := map[string]bool{
		"SELECT name from aaa":                                             true,
		"select /* back-quote num */ \"2a\" from t":                        true,
		"select * from t1 where k like 'Müller' collate latin1_german2_ci": true,

		"INSERT INTO aaa (name, age) VALUES ('a', 1)": true,
		"update /* bool in update */ a set b = true":  true,
		"delete /* where */ from a where a = b":       true,

		"ALTER TABLE a DETACH PARTITION b;": true,
		"create table a (\n\t\"a\" int\n)":  true,
		"drop view a":                       true,
		"SELECT name from aaa; drop view a": true,

		"I am a good select body gg where": false,
		"select":                           true,

		"":         true,
		"-- aa":    true,
		"/* aa */": true,
	}
	for sql, expect := range sqls {
		result := IsPostgresql(sql)
		if result != expect {
			t.Errorf("IsPostgresql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulPostgresql(t *testing.T) {
	sqls := map[string]bool{
		"SELECT name from aaa": true,
		"select":               false,
		"":                     false,
		"-- aa":                false,
		"# aaa":                false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulPostgresql(sql)
		if result != expect {
			t.Errorf("IsMeaningfulPostgresql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
