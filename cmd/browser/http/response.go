package http

import (
	"io"

	"golang.org/x/net/html"
)

type KitsuneClientResponse struct {
	Node *html.Node
}

func NewKitsuneClientResponse(reader io.Reader) (*KitsuneClientResponse, *error) {
	response := &KitsuneClientResponse{}
	var err error
	if response.Node, err = html.Parse(reader); err != nil {
		return nil, &err
	}

	return response, nil
}
