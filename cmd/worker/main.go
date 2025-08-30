package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Obtiene la URL de NATS desde la variable de entorno o usa el valor por defecto
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://nats:4222"
	}
	// Obtiene el tópico a suscribirse
	topic := os.Getenv("NATS_INGEST_TOPIC")
	if topic == "" {
		topic = "ingestions"
	}

	// Conexión a NATS
	nc, err := nats.Connect(url, nats.Name("worker"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	// Obtiene el contexto JetStream
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Suscripción durable al tópico
	_, err = js.Subscribe(topic, func(m *nats.Msg) {
		// Aquí se debe procesar el mensaje recibido
		log.Printf("processing message: %s", string(m.Data))
		// TODO: transformar payload, persistir en Postgres y tocar Redis
		m.Ack()
	}, nats.Durable("worker-durable"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("worker listening on %s", topic)
	<-context.Background().Done()
	time.Sleep(time.Second)
}
