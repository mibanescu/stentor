// Copyright © 2020 The Stentor Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
)

const (
	appName = "stentor"

	// exit codes
	successExitCode  = 0
	genericErrorCode = 1
)

var (
	version = "dev"
	commit  = "none"
	date    = "none"
)

func main() {
	// get the current wd
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to get current directory: %v\n", err)
		os.Exit(genericErrorCode)
	}

	os.Exit(New(wd, os.Args, os.Environ(), os.Stderr, os.Stdout).Run())
}
