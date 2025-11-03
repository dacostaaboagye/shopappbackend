package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/messages"
)

func SendResponse(w http.ResponseWriter, data *models.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)

	return json.NewEncoder(w).Encode(data)
}

func GenSuccessResponse(entity string, statusCode int, data interface{}) *models.Response {

	var messageCode = statusCode
	if statusCode == http.StatusNoContent || statusCode == codes.LOGIN_SUCCESS {
		statusCode = http.StatusOK
	}

	return &models.Response{
		Success: true,
		Message: messages.Success(entity, messageCode),
		Data:    data,
		Code:    statusCode,
	}
}

func GenErrorResponse(entity string, statusCode int, err error) *models.Response {
	appErr := errors.New(entity, statusCode, err)
	return &models.Response{
		Success: false,
		Message: appErr.UserMessage,
		Code:    statusCode,
	}
}

func GenAuthResponse(authCode, statusCode int) *models.Response {
	return &models.Response{
		Success: false,
		Message: messages.Error(entities.AUTHORIZATION, authCode),
		Code:    statusCode,
	}
}
