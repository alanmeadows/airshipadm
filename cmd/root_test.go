package cmd_test

import (
	"testing"

	"github.com/ian-howell/airshipadm/internal/test"
)

func TestRoot(t *testing.T) {
	tests := []test.CmdTest{{
		Name:    "default",
		Command: "",
	}}
	test.RunCmdTests(t, tests)
}
