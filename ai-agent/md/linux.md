# Linux 面试笔记

> Linux 是服务端主流操作系统，面试核心围绕常用命令、进程管理、网络、文件系统、Shell 脚本和性能排查。

---

## 一、文件与目录操作

| 命令 | 作用 | 常用示例 |
|---|---|---|
| `ls` | 列出目录内容 | `ls -la` 显示全部（含隐藏文件） |
| `cd` | 切换目录 | `cd ~` 回到家目录，`cd -` 回到上次目录 |
| `pwd` | 显示当前路径 | `pwd` |
| `mkdir` | 创建目录 | `mkdir -p a/b/c` 递归创建 |
| `rm` | 删除 | `rm -rf dir` 强制递归删除 |
| `cp` | 复制 | `cp -r src dst` 递归复制目录 |
| `mv` | 移动/重命名 | `mv old new` |
| `touch` | 创建空文件/更新修改时间 | `touch file.txt` |
| `cat` | 查看文件内容 | `cat file` |
| `head` | 看文件前 N 行 | `head -n 20 file` |
| `tail` | 看文件后 N 行 | `tail -f log.txt` 实时跟踪 |
| `find` | 查找文件 | `find . -name "*.go"` |
| `grep` | 文本搜索 | `grep -n "error" log.txt` |
| `wc` | 统计行数/字数 | `wc -l file` 统计行数 |
| `tar` | 打包压缩 | `tar -czvf a.tar.gz dir/` |
| `df` | 磁盘空间 | `df -h` 人类可读 |
| `du` | 目录大小 | `du -sh dir/` |

---

## 二、进程管理

| 命令 | 作用 |
|---|---|
| `ps` | 查看进程状态 |
| `ps aux` | 查看所有进程详情 |
| `ps aux \| grep nginx` | 查找特定进程 |
| `top` | 实时查看进程资源占用 |
| `htop` | top 的增强版（需安装） |
| `kill` | 发送信号终止进程 |
| `kill -9 PID` | 强制终止 |
| `killall` | 按名称终止 |
| `pkill` | 按名称匹配终止 |
| `nohup` | 后台运行，忽略挂起信号 |
| `&` | 放到后台运行 |
| `jobs` | 查看后台任务 |
| `fg` | 前台运行后台任务 |
| `bg` | 后台继续运行 |

```bash
# 后台运行并输出到 nohup.out
nohup python train.py &

# 查看进程树
pstree -p

# 查看某进程的线程
ps -T -p PID
```

---

## 三、网络相关

| 命令 | 作用 |
|---|---|
| `ifconfig` / `ip addr` | 查看网卡信息 |
| `ping` | 测试连通性 |
| `netstat` | 查看网络连接 |
| `ss` | netstat 的替代（更快） |
| `lsof` | 查看打开的文件和网络连接 |
| `curl` | HTTP 请求工具 |
| `wget` | 下载工具 |
| `scp` | 安全拷贝 |
| `ssh` | 远程登录 |
| `traceroute` | 追踪路由路径 |
| `nslookup` / `dig` | DNS 查询 |
| `tcpdump` | 抓包分析 |
| `nc` (netcat) | 网络瑞士军刀 |

```bash
# 查看所有监听端口
netstat -tlnp
ss -tlnp

# 查看某端口被谁占用
lsof -i :8080

# 发起 GET 请求
curl http://localhost:8080/health

# 下载文件
curl -O http://example.com/file.zip

# 测试端口连通
nc -zv localhost 3306
```

---

## 四、权限与用户

| 命令 | 作用 |
|---|---|
| `chmod` | 修改文件权限 |
| `chown` | 修改文件所有者 |
| `sudo` | 以超级用户执行 |
| `su` | 切换用户 |
| `useradd` | 添加用户 |
| `passwd` | 修改密码 |
| `groups` | 查看用户所属组 |

```bash
# 权限数字表示
# r=4, w=2, x=1
chmod 755 file     # rwxr-xr-x
chmod +x script.sh # 添加执行权限

# 修改所有者
chown user:group file
```

---

## 五、文本处理三剑客

### grep

```bash
grep "error" log.txt           # 基本搜索
grep -i "error" log.txt        # 忽略大小写
grep -v "error" log.txt        # 反向匹配（不含 error）
grep -n "error" log.txt        # 显示行号
grep -r "pattern" dir/         # 递归搜索
grep -E "a|b" log.txt          # 扩展正则
```

### awk

```bash
# 打印第 1、3 列
awk '{print $1, $3}' file.txt

# 按逗号分隔，打印第 2 列
awk -F',' '{print $2}' file.csv

# 统计行数
awk 'END{print NR}' file.txt

# 求和
awk '{sum += $1} END {print sum}' numbers.txt
```

### sed

```bash
# 替换文本
sed 's/old/new/g' file.txt

# 直接修改文件
sed -i 's/old/new/g' file.txt

# 删除第 3 行
sed '3d' file.txt

# 打印 5-10 行
sed -n '5,10p' file.txt
```

---

## 六、Shell 脚本基础

### 变量与判断

