package main
 
import (
	"context"
	"fmt"
	"log"
	"net"
 
	//"github.com/c-gargi/Kub-Inventory/pkg/inventory"
	"google.golang.org/grpc"
)
 
type server struct {}
 
func main() {
	log.Println("Server running ...")
 
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
 
	srv := grpc.NewServer()
	inventory.RegisterInventoryServiceServer(srv, &server{})
 
	log.Fatalln(srv.Serve(lis))
}
 
func (s *server) Inventory(ctx context.Context, request *inventory.InventoryRequest) (*inventory.InventoryResponse, error) {
	log.Println(fmt.Sprintf("Request: %g", request.GetAmount()))
 
	return &inventory.InventoryResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetAmount())}, nil
}