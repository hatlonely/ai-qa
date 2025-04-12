# Linux Mint 安装和运行 Steam 游戏指南

## 安装 Steam 客户端

在 Linux Mint 上安装 Steam 客户端推荐使用 Flatpak 方式：

1. **安装 Flatpak 支持（如果尚未安装）**
   ```bash
   sudo apt update
   sudo apt install flatpak
   sudo apt install gnome-software-plugin-flatpak
   flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
   ```
   安装后重启系统以确保 Flatpak 正常工作。

2. **通过 Flatpak 安装 Steam**
   ```bash
   flatpak install flathub com.valvesoftware.Steam
   ```
   或者在软件管理器中搜索 Steam，选择 Flatpak 版本进行安装。

3. **启动 Flatpak 版 Steam**
   ```bash
   flatpak run com.valvesoftware.Steam
   ```
   或从应用程序菜单启动。

使用 Flatpak 版 Steam 的优势：
- 自带最新的运行时依赖
- 与系统库隔离，减少兼容性问题
- 自动更新，无需手动维护

## 配置 Steam Play (Proton)

Steam Play 是 Valve 开发的一项技术，基于 Wine 和其他技术，允许在 Linux 上运行 Windows 游戏。

1. **启用 Steam Play**
   - 打开 Steam
   - 点击左上角的 "Steam" → "设置"
   - 在设置菜单中选择 "Steam Play"
   - 勾选 "为所有其他产品启用 Steam Play"
   - 从下拉菜单中选择 Proton 版本（建议使用最新的稳定版本）
   - 点击 "确定" 保存设置
   - Steam 将重启并应用新设置

2. **安装 Windows 游戏**
   - 启用 Steam Play 后，您的游戏库将显示 Windows 游戏
   - 像安装 Linux 原生游戏一样，点击 "安装" 按钮即可
   - Steam 会自动使用 Proton 处理 Windows 游戏的安装和运行

## 提高游戏兼容性

1. **使用 Proton-GE**
   - Proton-GE 是社区维护的 Proton 增强版，提供更好的兼容性
   - 安装 ProtonUp-Qt 工具：
     ```bash
     sudo apt install python3-pip
     pip3 install protonup-qt
     ```
   - 运行 ProtonUp-Qt 并安装 Proton-GE 版本
   - 在 Steam 的 Steam Play 设置中选择 Proton-GE 版本

2. **检查游戏兼容性**
   - 访问 [ProtonDB](https://www.protondb.com/) 网站
   - 搜索特定游戏了解其兼容性状态
   - 查看社区提供的优化提示和启动选项

3. **为特定游戏设置 Proton 版本**
   - 在 Steam 库中右键点击游戏
   - 选择 "属性"
   - 在 "兼容性" 标签中勾选 "强制使用特定 Steam Play 兼容性工具"
   - 选择一个 Proton 版本（可以尝试不同版本以解决特定问题）

## 性能优化

1. **启用 GameMode**
   ```bash
   sudo apt install gamemode
   ```
   然后在游戏启动选项中添加：
   ```
   gamemoderun %command%
   ```

2. **安装显卡驱动**
   - NVIDIA 显卡：
     ```bash
     sudo apt install nvidia-driver-xxx
     ```
   - AMD 显卡：
     ```bash
     sudo apt install mesa-vulkan-drivers
     ```

## 常见问题解决

1. **游戏无法启动**
   - 尝试不同的 Proton 版本
   - 查阅 ProtonDB 网站上的社区解决方案
   - 在游戏启动选项中添加 `PROTON_LOG=1 %command%` 来生成日志文件进行调试

2. **多人游戏反作弊问题**
   - 某些带有反作弊系统的游戏可能无法在 Proton 下正常运行
   - 检查 Valve 的 [Steam Deck 验证列表](https://www.steamdeck.com/en/verified)，支持 Steam Deck 的游戏通常在 Linux 上也能正常运行

3. **音频问题**
   - 如果出现音频问题，在游戏启动选项中添加：
     ```
     PULSE_LATENCY_MSEC=60 %command%
     ```

通过以上配置，大多数 Steam 上的 Windows 游戏都可以在 Linux Mint 上流畅运行，为您提供与 Windows 系统类似的游戏体验。

## 运行非 Steam 游戏

对于非 Steam 平台的游戏（如暴雪游戏、Epic 游戏商城游戏等），您可以参考《[Linux-Mint-运行炉石传说指南.md](/home/hatlonely/hatlonely/github.com/hatlonely/ai-qa/Linux-Mint-运行炉石传说指南.md)》了解更多信息。
