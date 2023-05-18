package csvcompare

import (
	"fmt"
	"reflect"
)

func dedupSlice(slice [][]string, opts *Options) [][]string {
	cleaned := [][]string{}
	for i, value := range slice {
		if opts.headers && i == 0 {
			cleaned = append(cleaned, value)
			// skip first line
			continue
		}
		if len(opts.idxHeader) > 0 {
			if !colsInSlice(value, cleaned, opts.idxHeader) {
				cleaned = append(cleaned, value)
			}

		} else {
			if !inSlice(value, cleaned, opts.headers) {
				cleaned = append(cleaned, value)
			}
		}
	}
	if opts.revert {
		fmt.Print("Go revert")
		cleaned = revert(cleaned, opts.headers)
	}
	return cleaned
}

// first array is the newer
func dedupSlices(slice1 [][]string, slice2 [][]string, opts *Options) ([][]string, error) {
	if opts.headers {
		if !reflect.DeepEqual(slice1[0], slice2[0]) {
			return nil, fmt.Errorf("headers not matching : %v vs %v ", slice1[0], slice2[0])
		}
		// skip headers
		slice2 = slice2[1:]
	}
	if len(slice2) > 0 {
		slice1 = append(slice1, slice2...)
	}
	return dedupSlice(slice1, opts), nil
}

func inSlice(row []string, slice [][]string, headers bool) bool {
	for i, v := range slice {
		if headers && i == 0 {
			// skip first line
			continue
		}
		if reflect.DeepEqual(row, v) {
			return true
		}
	}
	return false
}

func colsInSlice(row []string, list [][]string, cols []int) bool {
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

func revert(s [][]string, headers bool) [][]string {
	max := len(s)
	d := 1
	if headers {
		d = 0
	}
	var f = make([][]string, max)
	for i, r := range s {
		if headers && i == 0 {
			f[i] = r
		} else {
			f[max-i-d] = r
		}
	}
	return f
}
