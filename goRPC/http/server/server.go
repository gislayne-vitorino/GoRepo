package main

import (
	"fmt"
	"go-work/impl"
	"go-work/shared"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {
	crivo := new(impl.CrivoDeEratostenes)

	server := rpc.NewServer()
	server.RegisterName("CrivoDeEratostenes", crivo)

	server.HandleHTTP("/", "/debug")

	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "Servidor não inicializado")

	fmt.Printf("Servidor RPC pronto (RPC-HTTP) na porta %v...\n", shared.CrivoPort)
	http.Serve(l, nil)
}
