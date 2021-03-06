package main

import (
	"github.com/changmink/shafoo/api"
	"github.com/changmink/shafoo/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadFile("./config.json")
	engine := gin.Default()

	v1 := engine.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		api.SignUp(auth)
		api.Login(auth)
		api.Logout(auth)
	}

	party := v1.Group("/parties")
	{
		api.SearchParties(party)
		api.CreateParty(party)
		api.JoinPartyById(party)
		api.LeaveParty(party)
		api.GetPartyById(party)
	}

	profile := v1.Group("/profiles")
	{
		api.GetProfileById(profile)
		api.EditProfileById(profile)
	}

	engine.Run(":" + config.C.Port)
}
