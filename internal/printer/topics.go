package printer

type topics struct {
	text   string
	prefix string
}

func NewTopics(t ...interface{}) *topics {
	return &topics{text: doPrintbs(t...)}
}

func (this *topics) Prefix(s ...interface{}) *topics {
	this.prefix = doPrintbs(s...)

	return this
}

func (this *topics) Default() {
	stdout.WriteString(this.prefix + PREFIX_LIST_DEFAULT+" "+this.text + endl)
}

func (this *topics) Done() {
	stdout.WriteString(this.prefix + PREFIX_LIST_DONE+" "+this.text + endl)
}

func (this *topics) Danger() {
	stdout.WriteString(this.prefix + PREFIX_LIST_DANGER+" "+this.text + endl)
}

func (this *topics) Warning() {
	stdout.WriteString(this.prefix + PREFIX_LIST_WARNING+" "+this.text + endl)
}
