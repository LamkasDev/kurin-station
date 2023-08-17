package http

import "net/http"

type KitsuneClient struct {
	Http *http.Client
}

func NewKitsuneClient() (*KitsuneClient, *error) {
	client := &KitsuneClient{
		Http: &http.Client{},
	}

	return client, nil
}

func FulfillKitsuneRequest(client *KitsuneClient, request *KitsuneClientRequest) (*KitsuneClientResponse, *error) {
	httpResponse, httpErr := client.Http.Get(request.Address)
	if httpErr != nil {
		return nil, &httpErr
	}

	response, err := NewKitsuneClientResponse(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func FreeKitsuneClient(client *KitsuneClient) {

}
