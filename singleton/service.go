package singleton

type IdService struct {
	counter int
}

func NewIdService() *IdService {
	return &IdService{counter: 0}
}

func (s *IdService) Next() int {
	s.counter++
	return s.counter
}
