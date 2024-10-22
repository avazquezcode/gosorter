package sorting

type (
	SortOptions struct {
		IgnoreCase bool
	}

	Sorter interface {
		// this method might add side effects on the original list passed as an input, depending on the sorting algorithm being used
		// if you need to guarantee no side effects, please copy your list in advance and pass a copy
		Sort(items []string) []string
	}
)
