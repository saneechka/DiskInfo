package main
import (
    "fmt"
    "log"
    "syscall"
    "os"
    "path/filepath"
)

func DiskUsed(path string) {
    var statistic syscall.Statfs_t

    err := syscall.Statfs(path, &statistic)

    if err != nil {
        log.Println("Error getting information", err)
    }

    total := statistic.Blocks * uint64(statistic.Bsize)
    free := statistic.Bfree * uint64(statistic.Bsize)
    used := total - free
    percentUsed := float64(used) / float64(total) * 100
    fmt.Printf("Disk usage of %s:\n", path)
    fmt.Printf("Total: %d GB\n", total/1e9)
    fmt.Printf("Used: %d GB (%.2f%%)\n", used/1e9, percentUsed)
    fmt.Printf("Free: %d GB\n", free/1e9)
}

func getDirSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return nil
    })
    return size, err
}

func main() {
    path := "/home"
    if len(os.Args) > 1 {
        path = os.Args[1]
    }
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        log.Printf("Error: '%s' Path doesn't exist.\n", path)
        return
    } else if err != nil {
        log.Printf("Error occurred while accessing path %s: %v \n", path, err)
        return
    }
    
    DiskUsed(path)
    
    err = filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && subpath != path {
            size, err := getDirSize(subpath)
            if err != nil {
                log.Printf("Error calculating size for %s: %v\n", subpath, err)
                return nil
            }
            fmt.Printf("\nDirectory: %s\n", subpath)
            fmt.Printf("Size: %.2f GB\n", float64(size)/1e9)
        }
        return nil
    })
    
    if err != nil {
        log.Printf("Error walking through directories: %v\n", err)
    }
}





