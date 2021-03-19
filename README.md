## 打车系统

### 项目简介
使用tcp socket网路编程开发的一个简易的打车系统

### 功能特性
<ul>
<li>乘客和司机注册,登陆,退出</li>
<li>乘客下单,取消订单</li>
<li>司机接单,完成订单</li>
<li>乘客与司机互相沟通</li>
</ul>

### 部署步骤
<ol>
<li>
开启go module<br>
set GO111MODULE=on(windows)<br>
export GO111MODULE=on(linux)
</li>
<li>
下载go mod中指定的所有依赖<br>
go mod download
</li>
<li>
将.env.example修改为.env文件,并修改其相关信息
</li>
<li>启动相关服务<br>
go run trail_didi_3/server/main.go(启动服务端，开启监听)<br
go run trail_didi_3/driver_client/main.go(启动司机端，开始接单)<br>
go run trail_didi_3/user_client/main.go(启动乘客端，开始下单)<br>
</li>
</ol>

### 目录结构
* config(`相关配置信息初始化目录`)
* driver_client(`司机客户端`)
* pkg(`手动编写的包服务`)
* server(`服务端`)
* user_client(`乘客客户端`)
* go.mod(`项目的第三方包依赖信息文件`)



