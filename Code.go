package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "golang.org/x/crypto/openpgp"
    "golang.org/x/crypto/openpgp/armor"
)

// Configuration holds the configuration parameters for the backup tool
type Configuration struct {
    RootDir      string      // Root directory for the backup
    LoggerFormat string      // Format of the logger file in the backup directory
    Logger       *log.Logger // Logger instance
}

var (
    // Default configuration values
    defaultConfig = Configuration{
        RootDir:      "./backup",
        LoggerFormat: "2006-01-02T15:04:05",
        Logger:       log.New(os.Stdout, "", log.LstdFlags),
    }
)

func main() {
    // Define command-line flags
    backupCmd := flag.NewFlagSet("backup", flag.ExitOnError)
    shareCmd := flag.NewFlagSet("share", flag.ExitOnError)
    configCmd := flag.NewFlagSet("config", flag.ExitOnError)

    // Backup flags
    backupSrcDir := backupCmd.String("src", "", "Source directory to backup")
    encryptFlag := backupCmd.Bool("encrypt", false, "Encrypt files during backup")
    recursiveFlag := backupCmd.Bool("recursive", false, "Recursively encrypt files")
    selectiveEncrypt := backupCmd.String("selective", "", "Selectively encrypt files (comma-separated)")

    // Share flags
    shareDir := shareCmd.String("dir", "", "Directory to share")
    shareFiles := shareCmd.String("files", "", "Files to send (comma-separated)")
    sharePreviousVersions := shareCmd.Bool("prev-versions", false, "Share previous backup versions")

    // Configuration flags
    configRootDir := configCmd.String("root-dir", defaultConfig.RootDir, "Root directory for the backup")
    configLoggerFormat := configCmd.String("logger-format", defaultConfig.LoggerFormat, "Format of the logger file")

    // Parse command-line flags
    if len(os.Args) < 2 {
        fmt.Println("Please provide a subcommand: backup, share, or config")
        os.Exit(1)
    }

    // Parse command and flags
    switch os.Args[1] {
    case "backup":
        backupCmd.Parse(os.Args[2:])
    case "share":
        shareCmd.Parse(os.Args[2:])
    case "config":
        configCmd.Parse(os.Args[2:])
    default:
        fmt.Println("Invalid command. Please use 'backup', 'share', or 'config'")
        os.Exit(1)
    }

    // Load configuration
    config := Configuration{
        RootDir:      *configRootDir,
        LoggerFormat: *configLoggerFormat,
        Logger:       log.New(os.Stdout, "", log.LstdFlags),
    }

    // Execute appropriate subcommand
    switch os.Args[1] {
    case "backup":
        if *backupSrcDir == "" {
            fmt.Println("Source directory is required for backup")
            os.Exit(1)
        }
        if err := backup(*backupSrcDir, *encryptFlag, *recursiveFlag, *selectiveEncrypt, config); err != nil {
            config.Logger.Fatalf("Backup failed: %v", err)
        }
    case "share":
        if *shareDir == "" {
            fmt.Println("Directory to share is required for sharing")
            os.Exit(1)
        }
        if err := share(*shareDir, *shareFiles, *sharePreviousVersions, config); err != nil {
            config.Logger.Fatalf("Sharing failed: %v", err)
        }
    case "config":
        fmt.Println("Root Directory:", config.RootDir)
        fmt.Println("Logger Format:", config.LoggerFormat)
    }
}

func backup(srcDir string, encrypt bool, recursive bool, selectiveEncrypt string, config Configuration) error {
    // Backup logic
    // Walk through the source directory and copy files to the backup location
    return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Copy regular files
        // Implement encryption if required
        destPath := filepath.Join(config.RootDir, info.Name())
        if encrypt {
            if err := encryptFile(path, destPath); err != nil {
                return err
            }
        } else {
            srcFile, err := os.Open(path)
            if err != nil {
                return err
            }
            defer srcFile.Close()

            destFile, err := os.Create(destPath)
            if err != nil {
                return err
            }
            defer destFile.Close()

            _, err = io.Copy(destFile, srcFile)
            if err != nil {
                return err
            }
        }

        // Logging copied file
        config.Logger.Printf("Copied file: %s", path)
        return nil
    })
}

func encryptFile(srcPath, destPath string) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    // Create new OpenPGP entity for recipient
    entity, err := openpgp.NewEntity("Backup", "", "", nil)
    if err != nil {
        return err
    }

    w, err := armor.Encode(destFile, nil, nil)
    if err != nil {
        return err
    }
    defer w.Close()

    // Encrypt file
    plaintext, err := openpgp.Encrypt(w, []*openpgp.Entity{entity}, nil, nil, nil)
    if err != nil {
        return err
    }

    _, err = io.Copy(plaintext, srcFile)
    if err != nil {
        return err
    }

    err = plaintext.Close()
    if err != nil {
        return err
    }

    return nil
}

func share(shareDir, files string, prevVersions bool, config Configuration) error {
    // Sharing logic
    if prevVersions {
        // Create a new logger file for sharing previous backup versions
        logFile, err := os.Create(filepath.Join(shareDir, "backup_share.log"))
        if err != nil {
            return err
        }
        defer logFile.Close()

        // Read the original logger file
        originalLogFile, err := os.Open(filepath.Join(config.RootDir, "backup.log"))
        if err != nil {
            return err
        }
        defer originalLogFile.Close()

        // Copy the content of the original logger file to the new logger file for sharing
        _, err = io.Copy(logFile, originalLogFile)
        if err != nil {
            return err
        }

        return nil
    }

    return nil
}
