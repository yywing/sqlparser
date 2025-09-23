package parser

import "testing"

func TestIsFlink(t *testing.T) {
	sqls := map[string]bool{
		"CREATE TEMPORARY VIEW `test_tag` AS\nSELECT\n  *\nFROM `moss`.`test_tag`;\n": true,

		"I am a good select body gg where": false,
		"select":                           false,

		"":      true,
		";":     true,
		"-- aa": true,
		"# aaa": true,
	}
	for sql, expect := range sqls {
		result := IsFlink(sql)
		if result != expect {
			t.Errorf("IsFlink(%s) expect %v, got %v", sql, expect, result)
		}
	}
}

func TestIsMeaningfulFlink(t *testing.T) {
	sqls := map[string]bool{
		"":      false,
		";":     false,
		"-- aa": false,
		"# aaa": false,
	}
	for sql, expect := range sqls {
		result := IsMeaningfulFlink(sql)
		if result != expect {
			t.Errorf("IsMeaningfulFlink(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
