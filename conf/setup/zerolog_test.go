package setup

import (
	"runtime/debug"
	"testing"
)

func TestInit(t *testing.T) {
	Zerolog("common")
	Logger("common").Info().Interface("stack", string(debug.Stack())).Send()
}
