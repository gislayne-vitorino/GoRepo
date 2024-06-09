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
	var durations []int64

	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()

		err = client.Call("CrivoDeEratostenes.InvocaCrivoDeEratostenes", req, &rep)

		t2 := time.Now().Sub(t1).Nanoseconds()
		durations = append(durations, t2)

		fmt.Printf("RTT: %v: %v\n", t2, rep)
	}

	mean := calculateMean(durations)
	fmt.Printf("Mean duration: %v ns\n", mean)
}

func calculateMean(durations []int64) float64 {
	if len(durations) == 0 {
		return 0
	}
	var sum int64
	for _, duration := range durations {
		sum += duration
	}
	return float64(sum) / float64(len(durations))
}
