// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datstats

import (
	"path/filepath"
	"strings"

	"code.gitea.io/gitea/modules/log"
)

// IsDATFile reports whether name looks like a DAT file
// based on its name.
func IsDATFile(name string) bool {
	var extension = filepath.Ext(name)
	log.Debug("IsDATFile? %s", extension)
	return extension == strings.ToLower(".dat")
}
