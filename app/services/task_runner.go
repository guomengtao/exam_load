package services

import (
     "log"

    "github.com/robfig/cron/v3"
)

var (
    taskCron *cron.Cron
    taskIDs  = make(map[string]cron.EntryID)
)

// StartTasks initializes and starts all scheduled tasks including Redis writer, cleaner, and importer.
func StartAllTasks() {
    // Start Redis Writer task (assumed to be a long-running goroutine)
    StartRedisWriter()

    taskCron = cron.New()

    // Schedule daily cleanup of expired Redis processed data at 3 AM
    idCleaner, err := taskCron.AddFunc("0 3 * * *", func() {
        log.Println("[Cron] 清理旧 Redis processed 数据任务开始...")
        if err := CleanOldProcessedData(); err != nil {
            log.Println("[Cron] 清理任务失败:", err)
        } else {
            log.Println("[Cron] 清理任务完成")
        }
    })
    if err != nil {
        log.Println("[Cron] 添加清理任务失败:", err)
        return
    }
    taskIDs["cleaner"] = idCleaner

    // Schedule Redis import task every 5 minutes
    idImporter, err := taskCron.AddFunc("*/5 * * * *", func() {
        log.Println("[Cron] 执行 Redis 导入任务开始...")
        RunRedisImportOnce()
        log.Println("[Cron] 执行 Redis 导入任务完成")
    })
    if err != nil {
        log.Println("[Cron] 添加 Redis 导入任务失败:", err)
        return
    }
    taskIDs["importer"] = idImporter

    // Start the cron scheduler
    taskCron.Start()

    log.Println("✅ 定时任务已启动 (cron 计划中)")
}

// StopTasks stops all scheduled tasks.
func StopTasks() {
    if taskCron != nil {
        taskCron.Stop()
        log.Println("🛑 所有任务已停止")
    }
}

// StopTask stops a specific scheduled task by name.
func StopTask(name string) {
    if id, ok := taskIDs[name]; ok && taskCron != nil {
        taskCron.Remove(id)
        log.Printf("🛑 任务 %s 已停止\n", name)
    } else {
        log.Printf("⚠️ 无法停止任务（未找到）: %s\n", name)
    }
}

// RestartTask restarts a scheduled task with a new cron spec and function.
func RestartTask(name string, spec string, taskFunc func()) {
    if taskCron == nil {
        log.Println("❌ Cron 未初始化")
        return
    }

    id, err := taskCron.AddFunc(spec, taskFunc)
    if err != nil {
        log.Printf("❌ 无法重新启动任务 %s: %v\n", name, err)
        return
    }

    taskIDs[name] = id
    log.Printf("✅ 任务 %s 已重新启动\n", name)
}
