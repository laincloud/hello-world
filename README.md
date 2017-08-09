# hello-world

一个简单的 LAIN 应用，提供 HTTP 服务，请求 `/` 时返回 `Hello, LAIN. You are the ${n}th visitor.`。

> - ${n} 指调用次数。
> - 用到了 [LAIN Service](https://laincloud.gitbooks.io/white-paper/usermanual/service.html)。

## 分支说明

本仓库有多个分支，便于 LAIN 的逐步进阶演示。对应的演示步骤见 [LAIN Tutorial](https://laincloud.gitbooks.io/white-paper/content/tutorial/)。

`basic`: 基础版，只是展示 LAIN 最基本的功能。

`master`: 进阶版，包括了prepare，Redis service等功能。
