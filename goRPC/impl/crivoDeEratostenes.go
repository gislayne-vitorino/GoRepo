package impl

import "go-work/shared"

type CrivoDeEratostenes struct{}

func (c *CrivoDeEratostenes) InvocaCrivoDeEratostenes(req shared.Request, reply *shared.Reply) error {
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

	for p := 2; p <= n; p++ {
		if primes[p] {
			reply.Result = append(reply.Result, p)
		}
	}

	return nil
}
