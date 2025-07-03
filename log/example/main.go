package main

import (
	"errors"

	"prabha/go-util/log"
)

func main() {
	config := log.NewConfig("TestApp")
	l := log.New(config)
	l.Debug("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
	l.Debug("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
	l.Info("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
	l.Warn("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
	l.Error("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
	l.Fatalf("This is for testing. Str: %s, Int: %d, Err: %+v", "Hello", 100, errors.New("no rows"))
}
