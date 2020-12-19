// shippy-vessel-service/handler.go
package main

import (
	"context"

	pb "github.com/leplasmo/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

// FindAvailable
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	//Find the next available vessel
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = UnmarshalVessel(vessel)
	return nil
}

// Create
func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}