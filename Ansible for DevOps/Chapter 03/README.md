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

## Manage users and groups

Create an admin group
```bash
ansible app -b -m group -a "name=admin state=present"
```

Create and user that belongs to *admin* group.
```bash
ansible app -b -m user -a "name=johndoe group=admin createhome=yes"
```

Delete that user:
```bash
ansible app -b -m user -a "name=johndoe state=absent remove=yes"
```

## Manage packages

Install git:
```bash
ansible app -b -m package -a "name=git state=present"
```

## Manage files and directories

Get information about a file:
```bash
ansible multi -m stat -a "path=/etc/environment"
```

Retrieve a file from the servers:
```bash
ansible multi -b -m fetch -a "src=/etc/hosts dest=/tmp"
```

Create directory:
```bash
ansible multi -m file -a "dest=/tmp/test mode=644 state=directory"
```

Delete it:
```bash
ansible multi -m file -a "dest=/tmp/test state=absent"
```

## Check logs

A very simple example:
```bash
ansible multi -b -m shell -a "tail /var/log/syslog | grep ansible-command | wc -l"
```

## Cron task

Create cronjob:
```bash
ansible multi -b -m cron -a "name='daily-cron-all-servers' hour=4 job='/path/to/daily-script.sh'"
```

That will create and entry in root's crontab:
```
# crontab -l
#Ansible: daily-cron-all-servers
* 4 * * * /path/to/daily-script.sh
```

Remove cronjobs:

```bash
ansible multi -b -m cron -a "name='daily-cron-all-servers' state=absent"
```

Create cronfiles:
```bash
ansible multi -b -m cron -a "name='daily-cron-all-servers' hour=4 job='/path/to/daily-script.sh' user=root cron_file=daily-cron-all-servers"
```

Create cronfiles:
```bash
ansible multi -b -m cron -a "name='daily-cron-all-servers' hour=4 job='/path/to/daily-script.sh' user=root cron_file=daily-cron-all-servers"
```

Remove cronfiles:
```bash
ansible multi -b -m cron -a "name='daily-cron-all-servers' hour=4 job='/path/to/daily-script.sh' user=root cron_file=daily-cron-all-servers state=absent"
```
