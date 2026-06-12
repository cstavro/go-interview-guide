package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("go-interview-guide"))
	http.Handle("/", fs)

	// API endpoint to generate workspace from template
	http.HandleFunc("/api/generate", handleGenerate)

	// API endpoint to serve template content
	http.HandleFunc("/api/template", handleTemplate)

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

	// Validate template path (prevent directory traversal)
	if strings.Contains(req.Template, "..") {
		http.Error(w, "Invalid template path", http.StatusBadRequest)
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
	templateDir := filepath.Join("templates", filepath.Dir(req.Template))
	if err := copyDir(templateDir, workspaceDir); err != nil {
		http.Error(w, "Failed to copy template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := generateResponse{Path: workspaceDir, Exists: exists, Created: true}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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

	// Prevent directory traversal
	if strings.Contains(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	fullPath := filepath.Join("templates", path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(content)
}
