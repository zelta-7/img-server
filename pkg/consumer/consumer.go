package consumer

import (
	"github.com/zelta-7/img-server/pkg/producer"
	"github.com/zelta-7/img-server/pkg/util"
)

// TODO: While calling StoreCompressedImage, foldername passed in the parameter exists outside this folder, MAKE IT WORK

func Consume(queueName string) ([]string, error) {
	var compressedData [][]byte

	q, ch, err := util.Connect(producer.QueueName)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	for msg := range msgs {
		imageData := msg.Body

		compressedImg, err := util.Compress(imageData)
		if err != nil {
			return nil, err
		}
		compressedData = append(compressedData, compressedImg)
	}

	compressedPath, err := util.StoreCommpressedImage(compressedData, "compressedImg") // Folder exist outside the consumer folder

	return compressedPath, nil
}
