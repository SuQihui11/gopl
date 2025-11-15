package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(x *IntSet) {
	for i, word := range x.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if (word & (1 << uint(j))) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	num := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for word != 0 {
			word &= word - 1
			num++
		}
	}
	return num
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var cp IntSet
	cp.words = make([]uint64, len(s.words))
	copy(cp.words, s.words)
	return &cp
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &= word
		} else {
			s.words = append(s.words, 0)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	n := len(t.words)
	for i := 0; i < len(s.words); i++ {
		if i < n {
			s.words[i] &^= t.words[i]
		} else {
			break
		}
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if (word & (1 << uint(j))) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	elems := x.Elems()
	for _, elem := range elems {
		fmt.Println(elem)
	}
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	fmt.Println(x.Len())

	x.Remove(9)
	fmt.Println(x.Len())
	fmt.Println(x.String())

	x.Clear()
	fmt.Println(x.Len())
	fmt.Println(x.String())

	y = *x.Copy()
	fmt.Println(y.String())
	fmt.Println(y.Len())

}
