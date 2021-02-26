package branchstack

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_read(t *testing.T) {
	in := []byte(`hello
world
gophers
rock
`)
	branches, err := read(bytes.NewReader(in))
	if err != nil {
		t.Error(err)
	}

	numberOfBranches := len(branches)
	if numberOfBranches != 4 {
		t.Error(fmt.Sprintf("wanted 4, got %d\n", numberOfBranches))
	}
}

func Test_initializeStack(t *testing.T) {
	branches := []string{"foo", "bar", "fiz", "buz"}
	initializeStack(branches)
	top := branchStack.Peek()
	if top.Data.(string) != "buz" {
		t.Errorf("got %v, expected buz\n", top)
	}
}

func Test_write(t *testing.T) {
	branches := []string{"foo", "bar", "fiz", "buz"}
	initializeStack(branches)
	var b bytes.Buffer
	err := write(&b)
	if err != nil {
		t.Error(err)
	}
	if b.String() != `buz
fiz
bar
foo
` {
		t.Errorf("got: %s\n", b.String())
	}
}