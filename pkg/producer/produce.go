package producer

import (
	"github.com/streadway/amqp"
	"github.com/zelta-7/img-server/pkg/util"
)

var QueueName string = "imageQueue"

func QueueImage(url string) error {
	imageData, err := util.Download(url)

	if err != nil {
		return err
	}

	q, ch, err := util.Connect(QueueName)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "image/jpeg",
			Body:        imageData,
			//MessageId:   file.Name(),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
