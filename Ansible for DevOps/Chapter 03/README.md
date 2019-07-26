# Chapter 3

## Update inventory

Update inventory with *app* and *db* servers:
```yaml
[app]
test01
test02

[db]
test03

[multi:children]
app
db

[multi:vars]
ansible_python_interpreter=/usr/bin/python3
```

## Install dependeces

### app dependences

Install Django dependences:
```bash
ansible app -b -m apt -a "name=python3-pip state=present"
ansible app -b -m pip -a "name=django<2 state=present"
ansible app -b -m pip -a "name=PyMySQL state=present"
ansible app -a "python3 -c 'import django; print(django.get_version())'"
```

### db dependences

Install dependences:
```bash
ansible db -b -m apt -a "name=mariadb-server state=present"
ansible db -b -m service -a "name=mariadb state=started enabled=yes"
ansible db -b -a "iptables -F"
ansible db -b -a "iptables -A INPUT -s 10.22.0.0/24 -p tcp -m tcp --dport 3306 -j ACCEPT"
ansible db -b -m apt -a "name=python3-pip state=present"
ansible db -b -m pip -a "name=PyMySQL state=present"
ansible db -b -m mysql_user -a "name=django host=% password=12345 priv=*.*:ALL state=present login_unix_socket=/var/run/mysqld/mysqld.sock"
```
