# Linux Mint 运行炉石传说指南

炉石传说是暴雪的热门卡牌游戏，在 Linux Mint 上运行炉石传说有以下几种方法：

## 方法一：使用 Lutris（推荐）

Lutris 是专为 Linux 设计的游戏管理平台，对暴雪游戏有很好的支持。

### 安装步骤

1. **安装 Lutris**
   ```bash
   sudo apt update
   sudo apt install lutris
   ```
   或使用 Flatpak 安装：
   ```bash
   flatpak install flathub net.lutris.Lutris
   ```

2. **安装 Wine 依赖**
   ```bash
   sudo apt install wine64 winetricks
   ```

3. **安装显卡驱动（如果尚未安装）**
   
   对于 NVIDIA 显卡：
   ```bash
   sudo apt install nvidia-driver-xxx
   ```
   对于 AMD 显卡：
   ```bash
   sudo apt install mesa-vulkan-drivers
   ```

4. **通过 Lutris 安装炉石传说**
   - 打开 Lutris
   - 在搜索栏中搜索「炉石传说」或「Hearthstone」
   - 点击游戏条目，然后点击「安装」
   - 按照安装向导的指示完成安装
   - 安装过程中会自动下载并安装 Battle.net 客户端
   - 登录您的暴雪账号并安装炉石传说

## 方法二：使用 Bottles

Bottles 是一个新一代的 Wine 容器管理工具，界面友好，功能强大。

### 安装步骤

1. **安装 Bottles**
   ```bash
   flatpak install flathub com.usebottles.bottles
   ```

2. **创建新的 Windows 瓶子**
   - 打开 Bottles
   - 点击「新建瓶子」
   - 选择「Gaming」环境
   - 命名瓶子（例如：「Battle.net」）
   - 等待创建完成

3. **安装 Battle.net 客户端**
   - 在新创建的瓶子中，点击「运行可执行文件」
   - 下载并选择 Battle.net 安装程序
   - 完成 Battle.net 安装
   - 登录并安装炉石传说

## 方法三：使用 PlayOnLinux

PlayOnLinux 是一个经典的 Wine 前端，适合有一定 Linux 经验的用户。

### 安装步骤

1. **安装 PlayOnLinux**
   ```bash
   sudo apt install playonlinux
   ```

2. **安装 Battle.net**
   - 打开 PlayOnLinux
   - 点击「安装」
   - 搜索并选择「Battle.net」
   - 按照安装向导完成安装
   - 安装炉石传说

## 性能优化提示

1. **启用 DXVK**（在 Lutris 或 Bottles 中）
   - DXVK 可以将 DirectX 调用转换为 Vulkan，提高游戏性能
   - 在 Lutris 中：右键点击游戏 → 配置 → 运行器选项 → 启用 DXVK
   - 在 Bottles 中：进入瓶子设置 → 启用 DXVK

2. **降低图形设置**
   - 在游戏内降低图形质量可以提高流畅度
   - 建议设置为中等或低等图形质量

3. **关闭背景程序**
   - 确保没有其他占用资源的程序在后台运行
   - 使用系统监视器关闭不必要的进程

4. **使用 GameMode**
   - GameMode 可以自动优化系统性能
   ```bash
   sudo apt install gamemode
   ```
   - 启动游戏时使用：
   ```bash
   gamemoderun lutris
   ```

5. **优化 Wine 设置**
   - 在 Lutris 中：右键点击游戏 → 配置 → 运行器选项
   - 尝试不同的 Wine 版本
   - 建议使用 Wine-GE 或 Proton-GE 版本

## 常见问题解决

1. **游戏崩溃或无法启动**
   - 尝试更新 Wine 版本
   - 安装缺少的 Windows 依赖（通过 winetricks）
   ```bash
   winetricks corefonts vcrun2015 d3dx9
   ```
   - 检查显卡驱动是否是最新版本

2. **登录问题**
   - 如果无法登录 Battle.net，尝试在 Wine 配置中调整 IE 版本
   ```bash
   winetricks ie8
   ```
   - 确保系统时间同步正确

3. **黑屏问题**
   - 尝试切换全屏/窗口模式（Alt+Enter）
   - 在 Lutris 或 Bottles 中启用虚拟桌面
   - 更新显卡驱动

4. **游戏卡顿**
   - 启用 Fsync/Esync（在 Lutris 或 Bottles 设置中）
   - 关闭游戏内的垂直同步
   - 限制帧率为 60fps

## 其他提示

1. **定期更新**
   - 保持 Lutris/Bottles 和 Wine 版本更新
   - 关注 ProtonDB 或 Lutris 论坛上的最新兼容性信息

2. **备份配置**
   - 备份 Battle.net 和炉石传说的配置文件，以防重装
   - Lutris 配置通常位于 `~/.local/share/lutris`
   - Bottles 配置通常位于 `~/.var/app/com.usebottles.bottles`

3. **使用 Battle.net 的替代启动方式**
   - 有时直接启动炉石传说可能比通过 Battle.net 启动更稳定
   - 在 Lutris 中可以设置直接启动游戏的选项
