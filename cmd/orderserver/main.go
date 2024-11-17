// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package main

import (
	// Importing the package to automatically set GOMAXPROCS.
	_ "go.uber.org/automaxprocs/maxprocs"

	"github.com/costa92/micros-service/cmd/orderserver/app"
)

func main() {
	app.NewApp().Run()
}
