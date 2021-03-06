// Copyright 2014-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Standard usage for uroot tools so usage works well in busybox mode.
package uroot

import (
	"flag"
	"os"
)

func Usage(cmd string) {
	defUsage := flag.Usage
	flag.Usage = func() {
		os.Args[0] = cmd
		defUsage()
	}
}
