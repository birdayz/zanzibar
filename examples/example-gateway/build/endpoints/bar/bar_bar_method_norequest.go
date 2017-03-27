// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"
)

// HandleNoRequestRequest handles "/bar/no-request-path".
func HandleNoRequestRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}

	// Handle request body.
	clientResp, err := clients.Bar.NoRequest(ctx, h)
	if err != nil {
		req.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		res.Flush()
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	if !res.IsOKResponse(clientResp.StatusCode, []int{200}) {
		req.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	b, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		res.SendError(500, errors.Wrap(err, "could not read client response body:"))
		res.Flush()
		return
	}
	var clientRespBody bar.BarResponse
	if err := clientRespBody.UnmarshalJSON(b); err != nil {
		res.SendError(500, errors.Wrap(err, "could not unmarshal client response body:"))
		res.Flush()
		return
	}
	response := convertNoRequestClientResponse(&clientRespBody)
	res.WriteJSON(clientResp.StatusCode, response)
	res.Flush()
}

func convertNoRequestClientResponse(body *bar.BarResponse) *bar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := &bar.BarResponse{}
	return downstreamResponse
}
