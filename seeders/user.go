package seeders

import (
	"database/sql"
	"log"
)

func MigrateUser(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS user_data;
		create table IF NOT EXISTS user_data (
			user_id INTEGER PRIMARY KEY,
			user_name VARCHAR(150) NOT NULL,
			user_email VARCHAR(100) NOT NULL,
			user_role VARCHAR(100) NOT NULL,
			user_password VARCHAR(100) NOT NULL
		);
		
		insert into user_data (user_name, user_email, user_role, user_password) values ('Camilla Kinchlea', 'ckinchlea0@wikia.com', 'Security', 'XoNmsBIzbbae');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Dion Fazakerley', 'dfazakerley1@paginegialle.it', 'Admin', 'i2gRgmF');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Lambert Cready', 'lcready2@theatlantic.com', 'Admin', 'UNRzr9');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Pearline Winsiowiecki', 'pwinsiowiecki3@ed.gov', 'Security', 's1momnSv8');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Izzy Niche', 'iniche4@ustream.tv', 'Staff', 'Qu07OispugFs');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Sonny Mitchely', 'smitchely5@usa.gov', 'Staff', '7fiO5jWJ');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Chick Snowball', 'csnowball6@scribd.com', 'Staff', '9ieHX1qnc4VS');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Francisco Solano', 'fsolano7@liveinternet.ru', 'Staff', 'BF4PjA28');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Sax Cant', 'scant8@livejournal.com', 'Staff', 'Mccb2dXDwHc');
		insert into user_data (user_name, user_email, user_role, user_password) values ('Brigitte Tivenan', 'ariefuddinsatriadharma@gmail.com', 'staff', 'NeXoEnqcs');
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating User")
		panic(err)
	}
}
