package impl

import (
	gen "C:\\Users\\voxar\\Documents\\GoRepo\\gRPC\\proto"

	"golang.org/x/net/context"
)

type CrivoDeEratostenesRPC struct {
}

func (c *CrivoDeEratostenes) InvocaCrivoDeEratostenes(req *gen.Request) (reply *gen.Reply, error) {
	n := req.Number

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

	reply := &gen.Reply{}
	for p := 2; p <= n; p++ {
		if primes[p] {
			reply.Result = append(reply.Result, p)
		}
	}

	return nil
}
/*
func (t *CrivoDeEratostenesRPC) Add(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 + request.P2
	return reply, nil
}
func (t *CrivoDeEratostenesRPC) Sub(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 - request.P2
	return reply, nil
}
func (t *CrivoDeEratostenesRPC) Mul(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 * request.P2
	return reply, nil
}
func (t *CrivoDeEratostenesRPC) Div(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 / request.P2
	return reply, nil
}
*/