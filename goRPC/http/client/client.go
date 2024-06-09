package main

import (
	"fmt"
	"go-work/shared"
	"net/rpc"
	"strconv"
	"time"
)

func main() {
	ClientePerf()
}

func Cliente() {
	clientCrivo, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "O Servidor não está pronto")
	defer func(clientCrivo *rpc.Client) {
		var err = clientCrivo.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor do Crivo de Eratostenes...")
	}(clientCrivo)

	req := shared.Request{Number: 10}
	rep := shared.Reply{}
	err = clientCrivo.Call("CrivoDeEratostenes.InvocaCrivoDeEratostenes", req, &rep)
	shared.ChecaErro(err, "Erro na invocação do Crivo de Eratostenes remoto...")

	fmt.Printf("InvocaCrivoDeEratostenes(%v) = %v \n", req.Number, rep.Result)
}

func ClientePerf() {
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "O Servidor não está pronto")
	defer func(clientCrivo *rpc.Client) {
		var err = clientCrivo.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor do Crivo de Eratostenes...")
	}(client)

	req := shared.Request{Number: 10}
	rep := shared.Reply{}
	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			err = client.Call("CrivoDeEratostenes.InvocaCrivoDeEratostenes", req, &rep)
			shared.ChecaErro(err, "Erro na invocação do Crivo de Eratostenes remoto...")
		}
		fmt.Printf("http;%v: %v\n", time.Now().Sub(t1).Milliseconds(), rep)
	}
}
