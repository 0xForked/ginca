package config

func (config AppConfig) SetupAppLog() {
	//year, month, day := time.Now().Date()
	//logName := fmt.Sprintf(
	//	"%v-%d-%d-%d.log",
	//	viper.GetString(`SERVER_NAME`),
	//	year, month, day)
	// Disable Console Color.
	//gin.DisableConsoleColor()
	// Logging to a file.
	//f, _ := os.Create("/logs/" + logName)
	// write the logs to file and console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func logWithFile() {

}

func logWith3rdParty() {

}

//func TouchFile(path string, name string) {
//	if _, err := os.Stat(path + name); os.IsNotExist(err) {
//		if err := os.Chdir(path); err != nil {
//			panic(err)
//		}
//		file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
//		if err != nil {
//			panic(err)
//		}
//		if err := file.Close(); err != nil {
//			panic(err)
//		}
//	}
//}