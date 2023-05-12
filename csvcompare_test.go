package csvcompare

import (
	"testing"
)

func TestCompareWithEmptyOptions(t *testing.T) {
	opts := Options{}
	_, err := Compare("testdata/csv1.csv", "testdata/csv2.csv", opts)
	if err == nil {
		t.Errorf("should return an error")
	}
}

func TestCompareWithOptionsComma(t *testing.T) {
	opts := Options{Comma: 'd'}
	_, err := Compare("testdata/csv1.csv", "testdata/csv2.csv", opts)
	if err == nil {
		t.Errorf("should return an error")
	}
}

func TestSrcFileEmpty(t *testing.T) {
	opts := Options{Comma: ','}
	_, err := Compare("", "testdata/csv2.csv", opts)
	if err == nil {
		t.Errorf("should return srcFile is mandatory")
	}
}

func TestCompareReturnsDedup(t *testing.T) {
	opts := Options{Comma: ','}
	rows, _ := Compare("testdata/csv1.csv", "", opts)
	if len(rows) != 6 {
		t.Errorf("should return 6 lines. got=%d", len(rows))
	}
}

func TestCompareReturnsDiff(t *testing.T) {
	opts := Options{Comma: ','}
	rows, _ := Compare("testdata/csv1.csv", "testdata/csv2.csv", opts)
	if len(rows) != 7 {
		t.Errorf("should return 7 line. got=%d", len(rows))
	}
}
