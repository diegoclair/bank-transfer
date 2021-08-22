package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
//Only work doing 1 command per version, you cannot create two tables in the same script, need to create a new version
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Create tab_user",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_user (
					user_id INT NOT NULL AUTO_INCREMENT,
					user_uuid CHAR(36) NOT NULL,
					document_number VARBINARY(16) NOT NULL,
					name VARCHAR(450) NOT NULL,
					password VARCHAR(200) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					active TINYINT(1) NOT NULL DEFAULT 1,
					PRIMARY KEY (user_id),
					UNIQUE INDEX user_id_UNIQUE (user_id ASC) VISIBLE,
					UNIQUE INDEX document_number_UNIQUE (document_number ASC) VISIBLE
				)
				ENGINE = InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
