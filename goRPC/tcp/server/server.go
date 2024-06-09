package main

import (
	"fmt"
	"go-work/impl"
	"go-work/shared"
	"net"
	"net/rpc"
	"strconv"
)

func main() {
	mathService := new(impl.CrivoDeEratostenes)

	server := rpc.NewServer()
	err := server.Register(mathService)
	shared.ChecaErro(err, "Erro ao registrar o crivo de Eratostenes")

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "Erro ao iniciar o listener")
	defer listener.Close()

	fmt.Printf("Servidor RPC pronto (RPC-TCP) na porta %v...\n", shared.CrivoPort)

	server.Accept(listener)
}
