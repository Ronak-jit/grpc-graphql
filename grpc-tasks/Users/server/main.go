package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/Ronak-Searce/grpc-tasks/users/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var users []*pb.UserInfo

type userServer struct {
	pb.UnimplementedUsererviceServer
}

func main() {
	initUsers()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)

	}
	s := grpc.NewServer()

	pb.RegisterUsererviceServer(s, &userServer{})

	log.Printf("server is listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initUsers() {
	user1 := &pb.UserInfo{Id: "1", FirstName: "Ronak", LastName: "Babu"}
	user2 := &pb.UserInfo{Id: "2", FirstName: "Sanjeeb", LastName: "Kumar"}
	users = append(users, user1)
	users = append(users, user2)

}

func (s *userServer) GetUser(ctx context.Context, in *pb.Id) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.UserInfo{}

	for _, user := range users {
		if user.GetId() == in.GetValue() {
			res = user
			break
		}
	}
	return res, nil
}

func (s *userServer) CreatUser(ctx context.Context,
	in *pb.UserInfo) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)
	res := pb.UserInfo{}
	// res.Value = strconv.Itoa(rand.Intn(100000000))
	// in.Id = res.GetValue()
	res.Id = strconv.Itoa(rand.Intn(100000000))
	res.FirstName = in.FirstName
	res.LastName = in.LastName

	in.Id = res.GetId()
	in.FirstName = res.GetFirstName()
	in.LastName = res.GetLastName()
	users = append(users, in)
	return &res, nil
}

func (s *userServer) UpdateUser(ctx context.Context,
	in *pb.UserInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, user := range users {
		if user.GetId() == in.GetId() {
			users = append(users[:index], users[index+1:]...)
			in.Id = user.GetId()
			users = append(users, in)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func (s *userServer) DeleteUser(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, user := range users {
		if user.GetId() == in.GetValue() {
			users = append(users[:index], users[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}
