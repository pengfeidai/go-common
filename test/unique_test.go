package main

import (
	"slice/unique"
	"testing"
)

func TestUnique(t *testing.T) {
	ftruis := []interface{}{"🍌", "🌰", "🍉", "🍊", "🍌", "🍉", "🍊", "🌰", "🍉", "🍊"}
	data := unique.Unique(ftruis)
	if len(data) == 4 {
		t.Log("[🍌 🌰 🍉 🍊]")
	}
}
