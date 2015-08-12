// Copyright 2015 The Sconf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ini // import "gopkg.in/sconf/ini.v0"

import (
	"io/ioutil"

	"gopkg.in/sconf/internal.v0/internal_"
)

func File(name string) internal.IdemReader {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return internal.ErrIdemReader{Err: err}
	}
	return internal.BytesIdemReader(b)
}

func Text(data string) internal.IdemReader {
	return internal.BytesIdemReader([]byte(data))
}
