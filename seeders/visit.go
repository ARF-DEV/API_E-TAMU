package seeders

import (
	"database/sql"
	"log"
)

func MigrateVisit(db *sql.DB) {
	sqlStatement := `
	
		drop table IF EXISTS visit;
		create table IF NOT EXISTS visit (
			visit_id INTEGER PRIMARY KEY,
			user_visited_id INTEGER,
			guest_name VARCHAR(150) NOT NULL,
			guest_email VARCHAR(100) NOT NULL,
			visit_intention VARCHAR(100) NOT NULL,
			vaccine_certificate VARCHAR(100) NOT NULL,
			visit_date VARCHAR(10),
			visit_status VARCHAR(10) NOT NULL,
			guest_count INTEGER NOT NULL,
			visit_hour VARCHAR(10),
			transportation VARCHAR(10),
			confirmation VARCHAR(10),
			FOREIGN KEY (user_visited_id) 
				REFERENCES user_data(user_id) 
				ON DELETE SET NULL
			 
		);

	insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (6, 'Elayne Camous', 'ecamous0@surveymonkey.com', 'Displaced avulsion fracture of unspecified ischium, subsequent encounter for fracture with nonunion', 'http://dummyimage.com/114x200.png/ff4444/ffffff', 'belum datang', 1, '2022-06-21', '8:38', 'mobil', '0');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (8, 'Lennie McHan', 'lmchan1@marketwatch.com', 'Other specified injury of unspecified blood vessel of thorax, sequela', 'http://dummyimage.com/119x226.png/cc0000/ffffff', 'sedang berlangsung',4, '2022-01-29', '2:07', 'mobil', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (6, 'Carmela Stollberg', 'cstollberg2@amazon.de', 'Displaced fracture (avulsion) of lateral epicondyle of left humerus, subsequent encounter for fracture with malunion', 'http://dummyimage.com/137x116.png/cc0000/ffffff', 'sedang berlangsung',3, '2022-06-28', '16:42', 'motor', '0');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (7, 'Garv Order', 'gorder3@biglobe.ne.jp', 'Burn of unspecified degree of right toe(s) (nail), sequela', 'http://dummyimage.com/147x184.png/dddddd/000000', 'selesai',4, '2021-12-01', '22:54', 'mobil', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (8, 'Florry Gillson', 'fgillson4@nymag.com', 'Unspecified fracture of first lumbar vertebra, subsequent encounter for fracture with nonunion', 'http://dummyimage.com/207x203.png/cc0000/ffffff', 'selesai',2, '2022-09-05', '22:33', 'motor', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (7, 'Dorthy Perfili', 'dperfili5@google.com.hk', 'Infantile idiopathic scoliosis, cervicothoracic region', 'http://dummyimage.com/140x201.png/dddddd/000000', 'belum datang',4, '2022-05-03', '1:45', 'mobil', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (7, 'Moses Donati', 'mdonati6@mashable.com', 'Displaced fracture of cuboid bone of unspecified foot, initial encounter for closed fracture', 'http://dummyimage.com/164x169.png/5fa2dd/ffffff', 'belum datang',4, '2022-01-22', '0:16', 'mobil', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (9, 'Timmy Allender', 'tallender7@admin.ch', 'Nondisplaced bicondylar fracture of left tibia, subsequent encounter for open fracture type I or II with malunion', 'http://dummyimage.com/185x178.png/ff4444/ffffff', 'selesai',2, '2022-07-04', '10:07', 'motor', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (10, 'Adele Chetwind', 'achetwind8@w3.org', 'Other fractures of lower end of unspecified radius, subsequent encounter for open fracture type I or II with delayed healing', 'http://dummyimage.com/246x162.png/5fa2dd/ffffff', 'selesai',1, '2022-09-18', '14:53', 'motor', '1');
insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_date, visit_hour, transportation, confirmation) values (7, 'Odetta Symondson', 'osymondson9@bizjournals.com', 'Chronic gout due to renal impairment, unspecified wrist', 'http://dummyimage.com/186x159.png/cc0000/ffffff', 'selesai',1, '2022-04-05', '12:14', 'mobil', '1');
	
		
	`
	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Println("Error while Migrating Visit")
		panic(err)
	}
}
