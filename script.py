#!/usr/bin/env python3
import os
dir = os.getcwd()
bash_command = ["cd "+dir, "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
#is_change = False
for result in result_os.split('\n'):
  if result.find('modified') != -1:
    prepare_result = result.replace('\tmodified:   ', '')
    print(dir +"/" + prepare_result)
