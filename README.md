# Go mod sample

## Command line
    go mod init github/shishuwu/my-go-mod-sample



## Other
[Go mod sample](https://mp.weixin.qq.com/s/TvTlz3uKIBqgg1FjlAItPQ)

    go mod init 创建了一个新的模块，初始化 go.mod 文件并且生成相应的描述
    go build, go test 和其它构建代码包的命令，会在需要的时候在 go.mod 文件中添加新的依赖项
    go list -m all 列出了当前模块所有的依赖项
    go get 修改指定依赖项的版本（或者添加一个新的依赖项）
    go mod tidy 移除模块中没有用到的依赖项。
