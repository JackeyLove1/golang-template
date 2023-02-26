package sqlite

import (
    "context"
    "fmt"
    "time"

    "github.com/sony/sonyflake"
    "golang-template/config"
    "golang-template/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type DB struct {
    db        *gorm.DB
    snowFlake *sonyflake.Sonyflake
}

func Init(globalConfig *config.GlobalConfig) (*DB, error) {
    cfg := globalConfig.SQLiteConfig
    if cfg.Open == 0 {
        return nil, nil
    }
    db, err := gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to db err:%v", err)
    }
    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("failed to set mysql config, err:%v", err)
    }
    sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) //空闲连接数
    sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //最大连接数
    sqlDB.SetConnMaxLifetime(time.Minute)
    return &DB{
        db: db,
    }, nil
}

func (d *DB) MigrateModel(ctx context.Context, model any) error {
    // migration
    return d.db.WithContext(ctx).AutoMigrate(&model)
}

func (d *DB) Close() {
    sqlDB, _ := d.db.DB()
    _ = sqlDB.Close()
}

// 创建用户
func (d *DB) Create(ctx context.Context, user *models.User) error {
    if user == nil {
        return fmt.Errorf("user is nil")
    }
    result := d.db.WithContext(ctx).Create(user)
    if result.Error != nil && result.RowsAffected == 0 {
        return result.Error
    }
    return nil
}
