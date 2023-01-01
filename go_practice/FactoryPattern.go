package model

// define a struct
type student struct {
	Name  string
	score float64
}

// encapsulate
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

//then import this package and call NewStudent

func (s *student) GetScore() float64 {
	return s.score //ok , if this variable is private then use this method to get value(if not same directory)
}
