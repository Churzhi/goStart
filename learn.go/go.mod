module learn.go

go 1.17

require (
	github.com/Churzhi/learn.go.tools v0.0.0-20211219072447-fcc7b1f15165
	github.com/armstrongli/go-bmi v0.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/spf13/cobra v1.3.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2

)

// 使用相对文件夹的路径 .. 表示上一层文件夹
//replace learn.go.tools => ../learn.go.tools

replace (
	github.com/Churzhi/learn.go.tools => github.com/Churzhi/learn.go.tools v0.0.2
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
//learn.go.tools v0.0.0-00010101000000-000000000000 // indirect
)
