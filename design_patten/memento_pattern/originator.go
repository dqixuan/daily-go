package memento_pattern

type Originator struct {
	State string
}

// SetState 设置状态
func (o *Originator) SetState(s string) {
	o.State = s
}

//
func (o *Originator) CreateSnapshot() {

}