# https://visualsubnetcalc.com/index.html?c=1N4IgbiBcIEwgNCARlEBGADAOm7g9GgGwIgC2qAggOoDKJAzlKBgOZMgYAW7GAljwCseAax4AbdgDtUABQBOAEwAEAWiUB3AKYpEAY1QgAvojQTIoadHnK1WnSH3Qjx9KPOwzF2YtVKAhgAOASSOIEaIAMyeIJYg1r6BwXoGhqkmQu4wbqAALNGx8WoK9qHhIACs+d42SsUhKS4R2TEGyU6pLmj8mRm5zQAcVVZymgFyAPY1dvXtiACcQ3EjY5O+021haRXNmIvyoxM1iTObJqZS3gerascbzo29IIQ7MHvLh751d51RF8NXNS+DgaiAA7M1Yid7p1uJlurlHv0doRFgBFPxrbRQzqgtEY2xY74uOY7QZ-EDohJBbEmBbkyk3alElzlR6YZowDB4z4lFKIGDndyxBm1XntTpoCGtYHixrwp5srLsGB5ckAYXGpExYtOsEq6s12uxPw5KINWsZSRlupguPNVKtpQ6YLZOQ5ZKFqA1FtFNNgdM90G9PONJnKUugkj8xo6qSAA

locals {
  web_subnet_configuration = {
    prd-web-subnet-public-1a = {
      env               = "prd"
      cidr_block        = "10.0.0.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    prd-web-subnet-public-1b = {
      env               = "prd"
      cidr_block        = "10.0.8.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    preprod-web-subnet-public-1a = {
      env               = "preprod"
      cidr_block        = "10.0.64.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    preprod-web-subnet-public-1b = {
      env               = "preprod"
      cidr_block        = "10.0.72.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    qa-web-subnet-public-1a = {
      env               = "qa"
      cidr_block        = "10.0.128.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    qa-web-subnet-public-1b = {
      env               = "qa"
      cidr_block        = "10.0.136.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    com-web-subnet-public-1a = {
      env               = "com"
      cidr_block        = "10.0.192.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    com-web-subnet-public-1b = {
      env               = "com"
      cidr_block        = "10.0.200.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
  }
  app_subnet_configuration = {
    prd-app-subnet-private-1a = {
      env               = "prd"
      cidr_block        = "10.0.16.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    prd-app-subnet-private-1b = {
      env               = "prd"
      cidr_block        = "10.0.24.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    preprod-app-subnet-private-1a = {
      env               = "preprod"
      cidr_block        = "10.0.80.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    preprod-app-subnet-private-1b = {
      env               = "preprod"
      cidr_block        = "10.0.88.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    qa-app-subnet-private-1a = {
      env               = "qa"
      cidr_block        = "10.0.144.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    qa-app-subnet-private-1b = {
      env               = "qa"
      cidr_block        = "10.0.152.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    com-db-subnet-private-1a = {
      env               = "com"
      cidr_block        = "10.0.208.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    com-db-subnet-private-1b = {
      env               = "com"
      cidr_block        = "10.0.216.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
  }
  db_subnet_configuration = {
    prd-db-subnet-private-1a = {
      env               = "prd"
      cidr_block        = "10.0.32.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    prd-db-subnet-private-1b = {
      env               = "prd"
      cidr_block        = "10.0.40.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    preprod-db-subnet-private-1a = {
      env               = "preprod"
      cidr_block        = "10.0.96.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    preprod-db-subnet-private-1b = {
      env               = "preprod"
      cidr_block        = "10.0.104.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    qa-db-subnet-private-1a = {
      env               = "qa"
      cidr_block        = "10.0.160.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    qa-db-subnet-private-1b = {
      env               = "qa"
      cidr_block        = "10.0.168.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
    com-db-subnet-private-1a = {
      env               = "com"
      cidr_block        = "10.0.224.0/21"
      availability_zone = "ap-southeast-1a"
      usable_hosts      = 2043
    }
    com-db-subnet-private-1b = {
      env               = "com"
      cidr_block        = "10.0.232.0/21"
      availability_zone = "ap-southeast-1b"
      usable_hosts      = 2043
    }
  }
}

resource "aws_subnet" "web" {
  for_each          = local.web_subnet_configuration
  vpc_id            = aws_vpc.common.id
  cidr_block        = each.value.cidr_block
  availability_zone = each.value.availability_zone

  tags = {
    Environment = each.value.env
    Name        = each.key
    Zone        = "web"
    ProjectCode = var.project_code
    SubnetType  = "public"
    UsableHosts = each.value.usable_hosts
    VpcId       = aws_vpc.common.id
    VpcName     = local.vpc_name
  }
}

resource "aws_route_table_association" "web" {
  for_each       = local.web_subnet_configuration
  route_table_id = aws_route_table.web.id
  subnet_id      = aws_subnet.web[each.key].id
}

resource "aws_subnet" "app" {
  for_each          = local.app_subnet_configuration
  vpc_id            = aws_vpc.common.id
  cidr_block        = each.value.cidr_block
  availability_zone = each.value.availability_zone

  tags = {
    Environment = each.value.env
    Name        = each.key
    Zone        = "app"
    ProjectCode = var.project_code
    SubnetType  = "private"
    UsableHosts = each.value.usable_hosts
    VpcId       = aws_vpc.common.id
    VpcName     = local.vpc_name
  }
}

resource "aws_route_table_association" "app" {
  for_each       = local.app_subnet_configuration
  route_table_id = aws_route_table.app.id
  subnet_id      = aws_subnet.app[each.key].id
}

resource "aws_subnet" "db" {
  for_each          = local.db_subnet_configuration
  vpc_id            = aws_vpc.common.id
  cidr_block        = each.value.cidr_block
  availability_zone = each.value.availability_zone

  tags = {
    Environment = each.value.env
    Name        = each.key
    Zone        = "db"
    ProjectCode = var.project_code
    SubnetType  = "private"
    UsableHosts = each.value.usable_hosts
    VpcId       = aws_vpc.common.id
    VpcName     = local.vpc_name
  }
}

resource "aws_route_table_association" "db" {
  for_each       = local.db_subnet_configuration
  route_table_id = aws_route_table.db.id
  subnet_id      = aws_subnet.db[each.key].id
}
