# Ubuntu 24 自动挂载NAS服务指南

本文介绍如何在Ubuntu 24服务器中设置自动挂载NAS服务，让系统在启动时自动连接到网络存储设备。

## 准备工作

根据NAS使用的协议，安装必要的软件包：

```bash
# 如果使用NFS协议
sudo apt update
sudo apt install nfs-common

# 如果使用CIFS/SMB协议(用于Samba或Windows共享)
sudo apt update
sudo apt install cifs-utils
```

## 使用fstab自动挂载

### 1. 创建挂载点

```bash
sudo mkdir -p /mnt/nas
```

### 2. 配置fstab文件

编辑`/etc/fstab`文件：

```bash
sudo nano /etc/fstab
```

根据NAS的协议类型，添加相应的配置：

#### NFS协议配置

```
192.168.1.100:/share /mnt/nas nfs defaults,_netdev 0 0
```

#### CIFS/SMB协议配置

```
//192.168.1.100/share /mnt/nas cifs username=用户名,password=密码,vers=3.0,_netdev 0 0
```

### 3. 使用凭据文件（更安全）

创建凭据文件以避免密码明文存储：

```bash
sudo nano /etc/nas-credentials
```

添加凭据信息：
```
username=用户名
password=密码
```

设置权限：
```bash
sudo chmod 600 /etc/nas-credentials
```

更新fstab配置：
```
//192.168.1.100/share /mnt/nas cifs credentials=/etc/nas-credentials,vers=3.0,_netdev 0 0
```

### 4. 测试挂载

```bash
sudo mount -a
```

验证挂载状态：
```bash
df -h | grep nas
```

## 高级配置选项

### 网络稳定性选项

添加自动重连和超时设置：

```
//192.168.1.100/share /mnt/nas cifs credentials=/etc/nas-credentials,vers=3.0,_netdev,auto,x-systemd.automount,x-systemd.idle-timeout=60,x-systemd.device-timeout=5s,x-systemd.mount-timeout=5s 0 0
```

### 文件权限设置

指定文件所有权：

```
//192.168.1.100/share /mnt/nas cifs credentials=/etc/nas-credentials,vers=3.0,_netdev,uid=1000,gid=1000 0 0
```

### 只读挂载

设置为只读访问：

```
//192.168.1.100/share /mnt/nas cifs credentials=/etc/nas-credentials,vers=3.0,_netdev,ro 0 0
```

## 故障排除

如果挂载失败，检查以下几点：

1. 确认NAS服务器是否可达：`ping 192.168.1.100`
2. 检查共享名称是否正确
3. 验证用户名和密码是否正确
4. 查看系统日志获取详细错误信息：`sudo journalctl -xe`
