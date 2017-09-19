// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package examplegatewayServiceGenerated

import (
	"github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/services/example-gateway/module"
)

// DependenciesTree re-exported for convenience.
type DependenciesTree module.DependenciesTree

// CreateGateway creates a new instances of the example-gateway
// service with the specified config
func CreateGateway(
	config *zanzibar.StaticConfig,
	opts *zanzibar.Options,
) (*zanzibar.Gateway, interface{}, error) {
	gateway, err := zanzibar.CreateGateway(config, opts)
	if err != nil {
		return nil, nil, err
	}

	tree, dependencies := module.InitializeDependencies(gateway)
	registerErr := registerDeps(gateway, dependencies)
	if registerErr != nil {
		return nil, nil, registerErr
	}

	return gateway, (*DependenciesTree)(tree), nil
}

func registerDeps(g *zanzibar.Gateway, deps *module.Dependencies) error {
	var err error
	err = deps.Endpoint.Bar.Register(g)
	if err != nil {
		return err
	}
	err = deps.Endpoint.Baz.Register(g)
	if err != nil {
		return err
	}
	err = deps.Endpoint.BazTChannel.Register(g)
	if err != nil {
		return err
	}
	err = deps.Endpoint.Contacts.Register(g)
	if err != nil {
		return err
	}
	err = deps.Endpoint.Googlenow.Register(g)
	if err != nil {
		return err
	}
	return nil
}
