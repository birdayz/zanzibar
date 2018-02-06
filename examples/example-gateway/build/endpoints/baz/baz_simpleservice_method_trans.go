// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package bazEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
)

// SimpleServiceTransHandler is the handler for "/baz/trans"
type SimpleServiceTransHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewSimpleServiceTransHandler creates a handler
func NewSimpleServiceTransHandler(deps *module.Dependencies) *SimpleServiceTransHandler {
	handler := &SimpleServiceTransHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"baz", "trans",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *SimpleServiceTransHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/baz/trans",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/baz/trans".
func (h *SimpleServiceTransHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBazBaz.SimpleService_Trans_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := SimpleServiceTransEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		case *endpointsBazBaz.OtherAuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}

	res.WriteJSON(200, cliRespHeaders, response)
}

// SimpleServiceTransEndpoint calls thrift client Baz.Trans
type SimpleServiceTransEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w SimpleServiceTransEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_Trans_Args,
) (*endpointsBazBaz.TransStruct, zanzibar.Header, error) {
	clientRequest := convertToTransClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Baz.Trans(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertTransAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		case *clientsBazBaz.OtherAuthErr:
			serverErr := convertTransOtherAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceTransClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTransClientRequest(in *endpointsBazBaz.SimpleService_Trans_Args) *clientsBazBaz.SimpleService_Trans_Args {
	out := &clientsBazBaz.SimpleService_Trans_Args{}

	if in.Arg1 != nil {
		out.Arg1 = &clientsBazBase.TransStruct{}
		out.Arg1.Message = string(in.Arg1.Message)
		if in.Arg1.Driver != nil {
			out.Arg1.Driver = &clientsBazBase.NestedStruct{}
			out.Arg1.Driver.Msg = string(in.Arg1.Driver.Msg)
			out.Arg1.Driver.Check = (*int32)(in.Arg1.Driver.Check)
		} else {
			out.Arg1.Driver = nil
		}
		if in.Arg1.Rider != nil {
			out.Arg1.Rider = &clientsBazBase.NestedStruct{}
			out.Arg1.Rider.Msg = string(in.Arg1.Rider.Msg)
			out.Arg1.Rider.Check = (*int32)(in.Arg1.Rider.Check)
		} else {
			out.Arg1.Rider = nil
		}
	} else {
		out.Arg1 = nil
	}
	if in.Arg2 != nil {
		out.Arg2 = &clientsBazBase.TransStruct{}
		out.Arg2.Message = string(in.Arg2.Message)
		if in.Arg2.Driver != nil {
			out.Arg2.Driver = &clientsBazBase.NestedStruct{}
			out.Arg2.Driver.Msg = string(in.Arg2.Driver.Msg)
			out.Arg2.Driver.Check = (*int32)(in.Arg2.Driver.Check)
		} else {
			out.Arg2.Driver = nil
		}
		if in.Arg2.Rider != nil {
			out.Arg2.Rider = &clientsBazBase.NestedStruct{}
			if in.Arg1 != nil && in.Arg1.Driver != nil {
				out.Arg2.Rider.Msg = string(in.Arg1.Driver.Msg)
			}
			out.Arg2.Rider.Check = (*int32)(in.Arg2.Rider.Check)
		} else {
			out.Arg2.Rider = nil
		}
	} else {
		out.Arg2 = nil
	}

	return out
}

func convertTransAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertTransOtherAuthErr(
	clientError *clientsBazBaz.OtherAuthErr,
) *endpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertSimpleServiceTransClientResponse(in *clientsBazBase.TransStruct) *endpointsBazBaz.TransStruct {
	out := &endpointsBazBaz.TransStruct{}

	out.Message = string(in.Message)
	if in.Driver != nil {
		out.Driver = &endpointsBazBaz.NestedStruct{}
		out.Driver.Msg = string(in.Driver.Msg)
		out.Driver.Check = (*int32)(in.Driver.Check)
	} else {
		out.Driver = nil
	}
	if in.Rider != nil {
		out.Rider = &endpointsBazBaz.NestedStruct{}
		out.Rider.Msg = string(in.Message)
		out.Rider.Check = (*int32)(in.Rider.Check)
	} else {
		out.Rider = nil
	}

	return out
}
