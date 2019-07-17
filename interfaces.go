package itertools

// Caper gets a collection's capacity
type Caper interface {
	Cap() int
}

// Lener gets a collection's length
type Lener interface {
	Len() int
}

// Arrayer wraps an array.
type Arrayer interface {
	// At gets a pointer to the i'th element and writes that pointer into
	// the double-pointer: pp.
	At(i int, pp interface{})

	Caper
	Lener
}

// Slicer wraps a slice.
type Slicer interface {
	Arrayer

	// Append a slice of values to the Slicer.
	Append(ps interface{})
}

// Iterator walks through elements in a collection.
type Iterator interface {
	// Next gets the address of the next item in the iterator.  The first
	// call gets the first element.
	Next(pp interface{}) bool

	// Reset the iterator to start over.
	Reset()
}
