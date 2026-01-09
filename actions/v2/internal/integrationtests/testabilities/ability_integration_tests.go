package testabilities

import (
	"testing"

	"github.com/bsv-blockchain/spv-wallet/actions/testabilities"
)

func New(t testing.TB) (given IntegrationTestFixtures, when IntegrationTestAction, then IntegrationTestAssertion) {
	appFixture, appAssertions := testabilities.New(t)

	integrationFixture := newFixture(t, appFixture)
	when = newActions(t, integrationFixture)
	then = newAssertions(t, integrationFixture, appAssertions)

	return integrationFixture, when, then
}
