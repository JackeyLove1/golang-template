package mysql

import (
    "encoding/hex"

    "github.com/spaolacci/murmur3"
    "golang-template/models"
    "golang-template/pkg/snowflake"
)

const seed = 20200706
const secret = "JackyFan"

func encryptPassword(data []byte) (result string) {
    h := murmur3.New32WithSeed(seed)
    h.Write([]byte(secret))
    return hex.EncodeToString(h.Sum(data))
}

func (d *DB) Register(user *models.User) error {
    // TODO
    // 检查用户是否存在
    // 生成user_id
    _, err := snowflake.GetID()
    if err != nil {
        return ErrorGenIDFailed
    }
    // 生成加密密码
    _ = encryptPassword([]byte(user.Password))
    // 把用户插入数据库
    return nil
}

func (d *DB) Login(user *models.User) error {
    // 查询数据库出错
    // 用户不存在返回错误
    // 生成加密密码与查询到的密码比较
    return nil
}

func GetUserByID(idStr string) (*models.User, error) {
    // TODO
    return nil, nil
}
