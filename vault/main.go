package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/vault/api"
)

func main() {
	c, err := api.NewClient(&api.Config{
		//Address: "http://0.0.0.0:8200",
		Address: "https://dev-vault.financialapps.com",
	})
	check(err)
	//c.SetToken("b2338e00-598d-4c89-ad2f-a818cd35c513")
	c.SetToken("95815174-6ab0-208f-7246-a969e4f74668")

	//res, err := c.Sys().SealStatus()
	//check(err)
	//spew.Dump(res)

	//m, err := c.Sys().ListMounts()
	//check(err)
	//spew.Dump(m)

	//write(c)
	//writeJacob(c)

	//read(c, "secret/admin/5ed7dd5e-4b8a-4f3d-420d-76815e1deb74")
	//read(c, "secret/admin/46975b4c-7151-4bc9-5e28-3065e21478b3")
	//read(c, "secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a/pii")
	//read(c, "transit/keys/49fb918d-7e71-44dd-7378-58f19606df2a")
	//read(c, "secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a/operator/9cc36daa-a63a-48c9-b465-1ddd2c527128")
	//read(c, "secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a/operator/9cc36daa-a63a-48c9-b465-1ddd2c527128")
	//read(c, "secret/tenant/815b024d-1f47-4c73-737e-b2a74231b38b/consumer/e03c9554-6f7a-4af3-4b12-e8c22abc388f")
	list(c)
	//deleteAll(c, "secret/application")
	//delete(c, "secret/test")

	//fmt.Println(c.Sys().Seal())
}

func write(c *api.Client) {
	var s *api.Secret
	var err error

	fmt.Println("Write a secret")

	s, err = c.Logical().Write(
		"secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a",
		map[string]interface{}{
			"private_id": "cd09721a-2319-4009-9fc5-da00620f159d",
			"pii_key":    "some-32-bit-key???--------------",
		},
	)
	check(err)
	spew.Dump(s)
}

func writeJacob(c *api.Client) {
	var s *api.Secret
	var err error

	fmt.Println("Write admin jacob info")

	s, err = c.Logical().Write(
		"secret/admin/a793a0dd-ee7d-41ef-7b03-2204907a5986",
		map[string]interface{}{"private_id": "a869b274-5a33-4b64-a922-af296f8aeaf5"},
	)
	check(err)
	spew.Dump(s)

}

func read(c *api.Client, path string) {
	fmt.Println("Read a secret")
	s, err := c.Logical().Read(path)
	check(err)
	spew.Dump(s)
}

func delete(c *api.Client, path string) {
	s, err := c.Logical().Delete(path)
	check(err)
	spew.Dump(s)
}

func list(c *api.Client) {
	var s *api.Secret
	var err error

	fmt.Println("List all secrets")

	s, err = c.Logical().List("secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a/operator")
	check(err)

	fmt.Println(s.Data)

	for _, v := range s.Data["keys"].([]interface{}) {
		v := v.(string)

		if v[len(v)-1:] == "/" {
			continue
		}
		fmt.Print(v, ":\t")

		s, err = c.Logical().Read("secret/tenant/49fb918d-7e71-44dd-7378-58f19606df2a/operator/" + v)
		check(err)

		fmt.Printf("%+v\n", s.Data)
	}
}

func deleteAll(c *api.Client, path string) {
	var s *api.Secret
	var err error

	fmt.Println("Dleete all secrets under", path)

	s, err = c.Logical().List(path)
	check(err)

	fmt.Println(s.Data)

	for _, v := range s.Data["keys"].([]interface{}) {
		v := v.(string)

		if v[len(v)-1:] == "/" {
			continue
		}
		s, err = c.Logical().Delete(path + "/" + v)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
