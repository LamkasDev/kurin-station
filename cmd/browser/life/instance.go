package life

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/browser/contextgfx"
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/http"
)

type KitsuneInstance struct {
	Renderer *gfx.KitsuneRenderer
	Client   *http.KitsuneClient
	Context  *context.KitsuneContextRender
}

func NewKitsuneInstance() (*KitsuneInstance, *error) {
	instance := &KitsuneInstance{}
	var err *error
	if instance.Renderer, err = gfx.NewKitsuneRenderer(); err != nil {
		return nil, err
	}
	if instance.Client, err = http.NewKitsuneClient(); err != nil {
		return nil, err
	}
	if err := GoToWebsite(instance, "http://motherfuckingwebsite.com"); err != nil {
		return nil, err
	}

	return instance, nil
}

func GoToWebsite(instance *KitsuneInstance, address string) *error {
	request := &http.KitsuneClientRequest{
		Address: address,
	}
	response, err := http.FulfillKitsuneRequest(instance.Client, request)
	if err != nil {
		return err
	}
	if instance.Context, err = contextgfx.NewKitsuneContextRender(instance.Renderer, response); err != nil {
		return err
	}
	instance.Context.Icon = instance.Renderer.Icons.Icon
	instance.Renderer.Window.SetTitle(instance.Context.Title)

	return nil
}

func RunKitsuneInstance(instance *KitsuneInstance) {
	gfx.ClearKitsuneRenderer(instance.Renderer, instance.Context)
	gfx.RenderKitsuneRenderer(instance.Renderer, instance.Context)
	gfx.PresentKitsuneRenderer(instance.Renderer, instance.Context)
}

func FreeKitsuneInstance(instance *KitsuneInstance) {
	gfx.FreeKitsuneRenderer(instance.Renderer)
	http.FreeKitsuneClient(instance.Client)
}
