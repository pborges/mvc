package mvc

func NewViewModel() (m *ViewModel) {
	m = new(ViewModel)
	m.ViewBag = make(map[string]interface{})
	return
}

type ViewModel struct {
	Model   interface{}
	Errors  []error
	Prefix  string
	ViewBag map[string]interface{}
}
