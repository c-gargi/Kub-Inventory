package main
 
import (
	"context"
	"fmt"
	"log"
	"net"

	"flag"
	"path/filepath"
	_"time"
 
	"github.com/c-gargi/Kub-Inventory/pkg/inventory"
	"google.golang.org/grpc"

	_ "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)
 
type server struct {
	inventory.UnimplementedInventoryServiceServer
}
 
func (s *server) Inventory(ctx context.Context, request *inventory.InventoryRequest) (*inventory.InventoryResponse, error) {
	fmt.Printf("Request: %d\n", request.GetInterval())

	/********/

	var kubeconfig *string
	var respArray string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//panic(err.Error())
	}
	
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))	
	for _, pod := range pods.Items {
	   //err = podInterface.Delete(pod.Name, &amp;kapi.DeleteOptions{})
	   //fmt.Printf("Found pod %s in namespace %s\n", pod.Name, pod.Namespace)
	   respArray += "Found pod "+pod.Name+" in namespace "+pod.Namespace+" \n"
	   // if err != nil {
	   //    log.Printf(“Error: %s”, err)
	   // }
	}
	
	fmt.Printf(respArray)
	
	return &inventory.InventoryResponse{Items: respArray}, nil

}


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
 
