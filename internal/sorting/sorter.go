package sorting

type (
	SortOptions struct {
		IgnoreCase bool
	}

	Sorter interface {
		Sort(items []string) []string
	}
)
