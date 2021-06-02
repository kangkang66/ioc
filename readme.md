# 模块化编程-依赖注入

## wire的使用：
> https://segmentfault.com/a/1190000039185137

## go依赖注入思想
1. 一切都是以对象（结构体）的方法进行操作，把依赖的资源定义为属性
2. 定义一个对象的构造函数，在形参中声明对象依赖的资源，并返回对象的实例化
3. 第三方的包可以直接创建一个构造函数来返回对象的实例化，参考pkg中的redis

## 项目介绍：
- cmd ： 入口
- internal/api ：接口定义
- internal/module：模块化功能
- pkg ：通用的功能包

### 依赖关系
cmd -> api -> module（service -> model） -> pkg

因此在cmd/wire.go中定义下注入顺序
```
func InitApi() (*api.Api,error) {
	panic(wire.Build(pkg.Provider, module.Provider, api.Provider))
}
```

而每个Provider可以再包裹当前包内的子包，这样就不至于很分散，由父包管理对外暴露的小包。
例如module.provider中包裹了，如果有新的小包，只需要更改父包的provider。
```
var Provider = wire.NewSet(
	task.Provider,
	sign_in.Provider,
	luckywheel.Provider,
)
```


