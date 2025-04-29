# Linux CIFS挂载NAS共享指南

本文档提供了在Linux系统中挂载CIFS/Samba网络共享的完整步骤和常见问题解决方案。

## 前提条件

在开始之前，确保安装了必要的软件包：

```bash
sudo apt update
sudo apt install cifs-utils smbclient
```

`cifs-utils`提供了挂载CIFS共享所需的基本工具，而`smbclient`用于测试连接和浏览共享内容。

## 配置挂载

### 1. 创建挂载点

```bash
mkdir -p /home/用户名/挂载点
sudo chown 用户名:用户名 /home/用户名/挂载点
```

### 2. 创建凭证文件

为了安全存储访问凭证，创建一个专用文件：

```bash
sudo mkdir -p /etc/samba
sudo touch /etc/samba/credentials
sudo chmod 600 /etc/samba/credentials
```

编辑此文件并添加凭证：

```bash
sudo nano /etc/samba/credentials
```

文件内容必须严格按照以下格式：

```
username=您的用户名
password=您的密码
```

确保文件格式正确，不要有多余的空格，并且"username"和"password"保持小写。

## 手动挂载

### 基本挂载命令

```bash
sudo mount -t cifs //服务器IP/共享名 /挂载点 -o credentials=/etc/samba/credentials,vers=3.0,uid=$(id -u 用户名),gid=$(id -g 用户名),file_mode=0770,dir_mode=0770
```

参数说明：
- `vers=3.0`：指定SMB协议版本（可选值：1.0, 2.0, 2.1, 3.0, 3.1.1）
- `uid`和`gid`：设置挂载文件系统的所有者ID
- `file_mode`和`dir_mode`：设置文件和目录的权限模式

### 不使用凭证文件

如果需要直接在命令行中提供凭证（不推荐用于生产环境）：

```bash
sudo mount -t cifs //服务器IP/共享名 /挂载点 -o username=用户名,password=密码,vers=3.0,uid=$(id -u 用户名),gid=$(id -g 用户名),file_mode=0770,dir_mode=0770
```

## 自动挂载

编辑`/etc/fstab`文件，添加以下行：

```
//服务器IP/共享名 /挂载点 cifs credentials=/etc/samba/credentials,vers=3.0,uid=用户ID,gid=组ID,file_mode=0770,dir_mode=0770,_netdev,nofail 0 0
```

重要参数：
- `_netdev`：告诉系统这是网络设备，等待网络服务启动后再挂载
- `nofail`：如果挂载失败，不阻止系统启动

保存更改后，重新加载systemd配置并测试挂载：

```bash
sudo systemctl daemon-reload
sudo mount -a
```

## 故障排除

### 1. 检查网络连接

```bash
ping 服务器IP
```

### 2. 验证共享是否存在

使用smbclient检查可用共享：

```bash
smbclient -L //服务器IP -U 用户名
```

### 3. 测试直接连接到共享

```bash
smbclient //服务器IP/共享名 -U 用户名
```

### 4. 常见错误和解决方案

#### "cannot mount read-only"错误

- 尝试指定协议版本：`vers=3.0`或`vers=2.1`
- 确保提供了正确的凭证
- 检查用户是否有访问该共享的权限

#### "failed to connect to IPC (rc=-13)"错误

- 表示权限被拒绝，通常是认证问题
- 检查凭证文件格式是否正确
- 确保服务器允许您的IP地址访问

#### "SMB1 disabled"消息

- 服务器不支持旧的SMB1协议
- 使用`vers=3.0`或`vers=2.1`选项

### 5. 卸载共享

```bash
sudo umount /挂载点
```

## 安全注意事项

- 确保凭证文件的权限设置为600（只有root可读写）
- 考虑使用更安全的身份验证方法，如Kerberos（适用于域环境）
- 定期更改用户密码，增强安全性
- 仅在可信网络上使用CIFS共享
