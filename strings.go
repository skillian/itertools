package itertools

// Strings is a Slicer implementation whose underlying datatype is []string
type Strings []string

// At puts the address to the given index into the double-pointer, pp.
func (ss *Strings) At(i int, pp interface{}) {
	pps := pp.(**string)
	*pps = &(*ss)[i]
}

// Cap gets the capacity of the strings slice,
func (ss *Strings) Cap() int { return cap(*ss) }

// Len gets the number of items in the
func (ss *Strings) Len() int { return len(*ss) }

// Append the given slice of strings to the collection.
func (ss *Strings) Append(ps interface{}) {
	*ss = Strings(append([]string(*ss), ps.([]string)...))
}
