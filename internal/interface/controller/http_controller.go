package controller

import "github.com/Goalt/FileSharer/internal/usecase/interactor"

type HTTPController interface {
	AddHandler(hc HTTPContext) HTTPError
}

type httpController struct {
	calculatorInteractor interactor.CalculatorInteractor
}

func NewHTTPController(calculator interactor.CalculatorInteractor) HTTPController {
	return &httpController{calculatorInteractor: calculator}
}

func (hc *httpController) AddHandler(httpCtx HTTPContext) HTTPError {
	return nil
}
