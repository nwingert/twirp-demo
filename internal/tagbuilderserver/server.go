package tagbuilderserver

import (
	"context"

	pb "github.com/nwingert/twirp-demo/rpc/tagbuilder"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) GetTags(ctx context.Context, merchant *pb.Merchant) (container *pb.TagContainer, err error) {
	if merchant.MerchantId <= 0 {
		return nil, twirp.InvalidArgumentError("merchantId", "Merchant must have valid id > 0")
	}

	// junk strings for tags
	var tag1, tag2, tag3, tag4 pb.Tag
	tag1.Tag = "tag1"
	tag2.Tag = "tag2"
	tag3.Tag = "tag3"
	tag4.Tag = "tag4"

	return &pb.TagContainer{
		MerchantId: merchant.MerchantId,
		Tags:       []*pb.Tag{&tag1, &tag2, &tag3, &tag4},
	}, nil
}
