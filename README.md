# IaC to deploy infar to AWS via cdktf

Infra is writtem in Golang

Currently uses local backend for Terraform state. (Empty state file is in Repo)
AWS cli needs to be configured separately
If state is stored in Terraform Cloud, PAT and connection to TF cloud needs to be configured.

