package client

func NewCluster() *clientStruct {
	return &clientStruct{
		name: "client",
		num:  1,
	}
}

type clientStruct struct {
	name string
	num  int
}
