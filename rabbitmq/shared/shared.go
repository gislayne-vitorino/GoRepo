package shared

import (
	"log"
	"math/rand"
)

type Request struct {
	Number int
}

type Reply struct {
	Result []interface{}
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

const CrivoPort = 4040
const StatisticSample = 30
const SampleSize = 10
const GrpcPort = 80
const RequestQueue = "request_queue"
const ResponseQueue = "response_queue"
