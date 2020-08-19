package pongo2

// Options allow you to change the behavior of template-engine.
// You can change the options before calling the Execute method.
type Options struct {
	// If this is set to true the first newline after a block is removed (block, not variable tag!). Defaults to false.
	TrimBlocks bool

	// If this is set to true leading spaces and tabs are stripped from the start of a line to a block. Defaults to false
	LStripBlocks bool

	// If true, parse all numbers as float64 values. Otherwise, it depends
	// whether or not a dot is placed after the digits. Without the dot, the
	// numbers are parsed as integers.
	ForceFloat64 bool
}

func newOptions() *Options {
	return &Options{
		TrimBlocks:   false,
		LStripBlocks: false,
		ForceFloat64: false,
	}
}

// Update updates this options from another options.
func (opt *Options) Update(other *Options) *Options {
	opt.TrimBlocks = other.TrimBlocks
	opt.LStripBlocks = other.LStripBlocks
	opt.ForceFloat64 = other.ForceFloat64

	return opt
}
