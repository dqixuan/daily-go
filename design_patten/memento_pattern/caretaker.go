package memento_pattern

// CareTaker 负责人
type CareTaker struct {
	mementoArray []*Memento
}

func (ct *CareTaker) AddMemento(m *Memento) {
	ct.mementoArray = append(ct.mementoArray, m)
}

func (ct *CareTaker) GetMemento() *Memento {
	return ct.mementoArray[len(ct.mementoArray) - 1]
}
