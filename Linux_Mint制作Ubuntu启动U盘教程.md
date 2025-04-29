# Linux Mint 制作 Ubuntu 启动 U 盘教程

本文档总结了如何在 Linux Mint 系统中使用 Ubuntu 24.04.1 服务器版 ISO 文件制作 U 盘安装镜像的方法。

## 可用方法

### 方法一：使用图形界面工具 USB Image Writer

1. 插入 U 盘
2. 打开"USB Image Writer"工具（Mint 自带）
3. 选择 ISO 文件(`ubuntu-24.04.1-live-server-amd64.iso`)
4. 选择 U 盘设备
5. 点击"写入"按钮开始制作

### 方法二：使用 Etcher 工具

1. 安装 Etcher：
   ```bash
   sudo apt update
   sudo apt install etcher-electron
   ```
   如果无法通过包管理器安装，可以从官网下载：https://www.balena.io/etcher/

2. 打开 Etcher
3. 选择 ISO 文件
4. 选择 U 盘
5. 点击"Flash"开始制作

### 方法三：使用命令行（dd 命令）

1. 确定 U 盘设备名称：
   ```bash
   lsblk
   ```

2. 卸载 U 盘（如已挂载）：
   ```bash
   sudo umount /dev/sda1
   ```

3. 使用 dd 命令写入：
   ```bash
   sudo dd bs=4M if=/path/to/ubuntu-24.04.1-live-server-amd64.iso of=/dev/sda status=progress oflag=sync
   ```

## 实际操作示例

以下是针对具体情况的操作示例：

ISO 文件路径: `/home/hatlonely/Downloads/ubuntu-24.04.1-live-server-amd64.iso`

设备信息（通过 `lsblk` 命令获取）:
```
NAME        MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
sda           8:0    1 57.8G  0 disk 
└─sda1        8:1    1 57.7G  0 part /media/hatlonely/U PAN
```

执行步骤：
1. 卸载 U 盘：
   ```bash
   sudo umount /dev/sda1
   ```

2. 使用 dd 命令制作启动盘：
   ```bash
   sudo dd bs=4M if=/home/hatlonely/Downloads/ubuntu-24.04.1-live-server-amd64.iso of=/dev/sda status=progress oflag=sync
   ```

## 注意事项

- 制作前请备份 U 盘中的重要数据，因为此操作会清除 U 盘上所有数据
- 使用命令行方法时，确保正确识别 U 盘设备，避免覆盖系统磁盘
- 完成后，可以安全弹出 U 盘
- 制作完成后，重启电脑并从 U 盘启动即可开始安装 Ubuntu 服务器版
