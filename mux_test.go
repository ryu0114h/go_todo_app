package main

import (
	"testing"
)

func TestNewMux(t *testing.T) {
	// w := httptest.NewRecorder()
	// r := httptest.NewRequest(http.MethodGet, "/health", nil)
	// sut := NewMux()
	// sut.ServeHTTP(w, r)
	// resp := w.Result()
	// t.Cleanup(func() { _ = resp.Body.Close() })
	// if resp.StatusCode != http.StatusOK {
	// 	t.Error("want status code, but", resp.StatusCode)
	// }
	// got, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Fatalf("failed to read body: %v", err)
	// }
	// want := `{"status": "ok"}`
	// if string(got) != want {
	// 	t.Errorf("want %q, but got %q", want, got)
	// }
}
