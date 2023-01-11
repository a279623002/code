package prototype

type Shiro struct {
	Name string
}

func (s *Shiro) Clone() Client {
	ss := *s // 取s地址的值赋予新变量ss
	return &ss
}

type Zzq struct {
	Name string
}

func (z *Zzq) Clone() Client {
	zz := *z
	return &zz
}
