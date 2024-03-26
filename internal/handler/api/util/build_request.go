package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"strings"

	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
)

type BuildRequestResponse struct {
	Signature string
}

const (
	HEADER_CLIENT_ID     = "x-client-id"
	HEADER_CLIENT_SECRET = "x-client-secret"
	HEADER_TIMESTAMP     = "x-timestamp"
)

func (s *UtilHandler) BuildRequest(w http.ResponseWriter, r *http.Request) {
	clientID := r.Header.Get(HEADER_CLIENT_ID)
	if clientID == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	clientSecret := r.Header.Get(HEADER_CLIENT_SECRET)
	if clientSecret == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	timestamp := r.Header.Get(HEADER_TIMESTAMP)
	if timestamp == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var requestBody bytes.Buffer
	_, err := requestBody.ReadFrom(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	bodyString := requestBody.String()
	minifiedReq := minifyRequestBody(bodyString)

	res.Yay(w, r, http.StatusOK, &BuildRequestResponse{
		Signature: generateSignature(clientID, clientSecret, minifiedReq, timestamp),
	})
}

func generateSignature(clientID, clientSecret, body, createdTime string) string {
	data := clientID + body + createdTime
	hmac := hmac.New(sha512.New, []byte(clientSecret))

	// compute the HMAC
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)
	signature := hex.EncodeToString(dataHmac)

	return signature
}

func minifyRequestBody(requestBody string) string {
	// Remove leading and trailing white spaces
	minifiedBody := strings.TrimSpace(requestBody)

	// Remove new lines and tabs
	minifiedBody = strings.ReplaceAll(minifiedBody, "\n", "")
	minifiedBody = strings.ReplaceAll(minifiedBody, "\t", "")

	// Remove unnecessary white spaces between characters
	minifiedBody = strings.Join(strings.Fields(minifiedBody), "")

	return minifiedBody
}
