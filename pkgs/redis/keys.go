package redis

import (
	"Listen/pkgs"
	"fmt"
)

func EpochSubmissionsCount(epochId uint64) string {
	return fmt.Sprintf("%s.%d", pkgs.EpochSubmissionsCountKey, epochId)
}

func EpochSubmissionsKey(epochId uint64) string {
	return fmt.Sprintf("%s.%d", pkgs.EpochSubmissionsKey, epochId)
}

func TriggeredProcessLog(process, identifier string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.ProcessTriggerKey, process, identifier)
}

func SlotSubmissionKey(slotId, day string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.SlotSubmissionsKey, day, slotId)
}

func SlotRewardsForDay(day string, slotId string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.DailyRewardPointsKey, day, slotId)
}

func TotalSlotRewards(slotId string) string {
	return fmt.Sprintf("%s.%s", pkgs.TotalRewardPointsKey, slotId)
}
