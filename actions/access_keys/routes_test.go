package accesskeys

import (
	"github.com/bsv-blockchain/spv-wallet/config"
)

func (ts *TestSuite) TestRegisterRoutes() {
	ts.Run("test routes", func() {
		testCases := []struct {
			method string
			url    string
		}{
			{"GET", "/api/" + config.APIVersion + "/users/current/keys/:id"},
			{"POST", "/api/" + config.APIVersion + "/users/current/keys"},
			{"DELETE", "/api/" + config.APIVersion + "/users/current/keys/:id"},
			{"GET", "/api/" + config.APIVersion + "/users/current/keys"},
		}

		ts.Router.Routes()

		for _, testCase := range testCases {
			found := false
			for _, routeInfo := range ts.Router.Routes() {
				if testCase.url == routeInfo.Path && testCase.method == routeInfo.Method {
					ts.NotNil(routeInfo.HandlerFunc)
					found = true
					break
				}
			}
			ts.True(found)
		}
	})
}
