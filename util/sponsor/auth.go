package sponsor

import (
	"context"
	"fmt"
	"time"

	"github.com/thommyho/robotui/api/proto/pb"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/cloud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Subject, Token string
	ExpiresAt      time.Time
)

func IsAuthorized() bool {
	return len(Subject) > 0
}

// check and set sponsorship token
func ConfigureSponsorship(token string) error {
	host := util.Getenv("GRPC_URI", cloud.Host)
	conn, err := cloud.Connection(host)
	if err != nil {
		return err
	}

	client := pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.IsAuthorized(ctx, &pb.AuthRequest{Token: token})
	if err == nil && res.Authorized {
		Subject = res.Subject
		ExpiresAt = res.ExpiresAt.AsTime()
		Token = token
	}

	if err != nil {
		if s, ok := status.FromError(err); ok && s.Code() != codes.Unknown {
			Subject = "sponsorship unavailable"
			err = nil
		} else {
			err = fmt.Errorf("sponsortoken: %w", err)
		}
	}

	return err
}
