package helpers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/lodashventure/nlp/model"
)

func CheckPathFormat(path string, regexpFormat *string) bool {
	re := regexp.MustCompile(*regexpFormat)
	return re.MatchString(path)
}

func RequestBodyValidation(validate *validator.Validate, requestBody interface{}) []model.ErrorResponse {
	var ErrorResponse []model.ErrorResponse

	message := make(map[string]string)
	message["len"] = "Length has invalid."
	message["max"] = "Max has invalid length."
	message["min"] = "Min has invalid length."
	message["required"] = "This Field has required."
	message["alpha"] = "Value must be alpha format."
	message["alphanum"] = "Value must be alpha or number format."
	message["alphanumunicode"] = "Value must be alpha or number or unicode format."
	message["number"] = "Value must be number format."
	message["uuid4"] = "Value must be uuid v4 format."
	message["numeric"] = "Value must be numeric format."
	message["iso4217"] = "Value must be iso4217 format."
	message["iso3166_1_alpha2"] = "Value must be iso3166_1_alpha2 format."
	message["uppercase"] = "Value must be uppercase format."

	if err := validate.Struct(requestBody); err != nil {
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var element model.ErrorResponse
				element.FailedField = err.StructNamespace()
				element.Tag = err.Tag()
				element.Value = err.Param()
				element.ErrorMessage = message[err.Tag()]

				ErrorResponse = append(ErrorResponse, element)
			}
		}
	}

	if ErrorResponse != nil {
		return ErrorResponse
	}

	return nil
}
