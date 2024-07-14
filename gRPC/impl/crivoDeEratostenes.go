package impl

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/gislayne-vitorino/GoRepo/gRPC/proto/gen"
)

type CrivoDeEratostenesRPC struct {
}

func (c *CrivoDeEratostenesRPC) Crivo(ctx context.Context, req *gen.Request) (reply *gen.Reply, er error) {
	n := int(req.P1)

	reply = &gen.Reply{}

	// Error handling in gRPC method
	if n <= 0 {
		return nil, fmt.Errorf("Invalid input: P1 must be greater than zero")
	}

	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	// Determina o número de CPUs
	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup

	// Cria canal
	tasks := make(chan int, numCPU)

	// Função worker que marca como falso os múltiplos de cada número primo
	worker := func(tasks <-chan int, primes []bool, wg *sync.WaitGroup) {
		defer wg.Done()
		for p := range tasks {
			if primes[p] {
				for i := p * p; i <= n; i += p {
					primes[i] = false
				}
			}
		}
	}

	// Inicializa numCPU goroutines
	for w := 0; w < numCPU; w++ {
		wg.Add(1)
		go worker(tasks, primes, &wg)
	}

	// Envia números para o canal
	for p := 2; p*p <= n; p++ {
		tasks <- p
	}
	close(tasks)

	// Espera conclusão das goroutines
	wg.Wait()

	for p := 2; p <= n; p++ {
		if primes[p] {
			reply.N = append(reply.N, int32(p))
		}
	}

	return reply, nil
}
