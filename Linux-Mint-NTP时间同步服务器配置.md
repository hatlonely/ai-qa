# Linux Mint 设置 NTP 时间同步服务器

本文档总结了如何在 Linux Mint 系统中配置 NTP 时间同步服务器，特别是针对中国大陆地区用户推荐的 NTP 服务器。

## Linux Mint 配置 NTP 服务的方法

### 方法一：使用 timedatectl 命令

1. 检查 NTP 服务状态：
   ```bash
   timedatectl status
   ```

2. 启用 NTP 服务：
   ```bash
   sudo timedatectl set-ntp true
   ```

3. 修改 NTP 服务器配置：
   ```bash
   sudo nano /etc/systemd/timesyncd.conf
   ```

4. 在配置文件中添加或修改：
   ```
   [Time]
   NTP=ntp1.aliyun.com ntp2.aliyun.com time1.cloud.tencent.com ntp.ntsc.ac.cn
   FallbackNTP=ntp.tuna.tsinghua.edu.cn ntp.sjtu.edu.cn
   ```

5. 重启服务：
   ```bash
   sudo systemctl restart systemd-timesyncd
   ```

### 方法二：使用 chrony

1. 安装 chrony：
   ```bash
   sudo apt install chrony
   ```

2. 编辑配置文件：
   ```bash
   sudo nano /etc/chrony/chrony.conf
   ```

3. 添加服务器：
   ```
   server ntp1.aliyun.com iburst
   server ntp2.aliyun.com iburst
   server time1.cloud.tencent.com iburst
   server ntp.ntsc.ac.cn iburst
   ```

4. 重启服务：
   ```bash
   sudo systemctl restart chronyd
   ```

5. 验证同步状态：
   ```bash
   chronyc sources
   ```

### 方法三：使用图形界面

1. 打开"系统设置"或"首选项"
2. 选择"日期和时间"
3. 启用"自动设置日期和时间"选项

## 中国大陆推荐 NTP 服务器列表

### 阿里云 NTP 服务
- ntp.aliyun.com
- ntp1.aliyun.com 到 ntp7.aliyun.com

### 腾讯云 NTP 服务
- time1.cloud.tencent.com 到 time5.cloud.tencent.com

### 国家权威 NTP 服务
- ntp.ntsc.ac.cn (国家授时中心)
- cn.ntp.org.cn
- ntp.cas.cn (中国科学院)

### 教育机构 NTP 服务
- ntp.tuna.tsinghua.edu.cn (清华大学)
- ntp.sjtu.edu.cn (上海交通大学)
- ntp.neu.edu.cn (东北大学)
- ntp.bupt.edu.cn (北京邮电大学)

### 其他服务
- ntp.baidu.com (百度)

## 选择建议
1. 优先选择地理位置接近的服务器
2. 配置多个 NTP 服务器以提高可靠性
3. 企业环境建议使用大型云服务提供商的 NTP 服务
4. 政府或研究机构可优先考虑国家授时中心的服务
