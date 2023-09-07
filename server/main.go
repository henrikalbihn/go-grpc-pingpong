// ADAPTED FROM: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// * Package server implements the server side of the PingPong service.
package main

import (
	"context"
	"fmt"
	pb "go-grpc-pingpong/pb"
	ut "go-grpc-pingpong/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

// * server is used to implement pingpong.PingPongServer
type server struct {
	pb.UnimplementedPingPongServer // Embed this to satisfy the interface
}

// * Ping implements the PingPongServer interface
func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	ut.LogPayload(in)
	return &pb.PongResponse{Message: "Pong"}, nil
}

// * main starts the server
func main() {
	port := ut.GetPort()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		ut.LogF(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	log.Printf("Starting server on port %s", port)
	pb.RegisterPingPongServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		ut.LogF(err, "failed to serve")
	}
}
