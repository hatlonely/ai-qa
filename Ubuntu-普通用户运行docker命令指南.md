# Ubuntu 普通用户运行 docker 命令指南

本文档提供在 Ubuntu 系统中允许普通用户无需 sudo 权限运行 docker 命令的详细步骤。

## 问题描述

默认情况下，在 Ubuntu 系统中只有 root 用户或拥有 sudo 权限的用户才能执行 docker 命令。这是因为 Docker 守护进程绑定到一个 Unix 套接字，该套接字默认归属于 root 用户。

## 解决方案

要让普通用户运行 docker 命令而无需每次使用 sudo，需要将该用户添加到 docker 组中。

### 详细步骤

1. 首先确认 docker 组是否存在：
   ```bash
   getent group docker
   ```

2. 如果 docker 组不存在，创建该组：
   ```bash
   sudo groupadd docker
   ```

3. 将用户添加到 docker 组：
   ```bash
   sudo usermod -aG docker 用户名
   ```
   将"用户名"替换为实际用户名。

4. 应用新的组成员资格（无需重启系统）：
   ```bash
   newgrp docker
   ```
   或者注销并重新登录系统使变更生效。

5. 验证普通用户是否可以使用 docker 命令：
   ```bash
   docker run hello-world
   ```

### 临时解决方案

如果遇到权限问题，可以临时调整 Docker 守护进程的 socket 权限：
```bash
sudo chmod 666 /var/run/docker.sock
```
注意：此命令是临时解决方案，系统重启后权限会恢复默认设置。

## 安全注意事项

将用户添加到 docker 组本质上授予了该用户 root 等效权限，因为 Docker 允许挂载主机文件系统并修改系统配置。请确保只将受信任的用户添加到 docker 组。