package impl

import (
	"context"
	"fmt"

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

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	for p := 2; p <= n; p++ {
		if primes[p] {
			reply.N = append(reply.N, int32(p))
		}
	}

	return reply, nil
}
