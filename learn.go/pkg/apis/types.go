package apis

type PersonalInformation struct {
	name string // 注意小写的私有成员变量在序列化时会被忽略

	Name   string
	Sex    string
	Tall   float64
	Weight float64
	Age    int
}
