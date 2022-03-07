package main

type PersonalInformation struct {
	Id     string  `json:"id,omitempty" gorm:"id:name"`
	Name   string  `json:"name,omitempty" gorm:"column:name"`
	Sex    string  `json:"sex,omitempty" gorm:"column:sex"`
	Tall   float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age    int64   `json:"age,omitempty" gorm:"column:age"`
}

// 需要告诉数据库和 PersonalInformation 的链接关系
func (*PersonalInformation) TableName() string {
	return "" //最好复制表名
}
