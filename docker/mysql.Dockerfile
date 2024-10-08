FROM --platform=linux/amd64 mysql:8.0.34-debian

COPY docker/my.cnf /etc/mysql/conf.d/my.cnf

RUN mv /etc/apt/sources.list.d/mysql.list /etc/apt/sources.list.d/mysql.list.disabled

RUN apt-get update && apt-get install -y gnupg curl

RUN curl -fsSL https://repo.mysql.com/RPM-GPG-KEY-mysql | gpg --dearmor -o /usr/share/keyrings/mysql.gpg

RUN echo "deb [signed-by=/usr/share/keyrings/mysql.gpg] http://repo.mysql.com/apt/debian bullseye mysql-8.0" > /etc/apt/sources.list.d/mysql.list

RUN gpg --no-default-keyring --keyring /usr/share/keyrings/mysql.gpg --list-keys

ENV LC_ALL=ja_JP.UTF-8
ENV TZ=Asia/Tokyo
ENV LANG=ja_JP.UTF-8

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

CMD ["mysqld"]
EXPOSE 3306