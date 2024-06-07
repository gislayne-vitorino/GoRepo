package main

import (
	"shared"
	"encoding/json"
	"fmt"
	"net"
	"os"
	//"strconv"
)

func main() {

	//t1 := time.Now()
	CrivoDeEratostenesClientTCP(10)
	//CalculatorClientUDP(n)
	//tTotal := time.Now().Sub(t1)

	//fmt.Println(tTotal.Nanoseconds()/1000000)
	//CalculatorClientUDP(n)
}

func CrivoDeEratostenesClientTCP(n int) {
	var response shared.Reply

	// retorna o endereço do endpoint
	r, err := net.ResolveTCPAddr("tcp", "localhost:1314")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	/// connecta ao servidor (sem definir uma porta local)
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha conexõa
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// cria enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < n; i++ {

		// prepara request
		msgToServer := shared.Request{i} //request esta enviando o numero declarado na main

		// serializa e envia request para o servidor
		err = jsonEncoder.Encode(msgToServer)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do servidor
		err = jsonDecoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(response)
	}
}