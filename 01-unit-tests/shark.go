package hunt

import "fmt"

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func NewPrey(name string, speed int) Prey {
	return Prey{
		name:  name,
		speed: speed,
	}
}

func NewShark(hungry bool, tired bool, speed int) Shark {
	return Shark{
		hungry: hungry,
		tired:  tired,
		speed:  speed,
	}
}

func (s *Shark) Hunt(p *Prey) error {
	if p == nil {
		return fmt.Errorf("there is no prey to catch")
	}
	if s.tired {
		return fmt.Errorf("cannot hunt, i am really tired")
	}
	if !s.hungry {
		return fmt.Errorf("cannot hunt, i am not hungry")
	}
	if p.speed >= s.speed {
		s.tired = true
		return fmt.Errorf("could not catch it")
	}

	s.hungry = false
	s.tired = true
	return nil
}
