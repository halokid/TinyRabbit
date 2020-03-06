
package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/halokid/ColorfulRabbit"
)


var cfgFile string
var endpointID int
var debug bool

var logx = Logx {
	//DebugFlag: debug,
	DebugFlag: true,
	LogFlFlag: false,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ocx",
	Short: "The missing Portainer CLI",
	Long:  `ocx lets you manage multiple Docker API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },

	/*
	Run: func(cmd *cobra.Command, args []string) {
		tog, err := cmd.Flags().GetBool("debug")
		fmt.Println("debug=", tog, " err=", err)
		if tog == true {
			logx.DebugFlag = true
		}
		debug = tog

		fmt.Println(logx)
	},
	*/
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ocx.yaml)")
	rootCmd.PersistentFlags().IntVar(&endpointID, "endpointID", 0, "ID of the endpoint to operate on (defaults to all)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "调试模式")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.Flags().BoolP("debug", "d", false, "调试模式")

	//debug, err := rootCmd.Flags().GetBool("debug")
	//fmt.Println("debug mode -------------", debug, err)

	//tog, err := rootCmd.Flags().GetBool("debug")
	//fmt.Println("debug=", tog, " err=", err)
	//if tog == true {
	//	logx.DebugFlag = true
	//}

	//fmt.Println(logx)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".barge" (without extension).
		// fixme: 在这里设置配置文件的名称
		viper.AddConfigPath(home)
		viper.SetConfigName(".tinyrabbit")
	}

	// 假如环境变量已经设置了变量的话， 就读取环境变量
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("读取配置文件:", viper.ConfigFileUsed())
	}
}
