package imageProcessingService

type ImageProcessingClient struct {
	BaseUrl string
}

func NewImageProcessingClient(url string) *ImageProcessingClient {
	return &ImageProcessingClient{
		BaseUrl: url,
	}
}
