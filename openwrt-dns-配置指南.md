# OpenWRT DNS 配置指南

本文档总结了 OpenWRT 系统中 DNS 相关的配置方法，包括基本的 dnsmasq 配置以及使用 DNS over TLS (DoT) 的 Stubby 配置。

## dnsmasq 基本配置

OpenWRT 默认使用 dnsmasq 作为 DNS 服务器，主要配置文件位于 `/etc/config/dhcp`。一个典型的配置如下：

```bash
config dnsmasq
        option domainneeded '1'
        option boguspriv '1'
        option filterwin2k '0'
        option localise_queries '1'
        option rebind_protection '1'
        option rebind_localhost '1'
        option local '/lan/'
        option domain 'lan'
        option expandhosts '1'
        option nonegcache '0'
        option cachesize '1000'
        option authoritative '1'
        option readethers '1'
        option leasefile '/tmp/dhcp.leases'
        option resolvfile '/tmp/resolv.conf.d/resolv.conf.auto'
        option nonwildcard '1'
        option localservice '1'
        option ednspacket_max '1232'
        option filter_aaaa '0'
        option filter_a '0'
```

### 推荐优化配置

添加可靠的上游 DNS 服务器：

```bash
# 添加公共DNS服务器
list server '114.114.114.114'  # 国内快速DNS
list server '223.5.5.5'        # 阿里DNS
# 或国际DNS
list server '8.8.8.8'          # Google DNS
list server '1.1.1.1'          # Cloudflare DNS
```

DNS 安全增强配置：

```bash
# 启用DNSSEC提高安全性
option dnssec '1'
option dnsseccheckunsigned '1'
```

DNS 查询优化：

```bash
# 设置最小TTL以减少频繁查询
option min_ttl '3600'

# 如果需要更快的DNS解析，可以设置并发查询
option queryport '0'  # 使用随机端口
option concurrent '1' # 启用并发查询
```

## DNS over TLS 配置 (Stubby)

dnsmasq 本身不支持 DNS over TLS 协议，需要使用 Stubby 等工具实现。

### 安装 Stubby

```bash
opkg update
opkg install stubby
```

### 配置 Stubby 使用 Google DNS (8.8.8.8)

编辑 `/etc/stubby/stubby.yml` 文件：

```yaml
resolution_type: GETDNS_RESOLUTION_STUB
round_robin_upstreams: 1
appdata_dir: "/var/lib/stubby"
tls_authentication: GETDNS_AUTHENTICATION_REQUIRED
tls_query_padding_blocksize: 128
edns_client_subnet_private: 1
idle_timeout: 10000
listen_addresses:
  - 0.0.0.0@5453  # 监听所有接口
  - ::@5453       # IPv6
dns_transport_list:
  - GETDNS_TRANSPORT_TLS
upstream_recursive_servers:
  - address_data: 8.8.8.8
    tls_auth_name: "dns.google"
    tls_port: 853
  - address_data: 8.8.4.4
    tls_auth_name: "dns.google"
    tls_port: 853
  - address_data: 2001:4860:4860::8888
    tls_auth_name: "dns.google"
    tls_port: 853
  - address_data: 2001:4860:4860::8844
    tls_auth_name: "dns.google"
    tls_port: 853
```

注意：如果使用 UCI 系统管理 Stubby，请确保在 `/etc/config/stubby` 中设置：
```
option manual '1'
```

### 配置 dnsmasq 使用 Stubby

编辑 `/etc/config/dhcp` 文件，添加：

```
config dnsmasq
        ...
        option noresolv '1'       # 忽略resolv.conf文件
        list server '127.0.0.1#5453'  # 转发到Stubby
        ...
```

### 添加防火墙规则（如需从外部访问 Stubby）

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

### 启动服务

```bash
/etc/init.d/stubby enable
/etc/init.d/stubby restart
/etc/init.d/dnsmasq restart
```

### 验证配置

检查 Stubby 是否正常工作：

```bash
# 查看Stubby状态
/etc/init.d/stubby status

# 使用dig测试DNS解析
dig @127.0.0.1 -p 5453 google.com

# 检查日志
logread | grep stubby
```

## 最佳实践

1. 使用本地 DNS 缓存提高解析速度
2. 配置多个上游 DNS 服务器提高可靠性
3. 使用 DNS over TLS 保护隐私和安全
4. 根据网络环境选择适合的 DNS 服务器
5. 定期检查 DNS 配置以确保系统安全和高效运行
