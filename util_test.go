package csvcompare

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDedup(t *testing.T) {
	optsWithoutHeaders := NewOptions(false)
	optsWithoutHeaders.NoRevert()
	optsWithHeaders := NewOptions(true)
	optsWithHeaders.NoRevert()
	optsWithoutHeadersNoRevert := NewOptions(false)
	optsWithoutHeadersNoRevert.NoRevert()
	optsWithHeadersNoRevert := NewOptions(true)
	optsWithHeadersNoRevert.NoRevert()
	optsColWithHeaders := NewOptions(true)
	optsColWithHeaders.NoRevert()
	optsColWithHeaders.SetIndexes(0, 2)
	optsColWithoutHeaders := NewOptions(true)
	optsColWithoutHeaders.NoRevert()
	optsColWithoutHeaders.SetIndexes(0, 2)

	tests := []struct {
		name          string
		sliceIn1      [][]string
		sliceIn2      [][]string
		line          []string
		lineHeaders   []string
		lenAttendee   int
		lineToCompare int
		opts          *Options
		errAttendee   bool
	}{
		{
			name:          "Test 1",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc1", "rc2", "rc3"},
			lenAttendee:   1,
			lineToCompare: 1,
			opts:          optsWithoutHeaders,
		},
		{
			name:          "Test 2",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc1", "rc2", "rc3"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   2,
			lineToCompare: 2,
			opts:          optsWithHeaders,
		},
		{
			name:          "Test 3",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc4", "rc5", "rc6"},
			lenAttendee:   2,
			lineToCompare: 2,
			opts:          optsWithoutHeaders,
		},
		{
			name:          "Test 4",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   3,
			lineToCompare: 3,
			opts:          optsWithHeaders,
		},
		{
			name:          "Test 5",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}},
			sliceIn2:      [][]string{{"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{},
			lenAttendee:   1,
			lineToCompare: 1,
			opts:          optsWithoutHeaders,
			errAttendee:   false,
		},
		{
			name:          "Test 6",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}},
			sliceIn2:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   2,
			lineToCompare: 2,
			opts:          optsWithHeaders,
		},
		{
			name:          "Test 7",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}},
			sliceIn2:      [][]string{{"h1", "h2", "h4"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   3,
			lineToCompare: 3,
			opts:          optsWithHeaders,
			errAttendee:   true,
		},
		{
			name:          "Test 8",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{},
			lenAttendee:   4,
			lineToCompare: 3,
			opts:          optsColWithoutHeaders,
			errAttendee:   false,
		},
		{
			name:          "Test 9",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc4", "rc5", "rc6"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   5,
			lineToCompare: 4,
			opts:          optsColWithHeaders,
			errAttendee:   false,
		},
		{
			name:          "Test 11",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc1", "rc3", "rc3"},
			lineHeaders:   []string{},
			lenAttendee:   4,
			lineToCompare: 2,
			opts:          optsColWithoutHeaders,
			errAttendee:   false,
		},
		{
			name:          "Test 12",
			sliceIn1:      [][]string{{"h1", "h2", "h3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc1", "rc3", "rc3"},
			lineHeaders:   []string{"h1", "h2", "h3"},
			lenAttendee:   5,
			lineToCompare: 3,
			opts:          optsColWithHeaders,
			errAttendee:   false,
		},
		{
			name:          "Test 13",
			sliceIn1:      [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}},
			sliceIn2:      [][]string{},
			line:          []string{"rc4", "rc8", "rc6"},
			lineHeaders:   []string{},
			lenAttendee:   4,
			lineToCompare: 4,
			opts:          optsWithoutHeadersNoRevert,
			errAttendee:   false,
		},
	}

	for _, tt := range tests {
		fmt.Println("=========")
		fmt.Printf("%s\n", tt.name)
		fmt.Println("=========")
		var r [][]string
		var err error
		if len(tt.sliceIn2) == 0 {
			r = unique(tt.sliceIn1, tt.opts)
		} else {
			r, err = uniqueSlices(tt.sliceIn1, tt.sliceIn2, tt.opts)
		}
		if tt.errAttendee && err == nil {
			t.Errorf("%s should return an error. got=%v", tt.name, r)
		} else if !tt.errAttendee {
			if len(r) != tt.lenAttendee {
				t.Errorf("%s should return %v. got=%v,  len=%d", tt.name, tt.lenAttendee, r, len(r))
			}
			if !reflect.DeepEqual(tt.line, r[tt.lineToCompare-1]) {
				t.Errorf("%s, line %d should return '%v'. got=%v", tt.name, tt.lineToCompare, tt.line, r[tt.lineToCompare-1])
			}
			if tt.opts.headers && len(tt.lineHeaders) > 0 {
				if !reflect.DeepEqual(tt.lineHeaders, r[0]) {
					t.Errorf("%s headers should return '%v'. got=%v", tt.name, tt.lineHeaders, r[0])
				}
			}
		}

	}

}
