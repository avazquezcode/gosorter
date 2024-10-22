package sorting

type (
	SortOptions struct {
		IgnoreCase bool
	}

	// The implementation of this interface MUST guarantee no side effects on the original slice (input)
	Sorter interface {
		Sort(items []string) []string
	}
)
