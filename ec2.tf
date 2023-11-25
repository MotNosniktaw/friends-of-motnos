locals {
  resource_name_prefix = "${var.namespace}-ec2"
}

resource "aws_instance" "_" {
  ami                    = var.ami
  instance_type          = var.instance_type
  user_data              = var.user_data
  key_name               = var.key_name
  subnet_id              = var.subnet_id
  vpc_security_group_ids = [var.security_group_id]
  tags = {
    Name = "${local.resource_name_prefix}-instance"
  }

}

resource "aws_eip" "_" {
  instance = aws_instance._.id
  vpc      = true

  tags = {
    Name = "${local.resource_name_prefix}-eip"
  }

}

resource "tls_private_key" "-" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "_" {
  key_name   = var.key_name
  public_key = tls_private_key._.public_key_openssh

  tags = {
    Name = "${local.resource_name_prefix}-key"
  }
}
