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
	comma     rune
	dedup     bool // deduplication based on column unicity
	headers   bool
	idxHeader []int // a set of index on which deduplication will be processed
	revert    bool  //default value to true
}

// create a new Options, with Revert = true
func NewOptions(headers bool) *Options {
	return &Options{
		headers: headers,
		revert:  true,
		comma:   ',',
	}
}

func (o *Options) SetIndexes(idxes ...int) {
	o.idxHeader = idxes
}

func (o *Options) SetComma(c rune) {
	o.comma = c
}

func (o *Options) NoRevert() {
	o.revert = false
}

// header default value is false
func (o *Options) HasHeader() {
	o.headers = true
}
