# NVIDIA Linux 驱动安装指南

## 系统环境

- **系统**：Linux Mint 22.1 (基于Ubuntu)
- **内核**：6.8.0-58-generic
- **显卡**：NVIDIA GeForce RTX 4090
- **当前驱动**：NVIDIA Driver 550.120
- **显示管理器**：LightDM
- **桌面环境**：Cinnamon

## 安装步骤

### 1. 准备工作

1. 备份重要数据

2. 清理旧驱动
   ```bash
   sudo apt-get purge nvidia*
   sudo apt-get autoremove
   ```

3. 禁用Nouveau驱动
   ```bash
   sudo bash -c "echo blacklist nouveau > /etc/modprobe.d/blacklist-nvidia-nouveau.conf"
   sudo bash -c "echo options nouveau modeset=0 >> /etc/modprobe.d/blacklist-nvidia-nouveau.conf"
   sudo update-initramfs -u
   ```

4. 安装必要的构建工具
   ```bash
   sudo apt-get install build-essential dkms linux-headers-$(uname -r)
   ```

### 2. 停止图形界面

**方法一**：通过tty终端
1. 按`Ctrl+Alt+F1`进入文本控制台
2. 登录您的账户
3. 停止LightDM服务
   ```bash
   sudo systemctl stop lightdm
   ```

**方法二**：通过GRUB引导
1. 重启系统
2. 在GRUB菜单按`e`编辑启动选项
3. 找到以`linux`开始的行，在行末添加`3`或`systemd.unit=multi-user.target`
4. 按`F10`或`Ctrl+X`启动

### 3. 安装驱动

1. 进入驱动所在目录
   ```bash
   cd ~/下载  # 或您保存驱动文件的目录
   ```

2. 给安装文件添加执行权限
   ```bash
   chmod +x NVIDIA-Linux-x86_64-570.144.run
   ```

3. 执行安装程序
   ```bash
   sudo ./NVIDIA-Linux-x86_64-570.144.run --dkms
   ```

4. 安装选项设置：
   - 选择"Yes"继续安装
   - 选择"Yes"注册内核模块
   - 选择"Yes"安装32位兼容库
   - 选择"Yes"更新X配置文件

### 4. 完成安装

1. 重启系统
   ```bash
   sudo reboot
   ```

2. 验证驱动安装
   ```bash
   nvidia-smi
   ```
   正确安装后应显示570.144版本的驱动信息。

## 常见问题及解决方案

### 无法进入图形界面

如果安装后无法正常进入图形界面：
1. 按`Ctrl+Alt+F1`进入文本控制台
2. 登录您的账户
3. 重新配置X服务器
   ```bash
   sudo nvidia-xconfig
   ```
4. 重启系统
   ```bash
   sudo reboot
   ```

### 回退到系统驱动

如需恢复原来的系统驱动：
1. 进入文本控制台
2. 卸载NVIDIA官方驱动
   ```bash
   sudo ./NVIDIA-Linux-x86_64-570.144.run --uninstall
   ```
3. 重新安装系统提供的驱动
   ```bash
   sudo apt install nvidia-driver-550
   ```

### 分辨率或显示问题

安装后如出现分辨率不正确或显示异常：
1. 通过系统"显示"设置调整
2. 或使用NVIDIA设置工具
   ```bash
   sudo nvidia-settings
   ```

## 注意事项

1. 安装前请确保系统已更新到最新状态
2. 确保电源稳定，避免安装过程中断电
3. 使用`--dkms`参数可以确保在内核更新后驱动能自动重新编译
4. 如遇系统更新后驱动失效，可能需要重新安装驱动
