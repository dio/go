// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	// provide a dummy error handler so parsing doesn't stop after first error
	ast, err := ParseFile(*src_, func(error) {}, nil, 0)
	if err != nil {
		t.Error(err)
	}

	if ast != nil {
		Fprint(os.Stdout, ast, true)
		fmt.Println()
	}
}

func TestPrintString(t *testing.T) {
	for _, want := range []string{
		"package p",
		"package p; type _ = int; type T1 = struct{}; type ( _ = *struct{}; T2 = float32 )",
		// TODO(gri) expand
	} {
		ast, err := Parse(nil, strings.NewReader(want), nil, nil, nil, 0)
		if err != nil {
			t.Error(err)
			continue
		}
		if got := String(ast); got != want {
			t.Errorf("%q: got %q", want, got)
		}
	}
}
