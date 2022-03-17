#!/usr/bin/env python3

import sys, os
import yaml
import shutil
from jinja2 import Environment

with open('.version') as f:
    VERSION = f.read()

RECIPE_DIR = 'build/{}/recipes'.format(VERSION)
if os.path.exists(RECIPE_DIR):
    shutil.rmtree(RECIPE_DIR)

FAILED = []

for repository in os.listdir('recipes'):
    print("=> configuring {}".format(repository))
    os.makedirs(RECIPE_DIR + '/' + repository)
    for file in os.listdir('recipes/{}/'.format(repository)):
        if file.endswith('.yml'):
            try:
                with open('recipes/{}/{}'.format(repository, file)) as f:
                    data = f.read()
                    obj = yaml.full_load(data)
                    try:
                        if 'packages' in obj and len(obj['packages']) != 1:
                            print("failed in {} {}".format(file, e))
                            FAILED.append(file)
                        data = Environment().from_string(data).render(obj)
                        with open('{}/{}/{}'.format(RECIPE_DIR, repository, file),'w') as fw:
                            fw.write(data)
                    except Exception as e:
                        print("failed in {} {}".format(file, e))        
                        FAILED.append(file)
            except Exception as e:
                print("failed in {} {}".format(file, e))
                FAILED.append(file)

print("failed", FAILED)