```bash
#!/bin/bash

name="world"
echo "hello, $name"

# 判断文件是否存在
if [ -f "file.txt" ]; then
    echo "文件存在"
fi

# 判断字符串非空
if [ -n "$name" ]; then
    echo "name 不为空"
fi

# 数值比较
if [ $a -gt $b ]; then
    echo "a > b"
fi
```

### 循环

```bash
# for 循环
for i in 1 2 3 4 5; do
    echo $i
done

# 或
for ((i=0; i<5; i++)); do
    echo $i
done

# while 循环
count=0
while [ $count -lt 5 ]; do
    echo $count
    ((count++))
done
```

### 函数

```bash
#!/bin/bash

function greet() {
    echo "hello, $1"
}

greet "world"   # 输出 hello, world
```

---

## 七、性能排查

### CPU

```bash
# 查看 CPU 占用最高的进程
top

# 查看某进程的 CPU 和内存
ps -p PID -o pid,ppid,cmd,%cpu,%mem

# 查看 CPU 核数
nproc
cat /proc/cpuinfo | grep "processor" | wc -l
```

### 内存

```bash
# 查看内存使用
free -h

# 查看进程内存
cat /proc/PID/status | grep VmRSS
```

### 磁盘 IO

```bash
# 查看磁盘 IO
iostat -x 1

# 查看磁盘空间
df -h

# 找出大文件
du -ah /path | sort -rh | head -n 20
```

### 网络

```bash
# 查看网络连接状态
netstat -an | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'

# 查看网络流量
iftop

# 抓包
tcpdump -i eth0 port 8080
```

---

## 八、常见面试题

### Q1：查看某端口被哪个进程占用？

**答**：
```bash
lsof -i :8080
# 或
netstat -tlnp | grep 8080
# 或
ss -tlnp | grep 8080
```

### Q2：查找大文件并删除？

**答**：
```bash
# 查找大于 100MB 的文件
find / -type f -size +100M

# 找出目录下最大的 10 个文件
du -ah /path | sort -rh | head -n 10

# 清空日志文件（不删除文件）
> access.log
```

### Q3：怎么查看 Linux 系统负载？

**答**：
```bash
uptime
# 输出：load average: 0.52, 0.58, 0.59
# 分别表示 1分钟、5分钟、15分钟的平均负载

# 查看每个 CPU 的负载
cat /proc/loadavg
```

> 负载值 ≈ CPU 核数表示满负载，> 核数表示过载。

### Q4：进程和线程的区别？

**答**：
| 特性 | 进程 | 线程 |
|---|---|---|
| 资源 | 独立地址空间 | 共享进程地址空间 |
| 切换 | 开销大 | 开销小 |
| 通信 | IPC（管道、消息队列等） | 直接读写共享内存 |
| 崩溃 | 不影响其他进程 | 可能导致整个进程崩溃 |

### Q5：Linux 文件系统目录结构？

| 目录 | 作用 |
|---|---|
| `/` | 根目录 |
| `/bin` | 基本命令 |
| `/etc` | 配置文件 |
| `/home` | 用户主目录 |
| `/var` | 可变数据（日志、缓存） |
| `/tmp` | 临时文件 |
| `/usr` | 用户程序 |
| `/opt` | 可选软件包 |
| `/proc` | 进程信息（虚拟文件系统） |
| `/dev` | 设备文件 |

### Q6：软链接和硬链接的区别？

| 特性 | 硬链接 | 软链接（符号链接） |
|---|---|---|
| 本质 | 同一个 inode 的多个名字 | 指向另一个文件路径 |
| 跨文件系统 | 不支持 | 支持 |
| 指向目录 | 不支持 | 支持 |
| 原文件删除 | 仍然可用 | 失效（悬空） |

```bash
ln file hard_link      # 硬链接
ln -s file soft_link   # 软链接
```

### Q7：怎么排查服务启动失败？

**答**：
1. 看日志：`journalctl -u service_name` 或 `cat /var/log/xxx`
2. 看端口是否被占：`lsof -i :port`
3. 看权限：`ls -la` 检查文件权限
4. 看依赖：`ldd binary` 检查动态库
5. 直接运行：`./binary` 看终端输出

### Q8：crontab 定时任务格式？

```bash
# 格式：分 时 日 月 周 命令
# 每天 2 点执行备份
0 2 * * * /home/user/backup.sh

# 每 5 分钟执行
*/5 * * * * /home/user/check.sh

# 每周一早上 8 点
0 8 * * 1 /home/user/weekly.sh
```

---

## 九、一句话总结

- **文件操作**：ls/cd/cp/mv/rm/find/grep 天天用
- **进程管理**：ps/top/kill/nohup/& 要熟练
- **网络排查**：netstat/ss/lsof/curl/ping/tcpdump
- **文本处理**：grep/awk/sed 三剑客
- **性能排查**：top/free/df/iostat 定位瓶颈
- **权限管理**：chmod/chown/sudo 基础安全

> **面试口诀：ls 看文件，ps 看进程，netstat 看端口，grep 搜日志，awk 做统计，sed 做替换，top 找瓶颈**
