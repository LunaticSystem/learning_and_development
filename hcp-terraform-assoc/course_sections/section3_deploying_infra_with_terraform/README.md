## Authentication And Authorization

Authentication is the process of verifying who a user is.

Authorization is the process of verifying what they have access to.

## Terraform Code - First EC2 Instance
```
provider "aws" {
  region     = "us-east-1"
  access_key = "PUT-YOUR-ACCESS-KEY-HERE"
  secret_key = "PUT-YOUR-SECRET-KEY-HERE"
}

resource "aws_instance" "myec2" {
    ami = "ami-00c39f71452c08778"
    instance_type = "t2.micro"
}
```

## Provider Tiers

There are three primary types of provider tiers in terraform.
* Official: Owned and Maintained by HashiCorp. (Recommended)
  * Namespace: hashicorp
* Partner: Owned and Maintained by Technology Company that maintains direct partnership with HashiCorp.
  * Namespace: Third-party organization e.g. mongodb/monbodbatlas
* Community: Owned and Maintained by Individiual Contributors.
  * Namespace: Maintainer's individual or organization account, e.g DeviaVir/gsuite

***<span style="color:green">URL: registry.terraform.io/providers/{NAMESPACE}/***</span>

> [!IMPORTANT]
>  Terraform requires explicit source information for any providers that are not HashiCorp-maintained, using a new syntax in the required_providers nested block inside the terraform configuration block.

<b>Hashicorp Maintained Provider Block:</b>
```
provider "aws" {
  region     = "us-west-2"
  access_key = "PUT-YOUR-ACCESS-KEY-HERE"
  secret_key = "PUT-YOUR-SECRET-KEY-HERE"
}
```

<b>Non-HashiCorp Maintained:</b>
```
terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
  }
}

provider "digitalocean" {
  token = "PUT-YOUR-TOKEN-HERE"
}
```

## Create GitHub Repository through Terraform

Requirements:
* Authentication:
  * Token
    * Settings
    * Developer Settings
    * Personal Access Tokens: Fine Grained.
      * Permissions:
        * Administration: Read and Write Access
      * Everything else can be default.
* Provider Block:
  ```
  terraform {
    required_providers {
      github = {
        source  = "integrations/github"
        version = "~> 5.0"
      }
    }
  }

  # Configure the Github Provider
  provider "github" {
      token = "GITHUB_TOKEN"
  }
  ```

* Creating a repository in terraform:
  ```
  resource "github_repository" "example" {
    name        = "example"
    description = "My awesome codebase"

    visibility = "public"
  }
  ```
* Plan and Apply the IAC:
  ```
  terraform plan

  terraform apply
  ```

## Terraform Destroy

<b>Best Practice:</b>
>If you keep the infrastructure running, you will get charged for it.</br></br>
> Hence it is important for usto also know on how we can delete the infrastructure resources created via terraform  

<b>Approach #1: Destroy All</b></br></br>
Terraform destroy allows us to destroy all the resources that are created within the folder.
```
terraform destroy
```

<b>Approach #2: Destroy Some</b></br></br>
Terraform destory with the `-target` flag allows us to destroy specific resources.

Combination of: Resource type + Local Resource Name</br>
<b>Resource Types:</b>
* aws_instance
* github_repository

<b>Local Resource Name:</b>
* myec2
* example


```
terraform destroy -target {RESOURCE_TYPE}.{LOCAL_RESOURCE_NAME}
```

<b>Approach #3: Remove code to destroy resources</b></br></br>
As the title sounds this is the approach that you can either remove or comment out the blocks of code you wish to destroy. Then once the new terraform file has the resources removed and is applied to terraform. 


## Understanding Terraform State Files

>Terraform stores the state of the infrstructure that is being created from the TF files.
>
>This state allows terraform to map real world resources to your existing configuration.
>
>Terraform state can be refreshed via `terraform plan` and `terraform apply`.
>
>State is located in the `terraform.tfstate` file in the local project directory.
>
>The state file also holds information about the resource. e.g ec2 shows the public_ip, and security groups. 

## Understanding Desired & Current States (NEW)

<b>Desired State:</b> Terraform's primary function is to create, modify, and destroy infrstructure resources to match the desired state described in a Terraform configuration</br>
<b>Current State:</b> Is the actual state of a resource that is currently deployed.</br>

>IMPORTANT: Terraform tries to ensure that the deployed infrastructure is based on the desired state.
>
> If there is a difference between the two, terraform plan presents a description of the changes necessary to achieve the desired state. 

## Challenges with the current state on computed values (NEW)

## Terraform Provider Versioning

Provider Architecture:
* droplet.tf >> terraform <> digital ocean provider <> Digital Ocean (New Server)

Provider plugins are released seperately from Terraform itself.

They have differnt set of version numbers.

<b>Dependency Lock FIle</b></br></br>
Location: terraform.lock.hcl

Upgrading providers:
```
terraform init -upgrade
```

## Terraform Refresh

The terraform refresh command will check the latest state of your infrastructure and update the state file accordingly.

><b>[Points to Note]</b> You shouldn't typically need to use this command, because Terraform automatically performs the same refreshing actions as a part of creating a plan in both the terraform plan and terraform apply commands.

