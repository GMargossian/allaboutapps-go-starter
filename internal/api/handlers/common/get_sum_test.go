package common_test

import (
	"net/http"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/test"
	"github.com/stretchr/testify/require"
)

func TestGetSumCorrectness(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?count=5&mgmt-secret="+s.Config.Management.Secret, nil, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)
		require.Equal(t, res.Body.String(), "15\n")
	})
}

func TestGetSumNoCountArgument(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?mgmt-secret="+s.Config.Management.Secret, nil, nil)
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		require.Equal(t, res.Body.String(), "Please provide an integer.\n")
	})
}

func TestGetSumNegativeCountArgument(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?count=-1&mgmt-secret="+s.Config.Management.Secret, nil, nil)
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		require.Equal(t, res.Body.String(), "Please provide an integer.\n")
	})
}

func TestGetSumNoIntegerCountArgument(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?count=testText&mgmt-secret="+s.Config.Management.Secret, nil, nil)
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		require.Equal(t, res.Body.String(), "Please provide an integer.\n")
	})
}

func TestGetSumNoSecretToken(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?count=5", nil, nil)
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})
}

func TestGetSumIncorrectSecretToken(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/-/sum?count=5&mgmt-secret=wrong-password-for-sure", nil, nil)
		require.Equal(t, http.StatusUnauthorized, res.Result().StatusCode)
	})
}
