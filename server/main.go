// Copyright (c) 2022 Tobias Briones. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause
// This file is part of https://github.com/tobiasbriones/ep-file-system-server

package main

import (
	"fmt"
)

const (
	port    = 8080
	network = "tcp"
	bufSize = 1024
)

func main() {

}

func getServerAddress() string {
	return fmt.Sprintf("localhost:%v", port)
}