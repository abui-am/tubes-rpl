package initializers

import (
	"log"

	"gihub.com/abui-am/tubes-rpl/seeder"
)

func DBSeed() {
	log.Println("Seeding database...")
	seeder.SeedRoles(DB)
	log.Println("Database seeded!")
}
