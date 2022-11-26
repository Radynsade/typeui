package typeui

import "fmt"

const (
	ListTypeAlphabet = iota
	ListTypeNumbered
	ListTypeUnordered
)

type ListType = int

type List struct {
	Items []string
	Type  int
	// Dots if true, brackets if false
	DotsOrBrackets bool
}

func NewList(
	// Check constants starting as 'ListType' for available types.
	listType ListType,
	// Dots if true, brackets if false.
	dotsOrBrackets bool,
	items ...string,
) *List {
	return &List{
		Items:          items,
		DotsOrBrackets: dotsOrBrackets,
		Type:           listType,
	}
}

func (l *List) Print() {
	var dotOrBracket string

	if l.DotsOrBrackets {
		dotOrBracket = "."
	} else {
		dotOrBracket = ")"
	}

	if l.Type == ListTypeNumbered {
		for i, item := range l.Items {
			fmt.Printf("%d%s %s\n", i+1, dotOrBracket, item)
		}
	}

	if l.Type == ListTypeAlphabet {
		for i, item := range l.Items {
			fmt.Printf("%c%s %s\n", 'a'+i, dotOrBracket, item)
		}
	}

	if l.Type == ListTypeUnordered {
		for _, item := range l.Items {
			fmt.Printf("- %s\n", item)
		}
	}
}
