package main

import "testing"

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed main run")
	}
	//defer db.SQL.Close()
}
