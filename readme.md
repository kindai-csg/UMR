![](doc/image/umr.png)
# User ManageR
近畿大学電子計算機研究会部員管理アプリケーションです.
アカウントデータはOpenLDAPに保存されます.
UMRは2018年度会長の呼称です.

# 構成
![](doc/image/arichitecture.png)

# セットアップ
1. server/config.toml編集
```
cp server/config.toml.sample server/config.toml
```
2. openldap/slapd.conf編集
```
cp openldap/slapd.conf.sample openldap/slapd.conf
```
3. install.sh
```
./install.sh
```