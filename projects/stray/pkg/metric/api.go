package metric

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/grafana/regexp"
	"github.com/munnerz/goautoneg"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/annotations"
	"github.com/prometheus/prometheus/util/httputil"

	jsoniter "github.com/json-iterator/go"
)

type (
	EngineApi struct {
		promQuerier *PromQuerier
		promStorage *PromStorage

		mux         *http.ServeMux
		httpHandler http.Handler
		cors        *regexp.Regexp
		logger      *slog.Logger
		codecs      []Codec

		ready func(http.HandlerFunc) http.HandlerFunc
	}

	EngineApiOptions struct {
		ServeMux    *http.ServeMux
		HttpHandler http.Handler
		Querier     *PromQuerier
		Storage     *PromStorage
	}

	errorType string
	status    string

	apiError struct {
		typ errorType
		err error
	}

	apiFuncResult struct {
		data      interface{}
		err       *apiError
		warnings  annotations.Annotations
		finalizer func()
	}

	apiFunc func(r *http.Request) apiFuncResult

	Response struct {
		Status    status      `json:"status"`
		Data      interface{} `json:"data,omitempty"`
		ErrorType errorType   `json:"errorType,omitempty"`
		Error     string      `json:"error,omitempty"`
		Warnings  []string    `json:"warnings,omitempty"`
		Infos     []string    `json:"infos,omitempty"`
	}
)

const (
	errorNone          errorType = ""
	errorTimeout       errorType = "timeout"
	errorCanceled      errorType = "canceled"
	errorExec          errorType = "execution"
	errorBadData       errorType = "bad_data"
	errorInternal      errorType = "internal"
	errorUnavailable   errorType = "unavailable"
	errorNotFound      errorType = "not_found"
	errorNotAcceptable errorType = "not_acceptable"

	statusSuccess status = "success"
	statusError   status = "error"

	// Non-standard status code (originally introduced by nginx) for the case when a client closes
	// the connection while the server is still processing the request.
	statusClientClosedConnection = 499
)

func setUnavailStatusOnTSDBNotReady(r apiFuncResult) apiFuncResult {
	if r.err != nil && errors.Is(r.err.err, tsdb.ErrNotReady) {
		r.err.typ = errorUnavailable
	}
	return r
}

func (api *EngineApi) respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&Response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})
	if err != nil {
		api.logger.Error("error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var code int
	switch apiErr.typ {
	case errorBadData:
		code = http.StatusBadRequest
	case errorExec:
		code = http.StatusUnprocessableEntity
	case errorCanceled:
		code = statusClientClosedConnection
	case errorTimeout:
		code = http.StatusServiceUnavailable
	case errorInternal:
		code = http.StatusInternalServerError
	case errorNotFound:
		code = http.StatusNotFound
	case errorNotAcceptable:
		code = http.StatusNotAcceptable
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if n, err := w.Write(b); err != nil {
		api.logger.Error("error writing response", "bytesWritten", n, "err", err)
	}
}

func (api *EngineApi) negotiateCodec(req *http.Request, resp *Response) (Codec, error) {
	for _, clause := range goautoneg.ParseAccept(req.Header.Get("Accept")) {
		for _, codec := range api.codecs {
			if codec.ContentType().Satisfies(clause) && codec.CanEncode(resp) {
				return codec, nil
			}
		}
	}

	defaultCodec := api.codecs[0]
	if !defaultCodec.CanEncode(resp) {
		return nil, fmt.Errorf("cannot encode response as %s", defaultCodec.ContentType())
	}

	return defaultCodec, nil
}

func (api *EngineApi) respond(w http.ResponseWriter, req *http.Request, data interface{}, warnings annotations.Annotations, query string) {
	statusMessage := statusSuccess
	warn, info := warnings.AsStrings(query, 10, 10)

	resp := &Response{
		Status:   statusMessage,
		Data:     data,
		Warnings: warn,
		Infos:    info,
	}

	codec, err := api.negotiateCodec(req, resp)
	if err != nil {
		api.respondError(w, &apiError{errorNotAcceptable, err}, nil)
		return
	}

	b, err := codec.Encode(resp)
	if err != nil {
		api.logger.Error("error marshaling response", "url", req.URL, "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", codec.ContentType().String())
	w.WriteHeader(http.StatusOK)
	if n, err := w.Write(b); err != nil {
		api.logger.Error("error writing response", "url", req.URL, "bytesWritten", n, "err", err)
	}
}

func (api *EngineApi) Register() {
	wrap := func(f apiFunc) http.HandlerFunc {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httputil.SetCORS(w, api.cors, r)
			result := setUnavailStatusOnTSDBNotReady(f(r))
			if result.finalizer != nil {
				defer result.finalizer()
			}
			if result.err != nil {
				api.respondError(w, result.err, result.data)
				return
			}

			if result.data != nil {
				api.respond(w, r, result.data, result.warnings, r.FormValue("query"))
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		return httputil.CompressionHandler{Handler: hf}.ServeHTTP
	}

	apiRoute := route.New()
	apiRoute.Get("/query", wrap(api.query))
	api.mux.Handle("/", apiRoute)
}

func (api *EngineApi) query(r *http.Request) (result apiFuncResult) {
	// TODO
	return apiFuncResult{nil, nil, nil, nil}
}

func NewEngineApi(opts EngineApiOptions) *EngineApi {
	cors, _ := compileCORSRegexString(".*")
	engApi := &EngineApi{
		cors:        cors,
		httpHandler: opts.HttpHandler,
		mux:         opts.ServeMux,
		promQuerier: opts.Querier,
		promStorage: opts.Storage,
	}

	engApi.codecs = append(engApi.codecs, JSONCodec{})
	engApi.Register()
	// registerHttpHandler(engApi, opts.HttpHandler)
	return engApi
}
