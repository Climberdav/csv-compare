package csvcompare

import (
	"fmt"
	"reflect"
	"strings"
)

func unique(slice [][]string, opts *Options) [][]string {
	// Create a map to store unique lines
	seen := map[string]bool{}

	// Create a map to store unique lines
	cleaned := [][]string{}

	// go throw each lines
	for i, line := range slice {
		if i == 0 && opts.headers {
			// skip first header line
			cleaned = append(cleaned, line)
			continue
		}
		//Create a unique key for the line
		key := strings.Join(line, "_")

		// If key is unique, add it
		if _, ok := seen[key]; !ok {
			seen[key] = true
			cleaned = append(cleaned, line)
		}
	}

	if opts.revert {
		cleaned = revert(cleaned, opts.headers)
	}
	return cleaned
}

// first array is the newer
func uniqueSlices(slice1 [][]string, slice2 [][]string, opts *Options) ([][]string, error) {
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

	// Create a map to store unique lines
	seen := map[string]int{}

	// Create a map to store unique lines
	cleaned := [][]string{}

	// go throw each lines
	for i, line := range slice1 {
		if i == 0 && opts.headers {
			// skip first header line
			cleaned = append(cleaned, line)
			continue
		}
		//Create a unique key for the line
		key := strings.Join(line, "_")

		// If key is unique, add it
		_, ok := seen[key]
		if !ok {
			seen[key] = 1
		} else {
			seen[key]++
		}
	}

	for _, line := range slice1 {
		key := strings.Join(line, "_")
		if v, _ := seen[key]; v == 1 {
			cleaned = append(cleaned, line)
		}
	}

	if opts.revert {
		cleaned = revert(cleaned, opts.headers)
	}

	return unique(cleaned, opts), nil
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
