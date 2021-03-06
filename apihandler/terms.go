package apihandler

import (
	"time"

	"github.com/alexwbaule/give-help/v2/generated/models"
	"github.com/alexwbaule/give-help/v2/generated/restapi/operations/terms"
	handler "github.com/alexwbaule/give-help/v2/handlers/terms"
	runtimeApp "github.com/alexwbaule/give-help/v2/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rafaelfino/metrics"
)

//PutUserAcceptHandler
func TermsPutUserAcceptHandler(rt *runtimeApp.Runtime) terms.PutUserAcceptHandler {
	return &putUserAcceptHandler{rt: rt}
}

type putUserAcceptHandler struct {
	rt *runtimeApp.Runtime
}

func (ctx *putUserAcceptHandler) Handle(params terms.PutUserAcceptParams, principal *models.LoggedUser) middleware.Responder {
	start := time.Now()

	c := handler.New(ctx.rt.GetDatabase())

	err := c.Accept(params.TermID, *principal.UserID)

	defer ctx.rt.GetMetricProcessor().Send(metrics.NewMetric("TermsPutUserAcceptHandler.ElapsedTime", metrics.CounterType, nil, float64(time.Since(start).Milliseconds())))

	if err != nil {
		return terms.NewPutUserAcceptInternalServerError().WithPayload(&models.APIError{Message: "An unexpected error occurred"})
	}

	return terms.NewPutUserAcceptOK()
}

//GetTermsHandler
func TermsGetTermsHandler(rt *runtimeApp.Runtime) terms.GetTermsHandler {
	return &getTermsHandler{rt: rt}
}

type getTermsHandler struct {
	rt *runtimeApp.Runtime
}

func (ctx *getTermsHandler) Handle(params terms.GetTermsParams) middleware.Responder {
	start := time.Now()

	c := handler.New(ctx.rt.GetDatabase())
	ret, err := c.LoadTerms()

	defer ctx.rt.GetMetricProcessor().Send(metrics.NewMetric("TermsGetTermsHandler.ElapsedTime", metrics.CounterType, nil, float64(time.Since(start).Milliseconds())))

	if err != nil {
		return terms.NewGetTermsInternalServerError().WithPayload(&models.APIError{Message: "An unexpected error occurred"})
	}

	return terms.NewGetTermsOK().WithPayload(ret)
}

//GetUserAcceptedHandler
func TermsGetUserAcceptedHandler(rt *runtimeApp.Runtime) terms.GetUserAcceptedHandler {
	return &getUserAcceptedHandler{rt: rt}
}

type getUserAcceptedHandler struct {
	rt *runtimeApp.Runtime
}

func (ctx *getUserAcceptedHandler) Handle(params terms.GetUserAcceptedParams, principal *models.LoggedUser) middleware.Responder {
	start := time.Now()

	c := handler.New(ctx.rt.GetDatabase())
	ret, err := c.LoadUserAcceptedTerms(*principal.UserID)

	defer ctx.rt.GetMetricProcessor().Send(metrics.NewMetric("TermsGetUserAcceptedHandler.ElapsedTime", metrics.CounterType, nil, float64(time.Since(start).Milliseconds())))

	if err != nil {
		return terms.NewGetUserAcceptedInternalServerError().WithPayload(&models.APIError{Message: "An unexpected error occurred"})
	}

	return terms.NewGetUserAcceptedOK().WithPayload(ret)
}
