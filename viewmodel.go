package mvc

func NewViewModel() (m *ViewModel) {
	m = new(ViewModel)
	m.ViewBag = make(map[string]interface{})
	m.Config = Config.clone()
	return
}

type ViewModel struct {
	Config  *Configuration
	Model   interface{}
	Errors  []error
	ViewBag map[string]interface{}
}
