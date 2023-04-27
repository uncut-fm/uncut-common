package model

import proto_user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"

type OffsetPaginationInput struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

func (o OffsetPaginationInput) IsLimitSet() bool {
	return o.Limit != nil
}

func (o OffsetPaginationInput) IsOffsetSet() bool {
	return o.Offset != nil
}

func PrepareOffsetPagination(paginationReq *OffsetPaginationInput) *OffsetPaginationInput {
	if paginationReq == nil {
		paginationReq = new(OffsetPaginationInput)
	}

	if !paginationReq.IsLimitSet() {
		paginationReq.Limit = ValPointer(0)
	}

	if !paginationReq.IsOffsetSet() {
		paginationReq.Offset = ValPointer(0)
	}

	return paginationReq
}

func ParseOffsetPaginationToProto(offsetPagination *OffsetPaginationInput) *proto_user.OffsetPaginationRequest {
	if offsetPagination == nil {
		return nil
	}

	PrepareOffsetPagination(offsetPagination)
	return &proto_user.OffsetPaginationRequest{
		Limit:  uint64(*offsetPagination.Limit),
		Offset: uint64(*offsetPagination.Offset),
	}
}

func ParseProtoOffsetPagination(protoOffsetPagination *proto_user.OffsetPaginationRequest) *OffsetPaginationInput {
	if protoOffsetPagination == nil {
		return nil
	}

	return &OffsetPaginationInput{
		Limit:  ValPointer(int(protoOffsetPagination.Limit)),
		Offset: ValPointer(int(protoOffsetPagination.Offset)),
	}
}
