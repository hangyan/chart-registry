package storage

import "testing"

func TestParseChartName(t *testing.T) {
	input := "simple-pod-0.1.0.tgz"
	name, version := parseChartName(input)
	if name != "simple-pod" || version != "0.1.0" {
		t.Errorf("error parse: %s %s ", name, version)
	}

}
