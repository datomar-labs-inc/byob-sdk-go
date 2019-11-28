package byob

// Session is used to keep track of information pertaining to a user
type Session struct {
	Flaggable
	Stack  Stack  `json:"stack" binding:"required"`
}

type Frame struct {
	Module int64 `json:"m" binding:"required"`
	Node   int64 `json:"n" binding:"required"`
}

type Stack struct {
	Frames []Frame `json:"frames" binding:"required"`
}

func (s *Stack) Push(frame Frame) {
	s.Frames = append(s.Frames, frame)
}

func (s *Stack) Pop() Frame {
	if len(s.Frames) == 1 {
		panic("cannot be frameless")
	}

	frame := s.GetCurrent()
	s.Frames = s.Frames[:len(s.Frames)-1]
	return frame
}

func (s *Stack) GetCurrent() Frame {
	return s.Frames[len(s.Frames)-1]
}

func (s *Stack) SetCurrent(node int64) {
	frame := s.GetCurrent()
	s.Frames = s.Frames[:len(s.Frames)-1]
	frame.Node = node
	s.Push(frame)
}

func (s *Stack) IsInitialized() bool {
	return len(s.Frames) >= 1
}

func (s *Stack) IsOnMainGraph() bool {
	return len(s.Frames) == 1
}
