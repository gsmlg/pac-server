# -*- coding:utf-8 -*-

import re

with open("gfwlist.conf", 'w') as wp:
    for line in open("gfwlist.pac"):
        if re.match(r" {12}\".*\",?\n", line) != None:
            domain_valid = line.split('"')[1]
            zone_line = 'forward-zone:\n'
            zone_line += f'  name: "{domain_valid}"\n'
            zone_line += f'  forward-tls-upstream: yes\n'
            zone_line += f'  forward-addr: 8.8.8.8@853\n'
            zone_line += f'  forward-addr: 1.1.1.1@853\n'
            wp.writelines(zone_line)
