package shared

import "log"

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

const CrivoPort = 4040
const StatisticSample = 30
const SampleSize = 10000
