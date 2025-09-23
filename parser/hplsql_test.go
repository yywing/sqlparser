package parser

import "testing"

func TestIsHplsql(t *testing.T) {
	sqls := map[string]bool{
		"set ttp.dependency.parse.skip=true;\n\nSELECT  *\nFROM  xr.ttp_test\nWHERE   dt = '20260126';\n          ;\n          set ttp.job.id=1000014306;\n          set ttp.job.create.timea=2025-09-02;": true,
		"create table if not exists t (i int, d decimal) stored as orc;": true,
		"create view v1 as select c from t":                              true,
		"drop table if exists t":                                         true,
		"select a from table where date = '20250806'":                    true,

		"I am a good select body gg where": false,
		"select":                           true,

		"":      false,
		"-- aa": false,
		"# aaa": true,
	}
	for sql, expect := range sqls {
		result := IsHplsql(sql)
		if result != expect {
			t.Errorf("IsHplsql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulHplsql(t *testing.T) {
	sqls := map[string]bool{
		"":      false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulHplsql(sql)
		if result != expect {
			t.Errorf("IsMeaningfulHplsql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
