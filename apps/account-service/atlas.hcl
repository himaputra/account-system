data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./migrations/migrate.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  migration {
    dir = "file://migrations"
  }
  dev = "docker://postgres/16/dev?search_path=public"
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "local" {
  dir = "file://migrations"
  dev = "docker://postgres/16/dev?search_path=public"
  url = "postgres://account:pass123@localhost:5432/account?sslmode=disable"
}
