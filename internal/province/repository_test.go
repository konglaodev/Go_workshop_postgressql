package province_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/anousoneFS/go-workshop/config"
	"github.com/anousoneFS/go-workshop/internal/province"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var DB *gorm.DB

func TestMain(m *testing.M) {
	fmt.Println("call testmain")
	cfg, err := config.LoadConfig("../..")
	failOnError(err, "failed to load config")
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Asia/Vientiane", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestGetAllProvince(t *testing.T) {
	fmt.Println("call testgetallprovince")
	repo := province.NewRepository(DB)
	i, err := repo.GetAll()
	require.NoError(t, err)
	require.NotEmpty(t, i)
	fmt.Printf("result: %+v\n", i)
}

func TestUpdateProvince(t *testing.T) {
	fmt.Println("call testupdateprovince")
	repo := province.NewRepository(DB)
	p := province.Province{Name: "hahaha"}
	err := repo.Update(p, 1)
	require.NoError(t, err)
}

func TestCreateProvince(t *testing.T) {
	repo := province.NewRepository(DB)
	province := province.Province{
		Name:   "ວຽງຈັນ2",
		NameEn: "vientiane2",
	}
	err := repo.Create(&province)
	require.NoError(t, err)

	i, err := repo.GetByID(province.ID)
	require.NoError(t, err)
	require.NotEmpty(t, i)
	require.Equal(t, i.Name, province.Name)
	require.Equal(t, i.NameEn, province.NameEn)
	require.Equal(t, i.ID, province.ID)
}
