package main

import (
	"encoding/json"
	"fmt"
	"go-work/shared"
	"net"
	"os"
	"time"
)

//"strconv"

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
	var durations []int64

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

	// fecha conexão
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// cria enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	// prepara request
	msgToServer := shared.Request{Number: n} //request esta enviando o numero declarado na main

	for i := 0; i < 100; i++ {
		t1 := time.Now()
		err = jsonEncoder.Encode(msgToServer)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		err = jsonDecoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		primes := make([]int, len(response.Result))
		for i, v := range response.Result {
			primes[i] = int(v.(float64))
		}

		t2 := time.Now().Sub(t1).Nanoseconds()
		durations = append(durations, t2)

		fmt.Printf("RTT: %v: %v\n", t2, primes)
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
