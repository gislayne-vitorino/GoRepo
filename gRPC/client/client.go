package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	gen "github.com/gislayne-vitorino/GoRepo/gRPC/proto/gen"
	"github.com/gislayne-vitorino/GoRepo/gRPC/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Estabelece conexão com o servidor
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	endPoint := "localhost" + ":" + strconv.Itoa(shared.GrpcPort)
	conn, err := grpc.NewClient(endPoint, opt)
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor em "+endPoint)

	// Fecha conexões
	defer conn.Close()

	// Cria um proxy
	crivo := gen.NewCrivoClient(conn)
	fmt.Println("Proxy criado...")

	// Cria um arquivo .txt para salvar os resultados
	file, err := os.Create("../outputs_concurrency/output_1000_10000.txt")
	if err != nil {
		fmt.Println("Erro na criação do arquivo:", err)
		return
	}
	defer file.Close()

	var i int32
	var j int32
	for i = 0; i < shared.StatisticSample; i++ {
		for j = 0; j < shared.SampleSize; j++ {
			t1 := time.Now()
			// Cria um contexto para execução remota
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)

			// Invoca operação remota
			reqCrivo := gen.Request{P1: 1000}

			repCrivo, err := crivo.Crivo(ctx, &reqCrivo)
			shared.ChecaErro(err, "Erro ao invocar a operação remota.")

			cancel()
			t2 := time.Now().Sub(t1).Nanoseconds()

			fmt.Println(repCrivo.N)

			// Converte int64 pra string e salva no arquivo
			_, err = file.WriteString(strconv.FormatInt(t2, 10) + "\n")
			//_, err := fmt.Fprintf(file, "%d\n", strconv.FormatInt(t2, 10))
			if err != nil {
				fmt.Println("Erro na escrita do arquivo:", err)
			}
		}
	}
}
