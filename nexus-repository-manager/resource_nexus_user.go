package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNexusUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceNexusUserCreate,
		Read:   resourceNexusUserRead,
		Update: resourceNexusUserUpdate,
		Delete: resourceNexusUserDelete,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"email_address": {
				Type:     schema.TypeString,
				Required: true,
			},

			"source": {
				Type:     schema.TypeString,
				Required: true,
			},

			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},

			"status": {
				Type:     schema.TypeString,
				Required: true,
			},

			"roles": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"external_roles": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceNexusUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userId := d.Get("user_id").(string)
	firstName := d.Get("first_name").(string)
	lastName := d.Get("last_name").(string)
	emailAddress := d.Get("email_address").(string)
	source := d.Get("source").(string)
	password := d.Get("password").(string)
	status := d.Get("status").(string)
	roles := expandStringList(d.Get("roles").(*schema.Set).List())
	externalRoles := expandStringList(d.Get("external_roles").(*schema.Set).List())

	user := nexusUser{UserId: userId, FirstName: firstName, LastName: lastName, EmailAddress: emailAddress, Source: source, Password: password, Status: status, Roles: roles, ExternalRoles: externalRoles}
	userJson, _ := json.Marshal(user)

	_, err := userCreate(*config, bytes.NewBuffer(userJson))
	if err != nil {
		return fmt.Errorf("error decoding JSON: %s", err)
	}

	d.SetId(userId)

	return resourceNexusUserRead(d, meta)
}

func resourceNexusUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userId := d.Get("user_id").(string)

	_, _ = userRead(*config, userId)
	return nil
}

func resourceNexusUserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userId := d.Get("user_id").(string)
	firstName := d.Get("first_name").(string)
	lastName := d.Get("last_name").(string)
	emailAddress := d.Get("email_address").(string)
	source := d.Get("source").(string)
	password := d.Get("password").(string)
	status := d.Get("status").(string)
	roles := expandStringList(d.Get("roles").(*schema.Set).List())
	externalRoles := expandStringList(d.Get("external_roles").(*schema.Set).List())

	user := nexusUser{UserId: userId, FirstName: firstName, LastName: lastName, EmailAddress: emailAddress, Source: source, Password: password, Status: status, Roles: roles, ExternalRoles: externalRoles}
	userJson, _ := json.Marshal(user)

	_, err := userUpdate(*config, userId, bytes.NewBuffer(userJson))
	if err != nil {
		return fmt.Errorf("error decoding JSON: %s", err)
	}

	return resourceNexusUserRead(d, meta)
}

func resourceNexusUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userId := d.Get("user_id").(string)

	_, _ = userDelete(*config, userId)
	return nil
}
