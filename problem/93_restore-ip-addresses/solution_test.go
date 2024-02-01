package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_restore_ip_addresses(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		restoreIpAddresses,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"25525511135"
["255.255.11.135","255.255.111.35"]

"0000"
["0.0.0.0"]

"101023"
["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

`
