locals {
  resource_name_prefix = "${var.namespace}-rds"
}

resource "aws_db_subnet_group" "_" {
  name       = "${local.resource_name_prefix}-${var.identifier}-subnet-group"
  subnet_ids = var.subnet_ids
}

resource "aws_db_instance" "_" {
  identifier = "${local.resource_name_prefix}-${var.identifier}"

  allocated_storage       = 20
  backup_retention_period = 1
  backup_window           = "10:46-11:16"
  maintenance_window      = "Mon:00:00-Mon:03:00"
  db_subnet_group_name    = aws_db_subnet_group._.id
  engine                  = "postgres"
  engine_version          = "15.3"
  instance_class          = "db.t2.micro"
  username                = var.rds_username
  password                = var.rds_password
  port                    = var.rds_port
  publicly_accessible     = false
  storage_encrypted       = false

  vpc_security_group_ids = ["${aws_security_group._.id}"]

  allow_major_version_upgrade = false
  auto_minor_version_upgrade  = false

  final_snapshot_identifier = var.rds_final_snapshot_identifier
  snapshot_identifier       = null
  skip_final_snapshot       = true

  performance_insights_enabled = false

  tags = {
    Name = "${local.resource_name_prefix}-${var.identifier}"
  }

}
