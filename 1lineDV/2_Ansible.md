Задание-обучение на практику с Ansible
Задание-обучение на практику с Ansible
 
Создать 2 хоста с Linux (debian/ubuntu) в VirtualBox, для применения Ansible рецептов.
"server-one"
"server-two"
Сделать так, чтобы на хосты можно было ходить по ключам, чтобы выполнять рецепты ansible c управляющей машины (тоже на Linux).
 
 
1. Прописать в инвентори эти хосты, так чтобы они входили в группу "servers_group"
inventories/prod/hosts.yml
 
1.1. Добавить переменные на эту группу "servers_group" хостов:
inventories/prod/group_vars
var_exmpl_xx = Hello World
var_exmpl_yy = 123321
var_exmpl_zz = 000
 
Переопределить переменную для хоста "server-two" в директории  хостов: "inventories/prod/host_vars"
var_exmpl_zz = 111
 
 
2. Создать и запустить сценарий "scenario_one.yml" на хост "server-one" на создание директории "/tmp/example_dir"
С правами доступа
    owner: root
    group: root
При запуске  плейбука использовать файл инвентория inventories/prod/hosts.yml
Так же, в сценарии, указать установку пакетов: htop, tmux
 
 
3. Создать роль my_role. Выглядит как директория с файлами: files,tasks,templates, handlers
В роли будет описано :
 
a. Созание директорий (желательно с использованием "loop:"):
/opt/dir_01
/opt/dir_02
 
b. Копирование пустого файла "filename.txt" в "/opt/dir_01" и применить права 755.
 
c. Из шаблона, "roles/my_role/templates/my_file.txt.j2" с текстовым содержанием подставить переменные которые определили ранее:
This example use variables in template:
${var_exmpl_xx}
${var_exmpl_yy}
${var_exmpl_yy}
 
 
Скопировать этот файлик-шаблон в директории на сервера из группы "servers_group" в директорию /opt/dir_02
 
d. Опциональное задание. Добавить сценарий, добавляющий в крон-задание пользователя root запускать команду "date" каждую минуту.
Убедиться по логам, что это задание отрабатывает.
