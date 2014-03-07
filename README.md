weiyan
===========
基于studygolang模仿糗百的网站


# 本地搭建一个 Weiyan 社区 #

1、下载 weiyan 代码
	
	git clone https://github.com/philsong/studygolang

2、下载安装依赖库（如果依赖库下载不下来可以联系我）

	cd studygolang/websites/code/thirdparty
	// windows 下执行
	getpkg.bat
	// linux/mac 下执行
	sh getpkg

3、编译并运行 studygolang

先编译

	// 接着上一步
	cd ../weiyan/
	// windows 下执行
	install.bat
	// linux/mac 下执行
	sh install
	
这样便编译好了 studygolang，下面运行 studygolang。（运行前可以根据需要修改 config/config.json.example 配置，当然首先得改名为 config.json）

	// windows 下执行
	start.bat
	// linux/mac 下执行
	sh start

一切顺利的话，studygolang 应该就启动了。

4、浏览器中查看

在浏览器中输入：http://127.0.0.1:9090

应该就能看到了。

5、建立数据库

运行起来了，但没有建数据库。源码中有一个 databases 文件夹，里面有建表和初始化的sql语句。之前这些sql之前，在mysql数据库中建立一个数据库：studygolang，之后执行这些sql语句。

根据你的数据库设置，修改上面提到的 `config/config.json` 对应的配置，重新启动 studygolang.（通过restart脚本重新启动）
