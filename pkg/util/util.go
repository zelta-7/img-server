package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/streadway/amqp"
)

func Download(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func Connect(queueName string) (amqp.Queue, *amqp.Channel, error) {
	var q amqp.Queue
	var ch *amqp.Channel
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("CONNECTION FAILED: %v\n", err)
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		fmt.Printf("CHANNEL FAILED: %v\n", err)
		log.Fatal(err)
	}
	defer ch.Close()

	q, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("QUEUE CREATION FAILED: %v\n", err)
		log.Fatal(err)
	}
	return q, ch, nil
}

func Compress(data []byte) ([]byte, error) {
	var compressedData bytes.Buffer
	writer := gzip.NewWriter(&compressedData)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	return compressedData.Bytes(), nil
}

func generateUniqueName(basename, fileextension string) string {
	timestamp := time.Now().Format("20060102150405") //YYYYMMDDHHMMSS

	uniqueName := fmt.Sprintf("%s_%s%s", basename, timestamp, fileextension)
	return uniqueName
}

// Stores images to the mentioned folder name and returns the file path for stored image
func StoreCommpressedImage(data [][]byte, folderName string) ([]string, error) {
	var compressedPath []string

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			return nil, err
		}
	}
	for i, compressedImg := range data {
		fileName := fmt.Sprintf("api-c_img%d.jpg", i)
		filePath := filepath.Join(folderName, fileName)

		err := os.WriteFile(filePath, compressedImg, 0644)
		if err != nil {
			return nil, err
		}
		compressedPath = append(compressedPath, filePath)
	}
	return compressedPath, nil
}
