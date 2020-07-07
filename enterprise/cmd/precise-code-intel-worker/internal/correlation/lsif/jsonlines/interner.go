package jsonlines

import (
	"strconv"
	"sync"
)

type Interner struct {
	sync.RWMutex
	m map[string]int
}

func NewInterner() *Interner {
	return &Interner{
		m: map[string]int{},
	}
}

func (i *Interner) Intern(raw []byte) (int, error) {
	if len(raw) == 0 {
		return 0, nil
	}

	if raw[0] != '"' {
		return strconv.Atoi(string(raw))
	}

	s := string(raw[1 : len(raw)-1])
	if v, err := strconv.Atoi(s); err == nil {
		return v, nil
	}

	i.RLock()
	v, ok := i.m[s]
	i.RUnlock()
	if ok {
		return v, nil
	}

	i.Lock()
	defer i.Unlock()

	v, ok = i.m[s]
	if !ok {
		v = len(i.m) + 1
		i.m[s] = v
	}

	return v, nil
}
