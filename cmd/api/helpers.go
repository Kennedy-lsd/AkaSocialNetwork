package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// writeJSON sends a JSON response with the specified HTTP status code and data.
// It sets the `Content-Type` header to `application/json`.
//
// Parameters:
//   - w: The `http.ResponseWriter` to write the response to.
//   - status: The HTTP status code to send.
//   - data: The data to encode and send as a JSON response.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// idParser extracts and parses an integer ID from a URL path at the specified index.
//
// Parameters:
//   - r: The `*http.Request` containing the URL path.
//   - urlIndex: The index of the path segment to parse as an ID.
//
// Returns:
//   - The parsed `int64` ID if successful.
//   - An error if the URL format is invalid or the ID cannot be parsed.
//
// Example:
//
//	For a URL path `/api/user/123`, calling `idParser(r, 3)` will return the ID `123`.
func idParser(r *http.Request, urlIndex int) (int64, error) {
	pathFragments := strings.Split(r.URL.Path, "/")

	if len(pathFragments) < urlIndex {
		err := errors.New("Invalid URL format")
		return 0, err
	}

	idStr := pathFragments[urlIndex]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = errors.New("Invalid id format")
		return 0, err
	}

	return id, nil
}
