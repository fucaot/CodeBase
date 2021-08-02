### 规范以及约定
- 项目根目录下为各语言目录
- 各语言根目录（项目二级目录)下为基础语法
- 从三级目录开始对应包结构构建包实现


## 层级信息
<details>
<summary>展开查看</summary>
<pre><code>
.
├── C
│   └── include
├── Golang
│   ├── slice.go
│   ├── src
│   │   ├── bufio.go
│   │   ├── encoding
│   │   │   └── json.go
│   │   ├── fmt.go
│   │   ├── log.go
│   │   ├── net
│   │   │   ├── http.go
│   │   │   └── url.go
│   │   ├── os.go
│   │   └── strings.go
│   ├── string.go
│   └── struct.go
├── Java
│   ├── Lamdba.java
│   └── String.java
├── Python
│   └── bin
├── README.md
└── Shell
    └── date.sh
10 directories, 15 files
</code></pre>
</details>

### 关于Golang

Golang部分的代码为了方便阅读与理解，通过go语言自身的go_test进行编写，大部分方法可以通过工具直接运行方法来查看案例的情况，当然，也有例外。

为了保证所有的_test.go文件都可以顺利运行，Golang部分所有的代码通过小写开头，若对应API的包开头为大写，则使用双小写进行替代，例如:

`net/http.Request -> ./Golang/src/net/http/rrequest_test.go`