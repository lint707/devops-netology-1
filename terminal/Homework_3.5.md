# Домашнее задание к занятию "3.5. Файловые системы"

1. Ознакомился с материалом о [sparse](https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%B7%D1%80%D0%B5%D0%B6%D1%91%D0%BD%D0%BD%D1%8B%D0%B9_%D1%84%D0%B0%D0%B9%D0%BB) (разряженных) файлах.

1. `Hardlink` это ссылка на тот же самый файл и имеет тот же `inode`, соотвественно права будут одни и теже, в качестве примера выполнил прверку с его созданием:
    ```
    vagrant@vagrant:~/test2$ ln new.txt new.link
    vagrant@vagrant:~/test2$ ls -ilh
    total 20K
    1179655 -rw-rw-r-- 2 vagrant vagrant 24 Mar 25 07:20 new.link
    1179655 -rw-rw-r-- 2 vagrant vagrant 24 Mar 25 07:20 new.txt
    vagrant@vagrant:~/test2$
    vagrant@vagrant:~/test2$ chmod 0555 new.txt
    vagrant@vagrant:~/test2$ ls -ilh
    total 20K
    1179655 -r-xr-xr-x 2 vagrant vagrant 24 Mar 25 07:20 new.link
    1179655 -r-xr-xr-x 2 vagrant vagrant 24 Mar 25 07:20 new.txt
    vagrant@vagrant:~/test2$
    ```
1. Выполнил `vagrant destroy` на имеющийся инстанс Ubuntu. Заменил содержимое Vagrantfile следующим:

    ```bash
    Vagrant.configure("2") do |config|
      config.vm.box = "bento/ubuntu-20.04"
      config.vm.provider :virtualbox do |vb|
        lvm_experiments_disk0_path = "/tmp/lvm_experiments_disk0.vmdk"
        lvm_experiments_disk1_path = "/tmp/lvm_experiments_disk1.vmdk"
        vb.customize ['createmedium', '--filename', lvm_experiments_disk0_path, '--size', 2560]
        vb.customize ['createmedium', '--filename', lvm_experiments_disk1_path, '--size', 2560]
        vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 1, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk0_path]
        vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 2, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk1_path]
      end
    end
    ```

    Создал новую виртуальную машину с двумя дополнительными неразмеченными дисками по 2.5 Гб.
    ```    
    vagrant@vagrant:~$ lsblk
    NAME                      MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
    loop0                       7:0    0 70.3M  1 loop /snap/lxd/21029
    loop1                       7:1    0 55.4M  1 loop /snap/core18/2128
    loop2                       7:2    0 32.3M  1 loop /snap/snapd/12704
    loop3                       7:3    0 43.6M  1 loop /snap/snapd/15177
    loop4                       7:4    0 55.5M  1 loop /snap/core18/2344
    sda                         8:0    0   64G  0 disk
    ├─sda1                      8:1    0    1M  0 part
    ├─sda2                      8:2    0    1G  0 part /boot
    └─sda3                      8:3    0   63G  0 part
      └─ubuntu--vg-ubuntu--lv 253:0    0 31.5G  0 lvm  /    
    sdb                         8:16   0  2.5G  0 disk
    sdc                         8:32   0  2.5G  0 disk
    ```
    
1. Используя `fdisk`, разбил диск `sdb` на 2 раздела: 2 Гб и оставшееся пространство.
    ```
    vagrant@vagrant:~$ sudo fdisk /dev/sdb
    Welcome to fdisk (util-linux 2.34).
    Changes will remain in memory only, until you decide to write them.
    Be careful before using the write command.
    Device does not contain a recognized partition table.
    Created a new DOS disklabel with disk identifier 0xf6da8a78.
    
    Command (m for help): n
    Partition type
       p   primary (0 primary, 0 extended, 4 free)
       e   extended (container for logical partitions)
    Select (default p): p   
    Partition number (1-4, default 1): 1
    First sector (2048-5242879, default 2048): 2048
    Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-5242879, default 5242879): +2G
    Created a new partition 1 of type 'Linux' and of size 2 GiB.
    
    Command (m for help): n
    Partition type
       p   primary (1 primary, 0 extended, 3 free)
       e   extended (container for logical partitions)
    Select (default p): p
    Partition number (2-4, default 2): 2
    First sector (4196352-5242879, default 4196352): 4196352
    Last sector, +/-sectors or +/-size{K,M,G,T,P} (4196352-5242879, default 5242879):
    Created a new partition 2 of type 'Linux' and of size 511 MiB.
    
    Command (m for help): w
    The partition table has been altered.
    Calling ioctl() to re-read partition table.
    Syncing disks.
    ```
    Результат выполения:
    ```    
    vagrant@vagrant:~$ lsblk
    NAME                      MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
    sdb                         8:16   0  2.5G  0 disk
    ├─sdb1                      8:17   0    2G  0 part
    └─sdb2                      8:18   0  511M  0 part
    ```
1. Используя `sfdisk`, перенесите данную таблицу разделов на второй диск.

1. Соберите `mdadm` RAID1 на паре разделов 2 Гб.

1. Соберите `mdadm` RAID0 на второй паре маленьких разделов.

1. Создайте 2 независимых PV на получившихся md-устройствах.

1. Создайте общую volume-group на этих двух PV.

1. Создайте LV размером 100 Мб, указав его расположение на PV с RAID0.

1. Создайте `mkfs.ext4` ФС на получившемся LV.

1. Смонтируйте этот раздел в любую директорию, например, `/tmp/new`.

1. Поместите туда тестовый файл, например `wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`.

1. Прикрепите вывод `lsblk`.

1. Протестируйте целостность файла:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

1. Используя pvmove, переместите содержимое PV с RAID0 на RAID1.

1. Сделайте `--fail` на устройство в вашем RAID1 md.

1. Подтвердите выводом `dmesg`, что RAID1 работает в деградированном состоянии.

1. Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

1. Погасите тестовый хост, `vagrant destroy`.

 
 ---

## Как сдавать задания
