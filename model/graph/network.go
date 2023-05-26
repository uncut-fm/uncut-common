package model

import (
	"github.com/uncut-fm/uncut-common/pkg/proto/graph"
)

type NetworkUser struct {
	User         *User
	NFTsInCommon []*NFT
}

type UsersNetwork struct {
	Users []*NetworkUser
}

// convert UsersNetwork to graph.NetworkMembersInfoResponse
func (n UsersNetwork) ToProto() *graph.NetworkMembersInfoResponse {
	protoResponse := &graph.NetworkMembersInfoResponse{
		Users: make([]*graph.NetworkMemberUser, len(n.Users)),
	}

	for i, networkUser := range n.Users {
		protoResponse.Users[i] = &graph.NetworkMemberUser{}
		protoResponse.Users[i].User = networkUser.User.ToProto()
		protoResponse.Users[i].NftsInCommon = make([]*graph.Nft, len(networkUser.NFTsInCommon))

		for j, nft := range networkUser.NFTsInCommon {
			protoResponse.Users[i].NftsInCommon[j] = nft.ToProto()
		}
	}

	return protoResponse
}
