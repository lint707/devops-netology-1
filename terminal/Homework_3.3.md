# Домашнее задание к занятию "3.3. Операционные системы, лекция 1"

1. Выполнил команду `strace /bin/bash -c 'cd /tmp'`, системный вызов для `cd`: 
    ```
    chdir("/tmp")                           = 0
    ```
    
1. Выполнил команду `file` на объектах разных типов на файловой системе. Например:
    ```bash
    vagrant@vagrant:~$ file /dev/tty
    /dev/tty: character special (5/0)
    vagrant@vagrant:~$ file /dev/sda
    /dev/sda: block special (8/0)
    vagrant@vagrant:~$ file /dev/pts
    /dev/pts: directory
    vagrant@vagrant:~$ file /bin/bash
    /bin/bash: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=a6cb40078351e05121d46daa768e271846d5cc54, for GNU/Linux 3.2.0, stripped
    ```
    База данных `file` находится в: `"/usr/share/misc/magic.mgc"`.
    ```
    stat("/home/vagrant/.magic.mgc", 0x7ffc4a138800) = -1 ENOENT (No such file or directory)
    stat("/home/vagrant/.magic", 0x7ffc4a138800) = -1 ENOENT (No such file or directory)
    openat(AT_FDCWD, "/etc/magic.mgc", O_RDONLY) = -1 ENOENT (No such file or directory)
    stat("/etc/magic", {st_mode=S_IFREG|0644, st_size=111, ...}) = 0
    openat(AT_FDCWD, "/etc/magic", O_RDONLY) = 3
    fstat(3, {st_mode=S_IFREG|0644, st_size=111, ...}) = 0
    read(3, "# Magic local data for file(1) c"..., 4096) = 111
    read(3, "", 4096)                       = 0
    close(3)                                = 0
    openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3
    ```
    
1. Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).
    ```
    ХЗХЗХЗЗЗЗХЗХЗХЗХЗХЗХЗЗЗХХ
    ```

1. Зомби-процессы не занимают ресурсы в ОС (CPU, RAM, IO),  но не освобождают запись в таблице процессов.
    
1. В iovisor BCC есть утилита `opensnoop`:
    ```
    vagrant@vagrant:~$ dpkg -L bpfcc-tools | grep sbin/opensnoop
    /usr/sbin/opensnoop-bpfcc
    vagrant@vagrant:~$ sudo /usr/sbin/opensnoop-bpfcc
    PID    COMM               FD ERR PATH
    882    vminfo              4   0 /var/run/utmp
    655    dbus-daemon        -1   2 /usr/local/share/dbus-1/system-services
    655    dbus-daemon        19   0 /usr/share/dbus-1/system-services
    655    dbus-daemon        -1   2 /lib/dbus-1/system-services
    655    dbus-daemon        19   0 /var/lib/snapd/dbus-1/system-services/
    661    irqbalance          6   0 /proc/interrupts
    661    irqbalance          6   0 /proc/stat
    ```
    
1. Какой системный вызов использует `uname -a`? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в `/proc`, где можно узнать версию ядра и релиз ОС.
    ```
    ХЗХЗХЗЗЗЗХЗХЗХЗХЗХЗХЗЗЗХХ
    ```

1. Оператор `;` выполняет несколько команд одновременно последовательно и обеспечивает вывод без зависимости от успеха и отказа других команд.
    Оператор `&&` (AND оператор) выполнит вторую команду только в том случае, если команда 1 успешно выполнена.
    ```
    vagrant@vagrant:~$ test -d /tmp/some_dir; echo Hi
    Hi
    vagrant@vagrant:~$ test -d /tmp/some_dir&& echo Hi
    vagrant@vagrant:~$
    ```
    `set -e` останавливает выполнение скрипта, если команда или конвейер имеет ошибку, поэтому использование вместе с `&&`, не имеет смысла.


1. Из каких опций состоит режим bash `set -euxo pipefail` и почему его хорошо было бы использовать в сценариях?
    ```
    ХЗХЗХЗЗЗЗХЗХЗХЗХЗХЗХЗЗЗХХ
    ```

1. Используя `-o stat` для `ps`, определите, какой наиболее часто встречающийся статус у процессов в системе. В `man ps` ознакомьтесь (`/PROCESS STATE CODES`) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).
    ```
    ХЗХЗХЗЗЗЗХЗХЗХЗХЗХЗХЗЗЗХХ
    ```
 
 ---
