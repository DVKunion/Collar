package main

import (
	"errors"
	"fmt"
	"github.com/DVKunion/collar/pkg/utils"
	"os"
	"time"

	"github.com/DVKunion/collar/pkg/config"
	"github.com/DVKunion/collar/pkg/log"
	"github.com/DVKunion/collar/pkg/sdk"
	"github.com/spf13/cobra"
)

var (
	token   string
	auto    bool
	rootCmd = &cobra.Command{}
	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "auth rivers.chaitin.cn",
		RunE:  auth,
	}
	listCmd = &cobra.Command{
		Use:     "hosts",
		Short:   "list hosts",
		PreRunE: auth,
		RunE:    hosts,
	}

	topCmd = &cobra.Command{
		Use:     "top",
		Short:   "list host process && systemInfo",
		PreRunE: auth,
		RunE:    process,
	}

	shellCmd = &cobra.Command{
		Use:     "shell",
		Short:   "run a shell at host",
		PreRunE: auth,
		RunE:    terminal,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Collie - Collar: %s\n", config.Version)
		},
	}
)

func auth(cmd *cobra.Command, args []string) error {
	err := config.SingleConfig.Load()
	if err != nil {
		// 说明加载失败了，需要重新认证
		if token == "" {
			return errors.New("empty token")
		}
		config.SingleConfig = &config.Config{
			Token:    token,
			HostList: make([]config.Host, 0),
			UpdateAt: time.Now(),
		}
	}

	resp, errSDK := sdk.GetHostList()
	if errSDK != nil {
		return errSDK
	}

	if v, ok := resp.Body().([]config.Host); ok {
		config.SingleConfig.HostList = v
		// 说明加载的情况
		if err != nil {
			log.Infof(cmd.Context(), "success auth ! host total: %d", len(config.SingleConfig.HostList))
			log.Infof(cmd.Context(), "online: %d", len(config.SingleConfig.GetOnline()))
			log.Infof(cmd.Context(), "offline: %d", len(config.SingleConfig.HostList)-len(config.SingleConfig.GetOnline()))
		}
		return config.SingleConfig.Save()
	} else {
		return errors.New("get host list error")
	}
}

func hosts(cmd *cobra.Command, args []string) error {
	log.Info(cmd.Context(), config.SingleConfig.HostList)
	return nil
}

func process(cmd *cobra.Command, args []string) error {
	hostId := utils.Trans2HostId(args[0])
	if hostId == "" {
		return errors.New("host offline or not exists，use `collar hosts` to check host status。")
	}
	resp, err := sdk.GetProcessList(hostId)
	if err != nil {
		return err
	}
	log.Info(cmd.Context(), resp.Body())
	return nil
}

func terminal(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	if len(args) != 1 {
		return errors.New("must have at least one host arg")
	}
	log.Infof(ctx, "connecting to %s ......", args[0])
	hostId := utils.Trans2HostId(args[0])
	if hostId == "" {
		return errors.New("host offline or not exists，use `collar hosts` to check host status。")
	}
	user := ""
	if auto {
		resp, err := sdk.GetAutoLoginUser(hostId)
		if err != nil {
			return err
		}
		user = resp.Body().(string)
	}
	tm, err := sdk.GetTerminal(ctx, hostId, user)
	if err != nil {
		return err
	}
	if err = tm.Transfer(); err != nil {
		return err
	}
	return nil
}

func main() {
	authCmd.Flags().StringVarP(&token, "token", "t", "", "users token")
	shellCmd.Flags().BoolVarP(&auto, "auto", "a", false, "use auto login")

	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(topCmd)
	rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
