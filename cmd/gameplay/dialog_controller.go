package gameplay

type DialogRequest struct {
	Type string
	Data interface{}
}

type DialogController struct {
	OpenRequest  *DialogRequest
	CloseRequest bool
}

func NewDialogController() DialogController {
	return DialogController{}
}

func OpenDialog(request *DialogRequest) {
	GameInstance.DialogController.OpenRequest = request
}

func CloseDialog() {
	GameInstance.DialogController.CloseRequest = true
}
