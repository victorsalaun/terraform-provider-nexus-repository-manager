provider "nexus" {
  url           = "http://localhost:8081"
  auth_username = "admin"
  auth_password = "admin123"
}

resource "nexus_user" "user1" {
  user_id        = "user1"
  first_name     = "User"
  last_name      = "One"
  email_address  = "user1@example.com"
  source         = "default"
  password       = "admin123"
  status         = "active"
  roles          = [
    "nx-admin"
  ]
  external_roles = []
}

