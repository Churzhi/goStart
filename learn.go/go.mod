module learn.go

go 1.17

require (
	github.com/Churzhi/learn.go.tools v0.0.0-20211219072447-fcc7b1f15165
	github.com/spf13/cobra v1.3.0
)

// 使用相对文件夹的路径 .. 表示上一层文件夹
//replace learn.go.tools => ../learn.go.tools

replace github.com/Churzhi/learn.go.tools => github.com/Churzhi/learn.go.tools v0.0.2

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
//learn.go.tools v0.0.0-00010101000000-000000000000 // indirect
)
