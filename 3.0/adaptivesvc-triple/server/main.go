/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	"github.com/dubbogo/dubbo-go-benchmark/3.0/adaptivesvc-triple/api"
	"time"
)

var ErrNShouldGreaterThanZero = fmt.Errorf("n should greater than zero")

func main() {
	config.SetProviderService(&Provider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}

type Provider struct {
	api.UnimplementedProviderServer
}

func (p *Provider) Fibonacci(_ context.Context, req *api.FibonacciRequest) (*api.FibonacciResult, error) {
	ret, err := fibonacci(req.N)
	if err != nil {
		return nil, err
	}
	return &api.FibonacciResult{
		Result: ret,
	}, nil
}

func (p *Provider) Sleep(_ context.Context, req *api.SleepRequest) (*api.SleepResult, error) {
	time.Sleep(time.Duration(req.Time))
	return &api.SleepResult{
		Ret: 1,
	}, nil
}

func fibonacci(n int64) (int64, error) {
	if n < 0 {
		return 0, ErrNShouldGreaterThanZero
	}
	if n < 2 {
		return n, nil
	}

	f1, err := fibonacci(n - 1)
	if err != nil {
		return 0, err
	}
	f2, err := fibonacci(n - 2)
	if err != nil {
		return 0, err
	}

	return f1 + f2, nil
}