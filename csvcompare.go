/*
Package **csv-compare** provides un API to compare 1 or more csv file with same structure in order
to provide a diff list of unique rows
Different option can be provided
See documentation
*/
package csvcompare

import (
	"encoding/csv"
	"fmt"
	"os"
	"unicode"
)

// Compare self or more files and returns a map of all unique lines
func Compare(srcFile string, opts *Options, filesToCompare ...string) ([][]string, error) {
	if !unicode.IsPunct(opts.comma) {
		return nil, fmt.Errorf("'opts.Comma' must be a punctuation character")
	}
	if srcFile == "" {
		return nil, fmt.Errorf("srcFile must be defined")
	}

	// open files
	fs, err := os.Open(srcFile)
	if err != nil {
		return nil, fmt.Errorf("error openning file %s. err: %v", srcFile, err)
	}
	defer fs.Close()

	srcCsvReader := csv.NewReader(fs)
	if err != nil {
		return nil, fmt.Errorf("error reading csv file %s. err: %v", srcFile, err)
	}
	srcCsvReader.Comma = opts.comma
	srcRows, err := srcCsvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s. err: %v", srcFile, err)
	}

	var finalSlice [][]string

	if len(filesToCompare) > 0 {

		finalSlice = dedupSlice(srcRows, opts)
		// prevent slices to be mixed
		revert := opts.revert
		if revert {
			opts.revert = false
		}

		for _, fToCompare := range filesToCompare {
			var compRows [][]string

			fc, err := os.Open(fToCompare)
			if err != nil {
				return nil, fmt.Errorf("error while open file %s. err: %v", fToCompare, err)
			}
			defer fc.Close()

			compCsvReader := csv.NewReader(fc)
			if err != nil {
				return nil, fmt.Errorf("error while reading csv file %s. err: %v", fToCompare, err)
			}
			compCsvReader.Comma = opts.comma

			compRows, err = compCsvReader.ReadAll()
			if err != nil {
				return nil, fmt.Errorf("error parsing file %s. err: %v", fToCompare, err)
			}

			if opts.dedup {
				compRows = dedupSlice(compRows, opts)
			}
			if revert {
				opts.revert = true
			}
			finalSlice, err = dedupSlices(finalSlice, compRows, opts)
			if err != nil {
				return nil, err
			}
		}
	} else {
		opts.dedup = true
		finalSlice = dedupSlice(srcRows, opts)
	}
	return finalSlice, nil
}
