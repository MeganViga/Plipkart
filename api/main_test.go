package api

import (
	"testing"
	"time"

	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T,store db.Store)*Server{
	config := util.Config{
		HTTPServerAddress: "127.0.0.0.:4789",
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config,store)
	require.NoError(t,err)
	return server
}