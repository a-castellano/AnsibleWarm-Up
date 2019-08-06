# Inventories

Provision examples does not work because [dopy](https://github.com/Wiredcraft/dopy) seems to be not compatible with Python3.

## Dynamic Inventories

Compile **dynamic_inventory.go**:
```
go build dynamic_inventory.go
```

Check it changing default user:

Invalid user:
```bash
DYNAMIC_INVENTORY_ANSIBLE_USER=invaliduser ansible all -i dynamic_inventory -m ping
test02 | UNREACHABLE! => {
    "changed": false,
    "msg": "Failed to connect to the host via ssh: invaliduser@test02: Permission denied (publickey).",
    "unreachable": true
}
test01 | UNREACHABLE! => {
    "changed": false,
    "msg": "Failed to connect to the host via ssh: invaliduser@test01: Permission denied (publickey).",
    "unreachable": true
}
```

Correct user:
```bash
DYNAMIC_INVENTORY_ANSIBLE_USER=user ansible all -i dynamic_inventory -m ping
test02 | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
test01 | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
```
