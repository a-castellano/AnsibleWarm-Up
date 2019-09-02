package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	Hosts []string          `json:"hosts"`
	Vars  map[string]string `json:"vars"`
}

func main() {

	ansible_user := os.Getenv("DYNAMIC_INVENTORY_ANSIBLE_USER")

	hosts := Group{[]string{"test01", "test02"}, map[string]string{"some_variable": "some_value", "ansible_python_interpreter": "/usr/bin/python3"}}

	if len(ansible_user) > 0 {
		hosts.Vars["ansible_ssh_user"] = ansible_user
	}

	groups := map[string]Group{"group": hosts}

	js, _ := json.Marshal(groups)
	fmt.Print(string(js))
}
