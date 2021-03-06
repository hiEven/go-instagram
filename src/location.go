package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/auth"
	"github.com/hieven/go-instagram/src/utils/request"
)

type location struct {
	authManager    auth.AuthManager
	requestManager request.RequestManger
}

func (location *location) Feed(ctx context.Context, req *LocationFeedRequest) (*protos.LocationFeedResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.LocationFeedEndpoint, req.Pk)) // TODO: handle error
	query := urlStru.Query()
	query.Set("rank_token", location.authManager.GenerateUUID())

	urlStru.RawQuery = query.Encode()

	_, body, _ := location.requestManager.Get(ctx, urlStru.String()) // TODO: handle error

	result := &protos.LocationFeedResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}

func (location *location) Section(ctx context.Context, req *LocationSectionRequest) (*protos.LocationSectionResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.LocationSectionEndpoint, req.Pk)) // TODO: handle error

	internalReq := &protos.LocationSectionRequest{
		UUID: location.authManager.GenerateUUID(),
		Tab:  string(req.Tab),
	}

	_, body, _ := location.requestManager.Post(ctx, urlStru.String(), internalReq)

	result := &protos.LocationSectionResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}
