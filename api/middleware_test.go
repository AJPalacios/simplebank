package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/devspace/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	authorizationHeader := authorizationType + " " + token
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{{
		name: "OK",
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "user", time.Minute)
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorder.Code)
		},
	}, {
		name: "NoAuthorization",
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// Do nothing
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusUnauthorized, recorder.Code)
		},
	}, {
		name: "InvalidAuthorizationFormat",
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			authorizationHeader := "Bearer"
			request.Header.Set(authorizationHeaderKey, authorizationHeader)
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusUnauthorized, recorder.Code)
		},
	}, {
		name: "UnsupportedAuthorizationType",
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			token, err := tokenMaker.CreateToken("user1", time.Minute)
			require.NoError(t, err)
			require.NotEmpty(t, token)

			authorizationHeader := "Basic " + token
			request.Header.Set(authorizationHeaderKey, authorizationHeader)
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusUnauthorized, recorder.Code)
		},
	}, {
		name: "InvalidToken",
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			authorizationHeader := "Bearer invalid_token"
			request.Header.Set(authorizationHeaderKey, authorizationHeader)
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusUnauthorized, recorder.Code)
		},
	}}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)

			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(200, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
