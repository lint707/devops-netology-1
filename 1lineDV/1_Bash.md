Задание-навыки использования Bash
 
1. Используя цикл for вывести дату в формате "2021.04.19_20:12:17" 3 раза с шагом в 2 секунды.  Однострочник.
```
#!/bin/bash
for ((i=1; i<=3; i++))
do
echo `date +'20%y.%m.%d_%R:%S'`
sleep 2
done
```
ответ:
```
grigorijkazacekno@MacBook-Pro-Grigorij task % ./bash0  
2022.05.12_12:22:18
2022.05.12_12:22:20
2022.05.12_12:22:22
grigorijkazacekno@MacBook-Pro-Grigorij task % 
```
1 Вывод списка файлов существующей директории "/home/" сохранить в файл "/tmp/ls_exist_dir.txt"
Одновременно отображая вывод на экране, подсказка: tee.
```
grigorijkazacekno@MacBook-Pro-Grigorij ~ % ls -lha | tee /tmp/ls_exist_dir.txt
```
проверка:
```
grigorijkazacekno@MacBook-Pro-Grigorij ~ % cat /tmp/ls_exist_dir.txt
```

1 Попытку вывода списка файлов НЕ существующей директории "ls -al /noexist_dir" сохранить в файл "/tmp/ls_none_exist_dir.txt"
Подсказка, используя перенаправление вывода.
```
grigorijkazacekno@MacBook-Pro-Grigorij ~ % ls -al /noexist_dir >/tmp/ls_none_exist_dir.txt 2>&1
```
Проверка:
```
grigorijkazacekno@MacBook-Pro-Grigorij ~ % vim /tmp/ls_none_exist_dir.txt

ls: /noexist_dir: No such file or directory                                                                                        
~                                                                                                
"/tmp/ls_none_exist_dir.txt" 1L, 44B
``` 
1 Сгенерировать пары ключей, и научиться ходить с клиента на сервер по ssh без использования пароля.
Машинки можно поднять в VirtualBox.
Хостнейм для последующего задания можно указать "client, server-one"

Установил на Ubuntu ssh сервер и сгенерировал новый приватный ключ: </br>
![ssh](img/ssh1.jpg)  </br>

Скопировал свой публичный ключ на другой сервер и подключился к серверу по SSH-ключу: </br>
![ssh](img/ssh2.jpg) </br>

