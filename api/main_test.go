package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nuggetplum/simplebank/db/sqlc"
	"github.com/nuggetplum/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
