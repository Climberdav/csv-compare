package csvcompare

import (
	"reflect"
)

func dedup(rows [][]string) [][]string {
	cleaned := [][]string{}
	for _, value := range rows {
		if !sliceInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}

func sliceInSlice(row []string, list [][]string) bool {
	for _, v := range list {
		if reflect.DeepEqual(row, v) {
			return true
		}
	}
	return false
}

func dedup2(rows [][]string, rows2 [][]string) [][]string {
	rows = append(rows, rows2...)
	return dedup(rows)
}

func dedupOnCol(rows [][]string, cols []int) [][]string {
	cleaned := [][]string{}
	for _, value := range rows {
		if !sliceColInSlice(value, cleaned, cols) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}

func sliceColInSlice(row []string, list [][]string, cols []int) bool {
	// prepare a slice of unique values to compare
	u1 := []string{}
	for _, c := range cols {
		u1 = append(u1, row[c])
	}

	for _, v := range list {
		var u2 []string
		for _, c := range cols {
			u2 = append(u2, v[c])
		}
		if reflect.DeepEqual(u2, u1) {
			return true
		}
	}
	return false
}
