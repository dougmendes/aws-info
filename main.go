package main

import (
	"github.com/dougmendes/aws-info/services"

	"github.com/dougmendes/aws-info/config"
)

func main() {

	iamClient := config.LoadAWSConfig("us-east-1")

	services.ListUsers(&iamClient)
}
