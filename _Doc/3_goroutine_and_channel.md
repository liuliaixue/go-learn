# Goroutine + Channel 实践

ref: https://studygolang.com/articles/2423

## 一个简单的http请求后端机制
```
func (m *SomeController) PorcessSomeTask() {
    var task models.Task
    if err := task.Parse(m.Ctx.Request); err != nil {
        m.Data["json"] = err 
        m.ServeJson()
        return
    }
    task.Process()
    m.ServeJson()
}
```

`task.Process()` 如果非常耗时，则前端会一直收不到响应，如果请求不在意处理结果，只需要知道服务器收到了处理请求。
比如 进行复杂的io处理，请求或等待其他服务器的响应。

## 立即响应的后端机制

```
func (m *SomeController) PorcessSomeTask() {
    var task models.Task
    if err := task.Parse(m.Ctx.Request); err != nil {
        m.Data["json"] = err 
        m.ServeJson()
        return
    }
    // go here
    go task.Process()
    m.ServeJson()
}
```
