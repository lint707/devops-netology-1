#!/usr/bin/env python3
import socket
import time
import json
import yaml

srv = {'drive.google.com':'0.0.0.0', 'mail.google.com':'0.0.0.0', 'google.com':'0.0.0.0'}
while 1 == 1:
  for url, ip in srv.items():
    new_ip = socket.gethostbyname(url)
    if new_ip != ip:
      print(' [ERROR] ' + str(url) +' IP mistmatch: '+srv[url]+' '+new_ip)
      srv[url]=new_ip
    else:
     # srv[url] = new_ip  
      print(str(url) + ' - ' + ip) 
  with open('srv.json', 'w') as json_file:
    json_data= json.dumps(srv, indent=2)
    json_file.write(json_data) 
  with open('srv.yaml', 'w') as yaml_file:
    yaml_data= yaml.dump(srv, explicit_start=True, explicit_end=True)
    yaml_file.write(yaml_data) 
                                            
#  with open('srv.json', 'w') as json_file:
#    json_file.write(json.dumps(srv, indent=2))
#  with open('srv.yaml', 'w') as yaml_file:
#    yaml_file.write(yaml.dump(srv, indent=2))
    time.sleep(3)

