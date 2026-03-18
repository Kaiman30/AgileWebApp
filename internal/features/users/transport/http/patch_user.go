package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Kaiman30/AgileWebApp/internal/core/logger"
	core_http_request "github.com/Kaiman30/AgileWebApp/internal/core/transport/http/request"
	core_http_response "github.com/Kaiman30/AgileWebApp/internal/core/transport/http/response"
	core_http_types "github.com/Kaiman30/AgileWebApp/internal/core/transport/http/types"
	core_http_utils "github.com/Kaiman30/AgileWebApp/internal/core/transport/http/utils"
	"go.uber.org/zap"
)

type PatchUserRequest struct {
	FullName core_http_types.Nullable[string] `json:"full_name"`
	Email    core_http_types.Nullable[string] `json:"email"`
}

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	fmt.Println(userID)
	// TODO: Убрать

	var request PatchUserRequest
	log.Debug("request: ", zap.Any("req", r.Body))
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	log.Debug(
		fmt.Sprintf(
			"PatchUserRequest fields:\nFullName: '%v'\nEmail: '%v'\n",
			request.FullName,
			request.Email,
		),
	)

	rw.WriteHeader(http.StatusOK)
}
