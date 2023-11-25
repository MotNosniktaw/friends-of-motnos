variable "region" {
  type    = string
  default = "uk-south"
}

variable "route" {
  type = list(object({
    cidr_block     = string
    gateway_id     = string
    instance_id    = string
    nat_gateway_id = string
  }))
}

variable "subnet_ids" {
  type = list(string)
}

variable "enable_dns_support" {
  type    = bool
  default = true
}

variable "enable_dns_hostnames" {
  type    = bool
  default = true
}

variable "ami" {
  type = string
}

variable "instance_type" {
  type    = string
  default = "aws.t2.micro"
}

variable "user_data" {
  type    = string
  default = ""
}

variable "key_name" {
  type = string
}

variable "security_group_id" {
  type = string
}

variable "namespace" {
  type    = string
  default = "FoM"
}

variable "subnet_id" {
  type = string
}

variable "rds_vpc_id" {
  type = string
}

variable "identifier" {
  type = string
}

variable "rds_port" {
  type    = number
  default = 5432
}

variable "rds_username" {
  type    = string
  default = "FoM"
}

variable "rds_password" {
  type    = string
  default = "kirk4lyf"
}

variable "rds_final_snapshot_identifier" {
  type    = string
  default = "FoM-RDS-final-snapshot"
}
