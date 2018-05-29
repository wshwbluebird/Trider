package Trider

import (
	"testing"
)

func TestTrider_Run(t *testing.T) {
	trider := NewTrider(3)
	trider.Run()
}
