package base

import (
	"github.com/bsv-blockchain/spv-wallet/config"
)

func (ts *TestSuite) TestBaseRegisterRoutes() {
	ts.Run("test routes", func() {
		testCases := []struct {
			method string
			url    string
		}{
			{"GET", "/"},
			{"OPTIONS", "/"},
			{"HEAD", "/"},
			{"GET", "/" + config.HealthRequestPath},
			{"OPTIONS", "/" + config.HealthRequestPath},
			{"HEAD", "/" + config.HealthRequestPath},
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
