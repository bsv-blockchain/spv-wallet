package transactions

import (
	"github.com/bsv-blockchain/spv-wallet/config"
)

func (ts *TestSuite) TestTransactionRegisterRoutes() {
	ts.Run("test routes", func() {
		testCases := []struct {
			method string
			url    string
		}{
			{"GET", "/api/" + config.APIVersion + "/transactions/:id"},
			{"PATCH", "/api/" + config.APIVersion + "/transactions/:id"},
			{"GET", "/api/" + config.APIVersion + "/transactions"},
			{"POST", "/api/" + config.APIVersion + "/transactions/drafts"},
			{"POST", "/api/" + config.APIVersion + "/transactions"},
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
