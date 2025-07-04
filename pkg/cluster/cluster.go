package cluster

func NewCluster() *clusterStruct {
	return &clusterStruct{
		name: "cluster",
		num:  1,
	}
}

type clusterStruct struct {
	name string
	num  int
}
