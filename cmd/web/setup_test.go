package main

import (
	"os"
	"testing"

	"github.com/YoungsoonLee/design-pattern-go/config"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		App: config.New(nil),
	}

	os.Exit(m.Run())
}
