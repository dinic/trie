package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	//ascii test
	var ascii [256]byte
	for i := 0; i < 256; i++ {
		ascii[i] = byte(i)
	}

	tt := NewTrieTree(ascii[0:256])
	if _, err := tt.Insert([]byte("abcd"), 1); err != nil {
		t.Errorf("insert \\abcd\\ faild, error %s", err)
	}

	if _, err := tt.Insert([]byte("12345"), 2); err != nil {
		t.Errorf("insert \\abcdef\\ faild, error %s", err)
	}

	if node, err := tt.Insert([]byte("abcd"), 3); err != ErrDuplicate {
		t.Errorf("duplicate inster, but return %s", err)
	} else {
		node.Value = 3
	}

	//call BuileClue after insert all keys
	tt.BuileClue()

	if value := tt.Query([]byte("abcdefghijklmnopqrstuvwxyz01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")); value != nil {
		if v, ok := value.(int); !ok || v != 3 {
			t.Errorf("not correct value")
		}
	}

	//hex test
	hex := []byte("0123456789abcdef")
	tt = NewTrieTree(hex)
	if _, err := tt.Insert([]byte("1a2b3c4d5f"), "1a2b3c4d5f"); err != nil {
		t.Errorf("insert \\1a2b3c4d5f\\ faild, error %s", err)
	}

	if _, err := tt.Insert([]byte("xyz"), "xyz"); err != ErrNotInTable {
		t.Errorf("err insert but return %s", err)
	}
}
