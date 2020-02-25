package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
	v1 "worldserver/api/proto/v1"

	"google.golang.org/grpc"
)

const port int = 6000

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	grpcServer := grpc.NewServer()
	v1.RegisterWorldServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
func newServer() *server {
	s := &server{threadSafeSlice: threadSafeSlice{channels: make(map[int64]chan<- v1.Aircraft)}}
	return s
}

type server struct {
	threadSafeSlice threadSafeSlice
}

func (s *server) Update(_ context.Context, update *v1.UpdateMessage) (*v1.Response, error) {
	s.broadcast(*update.AircraftInfo)

	return &v1.Response{Status: "good"}, nil
}

func (s *server) Remove(_ context.Context, _ *v1.RemoveMessage) (*v1.Response, error) {
	panic("not implemented") // TODO: Implement
}

type threadSafeSlice struct {
	sync.Mutex
	channels map[int64]chan<- v1.Aircraft
}

func (slice *threadSafeSlice) Push(channel chan<- v1.Aircraft, key int64) {
	slice.Lock()
	defer slice.Unlock()

	slice.channels[key] = channel
	// slice.channels = append(slice.channels, channel)
}
func (slice *threadSafeSlice) Pop(channel chan<- v1.Aircraft, key int64) {
	slice.Lock()
	defer slice.Unlock()
	delete(slice.channels, key)
	// slice.channels = append(slice.channels, channel)
}
func (slice *threadSafeSlice) Iter(routine func(chan<- v1.Aircraft)) {
	slice.Lock()
	defer slice.Unlock()

	for _, channel := range slice.channels {
		routine(channel)
	}
}

func (s *server) Monitor(_ *v1.MonitorMessage, sender v1.WorldService_MonitorServer) error {
	fmt.Println("Client Connected")
	myChan := make(chan v1.Aircraft, 1)
	key := time.Now().UnixNano()
	s.threadSafeSlice.Push(myChan, key)
	defer s.threadSafeSlice.Pop(myChan, key)

	for {
		select {
		case aircraft := <-myChan:

			sender.Send(&aircraft)
		default:
		}
	}
}
func (s *server) broadcast(msg v1.Aircraft) {
	go func() {

		s.threadSafeSlice.Iter(func(c chan<- v1.Aircraft) {
			select {
			case c <- msg:
			default:
			}
		})

	}()
	return
}
