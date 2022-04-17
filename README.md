# chat-project 使用说明
## 命令行操作
```
# 启动服务端，不指定参数默认8080端口
./chat server -p 8080 

# 注册用户
./chat register -u "张三" -p "123456"
{Acccount: 3919}

./chat register -u "李四" -p "123456"
{Acccount: 23850}

# 登陆用户
./chat login -a "3919" -p "123456"

# 查看在线用户
./chat show

# 开启聊天
Flags:
  -n, --account string              发起聊天人账号
  -c, --chatting-with string        被聊天人的账号，以逗号分割 A,B,C (default "all")

./chat with -a "3919" -n "23850"
```

## 进入聊天功能后命令菜单
```
# send关键字给对方发送信息
$ send 你好

# list关键字查看登陆用户
$ list

# quit退出客户端
$ quit

# subscribe 查看谁找我聊天
$ subscribe
```


## 未完成功能
- 此程序进入聊天功能必须通过quit进行退出，不可通过CTRL + C等方式退出
- 一对多聊天功能暂未完善
- 单元测试未完成
