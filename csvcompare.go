package csvcompare

import (
	"encoding/csv"
	"fmt"
	"os"
	"unicode"
)

// Compare 2 files and returns a map of all unique lines in both files
// if comparedFile is empty, compares itself and deduplicate lines
func Compare(srcFile, comparedFile string, opts Options) ([][]string, error) {
	if !unicode.IsPunct(opts.Comma) {
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
	srcCsvReader.Comma = opts.Comma
	srcRows, err := srcCsvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s. err: %v", srcFile, err)
	}

	var compRows [][]string

	if comparedFile != "" {
		fc, err := os.Open(comparedFile)
		if err != nil {
			return nil, fmt.Errorf("error openning file %s. err: %v", srcFile, err)
		}
		defer fc.Close()

		compCsvReader := csv.NewReader(fc)
		if err != nil {
			return nil, fmt.Errorf("error reading csv file %s. err: %v", srcFile, err)
		}
		compCsvReader.Comma = opts.Comma

		compRows, err = compCsvReader.ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error parsing file %s. err: %v", srcFile, err)
		}

		if opts.Dedup {
			compRows = dedup(compRows)
		}

	} else {
		opts.Dedup = true
	}

	if opts.Dedup {
		srcRows = dedup(srcRows)
	}

	// find lines not present in comparedFile case
	if comparedFile != "" {
		return dedup2(srcRows, compRows), nil
	} else {
		return srcRows, nil
	}
}

// func CompareMutliple(srcFile string, opts Options, f ...any) ([]map[string]string, error) {
// 	var finalRows []map[string]string

// 	return finalRows, nil
// }
