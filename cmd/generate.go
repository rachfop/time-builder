package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Template input structures
type WorkflowInput struct {
	Workflow      string
	ChildWorkflow string
	LocalActivity string
	Signal        []string
	Query         []string
	Cloud         bool
	Config        string
}

// Flags
var workflowName string
var childWorkflowName string
var localActivityName string
var signalName string
var queryName string
var signalCount int
var queryCount int
var useCloud bool
var configPath string

// generateWorkflowCmd represents the generate command for workflow
var generateWorkflowCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Python code based on templates from 'python_templates' directory.",
	Run: func(cmd *cobra.Command, args []string) {
		generateWorkflowCode(workflowName, childWorkflowName, localActivityName, signalName, signalCount, queryName, queryCount, useCloud, configPath)
	},
}

func init() {
	rootCmd.AddCommand(generateWorkflowCmd)

	// Workflow flags
	generateWorkflowCmd.Flags().StringVar(&workflowName, "workflow", "YourWorkflow", "Name of the Workflow")
	generateWorkflowCmd.Flags().StringVar(&childWorkflowName, "child-workflow", "", "Name of the Child Workflow to execute")
	generateWorkflowCmd.Flags().StringVar(&localActivityName, "local-activity", "", "Name of the Local Activity to execute")
	generateWorkflowCmd.Flags().StringVar(&signalName, "signal", "SignalName", "Name of the Signal")
	generateWorkflowCmd.Flags().StringVar(&queryName, "query", "QueryName", "Name of the Query")
	generateWorkflowCmd.Flags().IntVar(&signalCount, "signal-count", 0, "Number of signals to add")
	generateWorkflowCmd.Flags().IntVar(&queryCount, "query-count", 0, "Number of queries to add")
	generateWorkflowCmd.Flags().BoolVar(&useCloud, "cloud", false, "Flag to indicate if connecting to a cloud Temporal service.")
	generateWorkflowCmd.Flags().StringVar(&configPath, "config", "", "Path to the configuration file")

}

func generateWorkflowCode(workflowName string, childWorkflow string, localActivityName string, signalName string, signals int, queryName string, queries int, useCloud bool, configPath string) {

	if configPath != "" {
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err == nil {
			workflowName = viper.GetString("workflowName")
			childWorkflowName = viper.GetString("childWorkflowName")
			useCloud = viper.GetBool("useCloud")
		} else {
			fmt.Printf("Unable to read config file at %s, using flags or defaults. Error: %s\n", configPath, err)
		}
	}
	input := WorkflowInput{
		Workflow:      workflowName,
		ChildWorkflow: childWorkflow,
		LocalActivity: localActivityName,
		Signal:        make([]string, signals),
		Query:         make([]string, queries),
		Cloud:         useCloud,
	}
	// Adds multiple Query or Signal
	if signals == 1 {
		input.Signal[0] = signalName
	} else {
		for i := 0; i < signals; i++ {
			input.Signal[i] = fmt.Sprintf("%s%d", signalName, i+1)
		}
	}
	if queries == 1 {
		input.Query[0] = queryName
	} else {
		for i := 0; i < queries; i++ {
			input.Query[i] = fmt.Sprintf("%s%d", queryName, i+1)
		}
	}
	if useCloud {
		generateEnvFile(input)
	}

	// Read all templates from the 'python_templates' directory
	files, err := os.ReadDir("python_templates")
	if err != nil {
		fmt.Println("Error reading the directory:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".template" {
			continue
		}
		if useCloud && file.Name() == ".env.template" {
			continue
		}
		templateData, err := os.ReadFile(filepath.Join("python_templates", file.Name()))
		if err != nil {
			fmt.Println("Error reading the template file:", err)
			continue
		}

		t := template.New(file.Name())
		template.Must(t.Parse(string(templateData)))

		var buf bytes.Buffer
		if err := t.Execute(&buf, input); err != nil {
			fmt.Println("Error generating code for", file.Name(), ":", err)
			continue
		}

		// Ensure the "yourproject" directory exists
		if _, err := os.Stat("yourproject"); os.IsNotExist(err) {
			os.Mkdir("yourproject", 0755)
		}

		outputFileName := file.Name()[:len(file.Name())-len(".template")]

		switch outputFileName {
		case ".env":
			continue
			// Do nothing as the file name is already as desired.
		case "README":
			outputFileName += ".md"
		case "pyproject":
			outputFileName += ".toml"
		default:
			outputFileName += ".py"
		}
		// Modify the path to save the output inside "yourproject" directory
		outputFile, err := os.Create("yourproject/" + outputFileName)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			continue
		}
		defer outputFile.Close()

		_, err = outputFile.WriteString(buf.String())
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			continue
		}

		fmt.Println("Generated Python code for", file.Name(), "saved to:", "yourproject/"+outputFileName)

	}
}

func generateEnvFile(input WorkflowInput) {
	tmpl, err := template.New(".env.template").ParseFiles(filepath.Join("python_templates", ".env.template"))
	if err != nil {
		log.Fatalf("Error parsing .env template: %s", err)
	}
	if _, err := os.Stat("yourproject"); os.IsNotExist(err) {
		os.Mkdir("yourproject", 0755)
	}
	file, err := os.Create("yourproject/" + ".env")
	if err != nil {
		log.Fatalf("Error creating .env file: %s", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, input)
	if err != nil {
		log.Fatalf("Error writing to .env file: %s", err)
	}

}
