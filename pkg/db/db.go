// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package db

import (
	"github.com/google/wire"
	redis "github.com/redis/go-redis/v9"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(NewMySQL, NewRedis, wire.Bind(new(redis.UniversalClient), new(*redis.Client)))
