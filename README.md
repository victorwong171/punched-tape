[![GoDoc](https://pkg.go.dev/badge/github.com/victorwong171/punched-tape?utm_source=godoc)](https://pkg.go.dev/github.com/victorwong171/punched-tape)
[![Go Report Card](https://goreportcard.com/badge/github.com/victorwong171/punched-tape)](https://goreportcard.com/report/github.com/victorwong171/punched-tape)
[![codecov](https://codecov.io/github/victorwong171/punched-tape/branch/master/graph/badge.svg)](https://codecov.io/github/victorwong171/punched-tape)
![GitHub License](https://img.shields.io/github/license/victorwong171/punched-tape)


# punched-tape

## 项目概述
`punched-tape` 是一个用于处理工单（ticket）的Go语言项目，它提供了工单模板的验证、工单审批等功能，支持联合签署、串行签署和任意人签署等多种签署方式。

## 项目结构

├─.github
│  └─workflows
├─.idea
├─internal
│  └─state_machine
├─models
└─ticket
├─template
└─ticket


### 主要目录和文件说明
- `models/`：包含项目的核心数据模型，如 `Ticket`、`Disposal`、`TicketTemplate` 等。
- `ticket/`：包含工单相关的处理逻辑，如模板验证和工单审批测试。
- `step_config.go`、`template.go`、`ticket.go`：提供了构建 `StepConfig`、`TicketTemplate` 和 `Ticket` 的构建器。

## 安装依赖
确保你已经安装了Go 1.20或更高版本。在项目根目录下运行以下命令来安装依赖：
```sh
go mod tidy
```
## 代码示例
- 创建工单模板

```go
package main

import (
    "github.com/victorwong171/punched-tape/models"
    "github.com/victorwong171/punched-tape/punched_tape"
)

func main() {
    // 创建 StepConfig
    nextStep := &models.NextStep{
        Step:      "next",
        Operation: "submit",
    }
    stepConfig := punched_tape.NewStepConfigBuilder().
        SetStep("start").
        SetOperator([]string{"user"}).
        SetNext([]*models.NextStep{nextStep}).
        Build()

    // 创建 TicketTemplate
    template := punched_tape.NewTemplateBuilder().
        SetStartStep("start").
        SetEndStep("end").
        AddConfig(stepConfig).
        Build()
}
```
- 验证工单模板
```go
package main

import (
    "github.com/victorwong171/punched-tape/models"
    "github.com/victorwong171/punched-tape/ticket/template"
    "github.com/victorwong171/go-utils/desc/set"
)

func main() {
    // 创建 Validator
    signTypeSet := set.Setify(models.JointlySign, models.SerialSign, models.AnyoneSign)
    v := &template.validator{
        signTypeSet: signTypeSet,
    }

    // 验证模板
    err := v.Validate(template)
    if err != nil {
        // 处理错误
    }
}
```
## 测试
在项目根目录下运行以下命令来执行测试：
```sh
go test ./...
```
## 贡献

如果你想为这个项目做出贡献，请遵循以下步骤：

Fork 这个仓库。

创建一个新的分支 (git checkout -b feature/your-feature)。

提交你的更改 (git commit -am 'Add some feature')。

将更改推送到你的分支 (git push origin feature/your-feature)。

提交一个 Pull Request。

许可证
本项目采用 [许可证名称] 许可证。有关详细信息，请参阅 LICENSE 文件。
