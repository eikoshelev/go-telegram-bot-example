package migrations

import "embed"

//go:embed *.sql
var Static embed.FS
