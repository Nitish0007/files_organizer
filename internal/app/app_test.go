package app

import (
	"testing"
	// "os"
	// "path/filepath"
	// "reflect"
)

func TestNewCliApp(t *testing.T) {
	app := NewCliApp()
	if app == nil {
		t.Errorf("NewCliApp() returned nil")
	}
	var _ App = app
}

func TestPerform(t *testing.T) {
	// app := NewCliApp()
	// app.Perform("/Users/nitish/Downloads", "/Users/nitish/Downloads")
}

// func TestAddToDirectory(t *testing.T) {
// 	addToDirectory(os.DirEntry{Name: "test.txt"}, "test")
// }