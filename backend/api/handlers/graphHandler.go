package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/helpers"
	"github.com/lodashventure/nlp/middlewares"
	"github.com/lodashventure/nlp/model"
)

type GraphClass struct {
	router *fiber.Router
}

func NewGraphHandler(router *fiber.Router) *GraphClass {
	return &GraphClass{
		router: router,
	}
}

func (property *GraphClass) NewGraphRoute() {
	(*property.router).Get("/img", middlewares.CheckToken(), middlewares.CheckQuota(), property.Graph)
}

func (property *GraphClass) Graph(c *fiber.Ctx) error {
	var ErrorResponse []model.ErrorResponse

	var requestGraph *model.RequestGraph

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	width := c.Query("width")
	height := c.Query("height")

	// err := c.BodyParser(&requestGraph)
	// if err != nil {
	// 	ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Request body invalid format or try again later."})
	// 	return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, nil, ErrorResponse))
	// }

	// errValidate := property.CheckGraphValidation(requestGraph)
	// if errValidate != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(infrastructure.Response(false, nil, errValidate))
	// }

	// var result *model.GraphResponse

	// tokenization, err := property.tokenizationService.GetByText(requestGraph.Text)
	// if err != nil {
	// 	cmd := exec.Command("python", "./handlers/english/tokenizationPython.py", requestGraph.Text)
	// 	out, err := cmd.Output()
	// 	if err != nil {
	// 		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Sorry, we were unable to process your request at the moment. Please try again later or contact customer support for assistance."})
	// 		return c.Status(http.StatusInternalServerError).JSON(infrastructure.Response(false, nil, ErrorResponse))
	// 	}

	// 	var resultPython model.PythonGraph
	// 	err = json.Unmarshal(out, &resultPython)
	// 	if err != nil {
	// 		ErrorResponse = append(ErrorResponse, model.ErrorResponse{ErrorMessage: "Sorry, we were unable to process your request at the moment. Please try again later or contact customer support for assistance."})
	// 		return c.Status(http.StatusInternalServerError).JSON(infrastructure.Response(false, nil, ErrorResponse))
	// 	}

	// 	result = &model.GraphResponse{
	// 		Results:  resultPython.Results,
	// 		Keywords: *helpers.KeyWordStr(&resultPython.Results),
	// 	}

	// 	tokenization := &model.Graph{
	// 		ID:        primitive.NewObjectID(),
	// 		Text:      requestGraph.Text,
	// 		Results:   result.Results,
	// 		Keywords:  result.Keywords,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	}
	// 	go func(tokenization *model.Graph) {
	// 		property.tokenizationService.Insert(tokenization)
	// 	}(tokenization)
	// } else {
	// 	result = &model.GraphResponse{
	// 		Results:  tokenization.Results,
	// 		Keywords: tokenization.Keywords,
	// 	}
	// }

	// return c.Status(http.StatusOK).JSON(infrastructure.Response(true, result, ErrorResponse))
	return c.Status(http.StatusOK).JSON("")
}

func (property *GraphClass) CheckGraphValidation(requestBody *model.RequestGraph) []model.ErrorResponse {
	var ErrorResponse []model.ErrorResponse
	validate := validator.New()

	err := helpers.RequestBodyValidation(validate, requestBody)
	if err != nil {
		ErrorResponse = append(ErrorResponse, err...)
	}

	if ErrorResponse != nil {
		return ErrorResponse
	}

	return nil
}
