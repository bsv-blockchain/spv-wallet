package chainmodels

// BHSConfig consists of AuthToken and URL used to communicate with BlockHeadersService (BHS)
type BHSConfig struct {
	AuthToken string //nolint:gosec // G117 auth token field in config struct
	URL       string
}
