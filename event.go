package main

type Event struct {
	FromT, ToT, Desc string
}

func isEqualEvents(e1, e2 []Event) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if e1 == nil || e2 == nil {
		return false
	}

	if len(e1) != len(e2) {
		return false
	}
	if len(e1) == 0 && len(e2) == 0 {
		return true
	}

	for i, e := range e1 {
		if e != e2[i] {
			return false
		}
	}

	return true
}
