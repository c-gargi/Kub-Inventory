# Kub-Inventory
### gRPC client/server to get kubernetes pod details

- Commands
1. Inside the root folder, run the following command to generate pb file.

    ```sh
    protoc -I pkg/inventory/ pkg/inventory/inventory.proto --go_out=plugins=grpc:pkg/inventory/
    ```
1. Inside server folder, run the following command to built and run gRPC server.

    ```sh
    go build && ./server
    ```


1. Inside client folder, run the following command to built and run gRPC client.

    ```sh
    go build && ./client
    ```
    
1. To start the minikube
   
    ```sh
    minikube start --driver=docker
    ```
