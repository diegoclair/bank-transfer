package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
//Only work doing 1 command per version, you cannot create two tables in the same script, need to create a new version
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Create tab_account",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_account (
					account_id INT NOT NULL AUTO_INCREMENT,
					account_uuid CHAR(36) NOT NULL,
					cpf VARBINARY(16) NOT NULL,
					name VARCHAR(450) NOT NULL,
					secret VARCHAR(200) NOT NULL,
					balance DECIMAL(7,2) NULL DEFAULT 200.00,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					active TINYINT(1) NOT NULL DEFAULT 1,
					PRIMARY KEY (account_id),
					UNIQUE INDEX account_id_UNIQUE (account_id ASC) VISIBLE,
					UNIQUE INDEX cpf_UNIQUE (cpf ASC) VISIBLE
				)
				ENGINE = InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
