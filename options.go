package csvcompare

type Comma int

const (
	COMMA Comma = iota
	SEMICOLON
	TAB
)

func (c Comma) String() rune {
	return []rune{',', ';', '\t'}[c]
}

type Options struct {
	Comma     rune
	Dedup     bool // deduplication based on column unicity
	NoHeader  bool
	IdxHeader []int // a set of index on which deduplication will be processed
}
