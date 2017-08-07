package api

type DeleteResponse struct {

  Objects []interface{}           `json:"objects,omitempty"`
  Message string                  `json:"message,omitempty"`
}
