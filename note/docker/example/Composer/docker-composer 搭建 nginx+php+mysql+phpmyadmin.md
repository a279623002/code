##### docker-composer 搭建 nginx+php

###### 目录

```
├── app
│   └── index.php
├── docker-compose.yml
├── mysql
│   ├── config
│   ├── data
│   └── log
├── nginx
│   ├── conf.d
│   │   └── default.conf
│   └── log
│       ├── access.log
│       └── error.log
└── php
    └── Dockerfile

```

###### nginx 配置文件 default.conf

```
server {
    listen  80 default_server;
    server_name  localhost;

    root  /app;
    location / {
        index index.html index.htm index.php;
    }

    location ~ \.php$ {
        fastcgi_pass   php:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

###### php Dockerfile

```
FROM php:7.4-fpm
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone #设置时区
RUN apt-get update \ 
    && rm -rf /var/lib/apt/lists/* \
    && docker-php-ext-install -j$(nproc) \
        mysqli #安装mysqli拓展

CMD ["php-fpm", "-F"]

# 下面是安装其他拓展例子
FROM php:7.4-fpm
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone
RUN apt-get update && apt-get install -y \
libwebp-dev \ #gd拓展需要的库
libpng-dev \ #gd拓展需要的库
libjpeg-dev \ #gd拓展需要的库
libfreetype6-dev \ #gd拓展需要的库
--no-install-recommends && rm -rf /var/lib/apt/lists/* \
&& docker-php-ext-install -j$(nproc) \
gettext mysqli pdo pdo_mysql standard \
&& docker-php-ext-configure gd --with-webp=/usr/include/webp --with-jpeg=/usr/include --with-freetype=/usr/include/freetype2/ \
&& docker-php-ext-install -j$(nproc) gd

CMD ["php-fpm", "-F"]
```

###### docker-compose.yml

```
version: '3'
services:
    nginx:
        image: nginx:latest
        container_name: 'shiro-nginx'
        ports:
            - '8099:80' #本地端口（映射）容器内端口
            - '8443:443'
        environment:
            - TZ=Asia/Shanghai
        # 依赖关系，先跑php
        depends_on:
            - 'php'
        volumes:
            - '~/docker/np/nginx/conf.d:/etc/nginx/conf.d' #本地目录（挂载）容器内目录
            - '~/docker/np/app:/app'
            - '~/docker/np/nginx/log:/var/log/nginx'

    php:
        build: ./php/
        container_name: 'shiro-php'
        ports:
            - '8090:9000'
        environment:
            - TZ=Asia/Shanghai
        links:
            - mysql:mysql
        volumes:
            - '~/docker/np/app:/app'

    mysql:
        hostname: mysql
        restart: always
        image: mysql:latest
        container_name: 'shiro-mysql'
        ports:
            - '9036:3306'
        volumes:
            - '~/docker/np/mysql/config:/etc/mysql'
            - '~/docker/np/mysql/log:/var/log/mysql'
            - '~/docker/np/mysql/data:/var/lib/mysql-files'
        environment:
            MYSQL_ROOT_PASSWORD: root
            MYSQL_USER: test
            MYSQL_PASSWORD: test

    phpmyadmin:
        image: phpmyadmin/phpmyadmin
        container_name: 'shiro-myadmin'
        ports:
            - '8089:80'
        links:
            - mysql:db

```

##### 测试

```
<?php
    //echo phpinfo();

	// 该地址使用宿主机地址
    $con = mysqli_connect('192.168.0.79:9036', 'root', 'root', 'demo');

    if (!$con) {

        die("could not connect to the db:\n" .  mysqli_connect_error());

    }

    $sql = 'select * from Persons';
    $res = mysqli_query($con, $sql);

    if (!$res) {
        die(mysqli_error());
    }

    while($row = mysqli_fetch_assoc($res)) {
        print_r($row);
    }

    mysqli_close($con);
    // 访问 127.0.0.1：8099
    // phpmyadmin 127.0.0.1:8089
```

