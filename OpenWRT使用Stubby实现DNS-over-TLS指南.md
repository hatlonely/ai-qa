# OpenWRT 使用 Stubby 实现 DNS over TLS 指南

本文档主要介绍在 OpenWRT 系统上使用 Stubby 实现 DNS over TLS (DoT) 的完整配置方法，提高 DNS 查询的安全性和隐私保护。

## Stubby 简介

Stubby 是一个开源的 DNS Privacy 守护进程，专为支持 DNS over TLS 而设计。在 OpenWRT 上使用 Stubby 可以加密 DNS 查询流量，防止 ISP 或网络中的其他实体监控或修改您的 DNS 查询。

## 安装 Stubby

在 OpenWRT 上安装 Stubby：

```bash
opkg update
opkg install stubby
```

## 配置 Stubby

Stubby 的主配置文件位于 `/etc/stubby/stubby.yml`，以下是使用 Google DNS (8.8.8.8) 的配置示例：

```yaml
# 确保在 /etc/config/stubby 中设置 option manual '1' 以使用此配置文件
resolution_type: GETDNS_RESOLUTION_STUB
round_robin_upstreams: 1
appdata_dir: "/var/lib/stopbby"
tls_authentication: GETDNS_AUTHENTICATION_REQUIRED
tls_query_padding_blocksize: 128
edns_client_subnet_private: 1
idle_timeout: 10000

# 监听地址设置
listen_addresses:
  - 0.0.0.0@5453    # 监听所有接口，允许局域网设备查询
  - ::@5453         # IPv6 支持

# 指定使用 TLS 传输
dns_transport_list:
  - GETDNS_TRANSPORT_TLS

# 上游 DNS 服务器配置（Google DNS）
upstream_recursive_servers:
  # IPv4 地址
  - address_data: 8.8.8.8
    tls_auth_name: "dns.google"
    tls_port: 853
  - address_data: 8.8.4.4
    tls_auth_name: "dns.google"
    tls_port: 853
    
  # IPv6 地址
  - address_data: 2001:4860:4860::8888
    tls_auth_name: "dns.google"
    tls_port: 853
  - address_data: 2001:4860:4860::8844
    tls_auth_name: "dns.google"
    tls_port: 853
```

### 配置说明

1. **监听设置**:
   - `0.0.0.0@5453`: 在所有接口上监听，端口为 5453
   - 如果只需在本地使用，可设为 `127.0.0.1@5453`

2. **TLS 认证**:
   - `tls_authentication: GETDNS_AUTHENTICATION_REQUIRED`: 要求 TLS 认证
   - `tls_auth_name`: 指定 DNS 服务器的 TLS 主机名

3. **上游服务器**:
   - Google DNS 的 TLS 主机名为 `dns.google`
   - 标准 DoT 端口为 853

## 启用 UCI 配置

如果您希望通过 OpenWRT 的 UCI 系统管理 Stubby 配置，请编辑 `/etc/config/stubby`：

```
config stubby 'global'
        option manual '1'
        # 其他设置...
```

设置 `option manual '1'` 将使 Stubby 使用 `/etc/stubby/stubby.yml` 文件的配置。

## 整合 dnsmasq 与 Stubby

要使 OpenWRT 的 DNS 系统使用 Stubby，需要修改 dnsmasq 配置，编辑 `/etc/config/dhcp`：

```
config dnsmasq
        # 现有设置保持不变...
        option noresolv '1'           # 忽略 resolv.conf
        list server '127.0.0.1#5453'  # 使用 Stubby 作为上游 DNS
```

## 添加防火墙规则

如果您需要从局域网其他设备直接查询 Stubby（而不是通过 dnsmasq），请添加防火墙规则：

```bash
uci add firewall rule
uci set firewall.@rule[-1].name='Allow-Stubby'
uci set firewall.@rule[-1].src='lan'
uci set firewall.@rule[-1].proto='tcp udp'
uci set firewall.@rule[-1].dest_port='5453'
uci set firewall.@rule[-1].target='ACCEPT'
uci commit firewall
/etc/init.d/firewall restart
```

## 启动服务

配置完成后，启用并重启服务：

```bash
# 启用 Stubby 开机自启动
/etc/init.d/stubby enable

# 重启 Stubby
/etc/init.d/stubby restart

# 重启 dnsmasq
/etc/init.d/dnsmasq restart
```

## 验证配置

通过以下方法验证 Stubby 是否正常工作：

```bash
# 检查 Stubby 服务状态
/etc/init.d/stubby status

# 在路由器上测试本地 DNS 解析
dig @127.0.0.1 -p 5453 google.com

# 从局域网设备测试（假设路由器 IP 为 192.168.0.6）
dig @192.168.0.6 -p 5453 google.com

# 查看日志
logread | grep stubby
```

## 常见问题排查

1. **连接被拒绝**：
   - 检查 Stubby 监听地址是否为 `0.0.0.0@5453`（而不是仅 127.0.0.1）
   - 确认防火墙规则已正确添加
   - 验证 Stubby 服务是否正在运行

2. **DNS 解析失败**：
   - 检查 Stubby 配置中的上游服务器信息是否正确
   - 确保路由器能够访问外网和 DNS 服务器

3. **TLS 握手错误**：
   - 检查系统时间是否正确（TLS 证书验证依赖于正确的时间）
   - 确认 `tls_auth_name` 配置正确

## 最佳实践

1. **部署建议**：使用 dnsmasq 转发到 Stubby，而不是直接从客户端连接 Stubby
2. **服务器选择**：除了 Google DNS，还可以考虑其他支持 DoT 的服务商，如 Cloudflare (1.1.1.1)
3. **定期更新**：保持 Stubby 和 OpenWRT 系统更新，以获取最新的安全修复
4. **监控**：定期检查日志，确保 DNS 解析正常工作
