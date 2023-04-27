#!/usr/bin/python3

# put this file in /usr/lib/zabbix/alertscripts/

import requests
import sys
import json
import logging.handlers


URL = "http://104.130.246.15:9090/api/new_event"


log = logging.getLogger(__name__)
log.setLevel(logging.DEBUG)
handler = logging.handlers.SysLogHandler(address='/dev/log')
formatter = logging.Formatter('%(module)s.%(funcName)s: %(message)s')
handler.setFormatter(formatter)
log.addHandler(handler)


def main(askey, subject, body, key):
    reqbody = {}
    itemLists = body.splitlines()
    for i in itemLists:
        reqbody[i.split(":")[0]] = i.split(":")[1]

    headers = {'X-Auth-Key': key}

    resp = requests.post(URL, headers=headers, data=json.dumps(reqbody))
    if resp.status_code != 200:
        print(resp.text)
        sys.exit(1)

if __name__ == "__main__":
    to = sys.argv[1]
    subject = sys.argv[2]
    body = sys.argv[3]
    key = sys.argv[4]
    print(to)
    print(subject)
    print(body)
    main(to, subject, body, key)
