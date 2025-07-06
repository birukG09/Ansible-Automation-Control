package main

import (
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

const playbookDir = "./playbooks"

func main() {
    r := gin.Default()

    // Ensure playbooks folder exists
    os.MkdirAll(playbookDir, os.ModePerm)

    r.GET("/playbooks", listPlaybooks)
    r.POST("/playbooks/upload", uploadPlaybook)
    r.POST("/playbooks/:name/run", runPlaybook)

    r.Run(":8080")
}

// List all playbooks (filenames)
func listPlaybooks(c *gin.Context) {
    files, err := os.ReadDir(playbookDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    var playbooks []string
    for _, f := range files {
        if !f.IsDir() {
            playbooks = append(playbooks, f.Name())
        }
    }
    c.JSON(http.StatusOK, gin.H{"playbooks": playbooks})
}

// Upload a playbook file
func uploadPlaybook(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
        return
    }

    // Save with UUID prefix to avoid overwrite conflicts
    uniqueName := uuid.New().String() + "-" + filepath.Base(file.Filename)
    dest := filepath.Join(playbookDir, uniqueName)

    if err := c.SaveUploadedFile(file, dest); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Uploaded", "filename": uniqueName})
}

// Run the given playbook and return output
func runPlaybook(c *gin.Context) {
    name := c.Param("name")
    playbookPath := filepath.Join(playbookDir, name)

    if _, err := os.Stat(playbookPath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Playbook not found"})
        return
    }

    // Execute ansible-playbook command
    cmd := exec.Command("ansible-playbook", playbookPath)

    output, err := cmd.CombinedOutput()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
            "output": string(output),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Playbook run successfully",
        "output": string(output),
    })
}
