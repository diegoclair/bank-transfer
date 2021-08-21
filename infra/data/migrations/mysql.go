package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
//Only work doing 1 command per version, you cannot create two tables in the same script, need to create a new version
var (
	Migrations = []darwin.Migration{}
)
