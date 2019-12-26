package v1

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes"

	v1 "github.com/laughtt/templatesGrpc/service+sql/api/protov1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//AuthServiceServer connection
type AuthServiceServer struct {
	db *sql.DB
}

//NewAuthServiceServer create a AUTHservice
func NewAuthServiceServer(dd *sql.DB) v1.AuthServiceServer {
	return &AuthServiceServer{db: dd}
}

//Config  datbase connection
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// DB Datastore parameters selsction
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

//Connect database
func (s *AuthServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *AuthServiceServer) SendMessage(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	
	c, err := s.connect(ctx)

	if err != nil {
		return &v1.CreateResponse{
			Id: 12,
			Message: "nop bitch",
			Error: err.Error(),
		}, err
	}
	defer c.Close()
	_ ,  err = ptypes.Timestamp(req.GetReminder())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}
	return &v1.CreateResponse{
		Id: 12,
		Message: "complete",
	}, nil
}
