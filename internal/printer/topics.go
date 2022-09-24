package printer

type _Topics struct {
	text   string
	prefix string
}

func NewTopics(t ...interface{}) *_Topics {
	return &_Topics{text: doPrintbs(t...)}
}

func (model *_Topics) Prefix(s ...interface{}) *_Topics {
	model.prefix = doPrintbs(s...)

	return model
}

func (model *_Topics) Default() {
	stdout.WriteString(model.prefix + PREFIX_LIST_DEFAULT+" "+model.text + "\n")
}

func (model *_Topics) Done() {
	stdout.WriteString(model.prefix + PREFIX_LIST_DONE+" "+model.text + "\n")
}

func (model *_Topics) Danger() {
	stdout.WriteString(model.prefix + PREFIX_LIST_DANGER+" "+model.text + "\n")
}

func (model *_Topics) Warning() {
	stdout.WriteString(model.prefix + PREFIX_LIST_WARNING+" "+model.text + "\n")
}
