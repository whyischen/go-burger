package main

import (
	"errors"
	"sync"
)

func sMap() error {
	sMap := sync.Map{}

	sMap.Store("a", "111")
	value, ok := sMap.Load("a")
	if !ok || value != "111" {
		return errors.New("load error")
	}
	return nil
}
