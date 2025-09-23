package parser

import (
	"testing"
)

func TestIsInfluxql(t *testing.T) {
	sqls := map[string]bool{
		`SELECT mean("value") FROM "io_hang_time" WHERE ("value" > 3000 and "value" < 86400000) AND $timeFilter GROUP BY time(30s), "Region","k8s_nodeip" ,"volume_name" , "cluster_name"  fill(0)`: true,
		`select`:   false,
		"-- aa":    false,
		"/* aa */": false,
	}
	for sql, expect := range sqls {
		result := IsInfluxql(sql)
		if result != expect {
			t.Errorf("IsInfluxql(%s) expect %v, got %v", sql, expect, result)
		}
	}
}
