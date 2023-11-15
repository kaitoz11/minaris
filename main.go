package main

import (
	"fmt"
	"minaris/pkg/shooting"
	"minaris/pkg/utils/setting"
	"os"

	"github.com/spf13/viper"
)

func main() {

	// fmt.Printf("%s\n", banner)
	// cmd.Execute()

	// test1()
	load()
}

func load() {
	viper.SetConfigName("vanieva.k4it0z11")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./raws")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(viper.AllSettings())

	var config shooting.Config

	viper.Unmarshal(&config)

	fmt.Println(config.Workflows)

	shooter := shooting.NewShooterWorkflow(&config)
	shooter.UseConfig(&setting.ShootingConfig{Proxy: "http://127.0.0.1:8080"})
	shooter.UseReader(new(shooting.RawReader))
	shooter.ShootWorkflow()
}

func test1() {
	rawReader := new(shooting.RawReader)
	requests, err := shooting.ReadFromFile("./raws/vanibella/vanibella_login.k4it0z11", rawReader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gun := new(shooting.Gun)

	gun.Init("http://127.0.0.1:8080")

	// for _, r := range requests {
	// 	fmt.Println(r)
	// }

	gun.Load(requests)

	gun.Shoot()
	fmt.Println(gun.PrintFacts())
}

// func test() {
// 	// raw, err := os.ReadFile("./raws/test")
// 	rawReader := new(shooting.RawReader)

// 	raw, err := rawReader.ReadFromFile("./raws/vanieva.k4it0z11")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	gun := new(shooting.Gun)

// 	gun.Init("http://127.0.0.1:8080")
// 	// gun.Init("")
// 	gun.LoadAmmos(raw)
// 	gun.Shoot()
// 	// rawHttpString := raw["noice"]

// 	// req, err := shooting.NewFromRaw(rawHttpString)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }
// 	// fmt.Println(req)
// 	// fmt.Printf("Method: %s Path: %s Version: %s\nHost: %s\n\n", req.Method, req.Url, req.HttpVersion, req.Host)
// 	// for h := range req.Headers {
// 	// 	fmt.Println(h, req.Headers[h])
// 	// }
// 	// fmt.Printf("%s", req.Data)
// 	// models.NewFromRaw(rawHttp)
// }
