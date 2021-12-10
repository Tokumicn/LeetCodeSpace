package stack

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() interface{} {
	this := *s
	if s.IsEmpty() {
		return nil
	}

	value := this[this.Len()-1]
	*s = this[:this.Len()-1]
	return value
}

func (s Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s[s.Len()-1]
}
