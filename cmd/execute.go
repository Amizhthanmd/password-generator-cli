package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	isUpperCase   bool
	isLowerCase   bool
	isNumbers     bool
	isSpecialChar bool
	length        int
)

var rootCmd = &cobra.Command{
	Use:   "Password Generator",
	Short: "A Password Generator CLI application",
	Long:  `This is a Password Generator CLI application built to secure your accounts with strong passwords.`,
}

func Execute() {
	rootCmd.AddCommand(generate)
	commandFlags()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func commandFlags() {
	generate.Flags().BoolVarP(&isUpperCase, "upper-case", "u", true, "Includes the upper case")
	generate.Flags().BoolVarP(&isLowerCase, "lower-case", "l", true, "Includes the lower case")
	generate.Flags().BoolVarP(&isNumbers, "numbers", "n", true, "Includes the numbers")
	generate.Flags().BoolVarP(&isSpecialChar, "special-characters", "s", true, "Includes the special characters")
	generate.Flags().IntVarP(&length, "length", "t", 12, "Length of the password")
}

var generate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a secure password with customizable length and complexity",
	Long: `The 'generate' command creates a new, secure password based on user-defined criteria.
You can specify the password length and choose to include upper and lower case letters, numbers, 
and symbols for added complexity. Ideal for creating strong passwords that 
enhance security for online accounts and applications.`,

	Run: func(cmd *cobra.Command, args []string) {
		includes := generateIncludes(isUpperCase, isLowerCase, isNumbers, isSpecialChar)
		password, err := generatePassword(length, includes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Generated Password: %s\n", password)
	},
}
