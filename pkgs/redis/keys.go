package redis

import (
	"Listen/pkgs"
	"fmt"
)

func EpochSubmissionCountsReceivedInSlotKey(dataMarketAddress string, slotId uint64, epochId uint64) string {
	return fmt.Sprintf("%s.%s.%d.%d", pkgs.EpochSubmissionCountsReceivedInSlotKey, dataMarketAddress, slotId, epochId)
}
