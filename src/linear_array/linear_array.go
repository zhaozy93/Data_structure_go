package linear_array

type Array struct{
	d []interface{}
	len int
}

func NewArray() *Array {
	return &Array{
		d: make([]interface{}, 0),
		len: 0,
	}
}

func (a *Array) Len()int{
	return a.len
}


func (a *Array) Push(val interface{})*Array{
	a.d = append(a.d, val)
	a.len += 1
	return a
}

func (a *Array) PushWithIndex(val interface{}, index int)*Array{
	if index >= a.len{
		return a.Push(index)
	}
	a.len += 1
	d := make([]interface{}, a.len)
	copy(d[:index], a.d[:index])
	d[index] = val
	copy(d[index+1:], a.d[index:])
	a.d = d 
	return a
}

func (a *Array) DeleteWithIndex(index int)(val interface{}){
	if index >= a.len{
		return a
	}
	val = a.d[index]
	a.len -= 1
	d := make([]interface{}, a.len)
	copy(d[:index], a.d[:index])
	copy(d[index:], a.d[index+1:])
	a.d = d 
	return
}