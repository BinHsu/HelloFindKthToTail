package main

import (
	"container/list"
	"fmt"
)

// simulator with sigle way list
func findKthToTail(l *list.List, k int) *list.Element {
	if nil == l {
		return nil
	}

	if k < 0 {
		return nil
	}

	pAhead := l.Front()

	i := 0
	for e := l.Front(); e != nil && i < k; e = e.Next() {
		pAhead = e
		i += 1
	}

	if i != k {
		return nil
	}

	pBehind := l.Front()
	for e := pAhead.Next(); e != nil; e = e.Next() {
		pBehind = pBehind.Next()
	}
	return pBehind
}

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack("str4")
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	//l *List := nil

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	eT2 := findKthToTail(l, 2)
	if eT2 != nil {
		fmt.Println(eT2.Value)
	}

	eT4 := findKthToTail(l, 4)
	if eT4 != nil {
		fmt.Println(eT4.Value)
	}

	eT5 := findKthToTail(l, 5)
	if eT5 != nil {
		fmt.Println("eT5 result is not expected")
	}

	eT0 := findKthToTail(l, 0)
	if eT0 != nil {
		fmt.Println(eT0.Value)
	}

	eTneg1 := findKthToTail(l, -1)
	if eTneg1 != nil {
		fmt.Println("eTneg1 result is not expected")
	}
}
