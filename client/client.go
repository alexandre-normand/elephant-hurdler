package client

import (
	nameproto "ClientNamenodeProtocol"
)

type DistributedFilesystemClient interface {
	List(source string) ([]string, error)
	Get(path string) (Reader, error)
	Put(path string, writer Writer) error
	PutWithReplication(path string, writer Writer, replication int) error
}

type HdfsService interface {
	GetListing(in *nameproto.GetListingRequestProto, out *nameproto.GetListingResponseProto) error
}

type HdfsClient struct {
	service HdfsService
}

func (client *HdfsClient) List(source string) (content []string, err error) {
	request = nameproto.GetListingRequestProto{}
	request.Src = Path
	request.StartAfter = ""
	request.NeedLocation = false

	response = nameproto.GetListingResponseProto{}
	listing = client.service.getListing(&request, &response)
}
