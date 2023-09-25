package member

import (
	"encoding/json"
	"io"

	"github.com/infracloudio/msbotbuilder-go/schema"
	"github.com/pkg/errors"
)

func requestToMember(data io.Reader) (*schema.MemberDetails, error) {

	memberdetails := &schema.MemberDetails{}

	err := json.NewDecoder(data).Decode(&memberdetails)

	return memberdetails, errors.Wrap(err, "Failed to parse data")
}
