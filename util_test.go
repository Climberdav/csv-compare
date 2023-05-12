package csvcompare

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDedup(t *testing.T) {
	map1 := [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}}
	r := dedup(map1)
	if len(r) != 1 {
		t.Errorf("dedup should return 1. got=%d", len(r))
	}
}

func TestDedup2(t *testing.T) {
	map1 := [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}}
	r := dedup(map1)
	if len(r) != 2 {
		t.Errorf("dedup should return 2. got=%d", len(r))
	}
}

func TestDedup3(t *testing.T) {
	map1 := [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}}
	map2 := [][]string{{"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}}
	r := dedup2(map1, map2)
	if len(r) != 2 {
		t.Errorf("dedup should return 2. got=%d", len(r))
	}
}

func TestCombinedDedup(t *testing.T) {
	map1 := [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}}
	map1 = dedup(map1)
	map2 := [][]string{{"rc1", "rc2", "rc3"}, {"rc1", "rc2", "rc3"}, {"rc4", "rc5", "rc6"}}
	map2 = dedup(map2)
	r := dedup2(map1, map2)
	fmt.Print(r)
	if len(r) != 2 {
		t.Errorf("dedup should return 2. got=%d", len(r))
	}
	if !reflect.DeepEqual(r[1], []string{"rc4", "rc5", "rc6"}) {
		t.Errorf("second line should return '[rc4 rc5 rc6]'. got=%v", r[1])
	}
}

func TestDedupCol(t *testing.T) {
	map1 := [][]string{{"hcol1", "hcol2", "hcol3"}, {"rc1", "rc2", "rc3"}, {"rc1", "rc3", "rc3"}, {"rc4", "rc5", "rc6"}, {"rc4", "rc8", "rc6"}}
	r := dedupOnCol(map1, []int{0, 2})
	if len(r) != 3 {
		t.Errorf("dedup should return 3 (header + 2 lines). got=%v, len=%d", r, len(r))
	}
}
