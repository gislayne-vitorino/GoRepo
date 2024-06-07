package main

import "go-work/shared"
import "go-work/impl"


import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	CrivoDeEratostenesServerTCP()

	//CalculatorServerUDP()

	fmt.Scanln()
}

func CrivoDeEratostenesServerTCP() {
	//  define o endpoint do servidor TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1314")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// cria um listener TCP
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Servidor TCP aguardando conexões na porta 1314...")

	for {
		// aguarda/aceita conexão
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// processa requests da conexão
		go HandleTCPConnection(conn)
	}
}


func HandleTCPConnection(conn net.Conn) {
	var msgFromClient shared.Request

	// Close connection
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	// Cria coder/decoder JSON
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		// recebe & unmarshall requests do cliente
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		// processa request
		r := impl.CrivoDeEratostenes{}.InvocaCrivoDeEratostenes(msgFromClient)

		// cria resposta
		msgToClient := shared.Reply{[]interface{}{r}}

		// serializa & envia resposta para o cliente
		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}