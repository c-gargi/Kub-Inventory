package main
 
import (
	"context"
	"fmt"
	"log"
	"net"

	"flag"
	"path/filepath"
	"time"
 
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
 
func (s *server) GetInventory(ctx context.Context, request *inventory.InventoryRequest) (*inventory.InventoryResponse, error) {
	log.Println(fmt.Sprintf("Request: %g", request.GetInterval()))

	/********/

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		respArray := [];
		for _, pod := range pods.Items {
		   //err = podInterface.Delete(pod.Name, &amp;kapi.DeleteOptions{})
		   fmt.Printf("Found pod %s in namespace %s\n", pod.Name, pod.Namespace)
		   eachPod := {Poddata: fmt.Sprintf("Found pod %s in namespace %s\n", pod.Name, pod.Namespace)}
		   respArray.push(eachPod)
		   // if err != nil {
		   //    log.Printf(“Error: %s”, err)
		   // }
		}

		hr := InventoryResponse{
			Items: respArray,
		}
		//fmt.Printf("There are %d pods in the cluster\n", pods.Items)

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		// namespace := "default"
		// pod := "example-xxxxx"
		// _, err = clientset.CoreV1().Pods("").Get(context.TODO(), pod, metav1.GetOptions{})
		// if errors.IsNotFound(err) {
		// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
		// 		pod, namespace, statusError.ErrStatus.Message)
		// } else if err != nil {
		// 	panic(err.Error())
		// } else {
		// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		// }

		time.Sleep(10 * time.Second)
	}





	/********/
 
	//return &inventory.InventoryResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetInterval())}, nil
	return &hr, nil
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
 
