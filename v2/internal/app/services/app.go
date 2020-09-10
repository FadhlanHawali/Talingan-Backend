package services

import (
	_ "fmt"
	config "github.com/Talingan-Backend/v2/configs"
	"github.com/Talingan-Backend/v2/constant"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type ServicesRest struct {
	Repos *Repos
	UseCases *Usecases

	Cfg config.Configuration
}

func NewServicesRest() (*ServicesRest, error){
	var err error
	var cfg config.Configuration

	cfgAddress := os.Getenv("BITLABS_WEBINAR_CONFIG_FILE_PATH")
	if cfgAddress == "" {
		cfg, err = readConfig(constant.ServicesConfigProjectFilepath)
	}else {
		cfg, err = readConfig(cfgAddress)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	db, err := initDB(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error connect db")
	}



	app := new(ServicesRest)

	app.Cfg = cfg

	app.Repos, err = newRepos(db)
	if err != nil {
		return nil, errors.Wrap(err, "errors invoking newRepos")
	}

	app.UseCases = newUsecases(app.Repos)

	return app, nil
}



func (a *ServicesRest) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.UseCases.Close()...)

	return errs
}

func readConfig(cfgPath string) (config.Configuration, error) {

	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig();err != nil{
		return config.Configuration{}, errors.Wrapf(err, "config file not found")
	}

	var cfg config.Configuration
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Configuration{}, errors.Wrapf(err, "error reading config")
	}

	return cfg, nil
}

func initDB(cfg config.Configuration) (*gorm.DB, error) {

	dbAddress := os.Getenv("DATABASE_URL")
	if dbAddress == "" {
		dbAddress = cfg.Database.ConnectionURI
	}
	// Initialize SQL Database

	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
