package api

type CreationResponse struct {

  Objects []interface{}   `json:"objects,omitempty"`
  Message string          `json:"message,omitempty"`
  Errors  interface{}     `json:"errors,omitempty"`
}
