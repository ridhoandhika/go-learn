package config

func Get() *Config {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("err when load env %s", err.Error())
	// }

	return &Config{
		Server{
			Host: "goservice", //os.Getenv("SERVER_HOST"),
			Port: "8080",      //os.Getenv("SERVER_PORT"),
		},
		Database{
			Host:     "postgresql", //os.Getenv("DATABASE_HOST"),
			Port:     "5432",       //os.Getenv("DATABASE_PORT"),
			User:     "postgres",   //os.Getenv("DATABASE_USER"),
			Password: "postgres",   //os.Getenv("DATABASE_PASSWORD"),
			Name:     "postgres",   //os.Getenv("DATABASE_NAME"),
		},
	}
}
