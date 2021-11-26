package invoker

import calc "go-test/ying/package"

//调用其他包步骤：
//1.修改goland设置：file->settings->go->go module
//勾选 Enable Go Modules (vgo) integration，proxy填：https://goproxy.cn
//2.在项目目录下执行
//go mod init
func invokeAdd() {
	calc.Add(3, 4)
}
