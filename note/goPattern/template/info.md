模板方法模式：这种模式通过常用于为某种特定的操作定义一个模板或者算法模型

* 使用场景

以一次性密码（OTP：One Time Password）为例。我们常见的一次性密码有两种：短信密码（SMS OTP）或者邮件密码（Email OTP）。不过不管是短信密码还是邮件密码，它们的处理步骤都是一样的，步骤如下：
* 生成一串随机字符串
* 将字符串保存进缓存用来执行后续的验证
* 准备通知内容
* 发送通知
* 记录统计信息
在以上的步骤中，除了第4项“发送通知”的具体方式不一样，其他步骤都是不变的。即使以后有了新的一次性密码发送方式，可以预见以上的步骤也是不变的。