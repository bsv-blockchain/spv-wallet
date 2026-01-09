package mappings

import (
	"github.com/iancoleman/strcase"

	"github.com/bsv-blockchain/spv-wallet/engine/datastore"
	"github.com/bsv-blockchain/spv-wallet/models/filter"
)

const (
	defaultPage     = 1
	defaultPageSize = 50
	defaultSortBy   = "created_at"
	defaultOrder    = "desc"
)

// MapToDbQueryParams converts filter.QueryParams from models to matching datastore.QueryParams
func MapToDbQueryParams(model *filter.Page) *datastore.QueryParams {
	if model == nil {
		return DefaultDBQueryParams()
	}
	return &datastore.QueryParams{
		Page:          getNumberOrDefault(model.Number, defaultPage),
		PageSize:      getNumberOrDefault(model.Size, defaultPageSize),
		OrderByField:  getStringOrDefaultToSnakeCase(model.SortBy, defaultSortBy),
		SortDirection: getStringOrDefalut(model.Sort, defaultOrder),
	}
}

// DefaultDBQueryParams returns the default db query params
func DefaultDBQueryParams() *datastore.QueryParams {
	return &datastore.QueryParams{
		Page:          defaultPage,
		PageSize:      defaultPageSize,
		OrderByField:  defaultSortBy,
		SortDirection: defaultOrder,
	}
}

func getNumberOrDefault(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func getStringOrDefaultToSnakeCase(value, defaultValue string) string {
	return strcase.ToSnake(
		getStringOrDefalut(value, defaultValue),
	)
}

func getStringOrDefalut(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
