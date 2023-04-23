package nats

var sc = openNATS()

func publisher(data []byte) {
	sc.Publish("msg", data)
}

func InitPublisher(data []byte) {
	publisher(data)
}
