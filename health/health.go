// Package health serves a small JSON readiness response for a voxigo bot.
package health

import (
	"encoding/json"
	"net/http"
)

// Response is what Handler writes on success.
type Response struct {
	Status  string `json:"status"`
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
}

// Handler returns an endpoint that answers 200 with {"status":"ok"} plus the
// optional service name and version. It is for a process probe, not a check of
// every downstream API key.
func Handler(service, version string) http.Handler {
	body, err := json.Marshal(Response{Status: "ok", Service: service, Version: version})
	if err != nil {
		// A struct of strings cannot fail to marshal; keep the handler honest.
		body = []byte(`{"status":"ok"}`)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	})
}
