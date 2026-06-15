package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("go-interview-guide"))
	http.Handle("/", fs)

	// API endpoint to generate workspace from template
	http.HandleFunc("/api/generate", handleGenerate)

	// API endpoint to list existing workspace directories
	http.HandleFunc("/api/workspaces", handleWorkspaces)

	// API endpoint to serve a single template file
	http.HandleFunc("/api/template", handleTemplate)

	// API endpoint to serve all template files in a directory as tabs
	http.HandleFunc("/api/template-dir", handleTemplateDir)

	port := "8080"
	fmt.Printf("Serving go-interview-guide on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type generateRequest struct {
	Template string `json:"template"`
	Section  string `json:"section"`
	Problem  string `json:"problem"`
	Overwrite bool  `json:"overwrite"`
}

type generateResponse struct {
	Path    string `json:"path"`
	Exists  bool   `json:"exists"`
	Created bool   `json:"created"`
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req generateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Determine the source template directory. Modern clients send the directory
	// path directly; for older clients that still send a file path, derive the
	// directory from it.
	templateDir := req.Template
	if filepath.Ext(filepath.Base(templateDir)) != "" {
		templateDir = filepath.Dir(templateDir)
	}

	fullTemplateDir, err := resolveTemplatePath(templateDir)
	if err != nil {
		http.Error(w, "Invalid template path", http.StatusBadRequest)
		return
	}

	info, err := os.Stat(fullTemplateDir)
	if err != nil || !info.IsDir() {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Determine workspace directory
	workspaceDir := filepath.Join("workspaces", req.Section+"-"+req.Problem)
	exists := false

	// Check if workspace already exists
	if _, err := os.Stat(workspaceDir); err == nil {
		exists = true
		if !req.Overwrite {
			resp := generateResponse{Path: workspaceDir, Exists: true, Created: false}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	// Create workspace directory
	if err := os.MkdirAll(workspaceDir, 0755); err != nil {
		http.Error(w, "Failed to create workspace: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy template files
	if err := copyDir(fullTemplateDir, workspaceDir); err != nil {
		http.Error(w, "Failed to copy template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := generateResponse{Path: workspaceDir, Exists: exists, Created: true}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func handleWorkspaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dirs []string
	entries, err := os.ReadDir("workspaces")
	if err != nil {
		if os.IsNotExist(err) {
			dirs = []string{}
		} else {
			http.Error(w, "Failed to read workspaces: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		for _, entry := range entries {
			if entry.IsDir() {
				dirs = append(dirs, entry.Name())
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dirs)
}

func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return err
			}
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		http.Error(w, "Missing path parameter", http.StatusBadRequest)
		return
	}

	fullPath, err := resolveTemplatePath(path)
	if err != nil {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(content)
}

type templateFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type templateDirResponse struct {
	Files []templateFile `json:"files"`
}

func handleTemplateDir(w http.ResponseWriter, r *http.Request) {
	dir := r.URL.Query().Get("dir")
	if dir == "" {
		http.Error(w, "Missing dir parameter", http.StatusBadRequest)
		return
	}

	fullPath, err := resolveTemplatePath(dir)
	if err != nil {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	info, err := os.Stat(fullPath)
	if err != nil || !info.IsDir() {
		http.Error(w, "Template directory not found", http.StatusNotFound)
		return
	}

	files, err := readTemplateDirFiles(fullPath)
	if err != nil {
		http.Error(w, "Failed to read template files", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templateDirResponse{Files: files})
}

func readTemplateDirFiles(dir string) ([]templateFile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var readme *templateFile
	var goFiles []templateFile

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		path := filepath.Join(dir, name)

		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		file := templateFile{Name: name, Content: string(content)}

		switch {
		case strings.EqualFold(name, "README.md"):
			readme = &file
		case strings.HasSuffix(name, ".go"):
			goFiles = append(goFiles, file)
		}
	}

	sort.Slice(goFiles, func(i, j int) bool {
		return goFiles[i].Name < goFiles[j].Name
	})

	var result []templateFile
	if readme != nil {
		result = append(result, *readme)
	}
	result = append(result, goFiles...)
	return result, nil
}

// resolveTemplatePath validates that subpath is a relative path with no ".."
// components and that it stays within the templates directory. It returns the
// absolute filesystem path for the validated subpath.
func resolveTemplatePath(subpath string) (string, error) {
	if subpath == "" {
		return "", fmt.Errorf("missing path")
	}
	if filepath.IsAbs(subpath) {
		return "", fmt.Errorf("invalid path")
	}

	// Reject any ".." component explicitly.
	for _, part := range strings.Split(filepath.ToSlash(subpath), "/") {
		if part == ".." {
			return "", fmt.Errorf("invalid path")
		}
	}

	basePath, err := filepath.Abs("templates")
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(basePath, subpath)

	// Defense in depth: ensure the resolved path is strictly inside templates.
	rel, err := filepath.Rel(basePath, fullPath)
	if err != nil || strings.HasPrefix(rel, "..") || rel == "." {
		return "", fmt.Errorf("invalid path")
	}

	return fullPath, nil
}
