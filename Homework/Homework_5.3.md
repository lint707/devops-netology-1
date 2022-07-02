# Домашнее задание к занятию "5.3. Введение. Экосистема. Архитектура. Жизненный цикл Docker контейнера"

---

## Задача 1

Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

```
vagrant@server1:~$ docker pull ubuntu/nginx:1.18-22.04_edge
1.18-22.04_edge: Pulling from ubuntu/nginx
afad979d238f: Pull complete 
7efc54ee67e0: Pull complete 
9185eb702abc: Pull complete 
70bc7ac6a482: Pull complete 
Digest: sha256:c798b60006732b86a43bca84d9e5059b465ecdcdb305dd3452ed5613b29e30ac
Status: Downloaded newer image for ubuntu/nginx:1.18-22.04_edge
docker.io/ubuntu/nginx:1.18-22.04_edge

vagrant@server1:~$ docker images
REPOSITORY     TAG               IMAGE ID       CREATED      SIZE
ubuntu/nginx   1.18-22.04_edge   c1d6caca3772   9 days ago   144MB

vagrant@server1:~/docker$ vim dockerfile

vagrant@server1:~/docker$ docker build -f dockerfile -t lint707/nginx_netology5.3  .
Sending build context to Docker daemon  2.048kB
Step 1/2 : FROM ubuntu/nginx:1.18-22.04_edge
 ---> c1d6caca3772
Step 2/2 : RUN echo '<html><head>Hey, Netology</head><body><h1>I am DevOps Engineer!</h1></body></html>' > /usr/share/nginx/html/index.html
 ---> Running in 49d94709c9e8
Removing intermediate container 49d94709c9e8
 ---> a61bb2bbbf69
Successfully built a61bb2bbbf69
Successfully tagged lint707/nginx_netology5.3:latest

vagrant@server1:~/docker$ docker images
REPOSITORY                  TAG               IMAGE ID       CREATED              SIZE
lint707/nginx_netology5.3   latest            a61bb2bbbf69   About a minute ago   144MB
ubuntu/nginx                1.18-22.04_edge   c1d6caca3772   9 days ago           144MB

vagrant@server1:~/docker$ docker login -u lint707
Password: 
Login Succeeded

vagrant@server1:~/docker$ docker push lint707/nginx_netology5.3
Using default tag: latest
The push refers to repository [docker.io/lint707/nginx_netology5.3]
760a4341f34a: Pushed 
e29410d289ba: Mounted from ubuntu/nginx 
9126dae3b3a7: Mounted from ubuntu/nginx 
39767fb85b82: Mounted from ubuntu/nginx 
a790f937a6ae: Mounted from ubuntu/nginx 
latest: digest: sha256:cee73146fb887ef3c7937a9c7a9c483e965b335b3b64acf047cb228ceb76e31a size: 1362
```
https://hub.docker.com/r/lint707/nginx_netology5.3
```
vagrant@server1:~/docker$ curl localhost:8080
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```
При этом в контейнере корректно запустилось: 
```
root@17ff1d0d7a4e:/usr/share/nginx/html# cat index.html 
<html><head>Hey, Netology</head><body><h1>I am DevOps Engineer!</h1></body></html>
root@17ff1d0d7a4e:/usr/share/nginx/html# 
```
Нужно понять где ошибка.

## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Ответ:
- Высоконагруженное монолитное java веб-приложение:
> Физичиские или виртуальные машины, если в приложениине заложено масштабированиетогла дучше физическая машина,
в случае мастрабируется и может взаимодействовать с балансировщиком, то виртуальная машина.
- Nodejs веб-приложение:
> Docker, удобство масштабирования, легковесность, простота развётрывания.
- Мобильное приложение c версиями для Android и iOS:
> Виртуальные машины, упрощает тестирование и размещение.
- Шина данных на базе Apache Kafka:
> Docker, возможность быстрого отката, в случае проблем на проде, изолированность приложений.
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana:
> 
> ??? дописать ответ
> 
- Мониторинг-стек на базе Prometheus и Grafana:
> Docker, удобство масштабирования, легковесность, простота развётрывания, есть готовые образы.
- MongoDB, как основное хранилище данных для java-приложения:
> Виртуальная машина, не подходящее решение хранить БД в контейнере. 
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry:
> Виртуальная машина, удобство бекапов и миграции.

## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.

## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе с остальными ответами к задачам.


---


Сценарий:



