package gameplay

type KurinDialogRequest struct {
	Type string
	Data interface{}
}

type KurinDialogController struct {
	OpenRequest  *KurinDialogRequest
	CloseRequest bool
}

func NewKurinDialogController() KurinDialogController {
	return KurinDialogController{}
}

func OpenKurinDialog(request *KurinDialogRequest) {
	GameInstance.DialogController.OpenRequest = request
}

func CloseKurinDialog() {
	GameInstance.DialogController.CloseRequest = true
}
