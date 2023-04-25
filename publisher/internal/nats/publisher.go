package nats

var sc = openNATS()

func Publisher(data []byte) {
	sc.Publish("msg", data)
}
