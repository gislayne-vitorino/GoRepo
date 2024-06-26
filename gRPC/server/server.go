// Para Quickstart do gRPC, acesse https://grpc.io/docs/languages/go/quickstart/
package main

import (
	"go-work/gRPC/shared"

	crivo "../impl"                                   //C:\\Users\\voxar\\Documents\\GoRepo\\gRPC3\\impl"
	gen1 "../proto/crivo/github.com/example/path/gen" //C:\\Users\\voxar\\Documents\\GoRepo\\gRPC3\\proto\\crivo\\github.com\\example\\path\\gen"

	//fibonacci "aulas/distribuida/fibonacci/impl"
	//gen2 "aulas/distribuida/fibonacci/proto"
	"fmt"
	"net"
	"strconv"

	"../shared" //C:\\Users\\voxar\\Documents\\GoRepo\\goRPC\\shared"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Cria listener
	endPoint := "localhost:" + strconv.Itoa(shared.GrpcPort)
	conn, err := net.Listen("tcp", endPoint)
	shared.ChecaErro(err, "Não foi possível criar o listener")

	// Cria um gRPC Server (“serviço de nomes” + servidor)”
	server := grpc.NewServer()

	// Registra a “Calculadora"/"Fibonacci" no servidor de nomes
	gen1.RegisterCrivoServer(server, &crivo.CrivoDeEratostenesRPC{})
	//gen2.RegisterFibonacciServer(server, &fibonacci.FibonacciRPC{})
	reflection.Register(server)

	fmt.Println("Servidor pronto ...")

	// Inicia servidor para atender requisções
	err = server.Serve(conn)
	shared.ChecaErro(err, "Falha ao iniciar servidor")
}
