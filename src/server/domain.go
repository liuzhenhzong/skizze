package server

import (
	pb "datamodel"

	"golang.org/x/net/context"
)

func (s *serverStruct) CreateDomain(ctx context.Context, in *pb.Domain) (*pb.Domain, error) {
	info := pb.NewEmptyInfo()
	info.Name = in.GetName()
	// FIXME: A Domain's info should have an array of properties for each Sketch (or just an array
	// of Sketches, like what the proto has). This is just a hack to choose the first Sketch and
	// use it's info for now
	info.Properties.MaxUniqueItems = uint(in.GetSketches()[0].GetProperties().GetMaxUniqueItems())
	info.Properties.Size = uint(in.GetSketches()[0].GetProperties().GetSize())
	err := s.manager.CreateDomain(info)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *serverStruct) ListDomains(ctx context.Context, in *pb.Empty) (*pb.ListDomainsReply, error) {
	res := s.manager.GetDomains()
	names := make([]string, len(res), len(res))
	for i, n := range res {
		names[i] = n[0]
	}
	doms := &pb.ListDomainsReply{
		Name: names,
	}
	return doms, nil
}

func (s *serverStruct) DeleteDomain(ctx context.Context, in *pb.Domain) (*pb.Empty, error) {
	return &pb.Empty{}, s.manager.DeleteDomain(in.GetName())
}

func (s *serverStruct) GetDomain(ctx context.Context, in *pb.Domain) (*pb.Domain, error) {
	return s.manager.GetDomain(in.GetName())
}