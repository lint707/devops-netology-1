# Домашнее задание к занятию "6.3. MySQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.  
![mysql](img/mysql_run.jpg)  

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-03-mysql/test_data) и восстановитесь из него.

Перейдите в управляющую консоль `mysql` внутри контейнера.  
```
mysql -uroot -p
```
Используя команду `\h` получите список управляющих команд.  
Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.  
![mysql](img/mysql_s.jpg) 

Подключитесь к восстановленной БД и получите список таблиц из этой БД.  
```
USE test_db;
```
**Приведите в ответе** количество записей с `price` > 300.  
![mysql](img/mysql_select.jpg)  

В следующих заданиях мы будем продолжать работу с данным контейнером.

## Задача 2
Создайте пользователя test в БД c паролем test-pass, используя:  
- плагин авторизации mysql_native_password  
- срок истечения пароля - 180 дней  
- количество попыток авторизации - 3   
- максимальное количество запросов в час - 100  
- аттрибуты пользователя:  
    - Фамилия "Pretty"  
    - Имя "James"  
![mysql](img/mysql_user.jpg)  

Предоставьте привелегии пользователю `test` на операции SELECT базы `test_db`.  
![mysql](img/mysql_grant.jpg)  

Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и  **приведите в ответе к задаче**.  
![mysql](img/mysql_slc.jpg)  

## Задача 3
Установите профилирование `SET profiling = 1`.  
Изучите вывод профилирования команд `SHOW PROFILES;`.  
Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.  
![mysql](img/mysql_inno.jpg)  

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:  
- на `MyISAM`  
- на `InnoDB`  
![mysql](img/mysql_show.jpg)  

## Задача 4 

Изучите файл `my.cnf` в директории /etc/mysql.
```
bash-4.4# cat my.cnf
# For advice on how to change settings please see
# http://dev.mysql.com/doc/refman/8.0/en/server-configuration-defaults.html

[mysqld]
#
# Remove leading # and set to the amount of RAM for the most important data
# cache in MySQL. Start at 70% of total RAM for dedicated server, else 10%.
# innodb_buffer_pool_size = 128M
#
# Remove leading # to turn on a very important data integrity option: logging
# changes to the binary log between backups.
# log_bin
#
# Remove leading # to set options mainly useful for reporting servers.
# The server defaults are faster for transactions and fast SELECTs.
# Adjust sizes as needed, experiment to find the optimal values.
# join_buffer_size = 128M
# sort_buffer_size = 2M
# read_rnd_buffer_size = 2M

# Remove leading # to revert to previous value for default_authentication_plugin,
# this will increase compatibility with older clients. For background, see:
# https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_default_authentication_plugin
# default-authentication-plugin=mysql_native_password
skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql

pid-file=/var/run/mysqld/mysqld.pid
[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/
```

Измените его согласно ТЗ (движок InnoDB):
- Скорость IO важнее сохранности данных
- Нужна компрессия таблиц для экономии места на диске
- Размер буффера с незакомиченными транзакциями 1 Мб
- Буффер кеширования 30% от ОЗУ
- Размер файла логов операций 100 Мб

Приведите в ответе измененный файл `my.cnf`.
```
[mysqld]
skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql

pid-file=/var/run/mysqld/mysqld.pid
[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/
??innodb_flush_log_at_trx_commit = 0 / innodb_flush_log_at_trx_commit = 2
??innodb_file_format=Barracuda
innodb_log_buffer_size= 1M
??max_binlog_size= 100M / innodb_log_file_size = 100M
??innodb_file_per_table = ON
??innodb_buffer_pool_size = 1G
```
---
