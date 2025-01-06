package config_test

import (
	"fmt"
	"testing"

	"github.com/AssassinRobot/author/config"
	"github.com/stretchr/testify/require"
)

func TestDevelopmentConfig(t *testing.T) {
	serverPort, DBUrl := config.GetConfigs()

	fmt.Println(serverPort,DBUrl)
	require.NotEmpty(t, serverPort)
	require.NotEmpty(t, DBUrl)
}
