// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package kratos_middleware

import (
	"errors"
	httpNet "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var ErrUnknownCodec = errors.New("codec not supported")

type response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Ts      string      `json:"ts"`
}

func ResponseEncoder() http.ServerOption {
	return http.ResponseEncoder(func(w httpNet.ResponseWriter,
		r *httpNet.Request,
		i interface{}) error {
		reply := &response{
			Code: 200,
			Data: i,
			Ts:   time.Now().Local().Format(time.RFC3339), // 使用更标准的时间格式
		}
		codec := encoding.GetCodec("json")
		if codec == nil {
			return ErrUnknownCodec
		}
		data, err := codec.Marshal(reply)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8") // 明确指定字符集
		_, err = w.Write(data)
		return err
	})
}
