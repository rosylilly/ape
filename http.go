package ape

import (
	"net/http"
	"strings"
	"time"
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
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, err.(error).Error(), http.StatusInternalServerError)
		}
	}()

	req := newRequestFromHTTPRequest(r)
	res := newResponse(w)

	data, err := app.Serve(req, res)

	if res.Body == nil {
		encoder := app.Encoders[req.Format]
		res.Body, err = encoder.Encode(data)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	} else {
		res.Header().Set("Content-Type", app.ContentTypes[req.Format])
		res.WriteHeader(res.StatusCode)
		if res.StatusCode != http.StatusNoContent {
			res.Write(res.Body)
		} else {
			res.Write([]byte{})
		}
	}
}

func (app *App) Serve(req *Request, res *Response) (Any, error) {
	var (
		marshalable Any
		err         error
	)

	logger := app.Logger()
	now := time.Now().UTC()
	defer func() {
		if logger != nil {
			logger.Printf(
				"Completed %d in %dms",
				res.StatusCode,
				time.Now().UTC().Sub(now).Nanoseconds()*10,
			)
		}
	}()

	req.Format = app.DefaultFormat
	req.Path = strings.TrimPrefix(req.Path, app.Prefix)

	if len(req.Path) == 0 {
		req.Path = "/"
	}

	if logger != nil {
		logger.Printf("Serve %s %s", req.Verb, req.Path)
	}

	routes := app.router.MatchedRoutes(
		req.Verb,
		req.Path,
	)

	if len(routes) == 0 {
		res.StatusCode = http.StatusNotFound
		return nil, nil
	}

	for _, filter := range app.beforeHandlers {
		func() {
			defer func() {
				if e := recover(); e != nil {
					err = e.(error)
				}
			}()

			_, err = filter.Serve(req, res)
		}()

		if err != nil {
			switch err.(type) {
			case *RequestHaltedError:
				return nil, nil
			case *RequestPassedError:
				continue
			default:
				app.ErrorHandler(req, res, err)
				return nil, err
			}
		}
	}

	for _, route := range routes {
		params := route.Params(req.Path)

		if fmt, ok := params["_format"]; ok {
			if _, ok := app.Encoders[fmt]; ok {
				req.Format = fmt
			}
		}

		req.RouteParams = params

		func() {
			defer func() {
				if e := recover(); e != nil {
					err = e.(error)
				}
			}()

			marshalable, err = route.Handler.Serve(req, res)
		}()

		if err != nil {
			switch err.(type) {
			case *RequestHaltedError:
				break
			case *RequestPassedError:
				continue
			default:
				app.ErrorHandler(req, res, err)
			}
		}

		break
	}

	for _, filter := range app.afterHandlers {
		func() {
			defer func() {
				if e := recover(); e != nil {
					err = e.(error)
				}
			}()

			_, err = filter.Serve(req, res)
		}()

		if err != nil {
			switch err.(type) {
			case *RequestHaltedError:
				return marshalable, nil
			case *RequestPassedError:
				continue
			default:
				app.ErrorHandler(req, res, err)
				return marshalable, err
			}
		}
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
