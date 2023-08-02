package csvcompare

import (
	"reflect"
	"testing"
)

var rowJohnson81 = []string{"johnson81", "4081", "30no86", "cj4081", "Craig", "Johnson", "Depot", "London", "2006-02-12"}
var rowJenkins46 = []string{"jenkins46", "9346", "14ju73", "mj9346", "Mary", "Jenkins", "Engineering", "Manchester2", "2016-07-10"}

// var rowJenkins46Manchester2 = []string{"jenkins46", "9346", "14ju73", "mj9346", "Mary", "Jenkins", "Engineering", "Manchester", "2016-07-10"}
var rowGandalf00 = []string{"gandalf00", "6486", "9sdfcq", "9sqdcd", "Mithrandir", "Gandalf", "Magical", "Isengard", "2001-12-10"}
var rowDoe80 = []string{"doe80", "6546", "65dfcsd", "5dsfcsd", "Jane", "Doe", "Marketing", "Glasgow", "2010-05-18"}
var rowBooker12 = []string{"booker12", "9012", "12se74", "rb9012", "Rachel", "Booker", "Sales", "Manchester", "2023-10-02"}
var rowHead = []string{"Username", "Identifier", "One-time password", "Recovery code", "First name", "Last name", "Department", "Location", "Date"}

func TestCompareWithEmptyOptions(t *testing.T) {
	opts := &Options{}
	_, err := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv")
	if err == nil {
		t.Errorf("should return an error")
	}
}

func TestCompareWithOptionsCommaNotAllowed(t *testing.T) {
	opts := NewOptions(false)
	opts.SetComma('d')
	_, err := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv")
	if err == nil {
		t.Errorf("should return an error")
	}
}

func TestSrcFileEmpty(t *testing.T) {
	opts := NewOptions(false)
	opts.NoRevert()
	_, err := Compare("", opts)
	if err == nil {
		t.Errorf("should return srcFile is mandatory")
	}
}

func TestCompareReturnsDedup(t *testing.T) {
	opts := NewOptions(true)
	opts.NoRevert()
	rows, _ := Compare("testdata/csv1.csv", opts)
	if len(rows) != 6 {
		t.Errorf("should return 6 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[0], rowHead) {
		t.Errorf("line 0 should return %v got=%v", rowHead, rows[0])
	}
	if !reflect.DeepEqual(rows[5], rowJohnson81) {
		t.Errorf("line 5 should return %v got=%v", rowJohnson81, rows[5])
	}
}

func TestCompareReturnsDiff(t *testing.T) {
	opts := NewOptions(true)
	opts.NoRevert()
	opts.headers = true
	rows, _ := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv")
	if len(rows) != 7 {
		t.Errorf("should return 7 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[6], rowJenkins46) {
		t.Errorf("line 6 should return %v got=%v", rowJenkins46, rows[6])
	}
}

func TestCompareWithOneFileNoHeaders(t *testing.T) {
	opts := NewOptions(true)
	opts.NoRevert()
	opts.headers = true
	_, err := Compare("testdata/csv1.csv", opts, "testdata/csv2_without_headers.csv")
	if err == nil {
		t.Errorf("should return an error because csv2_without_headers.csv as no headers")
	}
}

func TestCompareReturnsDiff2(t *testing.T) {
	opts := NewOptions(true)
	opts.NoRevert()
	rows, _ := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv", "testdata/csv3.csv")
	if len(rows) != 8 {
		t.Errorf("should return 8 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[7], rowDoe80) {
		t.Errorf("line 8 should return %v got=%v", rowDoe80, rows[7])
	}
}

func TestCompareReturnsDiff3(t *testing.T) {
	opts := NewOptions(true)
	opts.NoRevert()
	rows, _ := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv", "testdata/csv3.csv", "testdata/csv4.csv")
	if len(rows) != 10 {
		t.Errorf("should return 10 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[9], rowGandalf00) {
		t.Errorf("line 10 should return %v got=%v", rowGandalf00, rows[9])
	}
}

func TestCompareReturnsDedupRevert(t *testing.T) {
	opts := NewOptions(true)
	rows, _ := Compare("testdata/csv1.csv", opts)
	if len(rows) != 6 {
		t.Errorf("should return 6 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[0], rowHead) {
		t.Errorf("line 0 should return %v got=%v", rowHead, rows[0])
	}
	if !reflect.DeepEqual(rows[5], rowBooker12) {
		t.Errorf("line 5 should return %v got=%v", rowBooker12, rows[5])
	}
}

func TestCompareReturnsDiff2Revert(t *testing.T) {
	opts := NewOptions(true)
	rows, _ := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv", "testdata/csv3.csv")
	if len(rows) != 8 {
		t.Errorf("should return 8 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[6], rowJohnson81) {
		t.Errorf("line 8 should return %v got=%v", rowJohnson81, rows[7])
	}
}

func TestCompareReturnsDiff3Revert(t *testing.T) {
	opts := NewOptions(true)
	rows, _ := Compare("testdata/csv1.csv", opts, "testdata/csv2.csv", "testdata/csv3.csv", "testdata/csv4.csv")
	if len(rows) != 10 {
		t.Errorf("should return 10 lines. got=%v, len=%d", rows, len(rows))
	}
	if !reflect.DeepEqual(rows[9], rowDoe80) {
		t.Errorf("line 10 should return %v got=%v", rowDoe80, rows[9])
	}
}
