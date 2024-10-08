package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	common "GolandProjects/apaxos-gautamsardana/api_common"
	"GolandProjects/apaxos-gautamsardana/server_alice/config"
	"GolandProjects/apaxos-gautamsardana/server_alice/logic"
	"GolandProjects/apaxos-gautamsardana/server_alice/logic/inbound"
)

type Server struct {
	common.UnimplementedPaxosServer
	Config *config.Config
}

func (s *Server) ProcessTxn(ctx context.Context, req *common.ProcessTxnRequest) (*emptypb.Empty, error) {
	err := logic.ProcessTxn(ctx, req, s.Config)
	if err != nil {
		log.Printf("Error processing txn: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}

func (s *Server) Prepare(ctx context.Context, req *common.Prepare) (*emptypb.Empty, error) {
	err := inbound.Prepare(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error receiving prepare: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}

func (s *Server) Promise(ctx context.Context, req *common.Promise) (*emptypb.Empty, error) {
	err := inbound.Promise(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error receiving promise: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}

func (s *Server) Accept(ctx context.Context, req *common.Accept) (*emptypb.Empty, error) {
	err := inbound.Accept(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error receiving promise: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}

func (s *Server) Accepted(ctx context.Context, req *common.Accepted) (*emptypb.Empty, error) {
	err := inbound.Accepted(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error receiving promise: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}
