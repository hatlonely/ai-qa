# Ubuntu 24 安装卸载 Docker

本文档介绍了在 Ubuntu 24 系统上如何正确卸载和安装 Docker 的方法。

## Docker 卸载步骤

### 基本卸载方法

1. 停止 Docker 服务：
```bash
sudo systemctl stop docker
sudo systemctl stop docker.socket
sudo systemctl stop containerd
```

2. 卸载 Docker 相关软件包：
```bash
sudo apt remove --purge docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
```

3. 删除 Docker 数据目录和配置文件：
```bash
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd
sudo rm -rf /etc/docker
```

4. 移除 Docker 添加的 apt 源：
```bash
sudo rm -f /etc/apt/sources.list.d/docker.list
```

5. 更新 apt 缓存：
```bash
sudo apt update
```

### 常见问题：卸载后命令仍存在

如果执行 `apt remove` 后 Docker 命令仍然存在，可能原因如下：

1. Docker 安装方式多样，需要确认当前安装方式：
   - 检查命令位置：`which docker`
   - 检查已安装包：`dpkg -l | grep -i docker`

2. 针对不同安装方式的卸载命令：
   - Ubuntu 仓库版本：`sudo apt remove --purge docker.io`
   - Snap 版本：`sudo snap remove docker`
   - 手动安装：`sudo rm -f $(which docker)`

3. 彻底清理所有 Docker 相关包：
```bash
sudo apt remove --purge docker docker.io docker-engine docker-ce docker-ce-cli containerd.io docker-compose-plugin docker-buildx-plugin
sudo apt autoremove --purge
```

## Docker 安装步骤（最新版本）

1. 移除旧版本（如果有）：
```bash
sudo apt remove docker docker.io containerd runc
```

2. 安装必要的依赖：
```bash
sudo apt update
sudo apt install -y ca-certificates curl gnupg lsb-release
```

3. 添加 Docker 官方 GPG 密钥：
```bash
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

4. 添加 Docker 官方仓库：
```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

5. 安装 Docker：
```bash
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

6. 验证安装：
```bash
sudo docker --version
sudo docker run hello-world
```

7. 配置非 root 用户运行 Docker（可选）：
```bash
sudo usermod -aG docker $USER
newgrp docker  # 立即应用更改，或重新登录
```