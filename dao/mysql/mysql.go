package mysql

import (
    "context"
    "fmt"
    "time"

    "github.com/sony/sonyflake"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "golang-template/config"
    "golang-template/models"
)

type DB struct {
    db        *gorm.DB
    snowFlake *sonyflake.Sonyflake
}

func Init(globalConfig *config.GlobalConfig) (*DB, error) {
    cfg := globalConfig.MySQLConfig
    if cfg.Open == 0 {
        return nil, nil
    }
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
