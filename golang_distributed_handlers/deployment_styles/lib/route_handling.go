package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Handle[Req any, Resp any](handlerFn func(context.Context, Req) (Resp, error)) httprouter.Handle {

	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// TODO Allow non json requests (with params in the URL)
		req, err := decodeReq[Req](r)
		if err != nil {

			// TODO gracefully handle bad requests...
			err := encodeResp(w, r, http.StatusBadRequest, struct{
				Error string
			}{
				Error: "cannot decode req:" + err.Error(),
			})
			if err != nil {
				log.Fatal("error writing error resp for bad request")
			}

			return
		}

		resp, err := handlerFn(r.Context(), req)

		if err != nil {
			// TODO: handle errors better
			err := encodeResp(w, r, http.StatusInternalServerError, struct{
				Error string
			}{
				Error: err.Error(),
			})

			if err != nil {
				log.Fatal("error writing error resp")
			}

			return
		}

		err = encodeResp(w, r, http.StatusOK, resp)
		if err != nil {
			log.Fatal("error writing error resp for encoding resp")
		}

		return
	}
}

func encodeResp[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decodeReq[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

