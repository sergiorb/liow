package api

type ReadResponse struct {

  Objects []interface{}   `json:"objects,omitempty"`
  Message string          `json:"message,omitempty"`
}
