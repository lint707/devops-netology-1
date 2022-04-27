# Домашнее задание к занятию "3.9. Элементы безопасности информационных систем"

1. Установил Bitwarden плагин для браузера. Зарегестрировался и сохранил несколько паролей.
![BW](img/BW.jpg)

2. Установил Google authenticator на мобильный телефон. Настроил вход в Bitwarden акаунт через Google authenticator OTP.
![BWauthenticator ](img/BWtwo.jpg)

3. Установите apache2, сгенерируйте самоподписанный сертификат, настройте тестовый сайт для работы по HTTPS.


4. Проверил на TLS уязвимости произвольный сайт в интернете:
```
vagrant@vagrant:~/testssl.sh$ ./testssl.sh -U --sneaky https://www.google.com/

###########################################################
    testssl.sh       3.1dev from https://testssl.sh/dev/
    (ab33f6c 2022-04-25 10:18:49 -- )

      This program is free software. Distribution and
             modification under GPLv2 permitted.
      USAGE w/o ANY WARRANTY. USE IT AT YOUR OWN RISK!

       Please file bugs @ https://testssl.sh/bugs/

###########################################################

 Using "OpenSSL 1.0.2-chacha (1.0.2k-dev)" [~183 ciphers]
 on vagrant:./bin/openssl.Linux.x86_64
 (built: "Jan 18 17:12:17 2019", platform: "linux-x86_64")


Testing all IPv4 addresses (port 443): 64.233.164.106 64.233.164.105 64.233.164.103 64.233.164.104 64.233.164.147 64.233.164.99
---------------------------------------------------------------------------------
 Start 2022-04-27 07:29:45        -->> 64.233.164.106:443 (www.google.com) <<--

 Further IP addresses:   64.233.164.99 64.233.164.147 64.233.164.104 64.233.164.103 64.233.164.105 2a00:1450:4010:c0e::6a
                         2a00:1450:4010:c0e::93 2a00:1450:4010:c0e::67 2a00:1450:4010:c0e::68
 rDNS (64.233.164.106):  lf-in-f106.1e100.net.
 Service detected:       HTTP


 Testing vulnerabilities

 Heartbleed (CVE-2014-0160)                not vulnerable (OK), no heartbeat extension
 CCS (CVE-2014-0224)                       not vulnerable (OK)
 Ticketbleed (CVE-2016-9244), experiment.  not vulnerable (OK)
 ROBOT                                     not vulnerable (OK)
 Secure Renegotiation (RFC 5746)           supported (OK)
 Secure Client-Initiated Renegotiation     not vulnerable (OK)
 CRIME, TLS (CVE-2012-4929)                not vulnerable (OK)
 BREACH (CVE-2013-3587)                    potentially NOT ok, "br gzip" HTTP compression detected. - only supplied "/" tested
                                           Can be ignored for static pages or if no secrets in the page
 POODLE, SSL (CVE-2014-3566)               not vulnerable (OK)
 TLS_FALLBACK_SCSV (RFC 7507)              Downgrade attack prevention supported (OK)
 SWEET32 (CVE-2016-2183, CVE-2016-6329)    VULNERABLE, uses 64 bit block ciphers
 FREAK (CVE-2015-0204)                     not vulnerable (OK)
 DROWN (CVE-2016-0800, CVE-2016-0703)      not vulnerable on this host and port (OK)
                                           make sure you don't use this certificate elsewhere with SSLv2 enabled services
                                           https://censys.io/ipv4?q=A6E4CC389843F4694161666B1EB8E73821592E180657A8824D8D8DCC8D05ED37 could help you to find out
 LOGJAM (CVE-2015-4000), experimental      not vulnerable (OK): no DH EXPORT ciphers, no DH key detected with <= TLS 1.2
 BEAST (CVE-2011-3389)                     TLS1: ECDHE-ECDSA-AES128-SHA ECDHE-ECDSA-AES256-SHA ECDHE-RSA-AES128-SHA
                                                 ECDHE-RSA-AES256-SHA AES128-SHA AES256-SHA DES-CBC3-SHA
                                           VULNERABLE -- but also supports higher protocols  TLSv1.1 TLSv1.2 (likely mitigated)
 LUCKY13 (CVE-2013-0169), experimental     potentially VULNERABLE, uses cipher block chaining (CBC) ciphers with TLS. Check patches
 Winshock (CVE-2014-6321), experimental    not vulnerable (OK)
 RC4 (CVE-2013-2566, CVE-2015-2808)        no RC4 ciphers detected (OK)


 Done 2022-04-27 07:30:27 [  44s] -->> 64.233.164.106:443 (www.google.com) <<--

---------------------------------------------------------------------------------
 Start 2022-04-27 07:30:27        -->> 64.233.164.105:443 (www.google.com) <<--

 Further IP addresses:   64.233.164.99 64.233.164.147 64.233.164.104 64.233.164.103 64.233.164.106 2a00:1450:4010:c0e::6a
                         2a00:1450:4010:c0e::93 2a00:1450:4010:c0e::67 2a00:1450:4010:c0e::68
                        
^C
```

5. Установите на Ubuntu ssh сервер, сгенерируйте новый приватный ключ. Скопируйте свой публичный ключ на другой сервер. Подключитесь к серверу по SSH-ключу.
 
 
6. Переименуйте файлы ключей из задания 5. Настройте файл конфигурации SSH клиента, так чтобы вход на удаленный сервер осуществлялся по имени сервера.


7. Соберите дамп трафика утилитой tcpdump в формате pcap, 100 пакетов. Откройте файл pcap в Wireshark.
```bash
vagrant@vagrant:~$ sudo tcpdump -c 100 -w dump.pcap
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
36 packets captured
39 packets received by filte
0 packets dropped by kernel
```
 ---
## Задание для самостоятельной отработки (необязательно к выполнению)

8*. Просканируйте хост scanme.nmap.org. Какие сервисы запущены?


9*. Установите и настройте фаервол ufw на web-сервер из задания 3. Откройте доступ снаружи только к портам 22,80,443


 ---
