package trie

import (
    "errors"
)

const (
	trieMaxQueueSize = 300
)

var (
    ErrNotInTable = errors.New("not in character table predefined")
    ErrDuplicate  = errors.New("duplicate insert")
)

type TrieValue interface {
}

type TrieNode struct {
	Value        interface{}
	failed_clue  *TrieNode
	next         []*TrieNode
}

type TrieTree struct {
	root         TrieNode
    branch       uint16
	keynum       int
    table        [256]uint16
}

func NewTrieTree(table []byte) *TrieTree {
	tt := new(TrieTree)
    for i := 0; i < 256; i++ {
        tt.table[i] = 0xffff
    }

    for i, c := range table {
        tt.table[c] = uint16(i)
    }

    tt.branch = uint16(len(table))
	return tt
}

func (tt *TrieTree) Insert(key []byte, value interface{}) (*TrieNode, error) {
	p := &tt.root
	for _, c := range key {
        index := tt.table[c]
        if index >= tt.branch {
            return nil, ErrNotInTable
        }

		if p.next == nil {
			p.next = make([]*TrieNode, tt.branch)
		}

		if p.next[index] == nil {
			p.next[index] = new(TrieNode)
		}
		p = p.next[index]
	}

    if p.Value != nil {
        return p, ErrDuplicate
    }

    p.Value = value
    tt.keynum++
	return p, nil
}

func (tt *TrieTree) BuileClue() {
	root := &tt.root
	var head, tail int
	var q [trieMaxQueueSize]*TrieNode
	q[head] = root
	head++
	for head != tail {
		p := q[tail]
		tail = (tail + 1) % trieMaxQueueSize

		if p.next == nil {
			continue
		}

		for c, t := range p.next {
			if t == nil {
				continue
			}

			if p == root {
				t.failed_clue = root
				q[head] = t
				head = (head + 1) % trieMaxQueueSize
				continue
			}

			tt := p.failed_clue
			for tt != nil {
				if tt.next != nil && tt.next[c] != nil {
					t.failed_clue = tt.next[c]
					break
				}
				tt = tt.failed_clue
			}

			if tt == nil {
				t.failed_clue = root
			}

			q[head] = t
			head = (head + 1) % trieMaxQueueSize
		}
	}
}

func (tt *TrieTree) Query(block []byte) interface{} {
	root := &tt.root
	p := root

	for _, c := range block {
		index := tt.table[c]
        if index >= tt.branch {
            p = root
            continue
        }

        for p.next == nil {
			if p == root {
				break
			}
			p = p.failed_clue
		}

		p = p.next[index]
		if p == nil {
			p = root
		}

		return p.Value
	}

	return nil
}
