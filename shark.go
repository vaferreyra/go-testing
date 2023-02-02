package hunt

import (
	"errors"
)

var (
	ErrTiredShark     = errors.New("cannot hunt, i am really tired")
	ErrSharkNotHungry = errors.New("cannot hunt, i am not hungry")
	ErrCantCatch      = errors.New("could not catch it")
)

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func (s *Shark) Hunt(p *Prey) error {
	if s.tired {
		return ErrTiredShark
	}
	if !s.hungry {
		return ErrSharkNotHungry
	}
	if p.speed >= s.speed {
		s.tired = true
		return ErrCantCatch
	}

	s.hungry = false
	s.tired = true
	return nil
}
