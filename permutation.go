package itertools

// Permuter creates permutations of values.
type Permuter struct {
	// Iterators is the collection of iterators that the Permuter
	// pulls values from in big-endian order (i.e. The first Iterator in
	// the slice is the one incremented every call to Next).
	Iterators []Iterator

	// next holds a function called whenever (*Permuter).Next is called.
	// The first call to Next needs to initialize the Iterators and so
	// is initialized with the function, (*Permuter).initialNext and
	// after that calls (*Permuter).subsequentNext.
	next func(p *Permuter, pps []interface{}) bool
}

// NewPermuter creates a new "permuter."
func NewPermuter(iterators ...Iterator) *Permuter {
	dups := make([]Iterator, len(iterators))
	copy(dups, iterators)
	perm := &Permuter{Iterators: dups}
	perm.next = (*Permuter).initialNext
	return perm
}

// Next gets the next permutation of values.  It returns false when there are no
// more results.  The first call gets the initial values.
func (p *Permuter) Next(pps ...interface{}) bool {
	return p.next(p, pps)
}

// initialNext initializes all of the Permuter's Iterators and the input
// pointers.
func (p *Permuter) initialNext(pps []interface{}) bool {
	for i, it := range p.Iterators {
		if !it.Next(pps[i]) {
			return false
		}
	}
	p.next = (*Permuter).subsequentNext
	return true
}

// Next puts the next set of values into the pointers given in pps.  This
// effectively works as an odometer, rolling over the values in the pointers.
// the values in the targets of those pointers should start at zero values
// (as set by (*Permuter).initialNext).
func (p *Permuter) subsequentNext(pps []interface{}) bool {
	for i, it := range p.Iterators {
		pp := pps[i]
		if it.Next(pp) {
			return true
		}
		it.Reset()
		if !it.Next(pp) {
			panic("set is empty")
		}
	}
	return false
}

// Reset the permutation engine.
func (p *Permuter) Reset() {
	for _, it := range p.Iterators {
		it.Reset()
	}
	p.next = (*Permuter).initialNext
}
