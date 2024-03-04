package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go_learn/pkg/console"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 允许且必须传一个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// 从模板中创建文件（做好变量替换）
	createFileFromStub(filePath, "cmd", model)

	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.CMD" + model.StructName)
	console.Warning("Please edit main.go's app.Commands slice to register command")
}
