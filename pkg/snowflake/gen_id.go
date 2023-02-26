package snowflake

import (
    "math"
    "math/rand"
    "time"

    "github.com/sony/sonyflake"
)

// 需传入当前的机器ID
func Init() (*sonyflake.Sonyflake, error) {
    sonyMachineID := func() (uint16, error) {
        return uint16(rand.Intn(math.MaxUint16)), nil
    }
    t, _ := time.Parse("2006-01-02", "2020-01-01")
    settings := sonyflake.Settings{
        StartTime: t,
        MachineID: sonyMachineID,
    }
    sonyFlake := sonyflake.NewSonyflake(settings)
    return sonyFlake, nil
}
