package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha512(t *testing.T) {
	assert.Equal(t, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff", Sha512("test"))
	assert.Equal(t, "67fc42bd109b374ae5934df3461e5e8b112476f05da400b9f345222f6d2eecddba2dd8e03c0938d08df26de41188795dd02b10cde41c3e9710ed025ccaf712ec", Sha512("string to convert"))
}
