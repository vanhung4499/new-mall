package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"new-mall/pkg/e"
	"new-mall/pkg/utils/ctl"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.Response {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fieldError := range ve {
			field := fmt.Sprintf("Field.%s", fieldError.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", fieldError.Tag)
			return ctl.ResError(ctx, err, fmt.Sprintf("%s%s", field, tag))
		}
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return ctl.ResError(ctx, err, "JSON type mismatch")
	}

	return ctl.ResError(ctx, err, err.Error(), e.ERROR)
}
