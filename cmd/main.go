package main

import (
    "messenger/internal/bootstrap"
    "log"
)

func main() {
    if err := bootstrap.StartServer(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}