package member

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/infracloudio/msbotbuilder-go/connector/client"
	"github.com/infracloudio/msbotbuilder-go/schema"
	"github.com/pkg/errors"
)

const (
	// APIVersion for response URLs
	APIVersion = "v3"

	sendToConversationURL = "/%s/conversations/%s/members"
	activityResourceURL   = "/%s/conversations/%s/members/%s"
)

type Members struct {
	Client client.Client
}

func (m *Members) GetMemberDetails(ctx context.Context, activity schema.Activity) (*schema.MemberDetails, error) {
	u, err := url.Parse(activity.ServiceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse ServiceURL %s.", activity.ServiceURL)
	}

	respPath := fmt.Sprintf(activityResourceURL, APIVersion, activity.Conversation.ID, activity.From.ID)

	// Send activity to client
	u.Path = path.Join(u.Path, respPath)
	data, err := m.Client.Get(*u, activity)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get member details")
	}

	memberDetails, err := requestToMember(data)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to cast member details")
	}

	return memberDetails, nil
}
