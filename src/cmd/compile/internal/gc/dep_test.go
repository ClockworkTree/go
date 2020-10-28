// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"os/exec"
	"std/internal/testenv"
	"strings"
	"testing"
)

func TestDeps(t *testing.T) {
	out, err := exec.Command(testenv.GoToolPath(t), "list", "-f", "{{.Deps}}", "cmd/compile/internal/gc").Output()
	if err != nil {
		t.Fatal(err)
	}
	for _, dep := range strings.Fields(strings.Trim(string(out), "[]")) {
		switch dep {
		case "go/build", "go/token":
			t.Errorf("undesired dependency on %q", dep)
		}
	}
}
