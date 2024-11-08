package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	em "github.com/IiMDMiI/employeeManager/api/employeeManagment"
)

type TokenAuthorizer struct {
	handler http.Handler
}

func NewAuth(handlerToWrap http.Handler) *TokenAuthorizer {
	return &TokenAuthorizer{handlerToWrap}
}

func (t *TokenAuthorizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Basic ") {
		t.setUnauthorizedResponse(w)
		return
	}
	if err := t.isTockenCorrect(&token); err != nil {
		t.setUnauthorizedResponse(w)
		return
	}
	t.handler.ServeHTTP(w, r)
}

func (*TokenAuthorizer) setUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(em.NewUnauthorizedProblem())
}

// TODO: get tocken from DB
func (ta *TokenAuthorizer) isTockenCorrect(tocken *string) error {
	actualTockenPart := strings.Split(*tocken, " ")[1]
	creds := "user:pass"
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(creds))

	if actualTockenPart == encodedCreds {
		return nil
	} else {
		return errors.New("Unauthorized")
	}
}
