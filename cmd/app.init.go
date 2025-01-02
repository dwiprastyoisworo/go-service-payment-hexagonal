package main

//func initMongodb(cfg *config.Config) (*mongo.Database, error) {
//	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Url, cfg.MongoDB.Password)
//	client, err := mongo.Connect(options.Client().ApplyURI(url))
//	if err != nil {
//		panic("Mongodb connection error")
//	}
//	return client.Database(cfg.MongoDB.Database), nil
//}
