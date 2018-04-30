package parsers

// SpineStack is spine represented as stack
type SpineStack spineStack

type spine struct {
	value string
	next  *spine
}

type spineStack struct {
	top *spine
}

// Value returns current spine value
func (s *SpineStack) Value() string {
	top := s.top
	if top != nil {
		return top.value
	}

	return ""
}

// Next change current stack's value and returns true
// if current value exists and next was executed
func (s *SpineStack) Next() bool {
	top := s.top
	if top != nil {
		s.pop()
		return true
	}

	return false
}

func (s *SpineStack) push(v *spine) {
	old := s.top
	v.next = old
	s.top = v
}

func (s *SpineStack) pop() *spine {
	old := s.top
	if old != nil {
		next := old.next
		s.top = next
		return old
	}
	return nil
}

// NewSpineStack creates spine stack from array
func NewSpineStack(arr []string) *SpineStack {
	spin := new(SpineStack)

	for i := len(arr) - 1; i >= 0; i-- {
		spin.push(&spine{
			value: arr[i],
			next:  nil,
		})
	}

	return spin
}
