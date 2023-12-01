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
	"log"
	"math"
	"os"
	"unicode"
)

// Compare self or more files and returns a map of all unique lines.
// When used whi one file, returned finalSlice will have uniques rows
// When used with 2 or more files, if a line is present 2 files,
// line will not be in the result slice
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
	nbLineSrc := len(srcRows)
	if opts.headers {
		nbLineSrc--
	}

	var finalSlice [][]string

	if len(filesToCompare) > 0 {
		// prevent slices to be mixed
		revert := opts.revert
		if revert {
			opts.revert = false
		}
		finalSlice = unique(srcRows, opts)
		var compRows [][]string
		for i, fToCompare := range filesToCompare {
			var compRowsFile [][]string
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

			compRowsFile, err = compCsvReader.ReadAll()
			if err != nil {
				return nil, fmt.Errorf("error parsing file %s. err: %v", fToCompare, err)
			}

			if opts.dedup {
				compRowsFile = unique(compRowsFile, opts)
			}
			if opts.headers && i != 0 {
				compRowsFile = compRowsFile[1:]
			}

			compRows = append(compRows, compRowsFile...)

		}
		if revert {
			opts.revert = true
		}
		finalSlice, err = uniqueSlices(finalSlice, compRows, opts)
		if err != nil {
			return nil, err
		}

	} else {
		opts.dedup = true
		finalSlice = unique(srcRows, opts)
	}
	nbLineFinal := len(finalSlice)
	if opts.headers {
		nbLineFinal--
	}
	log.Printf("There are %d lines left out of the %d. %v%% deduplication.", nbLineFinal, nbLineSrc,
		math.Round(100*float64(nbLineFinal)/float64(nbLineSrc)))
	return finalSlice, nil
}
