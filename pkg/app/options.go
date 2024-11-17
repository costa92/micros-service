// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package app

import (
	cliflag "k8s.io/component-base/cli/flag"
)

type CliOptions interface {

	// Flags returns the flags for the command
	Flags() cliflag.NamedFlagSets

	// Complete completes the command
	Complete() error
	// Validate validates the command's configuration
	Validate() error
}
