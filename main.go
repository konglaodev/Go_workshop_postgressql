package main

import (
	"os"

	"github.com/anousoneFS/go-workshop/config"
	"github.com/anousoneFS/go-workshop/internal/district"
	"github.com/anousoneFS/go-workshop/internal/province"
	"github.com/anousoneFS/go-workshop/internal/village"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func failOnError(err error, msg string) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	// load config
	cfg, err := config.LoadConfig("./")
	failOnError(err, "failed to load config")
	// gorfdlk
	dsn := "postgres://ajlrhvob:hR5QQMvvokaydTHQ5ygiSW-OkBoeFNuX@tiny.db.elephantsql.com/ajlrhvob"
	// dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Asia/Vientiane", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(province.Province{}, District{}, Village{})
	// db.Migrator().DropTable(Village{}, District{}, province.Province{})
	// fiber
	app := fiber.New()
	// api := app.Group("/api/v1")

	// endpoint: province
	// api.Get("/provinces", GetAllProvince)
	// api.Get("/provinces/:id", GetProvinceByID)
	// api.Post("/provinces", CreateProvince)
	// api.Patch("/provinces", UpdateProvince)
	// api.Delete("/provinces/:id", DeleteProvince)

	// endpoint: district
	// api.Get("/districts", GetAllDistrict)
	// api.Get("/districts/:id", GetDistrictByID)
	// api.Post("/districts", CreateDistrict)
	// api.Patch("/districts", UpdateDistrictByID)
	// api.Delete("/districts/:id", DeleteDistric)

	// endpoint: village
	// api.Get("/villages", GetAllVillage)
	// api.Get("/villages/:id", GetVillageByID)
	// api.Post("/villages", CreateVillage)
	// api.Patch("/villages", UpdateVillage)
	// api.Delete("/villages/:id", DeleteVillage)

	// province
	provinceRepo := province.NewRepository(db)
	provinceUsecase := province.NewUsecase(provinceRepo)
	province.NewHandler(provinceUsecase, app)

	// district
	districtRepo := district.NewRepository(db)
	districtUsecase := district.NewUsecase(districtRepo)
	district.NewHandlerDistrict(districtUsecase, app)

	// village
	villageRepo := village.NewRepository(db)
	villageUsecase := village.NewUsecase(villageRepo)
	village.NewHandlerDistrict(villageUsecase, app)

	// endpoint: village assignment
	app.Listen(cfg.AppPort)
}
