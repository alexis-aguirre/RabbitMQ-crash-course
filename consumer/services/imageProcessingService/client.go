package imageProcessingService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (client *ImageProcessingClient) ProcessPlate(data interface{}) error {
	log.Println("Sending ", fmt.Sprint(data)+" to image processing service...")
	reqBody, err := json.Marshal(data)
	if err != nil {
		print(err)
	}
	endpoint := fmt.Sprintf("%s/process", client.BaseUrl)
	resp, err := http.Post(endpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
