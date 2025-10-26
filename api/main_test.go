package api

import (
	"os"
	"testing"

	db "github.com/devspace/simplebank/db/sqlc"
	"github.com/devspace/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymKey:         util.RandomString(32),
		AccessTokenDuration: "1m",
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
