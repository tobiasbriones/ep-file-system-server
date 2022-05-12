// Copyright (c) 2022 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/ep-file-system-server

package server

import (
	"fmt"
	"os"
)

func getFilePath(relPath string) string {
	return fmt.Sprintf("%v%v%v", fsRootPath, os.PathSeparator, relPath)
}
