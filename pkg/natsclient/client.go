package natsclient

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/nats-io/stan.go"
)

const Client = "NatsStreamingWebClient"

func FastPublish(host, cluster, queue, message string) error {
	nc, err := NatsConnection(host, cluster)
	if err != nil {
		return err
	}
	defer func() {
		nc.Close()
	}()

	err = nc.Publish(queue, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func NatsConnection(hostUrl, clusterName string) (stan.Conn, error) {
	clientId := Client+"-"+randomString(5)

	conn, err := stan.Connect(clusterName, clientId, stan.NatsURL(hostUrl))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func randomString(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Println(err)
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	rs = rs[0 : len(rs)-1]

	return rs
}
