package parser

import "testing"

func TestIsHive(t *testing.T) {
	sqls := map[string]bool{
		"create external table if not exists t (i int, d decimal) stored as orc;": true,
		"create view v1 as select c from t":                                       true,
		"drop table if exists t":                                                  true,
		"SELECT A + - B * C ^ D + E | F OR G AND H || I - (SELECT COUNT(*) FROM DB.TB) OR J BETWEEN K AND L OR M IS TRUE FROM DB.TB1": true,
		`WITH A AS (
			SELECT *, 1e1 as c, 1E1 as d, 1.0 as e FROM TB
		)
		SELECT A,B,C
		FROM A`: true,
		"show extended tables in db where t = 't'":    true,
		"select a from table where date = '20250806'": true,

		"I am a good select body gg where": false,
		"select":                           false,

		"":      false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsHive(sql)
		if result != expect {
			t.Errorf("IsHive(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulHive(t *testing.T) {
	sqls := map[string]bool{
		"":      false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulHive(sql)
		if result != expect {
			t.Errorf("IsMeaningfulHive(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
