package quota

func NewQuota() *quotaStruct {
	return &quotaStruct{
		name: "quota",
		num:  1,
	}
}

type quotaStruct struct {
	name string
	num  int
}
