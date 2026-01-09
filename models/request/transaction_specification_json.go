package request

import (
	"encoding/json"
	"errors"

	"github.com/bsv-blockchain/spv-wallet/models/request/opreturn"
	paymailreq "github.com/bsv-blockchain/spv-wallet/models/request/paymail"
)

// ErrUnsupportedOutputType is returned when an unsupported output type is provided
var ErrUnsupportedOutputType = errors.New("unsupported output type")

// unmarshalOutput used by TransactionSpecification unmarshalling to get Output object by type
// IMPORTANT: Every time a new output type is added, it must be handled here also.
func unmarshalOutput(rawOutput json.RawMessage, outputType string) (Output, error) {
	switch outputType {
	case "op_return":
		var out opreturn.Output
		if err := json.Unmarshal(rawOutput, &out); err != nil {
			return nil, err //nolint:wrapcheck // unmarshalOutput is run internally by json.Unmarshal, so we don't want to wrap the error
		}
		return out, nil
	case "paymail":
		var out paymailreq.Output
		if err := json.Unmarshal(rawOutput, &out); err != nil {
			return nil, err //nolint:wrapcheck // unmarshalOutput is run internally by json.Unmarshal, so we don't want to wrap the error
		}
		return out, nil
	default:
		return nil, ErrUnsupportedOutputType
	}
}

// expandOutputForMarshaling used by TransactionSpecification marshaling to expand Output object before marshaling.
// IMPORTANT: Every time a new output type is added, it must be handled here also.
func expandOutputForMarshaling(output Output) (any, error) {
	switch o := output.(type) {
	// unfortunately we must do the same for each and every type,
	// because go json is not handling unwrapping embedded type when using just Output interface.
	case opreturn.Output:
		return struct {
			*opreturn.Output

			Type string `json:"type"`
		}{
			Type:   o.GetType(),
			Output: &o,
		}, nil
	case paymailreq.Output:
		return struct {
			*paymailreq.Output

			Type string `json:"type"`
		}{
			Type:   o.GetType(),
			Output: &o,
		}, nil
	default:
		return nil, ErrUnsupportedOutputType
	}
}
