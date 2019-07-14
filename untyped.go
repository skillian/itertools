package itertools

// ArrayIterator iterates over an Array.
type ArrayIterator struct {
	// Array holds an array of values to iterate over.
	Array Arrayer

	// Index is the index of the next value to return.
	Index int
}

// Next implements Iterator.
func (i *ArrayIterator) Next(pp interface{}) bool {
	if i.Index >= i.Array.Len() {
		return false
	}
	i.Array.At(i.Index, pp)
	i.Index++
	return true
}

// Reset the iterator.
func (i *ArrayIterator) Reset() {
	i.Index = 0
}
