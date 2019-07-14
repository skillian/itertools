package itertools

// Permuter creates permutations of values.
type Permuter struct {
	// A little-endian collection of iterators.
	Iterators []Iterator
	next      func(p *Permuter, pps []interface{}) bool
}

// NewPermuter creates a new "permuter."
func NewPermuter(iterators ...Iterator) *Permuter {
	dups := make([]Iterator, len(iterators))
	copy(dups, iterators)
	perm := &Permuter{Iterators: iterators}
	perm.next = (*Permuter).initialNext
	return perm
}

// Next gets the next permutation of values.  It returns false when there are no
// more results.  The first call gets the initial values.
func (p *Permuter) Next(pps ...interface{}) bool {
	return p.next(p, pps)
}

// Init both initializes the target pointer values and the underlying iterators.
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
// the values in the targets of those pointers should start at zero values.
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
