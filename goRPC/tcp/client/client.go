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
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "Erro ao conectar ao servidor")
	defer client.Close()

	req := shared.Request{Number: 10}
	rep := shared.Reply{}
	err = client.Call("CrivoDeEratostenes.InvocaCrivoDeEratostenes", req, &rep)
	shared.ChecaErro(err, "Erro na invocação remota...")

	fmt.Printf("Add(%v,%v) = %v \n", req.Number, rep.Result)
}

func ClientePerf() {
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CrivoPort))
	shared.ChecaErro(err, "Erro ao conectar ao servidor")
	defer client.Close()

	req := shared.Request{Number: 10}
	rep := shared.Reply{}
	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			err = client.Call("CrivoDeEratostenes.InvocaCrivoDeEratostenes", req, &rep)
			shared.ChecaErro(err, "Erro na invocação remota...")
		}
		fmt.Printf("tcp;%v: %v\n", time.Now().Sub(t1).Milliseconds(), rep)
	}
}
