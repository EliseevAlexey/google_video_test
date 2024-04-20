### Create Google Project

https://console.cloud.google.com/welcome/new?hl=en&project=primeval-wind-420823

PROJECT_ID:

```
primeval-wind-420823
```


```shell
gcloud config get-value project
```

---
Install utils
```shell
brew install yq
```

---

### Install Google SDK

Brew installation:

```shell
brew install google-cloud-sdk
```

Check with:

```shell
gcloud -v
```

Response:

```text
Google Cloud SDK 472.0.0
bq 2.1.4
core 2024.04.12
gcloud-crc32c 1.0.0
gsutil 5.27
```

---

### Login with SDK

```shell
gcloud auth login
```

Set default project

```shell
gcloud config set project primeval-wind-420823 
```

Response:

```shell
Updated property [core/project].
```

---

### Init Git

Install git project

```shell
git init 
```

---

### Install terraform

https://formulae.brew.sh/formula/terraform

```shell
brew install terraform
```

Check:

```shell
terraform -v
```

Response:

```text
Terraform v1.8.1
on darwin_amd64
```

---

### Create a Server Instance with Terraform:


Create a folder & move there

```shell
mkdir -p terraform
```

Create a config file

```shell
touch terraform/main.tf
```

Init terraform

```shell
cd terraform && terraform init
```

Response:

```text
Initializing the backend...

Initializing provider plugins...
- Finding latest version of hashicorp/google...
- Installing hashicorp/google v5.25.0...
- Installed hashicorp/google v5.25.0 (signed by HashiCorp)

Terraform has created a lock file .terraform.lock.hcl to record the provider
selections it made above. Include this file in your version control repository
so that Terraform can guarantee to make the same selections by default when
you run "terraform init" in the future.

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

will create a `./terraform/.terraform.lock.hcl` file

Create a Plan

```shell
cd terraform && terraform plan
```

Response:

```text
Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # google_compute_instance.my_instance will be created
  + resource "google_compute_instance" "my_instance" {
      + allow_stopping_for_update = true
      + can_ip_forward            = false
        ...
        some code here
        ...
    }

Plan: 1 to add, 0 to change, 0 to destroy.
```

Apply the config
```shell
cd terraform && terraform apply
```
Response:
```text
Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # google_compute_instance.terraform_instance will be created
  + resource "google_compute_instance" "terraform_instance" {
      + allow_stopping_for_update = true
      + can_ip_forward            = false
      + cpu_platform              = (known after apply)
        ...
        some code here
        ...
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes 

google_compute_instance.terraform_instance: Creating...
google_compute_instance.terraform_instance: Still creating... [10s elapsed]
google_compute_instance.terraform_instance: Creation complete after 15s [id=projects/primeval-wind-420823/zones/us-west1-b/instances/terraform-instance]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

---

Cloud Function VS Cloud Run

https://cloud.google.com/blog/products/serverless/cloud-run-vs-cloud-functions-for-serverless

