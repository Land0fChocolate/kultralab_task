package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_resize(t *testing.T) {
	assert.Equal(t, load("images/otterresizepng.png"), resize(load("images/otterpng.png")))
}

func Test_rotate(t *testing.T) {
	assert.Equal(t, load("images/otterrotatepng.png"), rotate(load("images/otterpng.png")))
}
