package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecoverGoRoutineWithHandler(t *testing.T) {
	value := ""
	expectedValue := "foo"
	function := func() {
		defer RecoverGoRoutineWithHandler("foo", func() {
			value = expectedValue
		})
		panic("foo")
	}

	function()

	assert.Equal(t, expectedValue, value)
}

func TestRecoverGoRoutine(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	function := func() {
		defer RecoverGoRoutine("foo")
		panic("foo")
	}

	function()
}
