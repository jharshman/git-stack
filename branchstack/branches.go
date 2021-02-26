package branchstack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/jharshman/simplestack"
)

const dotfile = ".git-stack"

var branchStack *simplestack.Stack
var once sync.Once

func initializeStack(branches []string) {
	once.Do(func() {
		branchStack = &simplestack.Stack{}
		for _, b := range branches {
			branchStack.Push(b)
		}
	})
}

func MustRead() {
	var branches []string
	f, err := os.Open(dotfile)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		goto INIT
	} else if err != nil {
		panic(err)
	}
	branches, err = read(f)
	if err != nil {
		panic(err)
	}
INIT:
	initializeStack(branches)
}

func read(r io.Reader) ([]string, error) {
	var branches []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Text()
		branches = append(branches, b)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return branches, nil
}

func MustWrite() {
	f, err := os.OpenFile(dotfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = write(f)
	if err != nil {
		panic(err)
	}
}

func write(w io.Writer) error {
	var branches []string
	for {
		item := branchStack.Pop()
		if item == nil {
			break
		}
		branches = append(branches, item.Data.(string))
	}
	if len(branches) == 0 {
		return nil
	}
	writer := bufio.NewWriter(w)
	for _, branch := range branches {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", branch))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func PushBranch(branch string) {
	branchStack.Push(branch)
}

func PopBranch() string {
	item := branchStack.Pop()
	if item != nil {
		return item.Data.(string)
	}
	return ""
}

func PeekBranch() string {
	item := branchStack.Peek()
	return item.Data.(string)
}