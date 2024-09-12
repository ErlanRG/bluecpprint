package internal

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	srcPath     = "src"
	binPath     = "bin"
	includePath = "include"
)

//go:embed templates/.gitignore.tmpl
var gitignoreTmpl string

//go:embed templates/main.cpp.tmpl
var mainTmpl string

//go:embed templates/Makefile.tmpl
var makefileTmpl string

//go:embed templates/.clang-format.tmpl
var clangFormatTmpl string

//go:embed templates/README.md.tmpl
var readmeTmpl string

type Project struct {
	ProjectName  string
	AbsolutePath string
}

func (p *Project) CreateProjectStructure() error {
	// Check for AbsolutePath. Create if it does not exist
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	// Trim whitespaces from ProjectName
	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create the root project directory
	projectPath := filepath.Join(p.AbsolutePath, p.ProjectName)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		if err := os.MkdirAll(projectPath, 0o751); err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	// Create project structure (src, bin, include directories)
	if err := p.CreatePath(srcPath, projectPath); err != nil {
		log.Printf("Error creating src directory %v\n", err)
		return err
	}

	if err := p.CreatePath(includePath, projectPath); err != nil {
		log.Printf("Error creating include directory %v\n", err)
		return err
	}

	if err := p.CreatePath(binPath, projectPath); err != nil {
		log.Printf("Error creating bin directory %v\n", err)
		return err
	}

	// Initialize git repository
	if err := ExecuteCmd("git", []string{"init"}, projectPath); err != nil {
		log.Printf("Error initializing git repository %v\n", err)
		return err
	}

	// Create .gitignore file
	if err := p.CreateFileFromTemplate(gitignoreTmpl, projectPath, ".gitignore", nil); err != nil {
		log.Printf("Error creating .gitignore file: %v\n", err)
		return err
	}

	// Create main.cpp file
	mainPath := filepath.Join(projectPath, srcPath)
	if err := p.CreateFileFromTemplate(mainTmpl, mainPath, "main.cpp", nil); err != nil {
		log.Printf("Error creating main.cpp file: %v\n", err)
		return err
	}

	// Create Makefile
	if err := p.CreateFileFromTemplate(makefileTmpl, projectPath, "Makefile", nil); err != nil {
		log.Printf("Error creating Makefile: %v\n", err)
		return err
	}

	// Create clang-format file
	if err := p.CreateFileFromTemplate(clangFormatTmpl, projectPath, ".clang-format", nil); err != nil {
		log.Printf("Error creating clang-format: %v\n", err)
		return err
	}

	// Create README file
	if err := p.CreateFileFromTemplate(readmeTmpl, projectPath, "README.md", map[string]string{"ProjectName": p.ProjectName}); err != nil {
		log.Printf("Error creating readme file: %v\n", err)
		return err
	}

	// Initialize the project with a commit
	if err := ExecuteCmd("git", []string{"add", "-A"}, projectPath); err != nil {
		log.Printf("Error adding files to git repository: %v\n", err)
		return err
	}

	if err := ExecuteCmd("git", []string{"commit", "-m", "Initial commit"}, projectPath); err != nil {
		log.Printf("Error committing files to git repository: %v\n", err)
		return err
	}

	return nil
}

func (p *Project) CreateFileFromTemplate(tmplContent string, destinationPath string, filename string, data interface{}) error {
	// read template file
	tmpl, err := template.New(filename).Parse(tmplContent)
	if err != nil {
		log.Printf("Error parsing template file: %v\n", err)
		return err
	}

	// create the output file
	generatedFile, err := os.Create(filepath.Join(destinationPath, filename))
	if err != nil {
		log.Printf("Error creating %s file: %v\n", filename, err)
		return err
	}
	defer generatedFile.Close()

	err = tmpl.Execute(generatedFile, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}

	return nil
}

func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0o751); err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}
	}

	return nil
}

// [https://github.com/Melkeydev/go-blueprint/blob/main/cmd/utils/utils.go]:
// Melkeydev implementation of ExecuteCmd
func ExecuteCmd(name string, args []string, dir string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func CheckArgs() error {
	// Check if there are enough arguments
	if len(os.Args) != 2 {
		return errors.New("Usage: bluecpprint <project_name>")
	}

	return nil
}
