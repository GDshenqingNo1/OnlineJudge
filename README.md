1.用户注册、登录
用户注册验证邮箱



2.发布题目 

这里先从其他网站爬取题目，经过整理后发布在网站上

用户也可以自己发布题目（时间、作者、*难度、测试案例(应该包括至少一个案例的输入、输出、解释)）


3.获取题目、修改题目

用户可以获取官方提供的题目或者获取其他用户提供的题目

题目的发布者可以修改自己的题目



4.提交代码
用户提交的代码应该包括数据的输出，并打印出结果
将提交的代码写进测试.go文件
生成uuid保存在本地code文件夹下 格式:/code/uuid/main.go
用户线上提交代码，提交成功后在个人界面查看结果(compile error、run error、accept....)


5.评测代码

评测代码在docker中进行
container挂载.go文件
runner.go运行指定uuid中的main.go文件

6.获取评测结果
go build 没报错(?怎么检测)compile pass 报错 compile error
go run 得到输出的结果与正确结果(msyql)比对 正确 accept 错误 not accept