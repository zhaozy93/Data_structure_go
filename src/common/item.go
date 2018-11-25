package common

type Item struct {
	pre   *Item
	next  *Item
	value interface{}
}

func NewItem(data interface{}) *Item {
	return &Item{value: data}
}

func (item *Item) GetValue() interface{} {
	return item.value
}

func (item *Item) SetValue(data interface{}) *Item {
	item.value = data
	return item
}

func (item *Item) GetPre() *Item {
	return item.pre
}

func (item *Item) SetPre(pre *Item) *Item {
	item.pre = pre
	return item
}

func (item *Item) GetNext() *Item {
	return item.next
}

func (item *Item) SetNext(next *Item) *Item {
	item.next = next
	return item
}
