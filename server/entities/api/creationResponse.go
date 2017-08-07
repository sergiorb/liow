package api

type CreationResponse struct {

  Objects []interface{}           `json:"objects,omitempty"`
  Message string                  `json:"message,omitempty"`
  Errors  map[string]interface{}  `json:"errors,omitempty"`
}
