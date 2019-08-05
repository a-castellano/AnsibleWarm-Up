# Chapter 5

## Ansible vault

Encrypt file
```
ansible-vault encrypt secrets.yml
```

Try to run playbook without password:
```
ansible-playbook Ansible\ for\ DevOps/Chapter\ 05/playbook.yml
ERROR! Attempting to decrypt but no vault secrets found
```
