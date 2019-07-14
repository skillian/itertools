package itertools

// ArrayIterator iterates over an Arrayer.
type ArrayIterator struct {
	Array Arrayer
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
