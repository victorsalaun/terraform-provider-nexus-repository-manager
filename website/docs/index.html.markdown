---
layout: "nexus-repository-manager"
page_title: "Provider: Nexus Repository Manager"
sidebar_current: "docs-nexus-reposutory-manager-index"
description: |-
  The Nexus Repository Manager provider configures users, repositories... in Nexus Repository Manager
---

# Nexus Repository Manager Provider

The Nexus Repository Manager provider allows Terraform to create and configure Users,
Repositories... in [Nexus Repository Manager](https://www.sonatype.com/product-nexus-repository). Rundeck is a tool for managing binaries and build artifacts
across your software supply chain.

The provider configuration block accepts the following arguments:

* ``url`` - (Required) The root URL of a Nexus Repository Manager server. May alternatively be set via the
  ``NEXUS_REPOSITORY_MANAGER_URL`` environment variable.

* ``auth_username`` - (Required) The API auth username to use when making requests. May alternatively
  be set via the ``NEXUS_REPOSITORY_MANAGER_AUTH_USERNAME`` environment variable.
 
* ``auth_password`` - (Required) The API auth password to use when making requests. May alternatively
  be set via the ``NEXUS_REPOSITORY_MANAGER_AUTH_PASSWORD`` environment variable.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "nexus-repository-manager" {
  url            = "https://nexus.example.com/"
  auth_username  = "admin"
  auth_password  = "admin123"
}
```