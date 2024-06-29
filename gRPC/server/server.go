// Para Quickstart do gRPC, acesse https://grpc.io/docs/languages/go/quickstart/
package main

import (
	"github.com/gislayne-vitorino/GoRepo/gRPC/shared"

	crivo "github.com/gislayne-vitorino/GoRepo/gRPC/impl"
	gen "github.com/gislayne-vitorino/GoRepo/gRPC/proto/gen"

	"fmt"
	"net"
	"strconv"

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

	// Registra o Crivo de Eratóstenes no servidor de nomes
	gen.RegisterCrivoServer(server, &crivo.CrivoDeEratostenesRPC{})
	reflection.Register(server)

	fmt.Println("Servidor pronto ...")

	// Inicia servidor para atender requisções
	err = server.Serve(conn)
	shared.ChecaErro(err, "Falha ao iniciar servidor")
}
