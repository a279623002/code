#### 源代码安装
1. 下载 Nginx，下载地址：https://nginx.org/en/download.html
   ```
   # cd /usr/local/src/
   # wget http://nginx.org/download/nginx-1.6.2.tar.gz
   ```
2. 安装包
   ```
   # tar zxvf nginx-1.6.2.tar.gz
   ```
3. 编译安装
   ```
   # cd nginx-1.6.2
   # ./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
   # make
   # make install
   ```
4. 查看nginx版本
   ```
   # /usr/local/webserver/nginx/sbin/nginx -v
   ```
5. ln
   ```
   sudo ln -s /usr/local/webserver/nginx/sbin/nginx /usr/bin/
   ```