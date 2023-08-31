package config

import "os"

var QueueHost = os.Getenv("RABBIT_MQ_HOST")
var QueuePort = os.Getenv("RABBIT_MQ_PORT")
var QueueUser = os.Getenv("RABBIT_MQ_USER")
var QueuePassword = os.Getenv("RABBIT_MQ_PASSWORD")
