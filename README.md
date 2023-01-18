# 后端已经部署，访问地址



## 1.用户注册、登录

用户注册验证邮箱



## 2.发布题目 

1. 这里先从其他网站爬取题目，经过整理后发布在网站上，写了一个leetcode爬虫可以爬leetcode的题目https://github.com/GDshenqingNo1/Crawler.git
2. 用户也可以自己发布题目（时间、作者、*难度、测试案例(应该包括至少一个案例的输入、输出、解释)）


## 3.获取题目、修改题目

1. 用户可以获取官方提供的题目或者获取其他用户提供的题目

2. 题目的发布者可以修改自己的题目




## 4.提交代码

1. 用户提交的代码应该包括数据的输出，并打印出结果
2. 将提交的代码写进测试.go文件
3. 生成uuid保存在本地code文件夹下 格式:/code/uuid/main.go
4. 用户线上提交代码，提交成功后在个人界面查看结果(compile error、run error、accept....)



## 5.获取评测结果

1. go build 没报错)compile pass 报错 compile error
2. go run 得到输出的结果与正确结果(msyql)比对 正确 accept 错误 not accept



- [x] 用户登录&注册
- [x] 获取个人信息
- [ ] jwt验证
- [x] 邮箱验证码
- [x] 获取题目信息、发布题目
- [x] 提交代码
- [x] 部署
- [ ] 前端页面

## mysql库结构

![mysql](.\mysql.png)