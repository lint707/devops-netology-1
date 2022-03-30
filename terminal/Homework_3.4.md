# Домашнее задание к занятию "3.4. Операционные системы, лекция 2"

1. Установка node_exporter:
    ```
   vagrant@vagrant:~$ vagrant@vagrant:~$ wget https://github.com/prometheus/node_exporter/releases/download/v1.3.1/node_exporter-1.3.1.linux-386.tar.gz
   vagrant@vagrant:~$ tar zxvf node_exporter-1.3.1.linux-386.tar.gz
   node_exporter-1.3.1.linux-386/
   node_exporter-1.3.1.linux-386/LICENSE
   node_exporter-1.3.1.linux-386/NOTICE
   node_exporter-1.3.1.linux-386/node_exporter
   vagrant@vagrant:~$ ls -l
   total 11296
   drwxr-xr-x 2 vagrant vagrant    4096 Dec  5 11:15 node_exporter-1.3.1.linux-386
   -rw-rw-r-- 1 vagrant vagrant 8624587 Dec  8 08:52 node_exporter-1.3.1.linux-386.tar.gz
    ```
   Создание /node_exporter.service:
    ```
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ sudo cp node_exporter /usr/local/bin/
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ sudo useradd --no-create-home --shell /bin/false nodeusr
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ sudo chown -R nodeusr:nodeusr /usr/local/bin/node_exporter
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ sudo vim /etc/systemd/system/node_exporter.service
   [Unit]
   Description=Node Exporter metrics
   After=multi-user.target
   [Service]
   User=nodeusr
   EnvironmentFile=-/etc/default/node_exporter
   ExecStart=/usr/local/bin/node_exporter $EXTRA_OPTS
   KillMode=process
   [Install]
   WantedBy=multi-user.target
    ```   
    Выставляем права (Все пользователи имеют право чтения; владелец может редактировать):
    ``` 
    vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ sudo chmod 644 /etc/systemd/system/node_exporter.service
    ```        
   Перезапуск deamon, Настройка автозапуска, старт node_exporter
    ```
    vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ systemctl daemon-reload
   ==== AUTHENTICATION COMPLETE ===
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ systemctl enable node_exporter
   ==== AUTHENTICATION COMPLETE ===
   vagrant@vagrant:~/node_exporter-1.3.1.linux-386$ systemctl start node_exporter
   ==== AUTHENTICATION COMPLETE ===
    ```
    Проверка запуска node_exporter, до и после перезагрузки, автозагрузка работает:
    ```    
    vagrant@vagrant:~$ systemctl status node_exporter
   ● node_exporter.service - Node Exporter metrics
     Loaded: loaded (/etc/systemd/system/node_exporter.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2022-03-30 11:26:22 UTC; 2min 50s ago
   Main PID: 1854 (node_exporter)
      Tasks: 5 (limit: 2278)
     Memory: 2.1M
     CGroup: /system.slice/node_exporter.service
             └─1854 /usr/local/bin/node_exporter
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:115 level=info collect>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=node_exporter.go:199 level=info msg="Li>
   Mar 30 11:26:22 vagrant node_exporter[1854]: ts=2022-03-30T11:26:22.907Z caller=tls_config.go:195 level=info msg="TLS i>   
   lines 1-19/19 (END)
    ```    
        
1. Опции node_exporter с выводом `/metrics` через  `curl http://localhost:9100/metrics`. Примеры опций для базового мониторинга хоста по CPU, памяти, диску и сети.
    ```    
   | grep node_cpu
   node_cpu_guest_seconds_total
   node_cpu_seconds_total
   
   | grep node_memory
   node_memory_Active_bytes
   node_memory_Inactive_bytes
   node_memory_MemFree_bytes
   node_memory_MemTotal_bytes
   node_memory_SwapCached_bytes  
   node_memory_SwapFree_bytes 
   node_memory_SwapTotal_bytes
   node_memory_MemAvailable_bytes
   
   | grep node_disk
   node_disk_io_time_seconds_total
   node_disk_read_bytes_total
   node_disk_read_time_seconds_total
   node_disk_reads_completed_total
   node_disk_write_time_seconds_total
   node_disk_writes_completed_total
   node_disk_written_bytes_total
   
   | grep node_network
   node_network_address_assign_type
   node_network_device_id
   node_network_flags
   node_network_iface_id
   node_network_iface_link
   node_network_mtu_bytes
   node_network_receive_bytes_total
   node_network_speed_bytes
   node_network_transmit_bytes_total
   node_network_up
    ```    
    
1. Установите в свою виртуальную машину [Netdata](https://github.com/netdata/netdata). Воспользуйтесь [готовыми пакетами](https://packagecloud.io/netdata/netdata/install) для установки (`sudo apt install -y netdata`). После успешной установки:
    * в конфигурационном файле `/etc/netdata/netdata.conf` в секции [web] замените значение с localhost на `bind to = 0.0.0.0`,
    * добавьте в Vagrantfile проброс порта Netdata на свой локальный компьютер и сделайте `vagrant reload`:

    ```bash
    config.vm.network "forwarded_port", guest: 19999, host: 19999
    ```

    После успешной перезагрузки в браузере *на своем ПК* (не в виртуальной машине) вы должны суметь зайти на `localhost:19999`. Ознакомьтесь с метриками, которые по умолчанию собираются Netdata и с комментариями, которые даны к этим метрикам.

1. Можно ли по выводу `dmesg` понять, осознает ли ОС, что загружена не на настоящем оборудовании, а на системе виртуализации?
1. Как настроен sysctl `fs.nr_open` на системе по-умолчанию? Узнайте, что означает этот параметр. Какой другой существующий лимит не позволит достичь такого числа (`ulimit --help`)?
1. Запустите любой долгоживущий процесс (не `ls`, который отработает мгновенно, а, например, `sleep 1h`) в отдельном неймспейсе процессов; покажите, что ваш процесс работает под PID 1 через `nsenter`. Для простоты работайте в данном задании под root (`sudo -i`). Под обычным пользователем требуются дополнительные опции (`--map-root-user`) и т.д.
1. Найдите информацию о том, что такое `:(){ :|:& };:`. Запустите эту команду в своей виртуальной машине Vagrant с Ubuntu 20.04 (**это важно, поведение в других ОС не проверялось**). Некоторое время все будет "плохо", после чего (минуты) – ОС должна стабилизироваться. Вызов `dmesg` расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?

 
 ---
