package ape

import (
	"net/http"
	"strings"
)

func (app *App) Server() *http.Server {
	return &http.Server{
		Handler: app,
	}
}

func (app *App) ListenAndServe(addr string) {
	server := app.Server()
	server.Addr = addr
	server.ListenAndServe()
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := newRequestFromHTTPRequest(r)
	res := newResponse()

	data, err := app.Serve(req, res)

	encoder := app.Encoders[req.Format]
	res.Body, err = encoder.Encode(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", app.ContentTypes[req.Format])
		for header, values := range res.Header {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}
		w.WriteHeader(res.StatusCode)
		if res.StatusCode != http.StatusNoContent {
			w.Write(res.Body)
		} else {
			w.Write([]byte{})
		}
	}
}

func (app *App) Serve(req *Request, res *Response) (Any, error) {
	var (
		marshalable Any
		err         error
	)

	req.Format = app.DefaultFormat
	req.Path = strings.TrimPrefix(req.Path, app.Prefix)

	routes := app.router.MatchedRoutes(
		req.Verb,
		req.Path,
	)

	if len(routes) == 0 {
		res.StatusCode = http.StatusNotFound
		return nil, nil
	}

	for _, route := range routes {
		params := route.Params(req.Path)

		if fmt, ok := params["_format"]; ok {
			if _, ok := app.Encoders[fmt]; ok {
				req.Format = fmt
			}
		}

		marshalable, err = route.Handler.Serve(req, res)

		if err != nil {
			switch err.(type) {
			default:
				app.ErrorHandler(req, res, err)
			}
		}

		break
	}

	if res.StatusCode == 0 {
		switch req.Verb {
		default:
			res.StatusCode = 200
		case VerbHead:
			res.StatusCode = 204
		case VerbPost:
			res.StatusCode = 201
		}
	}

	return marshalable, nil
}
