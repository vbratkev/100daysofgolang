// I had some success with this (via https://github.com/fatih/color) by:
//
// rootCmd.SetOutput(color.Output)
// Augmenting the UsageTemplate via something like the following:
// 	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
// 	usageTemplate := rootCmd.UsageTemplate()
// 	usageTemplate = strings.NewReplacer(
// 		`Usage:`, `{{StyleHeading "Usage:"}}`,
// 		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
// 		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
// 		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
// 		// The following one steps on "Global Flags:"
// 		// `Flags:`, `{{StyleHeading "Flags:"}}`,
// 	).Replace(usageTemplate)
// 	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
// 	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
// 	rootCmd.SetUsageTemplate(usageTemplate)
