// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pointerdb

import (
	"context"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	p "storj.io/storj/pkg/paths"
	pb "storj.io/storj/protos/pointerdb"
)

var (
	mon = monkit.Package()
)

// PointerDB creates a grpcClient
type PointerDB struct {
	grpcClient pb.PointerDBClient
}

// Client defines the methods for interacting with a pointerdb client
type Client interface {
	Put(ctx context.Context, bucket string, path p.Path, pointer *pb.Pointer, APIKey []byte) error
	Get(ctx context.Context, bucket string, path p.Path, APIKey []byte) (*pb.Pointer, error)
	List(ctx context.Context, bucket string, startingPath []byte, limit int64, APIKey []byte) (
		paths []byte, truncated bool, err error)
	Delete(ctx context.Context, bucket string, path p.Path, APIKey []byte) error
}

// NewClient initializes a new pointerdb client
func NewClient(address string) (*PointerDB, error) {
	c, err := clientConnection(address, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	return &PointerDB{
		grpcClient: c,
	}, nil
}

// ClientConnection makes a server connection
func clientConnection(serverAddr string, opts ...grpc.DialOption) (pb.PointerDBClient, error) {
	conn, err := grpc.Dial(serverAddr, opts...)

	if err != nil {
		return nil, err
	}
	return pb.NewPointerDBClient(conn), nil
}

// Put makes a request to put a pointer at a given path in the pointerdb
func (pdb *PointerDB) Put(ctx context.Context, b string, path p.Path, pointer *pb.Pointer, APIKey []byte) (err error) {
	defer mon.Task()(&ctx)(&err)

	_, err = pdb.grpcClient.Put(ctx, &pb.PutRequest{Bucket: b, Path: path.Bytes(), Pointer: pointer, APIKey: APIKey})

	return err
}

// Get makes a request to get a pointer from a given path in the pointerdb
func (pdb *PointerDB) Get(ctx context.Context, b string, path p.Path, APIKey []byte) (pointer *pb.Pointer, err error) {
	defer mon.Task()(&ctx)(&err)

	res, err := pdb.grpcClient.Get(ctx, &pb.GetRequest{Bucket: b, Path: path.Bytes(), APIKey: APIKey})
	if err != nil {
		return nil, err
	}

	pointer = &pb.Pointer{}
	err = proto.Unmarshal(res.GetPointer(), pointer)

	return pointer, nil
}

// List makes a request for a list of paths from the pointerdb, given a starting path and a specific limit
func (pdb *PointerDB) List(ctx context.Context, startingPath p.Path, limit int64, APIKey []byte) (paths [][]byte, truncated bool, err error) {
	defer mon.Task()(&ctx)(&err)
	res, err := pdb.grpcClient.List(ctx, &pb.ListRequest{StartingPathKey: startingPath.Bytes(), Limit: limit, APIKey: APIKey})

	if err != nil {
		return nil, false, err
	}

	return res.Paths, res.Truncated, nil
}

// Delete makes a request to delete a path and pointer from the pointerdb
func (pdb *PointerDB) Delete(ctx context.Context, b string, path p.Path, APIKey []byte) (err error) {
	defer mon.Task()(&ctx)(&err)

	_, err = pdb.grpcClient.Delete(ctx, &pb.DeleteRequest{Bucket: b, Path: path.Bytes(), APIKey: APIKey})

	return err
}
