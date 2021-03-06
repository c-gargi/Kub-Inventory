package main
 
import (
	"context"
	"log"
	"time"
 
	"github.com/c-gargi/Kub-Inventory/pkg/inventory"
	"google.golang.org/grpc"
)
 
func main() {
	log.Println("Client running ...")
 
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
 
	client := inventory.NewInventoryServiceClient(conn)
 
	request := &inventory.InventoryRequest{Interval: 1}
 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
 
	
 	
 	for {
 		response, err := client.Inventory(ctx, request)
		if err != nil {
			log.Fatalln(err)
		}
 		log.Println(response.GetItems())
 		time.Sleep(10 * time.Second)
 	}
	
}
