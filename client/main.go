// ADAPTED FROM: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
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

// * Package client implements the client side of the PingPong service.
package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-pingpong/pb"
	ut "go-grpc-pingpong/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ut.GetAddr()
	// Set up a connection to the server.
	log.Printf("Establishing gRPC connection - %s...", addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ut.LogF(err, "could not connect")
	}
	defer conn.Close()
	c := pb.NewPingPongClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		ut.LogF(err, "could not ping")
	}
	log.Printf("Response: %s", r.GetMessage())
}
