package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type VMDB_Object struct {
	end_point                 *ConnectionParameters_t
	href_slug                 string
	attributes                map[string]interface{}
	include_custom_attributes bool
	include_tags              bool
}

func (vmdb_object *VMDB_Object) Fetch() error {
	var href string
	if vmdb_object.include_custom_attributes {
		href = vmdb_object.href_slug + "?expand=custom_attributes"
	} else {
		href = vmdb_object.href_slug
	}
	body, _ := vmdb_object.end_point.Get(href)

	if err := json.Unmarshal(body, &vmdb_object.attributes); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func NewVMDB_Object(end_point *ConnectionParameters_t, href_slug string) *VMDB_Object {
	var vmdb VMDB_Object
	vmdb.attributes = make(map[string]interface{})
	vmdb.end_point = end_point
	vmdb.href_slug = href_slug
	vmdb.include_tags = false
	vmdb.include_custom_attributes = false
	return &vmdb
}

func (vmdb_object *VMDB_Object) CustomAttributes(enabled bool) {
	vmdb_object.include_custom_attributes = enabled
	return
}

func (vmdb_object *VMDB_Object) Dump() {
	fmt.Println("VMDB Attribute List", vmdb_object.GetAttributeList())
	fmt.Println("href ", vmdb_object.GetAttribute("href").(string))
	fmt.Println("actions ", vmdb_object.GetAttribute("actions"))
	fmt.Println("custom_attributes", vmdb_object.GetAttribute("custom_attributes"))
	return
}

func (vmdb *VMDB_Object) GetAttribute(attribute_name string) interface{} {
	return vmdb.attributes[attribute_name]
}

func (vmdb *VMDB_Object) GetAttributeList() []string {
	var keys []string

	for key, _ := range vmdb.attributes {
		keys = append(keys, key)
	}
	return keys
}

func (vmdb *VMDB_Object) AddCustomAttribute(name string, value string) error {
	var payload = make(map[string]interface{})
	var resources = make([]map[string]string, 1, 1)
	resources[0] = make(map[string]string)
	resources[0]["name"] = name
	resources[0]["value"] = value
	payload["action"] = "add"
	payload["resources"] = resources
	b, _ := json.Marshal(payload)
	href := vmdb.href_slug + "/custom_attributes"

	vmdb.end_point.Post(b, href)
	fmt.Println(string(b))
	return nil
}

func (vmdb *VMDB_Object) EditCustomAttribute(name string, value string) error {
	var payload = make(map[string]interface{})
	var resources = make([]map[string]string, 1, 1)
	resources[0] = make(map[string]string)
	resources[0]["name"] = name
	resources[0]["value"] = value
	payload["action"] = "edit"
	payload["resources"] = resources
	b, _ := json.Marshal(payload)
	href := vmdb.href_slug + "/custom_attributes"

	vmdb.end_point.Post(b, href)
	fmt.Println(string(b))
	return nil
}

func (vmdb *VMDB_Object) DeleteCustomAttribute(name string) error {
	var payload = make(map[string]interface{})
	var resources = make([]map[string]string, 1, 1)
	resources[0] = make(map[string]string)
	resources[0]["name"] = name
	payload["action"] = "delete"
	payload["resources"] = resources
	b, _ := json.Marshal(payload)
	href := vmdb.href_slug + "/custom_attributes"

	vmdb.end_point.Post(b, href)
	fmt.Println(string(b))
	return nil
}
