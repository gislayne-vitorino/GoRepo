package main

import (
	gen1 "aulas/distribuida/calculadora/grpc/proto2"
	//gen2 "aulas/distribuida/fibonacci/proto"
	"aulas/distribuida/shared"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Estabelece conexão com o servidor
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	endPoint := "localhost" + ":" + strconv.Itoa(shared.GrpcPort)
	conn, err := grpc.Dial(endPoint, opt)
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor em"+endPoint)

	// fecha conexões
	defer conn.Close()

	// cria um proxy
	crivo := gen1.NewCrivoClient(conn)

	// cria um contexto para execução remota
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var i int32
	for i = 0; i < shared.SampleSize; i++ {
		// invoca operação remota
		reqCrivo := gen1.Request{P1: i}

		repCrivo, err := crivo.Crivo(ctx, &reqCrivo)
		shared.ChecaErro(err, "Erro ao invocar a operação remota.")

		fmt.Printf("Add(%v,%v)=%v", reqCrivo.P1, repCrivo.Primes)
	}
}
