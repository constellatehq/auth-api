package utilities

import (
	"log"

	"github.com/sony/sonyflake"
)

var (
	Sonyflake *sonyflake.Sonyflake
)

func InitSonyflake() {
	st := sonyflake.Settings{}
	// st.MachineID = awsutil.AmazonEC2MachineID

	Sonyflake = sonyflake.NewSonyflake(st)
	if Sonyflake == nil {
		log.Fatal("failed to initialize sonyflake")
	}
}
