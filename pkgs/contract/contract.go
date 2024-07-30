// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PowerloomProtocolStateRequest is an auto generated low-level Go binding around an user-defined struct.
type PowerloomProtocolStateRequest struct {
	SlotId      *big.Int
	Deadline    *big.Int
	SnapshotCid string
	EpochId     *big.Int
	ProjectId   string
}

// PowerloomProtocolStateSlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PowerloomProtocolStateSlotInfo struct {
	SlotId                  *big.Int
	SnapshotterAddress      common.Address
	TimeSlot                *big.Int
	RewardPoints            *big.Int
	CurrentStreak           *big.Int
	CurrentStreakBonus      *big.Int
	CurrentDaySnapshotCount *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"slotsPerDay\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"externalStateAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DailyTaskCompletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DayStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"DelayedAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"snapshotterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedSnapshotSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enableEpochId\",\"type\":\"uint256\"}],\"name\":\"ProjectTypeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SequencersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"SnapshotBatchAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"snapshotterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"allSnapshottersUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DailySnapshotQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeploymentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_SIZE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOTS_PER_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USE_BLOCK_NUMBER_AS_EPOCH_ID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"internalType\":\"structPowerloomProtocolState.Request\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"acceptSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allSnapshotters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectType\",\"type\":\"string\"}],\"name\":\"allowedProjectTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"attestationFinalizedStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"attestationsReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"attestationsReceivedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"batchIdAttestationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchIdToProjects\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectids\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"checkDynamicConsensusAttestations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"checkSlotTaskStatusForDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dayCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochsInADay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"externalState\",\"outputs\":[{\"internalType\":\"contractIExternalState\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"forceCompleteConsensusAttestations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"forceSkipEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceStartDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStreak\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStreakBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentDaySnapshotCount\",\"type\":\"uint256\"}],\"internalType\":\"structPowerloomProtocolState.SlotInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotStreak\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slotId\",\"type\":\"uint256\"}],\"name\":\"getSnapshotterTimeSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSequencersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSnapshotterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastSnapshotterAddressUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dayCounter\",\"type\":\"uint256\"}],\"name\":\"loadCurrentDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"snapshotCount\",\"type\":\"uint256\"}],\"name\":\"loadSlotSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"maxAttestationFinalizedRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"maxAttestationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCidMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAttestationsForConsensus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionsForConsensus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"projectFirstEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"releaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBasePoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slotRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"slotSnapshotterMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"slotStreakCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"streak\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"slotSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"snapshotStatus\",\"outputs\":[{\"internalType\":\"enumPowerloomProtocolState.SnapshotStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshotSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"snapshotsReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"snapshotsReceivedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"snapshotsReceivedSlot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"streakBonusPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"projectIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"snapshotCids\",\"type\":\"string[]\"}],\"name\":\"submitSubmissionBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeSlotCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"timeSlotPreference\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleTimeSlotCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalSnapshotsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPowerloomProtocolState.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_projectType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateAllowedProjectType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newattestationSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateAttestationSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newbatchSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateBatchSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dailySnapshotQuota\",\"type\":\"uint256\"}],\"name\":\"updateDailySnapshotQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateEpochManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"externalStateAddress\",\"type\":\"address\"}],\"name\":\"updateExternalStateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAttestationsForConsensus\",\"type\":\"uint256\"}],\"name\":\"updateMinAttestationsForConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSubmissionsForConsensus\",\"type\":\"uint256\"}],\"name\":\"updateMinSnapshottersForConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"slotIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"submissionsList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newsnapshotSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateSnapshotSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newrewardBasePoints\",\"type\":\"uint256\"}],\"name\":\"updaterewardBasePoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newstreakBonusPoints\",\"type\":\"uint256\"}],\"name\":\"updatestreakBonusPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"internalType\":\"structPowerloomProtocolState.Request\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0xc7c7077e.
//
// Solidity: function DailySnapshotQuota() view returns(uint256)
func (_Contract *ContractCaller) DailySnapshotQuota(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DailySnapshotQuota")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0xc7c7077e.
//
// Solidity: function DailySnapshotQuota() view returns(uint256)
func (_Contract *ContractSession) DailySnapshotQuota() (*big.Int, error) {
	return _Contract.Contract.DailySnapshotQuota(&_Contract.CallOpts)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0xc7c7077e.
//
// Solidity: function DailySnapshotQuota() view returns(uint256)
func (_Contract *ContractCallerSession) DailySnapshotQuota() (*big.Int, error) {
	return _Contract.Contract.DailySnapshotQuota(&_Contract.CallOpts)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractCaller) DeploymentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DeploymentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractSession) DeploymentBlockNumber() (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xb1288f71.
//
// Solidity: function DeploymentBlockNumber() view returns(uint256)
func (_Contract *ContractCallerSession) DeploymentBlockNumber() (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractCaller) EPOCHSIZE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EPOCH_SIZE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractSession) EPOCHSIZE() (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_Contract *ContractCallerSession) EPOCHSIZE() (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts)
}

// SLOTSPERDAY is a free data retrieval call binding the contract method 0xe479b88d.
//
// Solidity: function SLOTS_PER_DAY() view returns(uint256)
func (_Contract *ContractCaller) SLOTSPERDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SLOTS_PER_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLOTSPERDAY is a free data retrieval call binding the contract method 0xe479b88d.
//
// Solidity: function SLOTS_PER_DAY() view returns(uint256)
func (_Contract *ContractSession) SLOTSPERDAY() (*big.Int, error) {
	return _Contract.Contract.SLOTSPERDAY(&_Contract.CallOpts)
}

// SLOTSPERDAY is a free data retrieval call binding the contract method 0xe479b88d.
//
// Solidity: function SLOTS_PER_DAY() view returns(uint256)
func (_Contract *ContractCallerSession) SLOTSPERDAY() (*big.Int, error) {
	return _Contract.Contract.SLOTSPERDAY(&_Contract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINBLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINID() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractCaller) USEBLOCKNUMBERASEPOCHID(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "USE_BLOCK_NUMBER_AS_EPOCH_ID")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_Contract *ContractCallerSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts)
}

// AcceptSnapshot is a free data retrieval call binding the contract method 0x154a6456.
//
// Solidity: function acceptSnapshot(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature) view returns(bool)
func (_Contract *ContractCaller) AcceptSnapshot(opts *bind.CallOpts, slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "acceptSnapshot", slotId, snapshotCid, epochId, projectId, request, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptSnapshot is a free data retrieval call binding the contract method 0x154a6456.
//
// Solidity: function acceptSnapshot(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature) view returns(bool)
func (_Contract *ContractSession) AcceptSnapshot(slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte) (bool, error) {
	return _Contract.Contract.AcceptSnapshot(&_Contract.CallOpts, slotId, snapshotCid, epochId, projectId, request, signature)
}

// AcceptSnapshot is a free data retrieval call binding the contract method 0x154a6456.
//
// Solidity: function acceptSnapshot(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature) view returns(bool)
func (_Contract *ContractCallerSession) AcceptSnapshot(slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte) (bool, error) {
	return _Contract.Contract.AcceptSnapshot(&_Contract.CallOpts, slotId, snapshotCid, epochId, projectId, request, signature)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractCaller) AllSnapshotters(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allSnapshotters", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractSession) AllSnapshotters(addr common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, addr)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractCallerSession) AllSnapshotters(addr common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, addr)
}

// AllowedProjectTypes is a free data retrieval call binding the contract method 0x04d76d59.
//
// Solidity: function allowedProjectTypes(string projectType) view returns(bool)
func (_Contract *ContractCaller) AllowedProjectTypes(opts *bind.CallOpts, projectType string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowedProjectTypes", projectType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedProjectTypes is a free data retrieval call binding the contract method 0x04d76d59.
//
// Solidity: function allowedProjectTypes(string projectType) view returns(bool)
func (_Contract *ContractSession) AllowedProjectTypes(projectType string) (bool, error) {
	return _Contract.Contract.AllowedProjectTypes(&_Contract.CallOpts, projectType)
}

// AllowedProjectTypes is a free data retrieval call binding the contract method 0x04d76d59.
//
// Solidity: function allowedProjectTypes(string projectType) view returns(bool)
func (_Contract *ContractCallerSession) AllowedProjectTypes(projectType string) (bool, error) {
	return _Contract.Contract.AllowedProjectTypes(&_Contract.CallOpts, projectType)
}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractCaller) AttestationFinalizedStatus(opts *bind.CallOpts, epochId *big.Int, projectId string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationFinalizedStatus", epochId, projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractSession) AttestationFinalizedStatus(epochId *big.Int, projectId string) (bool, error) {
	return _Contract.Contract.AttestationFinalizedStatus(&_Contract.CallOpts, epochId, projectId)
}

// AttestationFinalizedStatus is a free data retrieval call binding the contract method 0xbfefe303.
//
// Solidity: function attestationFinalizedStatus(uint256 epochId, string projectId) view returns(bool finalized)
func (_Contract *ContractCallerSession) AttestationFinalizedStatus(epochId *big.Int, projectId string) (bool, error) {
	return _Contract.Contract.AttestationFinalizedStatus(&_Contract.CallOpts, epochId, projectId)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_Contract *ContractCaller) AttestationSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_Contract *ContractSession) AttestationSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.AttestationSubmissionWindow(&_Contract.CallOpts)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_Contract *ContractCallerSession) AttestationSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.AttestationSubmissionWindow(&_Contract.CallOpts)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractCaller) AttestationsReceived(opts *bind.CallOpts, batchId *big.Int, validatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceived", batchId, validatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractSession) AttestationsReceived(batchId *big.Int, validatorAddr common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, batchId, validatorAddr)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x33420a0c.
//
// Solidity: function attestationsReceived(uint256 batchId, address validatorAddr) view returns(bool)
func (_Contract *ContractCallerSession) AttestationsReceived(batchId *big.Int, validatorAddr common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, batchId, validatorAddr)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractCaller) AttestationsReceivedCount(opts *bind.CallOpts, batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceivedCount", batchId, rootHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractSession) AttestationsReceivedCount(batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, batchId, rootHash)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x66f94346.
//
// Solidity: function attestationsReceivedCount(uint256 batchId, bytes32 rootHash) view returns(uint256 count)
func (_Contract *ContractCallerSession) AttestationsReceivedCount(batchId *big.Int, rootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, batchId, rootHash)
}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractCaller) BatchIdAttestationStatus(opts *bind.CallOpts, batchId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchIdAttestationStatus", batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractSession) BatchIdAttestationStatus(batchId *big.Int) (bool, error) {
	return _Contract.Contract.BatchIdAttestationStatus(&_Contract.CallOpts, batchId)
}

// BatchIdAttestationStatus is a free data retrieval call binding the contract method 0xa047977d.
//
// Solidity: function batchIdAttestationStatus(uint256 batchId) view returns(bool status)
func (_Contract *ContractCallerSession) BatchIdAttestationStatus(batchId *big.Int) (bool, error) {
	return _Contract.Contract.BatchIdAttestationStatus(&_Contract.CallOpts, batchId)
}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractCaller) BatchIdToProjects(opts *bind.CallOpts, batchId *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchIdToProjects", batchId, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractSession) BatchIdToProjects(batchId *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.BatchIdToProjects(&_Contract.CallOpts, batchId, arg1)
}

// BatchIdToProjects is a free data retrieval call binding the contract method 0x510c9f4c.
//
// Solidity: function batchIdToProjects(uint256 batchId, uint256 ) view returns(string projectids)
func (_Contract *ContractCallerSession) BatchIdToProjects(batchId *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.BatchIdToProjects(&_Contract.CallOpts, batchId, arg1)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_Contract *ContractCaller) BatchSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_Contract *ContractSession) BatchSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.BatchSubmissionWindow(&_Contract.CallOpts)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_Contract *ContractCallerSession) BatchSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.BatchSubmissionWindow(&_Contract.CallOpts)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0xac23e9ff.
//
// Solidity: function checkDynamicConsensusAttestations(uint256 batchId, uint256 epochId) view returns(bool)
func (_Contract *ContractCaller) CheckDynamicConsensusAttestations(opts *bind.CallOpts, batchId *big.Int, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkDynamicConsensusAttestations", batchId, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0xac23e9ff.
//
// Solidity: function checkDynamicConsensusAttestations(uint256 batchId, uint256 epochId) view returns(bool)
func (_Contract *ContractSession) CheckDynamicConsensusAttestations(batchId *big.Int, epochId *big.Int) (bool, error) {
	return _Contract.Contract.CheckDynamicConsensusAttestations(&_Contract.CallOpts, batchId, epochId)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0xac23e9ff.
//
// Solidity: function checkDynamicConsensusAttestations(uint256 batchId, uint256 epochId) view returns(bool)
func (_Contract *ContractCallerSession) CheckDynamicConsensusAttestations(batchId *big.Int, epochId *big.Int) (bool, error) {
	return _Contract.Contract.CheckDynamicConsensusAttestations(&_Contract.CallOpts, batchId, epochId)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractCaller) CheckSlotTaskStatusForDay(opts *bind.CallOpts, slotId *big.Int, day *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkSlotTaskStatusForDay", slotId, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractSession) CheckSlotTaskStatusForDay(slotId *big.Int, day *big.Int) (bool, error) {
	return _Contract.Contract.CheckSlotTaskStatusForDay(&_Contract.CallOpts, slotId, day)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractCallerSession) CheckSlotTaskStatusForDay(slotId *big.Int, day *big.Int) (bool, error) {
	return _Contract.Contract.CheckSlotTaskStatusForDay(&_Contract.CallOpts, slotId, day)
}

// CurrentBatchId is a free data retrieval call binding the contract method 0x0a763da1.
//
// Solidity: function currentBatchId() view returns(uint256)
func (_Contract *ContractCaller) CurrentBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBatchId is a free data retrieval call binding the contract method 0x0a763da1.
//
// Solidity: function currentBatchId() view returns(uint256)
func (_Contract *ContractSession) CurrentBatchId() (*big.Int, error) {
	return _Contract.Contract.CurrentBatchId(&_Contract.CallOpts)
}

// CurrentBatchId is a free data retrieval call binding the contract method 0x0a763da1.
//
// Solidity: function currentBatchId() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentBatchId() (*big.Int, error) {
	return _Contract.Contract.CurrentBatchId(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch")

	outstruct := new(struct {
		Begin   *big.Int
		End     *big.Int
		EpochId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCallerSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_Contract *ContractCaller) DayCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dayCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_Contract *ContractSession) DayCounter() (*big.Int, error) {
	return _Contract.Contract.DayCounter(&_Contract.CallOpts)
}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_Contract *ContractCallerSession) DayCounter() (*big.Int, error) {
	return _Contract.Contract.DayCounter(&_Contract.CallOpts)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCaller) EpochInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochInfo", arg0)

	outstruct := new(struct {
		Timestamp   *big.Int
		Blocknumber *big.Int
		EpochEnd    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Blocknumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, arg0)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCallerSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, arg0)
}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_Contract *ContractCaller) EpochsInADay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochsInADay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_Contract *ContractSession) EpochsInADay() (*big.Int, error) {
	return _Contract.Contract.EpochsInADay(&_Contract.CallOpts)
}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_Contract *ContractCallerSession) EpochsInADay() (*big.Int, error) {
	return _Contract.Contract.EpochsInADay(&_Contract.CallOpts)
}

// ExternalState is a free data retrieval call binding the contract method 0x1b9e4b56.
//
// Solidity: function externalState() view returns(address)
func (_Contract *ContractCaller) ExternalState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "externalState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExternalState is a free data retrieval call binding the contract method 0x1b9e4b56.
//
// Solidity: function externalState() view returns(address)
func (_Contract *ContractSession) ExternalState() (common.Address, error) {
	return _Contract.Contract.ExternalState(&_Contract.CallOpts)
}

// ExternalState is a free data retrieval call binding the contract method 0x1b9e4b56.
//
// Solidity: function externalState() view returns(address)
func (_Contract *ContractCallerSession) ExternalState() (common.Address, error) {
	return _Contract.Contract.ExternalState(&_Contract.CallOpts)
}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_Contract *ContractCaller) GetEpochManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_Contract *ContractSession) GetEpochManager() (common.Address, error) {
	return _Contract.Contract.GetEpochManager(&_Contract.CallOpts)
}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_Contract *ContractCallerSession) GetEpochManager() (common.Address, error) {
	return _Contract.Contract.GetEpochManager(&_Contract.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_Contract *ContractCaller) GetSequencers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSequencers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_Contract *ContractSession) GetSequencers() ([]common.Address, error) {
	return _Contract.Contract.GetSequencers(&_Contract.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_Contract *ContractCallerSession) GetSequencers() ([]common.Address, error) {
	return _Contract.Contract.GetSequencers(&_Contract.CallOpts)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Contract *ContractCaller) GetSlotInfo(opts *bind.CallOpts, slotId *big.Int) (PowerloomProtocolStateSlotInfo, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSlotInfo", slotId)

	if err != nil {
		return *new(PowerloomProtocolStateSlotInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PowerloomProtocolStateSlotInfo)).(*PowerloomProtocolStateSlotInfo)

	return out0, err

}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Contract *ContractSession) GetSlotInfo(slotId *big.Int) (PowerloomProtocolStateSlotInfo, error) {
	return _Contract.Contract.GetSlotInfo(&_Contract.CallOpts, slotId)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Contract *ContractCallerSession) GetSlotInfo(slotId *big.Int) (PowerloomProtocolStateSlotInfo, error) {
	return _Contract.Contract.GetSlotInfo(&_Contract.CallOpts, slotId)
}

// GetSlotStreak is a free data retrieval call binding the contract method 0x2252c7b8.
//
// Solidity: function getSlotStreak(uint256 slotId) view returns(uint256)
func (_Contract *ContractCaller) GetSlotStreak(opts *bind.CallOpts, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSlotStreak", slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlotStreak is a free data retrieval call binding the contract method 0x2252c7b8.
//
// Solidity: function getSlotStreak(uint256 slotId) view returns(uint256)
func (_Contract *ContractSession) GetSlotStreak(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSlotStreak(&_Contract.CallOpts, slotId)
}

// GetSlotStreak is a free data retrieval call binding the contract method 0x2252c7b8.
//
// Solidity: function getSlotStreak(uint256 slotId) view returns(uint256)
func (_Contract *ContractCallerSession) GetSlotStreak(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSlotStreak(&_Contract.CallOpts, slotId)
}

// GetSnapshotterTimeSlot is a free data retrieval call binding the contract method 0x2db9ace2.
//
// Solidity: function getSnapshotterTimeSlot(uint256 _slotId) view returns(uint256)
func (_Contract *ContractCaller) GetSnapshotterTimeSlot(opts *bind.CallOpts, _slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSnapshotterTimeSlot", _slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSnapshotterTimeSlot is a free data retrieval call binding the contract method 0x2db9ace2.
//
// Solidity: function getSnapshotterTimeSlot(uint256 _slotId) view returns(uint256)
func (_Contract *ContractSession) GetSnapshotterTimeSlot(_slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSnapshotterTimeSlot(&_Contract.CallOpts, _slotId)
}

// GetSnapshotterTimeSlot is a free data retrieval call binding the contract method 0x2db9ace2.
//
// Solidity: function getSnapshotterTimeSlot(uint256 _slotId) view returns(uint256)
func (_Contract *ContractCallerSession) GetSnapshotterTimeSlot(_slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSnapshotterTimeSlot(&_Contract.CallOpts, _slotId)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalSequencersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalSequencersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalSequencersCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSequencersCount(&_Contract.CallOpts)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalSequencersCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSequencersCount(&_Contract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalSnapshotterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalSnapshotterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Contract *ContractCallerSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCaller) LastFinalizedSnapshot(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastFinalizedSnapshot", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, projectId)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCallerSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, projectId)
}

// LastSnapshotterAddressUpdate is a free data retrieval call binding the contract method 0x25d6ad01.
//
// Solidity: function lastSnapshotterAddressUpdate(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) LastSnapshotterAddressUpdate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastSnapshotterAddressUpdate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSnapshotterAddressUpdate is a free data retrieval call binding the contract method 0x25d6ad01.
//
// Solidity: function lastSnapshotterAddressUpdate(uint256 ) view returns(uint256)
func (_Contract *ContractSession) LastSnapshotterAddressUpdate(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.LastSnapshotterAddressUpdate(&_Contract.CallOpts, arg0)
}

// LastSnapshotterAddressUpdate is a free data retrieval call binding the contract method 0x25d6ad01.
//
// Solidity: function lastSnapshotterAddressUpdate(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) LastSnapshotterAddressUpdate(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.LastSnapshotterAddressUpdate(&_Contract.CallOpts, arg0)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractCaller) MaxAttestationFinalizedRootHash(opts *bind.CallOpts, batchId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationFinalizedRootHash", batchId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractSession) MaxAttestationFinalizedRootHash(batchId *big.Int) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, batchId)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xcc542228.
//
// Solidity: function maxAttestationFinalizedRootHash(uint256 batchId) view returns(bytes32 rootHash)
func (_Contract *ContractCallerSession) MaxAttestationFinalizedRootHash(batchId *big.Int) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, batchId)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractCaller) MaxAttestationsCount(opts *bind.CallOpts, batchId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationsCount", batchId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractSession) MaxAttestationsCount(batchId *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, batchId)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x30c8ecb2.
//
// Solidity: function maxAttestationsCount(uint256 batchId) view returns(uint256 count)
func (_Contract *ContractCallerSession) MaxAttestationsCount(batchId *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, batchId)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractCaller) MaxSnapshotsCid(opts *bind.CallOpts, projectId string, epochId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxSnapshotsCid", projectId, epochId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractSession) MaxSnapshotsCid(projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, projectId, epochId)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractCallerSession) MaxSnapshotsCid(projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, projectId, epochId)
}

// MaxSnapshotsCidMap is a free data retrieval call binding the contract method 0x80671835.
//
// Solidity: function maxSnapshotsCidMap(string projectId, uint256 epochId) view returns(string snapshotCid)
func (_Contract *ContractCaller) MaxSnapshotsCidMap(opts *bind.CallOpts, projectId string, epochId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxSnapshotsCidMap", projectId, epochId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MaxSnapshotsCidMap is a free data retrieval call binding the contract method 0x80671835.
//
// Solidity: function maxSnapshotsCidMap(string projectId, uint256 epochId) view returns(string snapshotCid)
func (_Contract *ContractSession) MaxSnapshotsCidMap(projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCidMap(&_Contract.CallOpts, projectId, epochId)
}

// MaxSnapshotsCidMap is a free data retrieval call binding the contract method 0x80671835.
//
// Solidity: function maxSnapshotsCidMap(string projectId, uint256 epochId) view returns(string snapshotCid)
func (_Contract *ContractCallerSession) MaxSnapshotsCidMap(projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCidMap(&_Contract.CallOpts, projectId, epochId)
}

// MaxSnapshotsCount is a free data retrieval call binding the contract method 0xdbb54182.
//
// Solidity: function maxSnapshotsCount(string , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) MaxSnapshotsCount(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxSnapshotsCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSnapshotsCount is a free data retrieval call binding the contract method 0xdbb54182.
//
// Solidity: function maxSnapshotsCount(string , uint256 ) view returns(uint256)
func (_Contract *ContractSession) MaxSnapshotsCount(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxSnapshotsCount(&_Contract.CallOpts, arg0, arg1)
}

// MaxSnapshotsCount is a free data retrieval call binding the contract method 0xdbb54182.
//
// Solidity: function maxSnapshotsCount(string , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) MaxSnapshotsCount(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.MaxSnapshotsCount(&_Contract.CallOpts, arg0, arg1)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_Contract *ContractCaller) MinAttestationsForConsensus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minAttestationsForConsensus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_Contract *ContractSession) MinAttestationsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinAttestationsForConsensus(&_Contract.CallOpts)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_Contract *ContractCallerSession) MinAttestationsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinAttestationsForConsensus(&_Contract.CallOpts)
}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractCaller) MinSubmissionsForConsensus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minSubmissionsForConsensus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractSession) MinSubmissionsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinSubmissionsForConsensus(&_Contract.CallOpts)
}

// MinSubmissionsForConsensus is a free data retrieval call binding the contract method 0x66752e1e.
//
// Solidity: function minSubmissionsForConsensus() view returns(uint256)
func (_Contract *ContractCallerSession) MinSubmissionsForConsensus() (*big.Int, error) {
	return _Contract.Contract.MinSubmissionsForConsensus(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCaller) ProjectFirstEpochId(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "projectFirstEpochId", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, projectId)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_Contract *ContractCallerSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, projectId)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractCaller) RecoverAddress(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "recoverAddress", messageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractSession) RecoverAddress(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Contract.Contract.RecoverAddress(&_Contract.CallOpts, messageHash, signature)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 messageHash, bytes signature) pure returns(address)
func (_Contract *ContractCallerSession) RecoverAddress(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _Contract.Contract.RecoverAddress(&_Contract.CallOpts, messageHash, signature)
}

// RewardBasePoints is a free data retrieval call binding the contract method 0x7e9834d7.
//
// Solidity: function rewardBasePoints() view returns(uint256)
func (_Contract *ContractCaller) RewardBasePoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardBasePoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBasePoints is a free data retrieval call binding the contract method 0x7e9834d7.
//
// Solidity: function rewardBasePoints() view returns(uint256)
func (_Contract *ContractSession) RewardBasePoints() (*big.Int, error) {
	return _Contract.Contract.RewardBasePoints(&_Contract.CallOpts)
}

// RewardBasePoints is a free data retrieval call binding the contract method 0x7e9834d7.
//
// Solidity: function rewardBasePoints() view returns(uint256)
func (_Contract *ContractCallerSession) RewardBasePoints() (*big.Int, error) {
	return _Contract.Contract.RewardBasePoints(&_Contract.CallOpts)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_Contract *ContractCaller) RewardsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_Contract *ContractSession) RewardsEnabled() (bool, error) {
	return _Contract.Contract.RewardsEnabled(&_Contract.CallOpts)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_Contract *ContractCallerSession) RewardsEnabled() (bool, error) {
	return _Contract.Contract.RewardsEnabled(&_Contract.CallOpts)
}

// SlotCounter is a free data retrieval call binding the contract method 0xe59a4105.
//
// Solidity: function slotCounter() view returns(uint256)
func (_Contract *ContractCaller) SlotCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotCounter is a free data retrieval call binding the contract method 0xe59a4105.
//
// Solidity: function slotCounter() view returns(uint256)
func (_Contract *ContractSession) SlotCounter() (*big.Int, error) {
	return _Contract.Contract.SlotCounter(&_Contract.CallOpts)
}

// SlotCounter is a free data retrieval call binding the contract method 0xe59a4105.
//
// Solidity: function slotCounter() view returns(uint256)
func (_Contract *ContractCallerSession) SlotCounter() (*big.Int, error) {
	return _Contract.Contract.SlotCounter(&_Contract.CallOpts)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SlotRewardPoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotRewardPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SlotRewardPoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewardPoints(&_Contract.CallOpts, arg0)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SlotRewardPoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewardPoints(&_Contract.CallOpts, arg0)
}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractCaller) SlotSnapshotterMapping(opts *bind.CallOpts, slotId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotSnapshotterMapping", slotId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractSession) SlotSnapshotterMapping(slotId *big.Int) (common.Address, error) {
	return _Contract.Contract.SlotSnapshotterMapping(&_Contract.CallOpts, slotId)
}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractCallerSession) SlotSnapshotterMapping(slotId *big.Int) (common.Address, error) {
	return _Contract.Contract.SlotSnapshotterMapping(&_Contract.CallOpts, slotId)
}

// SlotStreakCounter is a free data retrieval call binding the contract method 0x7c5f1557.
//
// Solidity: function slotStreakCounter(uint256 slotId) view returns(uint256 streak)
func (_Contract *ContractCaller) SlotStreakCounter(opts *bind.CallOpts, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotStreakCounter", slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotStreakCounter is a free data retrieval call binding the contract method 0x7c5f1557.
//
// Solidity: function slotStreakCounter(uint256 slotId) view returns(uint256 streak)
func (_Contract *ContractSession) SlotStreakCounter(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotStreakCounter(&_Contract.CallOpts, slotId)
}

// SlotStreakCounter is a free data retrieval call binding the contract method 0x7c5f1557.
//
// Solidity: function slotStreakCounter(uint256 slotId) view returns(uint256 streak)
func (_Contract *ContractCallerSession) SlotStreakCounter(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotStreakCounter(&_Contract.CallOpts, slotId)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 day) view returns(uint256 count)
func (_Contract *ContractCaller) SlotSubmissionCount(opts *bind.CallOpts, slotId *big.Int, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotSubmissionCount", slotId, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 day) view returns(uint256 count)
func (_Contract *ContractSession) SlotSubmissionCount(slotId *big.Int, day *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotSubmissionCount(&_Contract.CallOpts, slotId, day)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 day) view returns(uint256 count)
func (_Contract *ContractCallerSession) SlotSubmissionCount(slotId *big.Int, day *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotSubmissionCount(&_Contract.CallOpts, slotId, day)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCaller) SnapshotStatus(opts *bind.CallOpts, projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotStatus", projectId, epochId)

	outstruct := new(struct {
		Status      uint8
		SnapshotCid string
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SnapshotCid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractSession) SnapshotStatus(projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, projectId, epochId)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCallerSession) SnapshotStatus(projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, projectId, epochId)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractCaller) SnapshotSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_Contract *ContractCallerSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts)
}

// SnapshotsReceived is a free data retrieval call binding the contract method 0x04dbb717.
//
// Solidity: function snapshotsReceived(string , uint256 , address ) view returns(bool)
func (_Contract *ContractCaller) SnapshotsReceived(opts *bind.CallOpts, arg0 string, arg1 *big.Int, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotsReceived", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SnapshotsReceived is a free data retrieval call binding the contract method 0x04dbb717.
//
// Solidity: function snapshotsReceived(string , uint256 , address ) view returns(bool)
func (_Contract *ContractSession) SnapshotsReceived(arg0 string, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _Contract.Contract.SnapshotsReceived(&_Contract.CallOpts, arg0, arg1, arg2)
}

// SnapshotsReceived is a free data retrieval call binding the contract method 0x04dbb717.
//
// Solidity: function snapshotsReceived(string , uint256 , address ) view returns(bool)
func (_Contract *ContractCallerSession) SnapshotsReceived(arg0 string, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _Contract.Contract.SnapshotsReceived(&_Contract.CallOpts, arg0, arg1, arg2)
}

// SnapshotsReceivedCount is a free data retrieval call binding the contract method 0x8e07ba42.
//
// Solidity: function snapshotsReceivedCount(string , uint256 , string ) view returns(uint256)
func (_Contract *ContractCaller) SnapshotsReceivedCount(opts *bind.CallOpts, arg0 string, arg1 *big.Int, arg2 string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotsReceivedCount", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SnapshotsReceivedCount is a free data retrieval call binding the contract method 0x8e07ba42.
//
// Solidity: function snapshotsReceivedCount(string , uint256 , string ) view returns(uint256)
func (_Contract *ContractSession) SnapshotsReceivedCount(arg0 string, arg1 *big.Int, arg2 string) (*big.Int, error) {
	return _Contract.Contract.SnapshotsReceivedCount(&_Contract.CallOpts, arg0, arg1, arg2)
}

// SnapshotsReceivedCount is a free data retrieval call binding the contract method 0x8e07ba42.
//
// Solidity: function snapshotsReceivedCount(string , uint256 , string ) view returns(uint256)
func (_Contract *ContractCallerSession) SnapshotsReceivedCount(arg0 string, arg1 *big.Int, arg2 string) (*big.Int, error) {
	return _Contract.Contract.SnapshotsReceivedCount(&_Contract.CallOpts, arg0, arg1, arg2)
}

// SnapshotsReceivedSlot is a free data retrieval call binding the contract method 0x37d080f0.
//
// Solidity: function snapshotsReceivedSlot(string , uint256 , uint256 ) view returns(bool)
func (_Contract *ContractCaller) SnapshotsReceivedSlot(opts *bind.CallOpts, arg0 string, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotsReceivedSlot", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SnapshotsReceivedSlot is a free data retrieval call binding the contract method 0x37d080f0.
//
// Solidity: function snapshotsReceivedSlot(string , uint256 , uint256 ) view returns(bool)
func (_Contract *ContractSession) SnapshotsReceivedSlot(arg0 string, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Contract.Contract.SnapshotsReceivedSlot(&_Contract.CallOpts, arg0, arg1, arg2)
}

// SnapshotsReceivedSlot is a free data retrieval call binding the contract method 0x37d080f0.
//
// Solidity: function snapshotsReceivedSlot(string , uint256 , uint256 ) view returns(bool)
func (_Contract *ContractCallerSession) SnapshotsReceivedSlot(arg0 string, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Contract.Contract.SnapshotsReceivedSlot(&_Contract.CallOpts, arg0, arg1, arg2)
}

// StreakBonusPoints is a free data retrieval call binding the contract method 0x36f71682.
//
// Solidity: function streakBonusPoints() view returns(uint256)
func (_Contract *ContractCaller) StreakBonusPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "streakBonusPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StreakBonusPoints is a free data retrieval call binding the contract method 0x36f71682.
//
// Solidity: function streakBonusPoints() view returns(uint256)
func (_Contract *ContractSession) StreakBonusPoints() (*big.Int, error) {
	return _Contract.Contract.StreakBonusPoints(&_Contract.CallOpts)
}

// StreakBonusPoints is a free data retrieval call binding the contract method 0x36f71682.
//
// Solidity: function streakBonusPoints() view returns(uint256)
func (_Contract *ContractCallerSession) StreakBonusPoints() (*big.Int, error) {
	return _Contract.Contract.StreakBonusPoints(&_Contract.CallOpts)
}

// TimeSlotCheck is a free data retrieval call binding the contract method 0x1a0adc24.
//
// Solidity: function timeSlotCheck() view returns(bool)
func (_Contract *ContractCaller) TimeSlotCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "timeSlotCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimeSlotCheck is a free data retrieval call binding the contract method 0x1a0adc24.
//
// Solidity: function timeSlotCheck() view returns(bool)
func (_Contract *ContractSession) TimeSlotCheck() (bool, error) {
	return _Contract.Contract.TimeSlotCheck(&_Contract.CallOpts)
}

// TimeSlotCheck is a free data retrieval call binding the contract method 0x1a0adc24.
//
// Solidity: function timeSlotCheck() view returns(bool)
func (_Contract *ContractCallerSession) TimeSlotCheck() (bool, error) {
	return _Contract.Contract.TimeSlotCheck(&_Contract.CallOpts)
}

// TimeSlotPreference is a free data retrieval call binding the contract method 0x1cdd886b.
//
// Solidity: function timeSlotPreference(uint256 slot) view returns(uint256)
func (_Contract *ContractCaller) TimeSlotPreference(opts *bind.CallOpts, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "timeSlotPreference", slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeSlotPreference is a free data retrieval call binding the contract method 0x1cdd886b.
//
// Solidity: function timeSlotPreference(uint256 slot) view returns(uint256)
func (_Contract *ContractSession) TimeSlotPreference(slot *big.Int) (*big.Int, error) {
	return _Contract.Contract.TimeSlotPreference(&_Contract.CallOpts, slot)
}

// TimeSlotPreference is a free data retrieval call binding the contract method 0x1cdd886b.
//
// Solidity: function timeSlotPreference(uint256 slot) view returns(uint256)
func (_Contract *ContractCallerSession) TimeSlotPreference(slot *big.Int) (*big.Int, error) {
	return _Contract.Contract.TimeSlotPreference(&_Contract.CallOpts, slot)
}

// TotalSnapshotsReceived is a free data retrieval call binding the contract method 0x797bbd58.
//
// Solidity: function totalSnapshotsReceived(string , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) TotalSnapshotsReceived(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSnapshotsReceived", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSnapshotsReceived is a free data retrieval call binding the contract method 0x797bbd58.
//
// Solidity: function totalSnapshotsReceived(string , uint256 ) view returns(uint256)
func (_Contract *ContractSession) TotalSnapshotsReceived(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TotalSnapshotsReceived(&_Contract.CallOpts, arg0, arg1)
}

// TotalSnapshotsReceived is a free data retrieval call binding the contract method 0x797bbd58.
//
// Solidity: function totalSnapshotsReceived(string , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) TotalSnapshotsReceived(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TotalSnapshotsReceived(&_Contract.CallOpts, arg0, arg1)
}

// Verify is a free data retrieval call binding the contract method 0x58939f83.
//
// Solidity: function verify(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractCaller) Verify(opts *bind.CallOpts, slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verify", slotId, snapshotCid, epochId, projectId, request, signature, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x58939f83.
//
// Solidity: function verify(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractSession) Verify(slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, slotId, snapshotCid, epochId, projectId, request, signature, signer)
}

// Verify is a free data retrieval call binding the contract method 0x58939f83.
//
// Solidity: function verify(uint256 slotId, string snapshotCid, uint256 epochId, string projectId, (uint256,uint256,string,uint256,string) request, bytes signature, address signer) view returns(bool)
func (_Contract *ContractCallerSession) Verify(slotId *big.Int, snapshotCid string, epochId *big.Int, projectId string, request PowerloomProtocolStateRequest, signature []byte, signer common.Address) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, slotId, snapshotCid, epochId, projectId, request, signature, signer)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x21a9b3f9.
//
// Solidity: function forceCompleteConsensusAttestations(uint256 batchId, uint256 epochId) returns()
func (_Contract *ContractTransactor) ForceCompleteConsensusAttestations(opts *bind.TransactOpts, batchId *big.Int, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceCompleteConsensusAttestations", batchId, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x21a9b3f9.
//
// Solidity: function forceCompleteConsensusAttestations(uint256 batchId, uint256 epochId) returns()
func (_Contract *ContractSession) ForceCompleteConsensusAttestations(batchId *big.Int, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceCompleteConsensusAttestations(&_Contract.TransactOpts, batchId, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x21a9b3f9.
//
// Solidity: function forceCompleteConsensusAttestations(uint256 batchId, uint256 epochId) returns()
func (_Contract *ContractTransactorSession) ForceCompleteConsensusAttestations(batchId *big.Int, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceCompleteConsensusAttestations(&_Contract.TransactOpts, batchId, epochId)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ForceSkipEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceSkipEpoch", begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, begin, end)
}

// ForceStartDay is a paid mutator transaction binding the contract method 0xcd529926.
//
// Solidity: function forceStartDay() returns()
func (_Contract *ContractTransactor) ForceStartDay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceStartDay")
}

// ForceStartDay is a paid mutator transaction binding the contract method 0xcd529926.
//
// Solidity: function forceStartDay() returns()
func (_Contract *ContractSession) ForceStartDay() (*types.Transaction, error) {
	return _Contract.Contract.ForceStartDay(&_Contract.TransactOpts)
}

// ForceStartDay is a paid mutator transaction binding the contract method 0xcd529926.
//
// Solidity: function forceStartDay() returns()
func (_Contract *ContractTransactorSession) ForceStartDay() (*types.Transaction, error) {
	return _Contract.Contract.ForceStartDay(&_Contract.TransactOpts)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_Contract *ContractTransactor) LoadCurrentDay(opts *bind.TransactOpts, _dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "loadCurrentDay", _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_Contract *ContractSession) LoadCurrentDay(_dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadCurrentDay(&_Contract.TransactOpts, _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_Contract *ContractTransactorSession) LoadCurrentDay(_dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadCurrentDay(&_Contract.TransactOpts, _dayCounter)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractTransactor) LoadSlotSubmissions(opts *bind.TransactOpts, slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "loadSlotSubmissions", slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractSession) LoadSlotSubmissions(slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadSlotSubmissions(&_Contract.TransactOpts, slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractTransactorSession) LoadSlotSubmissions(slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadSlotSubmissions(&_Contract.TransactOpts, slotId, dayId, snapshotCount)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ReleaseEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "releaseEpoch", begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, begin, end)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactor) SubmitBatchAttestation(opts *bind.TransactOpts, batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitBatchAttestation", batchId, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractSession) SubmitBatchAttestation(batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, batchId, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x70dfe736.
//
// Solidity: function submitBatchAttestation(uint256 batchId, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactorSession) SubmitBatchAttestation(batchId *big.Int, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, batchId, epochId, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractTransactor) SubmitSubmissionBatch(opts *bind.TransactOpts, batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitSubmissionBatch", batchCid, batchId, epochId, projectIds, snapshotCids)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractSession) SubmitSubmissionBatch(batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, batchCid, batchId, epochId, projectIds, snapshotCids)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xca467e62.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 batchId, uint256 epochId, string[] projectIds, string[] snapshotCids) returns()
func (_Contract *ContractTransactorSession) SubmitSubmissionBatch(batchCid string, batchId *big.Int, epochId *big.Int, projectIds []string, snapshotCids []string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, batchCid, batchId, epochId, projectIds, snapshotCids)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_Contract *ContractTransactor) ToggleRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "toggleRewards")
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_Contract *ContractSession) ToggleRewards() (*types.Transaction, error) {
	return _Contract.Contract.ToggleRewards(&_Contract.TransactOpts)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_Contract *ContractTransactorSession) ToggleRewards() (*types.Transaction, error) {
	return _Contract.Contract.ToggleRewards(&_Contract.TransactOpts)
}

// ToggleTimeSlotCheck is a paid mutator transaction binding the contract method 0x75f2d951.
//
// Solidity: function toggleTimeSlotCheck() returns()
func (_Contract *ContractTransactor) ToggleTimeSlotCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "toggleTimeSlotCheck")
}

// ToggleTimeSlotCheck is a paid mutator transaction binding the contract method 0x75f2d951.
//
// Solidity: function toggleTimeSlotCheck() returns()
func (_Contract *ContractSession) ToggleTimeSlotCheck() (*types.Transaction, error) {
	return _Contract.Contract.ToggleTimeSlotCheck(&_Contract.TransactOpts)
}

// ToggleTimeSlotCheck is a paid mutator transaction binding the contract method 0x75f2d951.
//
// Solidity: function toggleTimeSlotCheck() returns()
func (_Contract *ContractTransactorSession) ToggleTimeSlotCheck() (*types.Transaction, error) {
	return _Contract.Contract.ToggleTimeSlotCheck(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractTransactor) UpdateAddresses(opts *bind.TransactOpts, role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAddresses", role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractSession) UpdateAddresses(role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAddresses(&_Contract.TransactOpts, role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractTransactorSession) UpdateAddresses(role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAddresses(&_Contract.TransactOpts, role, _addresses, _status)
}

// UpdateAllowedProjectType is a paid mutator transaction binding the contract method 0x562209f3.
//
// Solidity: function updateAllowedProjectType(string _projectType, bool _status) returns()
func (_Contract *ContractTransactor) UpdateAllowedProjectType(opts *bind.TransactOpts, _projectType string, _status bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAllowedProjectType", _projectType, _status)
}

// UpdateAllowedProjectType is a paid mutator transaction binding the contract method 0x562209f3.
//
// Solidity: function updateAllowedProjectType(string _projectType, bool _status) returns()
func (_Contract *ContractSession) UpdateAllowedProjectType(_projectType string, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAllowedProjectType(&_Contract.TransactOpts, _projectType, _status)
}

// UpdateAllowedProjectType is a paid mutator transaction binding the contract method 0x562209f3.
//
// Solidity: function updateAllowedProjectType(string _projectType, bool _status) returns()
func (_Contract *ContractTransactorSession) UpdateAllowedProjectType(_projectType string, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAllowedProjectType(&_Contract.TransactOpts, _projectType, _status)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateAttestationSubmissionWindow(opts *bind.TransactOpts, newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAttestationSubmissionWindow", newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateAttestationSubmissionWindow(newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttestationSubmissionWindow(&_Contract.TransactOpts, newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateAttestationSubmissionWindow(newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttestationSubmissionWindow(&_Contract.TransactOpts, newattestationSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateBatchSubmissionWindow(opts *bind.TransactOpts, newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateBatchSubmissionWindow", newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateBatchSubmissionWindow(newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBatchSubmissionWindow(&_Contract.TransactOpts, newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateBatchSubmissionWindow(newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBatchSubmissionWindow(&_Contract.TransactOpts, newbatchSubmissionWindow)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 dailySnapshotQuota) returns()
func (_Contract *ContractTransactor) UpdateDailySnapshotQuota(opts *bind.TransactOpts, dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDailySnapshotQuota", dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 dailySnapshotQuota) returns()
func (_Contract *ContractSession) UpdateDailySnapshotQuota(dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDailySnapshotQuota(&_Contract.TransactOpts, dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 dailySnapshotQuota) returns()
func (_Contract *ContractTransactorSession) UpdateDailySnapshotQuota(dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDailySnapshotQuota(&_Contract.TransactOpts, dailySnapshotQuota)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_Contract *ContractTransactor) UpdateEpochManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateEpochManager", _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_Contract *ContractSession) UpdateEpochManager(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateEpochManager(&_Contract.TransactOpts, _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_Contract *ContractTransactorSession) UpdateEpochManager(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateEpochManager(&_Contract.TransactOpts, _address)
}

// UpdateExternalStateAddress is a paid mutator transaction binding the contract method 0x8bf08485.
//
// Solidity: function updateExternalStateAddress(address externalStateAddress) returns()
func (_Contract *ContractTransactor) UpdateExternalStateAddress(opts *bind.TransactOpts, externalStateAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateExternalStateAddress", externalStateAddress)
}

// UpdateExternalStateAddress is a paid mutator transaction binding the contract method 0x8bf08485.
//
// Solidity: function updateExternalStateAddress(address externalStateAddress) returns()
func (_Contract *ContractSession) UpdateExternalStateAddress(externalStateAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExternalStateAddress(&_Contract.TransactOpts, externalStateAddress)
}

// UpdateExternalStateAddress is a paid mutator transaction binding the contract method 0x8bf08485.
//
// Solidity: function updateExternalStateAddress(address externalStateAddress) returns()
func (_Contract *ContractTransactorSession) UpdateExternalStateAddress(externalStateAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateExternalStateAddress(&_Contract.TransactOpts, externalStateAddress)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractTransactor) UpdateMinAttestationsForConsensus(opts *bind.TransactOpts, _minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMinAttestationsForConsensus", _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractSession) UpdateMinAttestationsForConsensus(_minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinAttestationsForConsensus(&_Contract.TransactOpts, _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractTransactorSession) UpdateMinAttestationsForConsensus(_minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinAttestationsForConsensus(&_Contract.TransactOpts, _minAttestationsForConsensus)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractTransactor) UpdateMinSnapshottersForConsensus(opts *bind.TransactOpts, _minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMinSnapshottersForConsensus", _minSubmissionsForConsensus)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractSession) UpdateMinSnapshottersForConsensus(_minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinSnapshottersForConsensus(&_Contract.TransactOpts, _minSubmissionsForConsensus)
}

// UpdateMinSnapshottersForConsensus is a paid mutator transaction binding the contract method 0x38deacd3.
//
// Solidity: function updateMinSnapshottersForConsensus(uint256 _minSubmissionsForConsensus) returns()
func (_Contract *ContractTransactorSession) UpdateMinSnapshottersForConsensus(_minSubmissionsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinSnapshottersForConsensus(&_Contract.TransactOpts, _minSubmissionsForConsensus)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x153b56c2.
//
// Solidity: function updateRewards(uint256[] slotIds, uint256[] submissionsList, uint256 day) returns()
func (_Contract *ContractTransactor) UpdateRewards(opts *bind.TransactOpts, slotIds []*big.Int, submissionsList []*big.Int, day *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRewards", slotIds, submissionsList, day)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x153b56c2.
//
// Solidity: function updateRewards(uint256[] slotIds, uint256[] submissionsList, uint256 day) returns()
func (_Contract *ContractSession) UpdateRewards(slotIds []*big.Int, submissionsList []*big.Int, day *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewards(&_Contract.TransactOpts, slotIds, submissionsList, day)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x153b56c2.
//
// Solidity: function updateRewards(uint256[] slotIds, uint256[] submissionsList, uint256 day) returns()
func (_Contract *ContractTransactorSession) UpdateRewards(slotIds []*big.Int, submissionsList []*big.Int, day *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewards(&_Contract.TransactOpts, slotIds, submissionsList, day)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateSnapshotSubmissionWindow(opts *bind.TransactOpts, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSnapshotSubmissionWindow", newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, newsnapshotSubmissionWindow)
}

// UpdaterewardBasePoints is a paid mutator transaction binding the contract method 0x801939d7.
//
// Solidity: function updaterewardBasePoints(uint256 newrewardBasePoints) returns()
func (_Contract *ContractTransactor) UpdaterewardBasePoints(opts *bind.TransactOpts, newrewardBasePoints *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updaterewardBasePoints", newrewardBasePoints)
}

// UpdaterewardBasePoints is a paid mutator transaction binding the contract method 0x801939d7.
//
// Solidity: function updaterewardBasePoints(uint256 newrewardBasePoints) returns()
func (_Contract *ContractSession) UpdaterewardBasePoints(newrewardBasePoints *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdaterewardBasePoints(&_Contract.TransactOpts, newrewardBasePoints)
}

// UpdaterewardBasePoints is a paid mutator transaction binding the contract method 0x801939d7.
//
// Solidity: function updaterewardBasePoints(uint256 newrewardBasePoints) returns()
func (_Contract *ContractTransactorSession) UpdaterewardBasePoints(newrewardBasePoints *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdaterewardBasePoints(&_Contract.TransactOpts, newrewardBasePoints)
}

// UpdatestreakBonusPoints is a paid mutator transaction binding the contract method 0xa8939fa5.
//
// Solidity: function updatestreakBonusPoints(uint256 newstreakBonusPoints) returns()
func (_Contract *ContractTransactor) UpdatestreakBonusPoints(opts *bind.TransactOpts, newstreakBonusPoints *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updatestreakBonusPoints", newstreakBonusPoints)
}

// UpdatestreakBonusPoints is a paid mutator transaction binding the contract method 0xa8939fa5.
//
// Solidity: function updatestreakBonusPoints(uint256 newstreakBonusPoints) returns()
func (_Contract *ContractSession) UpdatestreakBonusPoints(newstreakBonusPoints *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdatestreakBonusPoints(&_Contract.TransactOpts, newstreakBonusPoints)
}

// UpdatestreakBonusPoints is a paid mutator transaction binding the contract method 0xa8939fa5.
//
// Solidity: function updatestreakBonusPoints(uint256 newstreakBonusPoints) returns()
func (_Contract *ContractTransactorSession) UpdatestreakBonusPoints(newstreakBonusPoints *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdatestreakBonusPoints(&_Contract.TransactOpts, newstreakBonusPoints)
}

// ContractDailyTaskCompletedEventIterator is returned from FilterDailyTaskCompletedEvent and is used to iterate over the raw logs and unpacked data for DailyTaskCompletedEvent events raised by the Contract contract.
type ContractDailyTaskCompletedEventIterator struct {
	Event *ContractDailyTaskCompletedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDailyTaskCompletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDailyTaskCompletedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDailyTaskCompletedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDailyTaskCompletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDailyTaskCompletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDailyTaskCompletedEvent represents a DailyTaskCompletedEvent event raised by the Contract contract.
type ContractDailyTaskCompletedEvent struct {
	SnapshotterAddress common.Address
	SlotId             *big.Int
	DayId              *big.Int
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDailyTaskCompletedEvent is a free log retrieval operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDailyTaskCompletedEvent(opts *bind.FilterOpts) (*ContractDailyTaskCompletedEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DailyTaskCompletedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractDailyTaskCompletedEventIterator{contract: _Contract.contract, event: "DailyTaskCompletedEvent", logs: logs, sub: sub}, nil
}

// WatchDailyTaskCompletedEvent is a free log subscription operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDailyTaskCompletedEvent(opts *bind.WatchOpts, sink chan<- *ContractDailyTaskCompletedEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DailyTaskCompletedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDailyTaskCompletedEvent)
				if err := _Contract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyTaskCompletedEvent is a log parse operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDailyTaskCompletedEvent(log types.Log) (*ContractDailyTaskCompletedEvent, error) {
	event := new(ContractDailyTaskCompletedEvent)
	if err := _Contract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDayStartedEventIterator is returned from FilterDayStartedEvent and is used to iterate over the raw logs and unpacked data for DayStartedEvent events raised by the Contract contract.
type ContractDayStartedEventIterator struct {
	Event *ContractDayStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDayStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDayStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDayStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDayStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDayStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDayStartedEvent represents a DayStartedEvent event raised by the Contract contract.
type ContractDayStartedEvent struct {
	DayId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDayStartedEvent is a free log retrieval operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDayStartedEvent(opts *bind.FilterOpts) (*ContractDayStartedEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DayStartedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractDayStartedEventIterator{contract: _Contract.contract, event: "DayStartedEvent", logs: logs, sub: sub}, nil
}

// WatchDayStartedEvent is a free log subscription operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDayStartedEvent(opts *bind.WatchOpts, sink chan<- *ContractDayStartedEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DayStartedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDayStartedEvent)
				if err := _Contract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDayStartedEvent is a log parse operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDayStartedEvent(log types.Log) (*ContractDayStartedEvent, error) {
	event := new(ContractDayStartedEvent)
	if err := _Contract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedAttestationSubmittedIterator is returned from FilterDelayedAttestationSubmitted and is used to iterate over the raw logs and unpacked data for DelayedAttestationSubmitted events raised by the Contract contract.
type ContractDelayedAttestationSubmittedIterator struct {
	Event *ContractDelayedAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedAttestationSubmitted represents a DelayedAttestationSubmitted event raised by the Contract contract.
type ContractDelayedAttestationSubmitted struct {
	BatchId       *big.Int
	EpochId       *big.Int
	Timestamp     *big.Int
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedAttestationSubmitted is a free log retrieval operation binding the contract event 0xc8ab56999aec6b2fb5794b180b189d9d51ffdaa8b6041cd424c9e941f6832dc7.
//
// Solidity: event DelayedAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) FilterDelayedAttestationSubmitted(opts *bind.FilterOpts, epochId []*big.Int, validatorAddr []common.Address) (*ContractDelayedAttestationSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedAttestationSubmittedIterator{contract: _Contract.contract, event: "DelayedAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedAttestationSubmitted is a free log subscription operation binding the contract event 0xc8ab56999aec6b2fb5794b180b189d9d51ffdaa8b6041cd424c9e941f6832dc7.
//
// Solidity: event DelayedAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) WatchDelayedAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedAttestationSubmitted, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedAttestationSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedAttestationSubmitted is a log parse operation binding the contract event 0xc8ab56999aec6b2fb5794b180b189d9d51ffdaa8b6041cd424c9e941f6832dc7.
//
// Solidity: event DelayedAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) ParseDelayedAttestationSubmitted(log types.Log) (*ContractDelayedAttestationSubmitted, error) {
	event := new(ContractDelayedAttestationSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedBatchSubmittedIterator is returned from FilterDelayedBatchSubmitted and is used to iterate over the raw logs and unpacked data for DelayedBatchSubmitted events raised by the Contract contract.
type ContractDelayedBatchSubmittedIterator struct {
	Event *ContractDelayedBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedBatchSubmitted represents a DelayedBatchSubmitted event raised by the Contract contract.
type ContractDelayedBatchSubmitted struct {
	BatchId   *big.Int
	BatchCid  string
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedBatchSubmitted is a free log retrieval operation binding the contract event 0xb11947be3643fd6e9120edf6f6f24c12f6b7a383503d1143fac6908f0cd1df07.
//
// Solidity: event DelayedBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDelayedBatchSubmitted(opts *bind.FilterOpts, epochId []*big.Int) (*ContractDelayedBatchSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedBatchSubmittedIterator{contract: _Contract.contract, event: "DelayedBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedBatchSubmitted is a free log subscription operation binding the contract event 0xb11947be3643fd6e9120edf6f6f24c12f6b7a383503d1143fac6908f0cd1df07.
//
// Solidity: event DelayedBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDelayedBatchSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedBatchSubmitted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedBatchSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedBatchSubmitted is a log parse operation binding the contract event 0xb11947be3643fd6e9120edf6f6f24c12f6b7a383503d1143fac6908f0cd1df07.
//
// Solidity: event DelayedBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDelayedBatchSubmitted(log types.Log) (*ContractDelayedBatchSubmitted, error) {
	event := new(ContractDelayedBatchSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedSnapshotSubmittedIterator is returned from FilterDelayedSnapshotSubmitted and is used to iterate over the raw logs and unpacked data for DelayedSnapshotSubmitted events raised by the Contract contract.
type ContractDelayedSnapshotSubmittedIterator struct {
	Event *ContractDelayedSnapshotSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedSnapshotSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedSnapshotSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedSnapshotSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedSnapshotSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedSnapshotSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedSnapshotSubmitted represents a DelayedSnapshotSubmitted event raised by the Contract contract.
type ContractDelayedSnapshotSubmitted struct {
	SnapshotterAddr common.Address
	SlotId          *big.Int
	SnapshotCid     string
	EpochId         *big.Int
	ProjectId       string
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelayedSnapshotSubmitted is a free log retrieval operation binding the contract event 0x8eb09d3de43d012d958db78fcc692d37e9f9c0cc2af1cdd8c58617832af17d31.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDelayedSnapshotSubmitted(opts *bind.FilterOpts, snapshotterAddr []common.Address, epochId []*big.Int) (*ContractDelayedSnapshotSubmittedIterator, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedSnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedSnapshotSubmittedIterator{contract: _Contract.contract, event: "DelayedSnapshotSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedSnapshotSubmitted is a free log subscription operation binding the contract event 0x8eb09d3de43d012d958db78fcc692d37e9f9c0cc2af1cdd8c58617832af17d31.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDelayedSnapshotSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedSnapshotSubmitted, snapshotterAddr []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedSnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedSnapshotSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedSnapshotSubmitted is a log parse operation binding the contract event 0x8eb09d3de43d012d958db78fcc692d37e9f9c0cc2af1cdd8c58617832af17d31.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDelayedSnapshotSubmitted(log types.Log) (*ContractDelayedSnapshotSubmitted, error) {
	event := new(ContractDelayedSnapshotSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEpochReleasedIterator is returned from FilterEpochReleased and is used to iterate over the raw logs and unpacked data for EpochReleased events raised by the Contract contract.
type ContractEpochReleasedIterator struct {
	Event *ContractEpochReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEpochReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEpochReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEpochReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEpochReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEpochReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEpochReleased represents a EpochReleased event raised by the Contract contract.
type ContractEpochReleased struct {
	EpochId   *big.Int
	Begin     *big.Int
	End       *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEpochReleased is a free log retrieval operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) FilterEpochReleased(opts *bind.FilterOpts, epochId []*big.Int) (*ContractEpochReleasedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEpochReleasedIterator{contract: _Contract.contract, event: "EpochReleased", logs: logs, sub: sub}, nil
}

// WatchEpochReleased is a free log subscription operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) WatchEpochReleased(opts *bind.WatchOpts, sink chan<- *ContractEpochReleased, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEpochReleased)
				if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochReleased is a log parse operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) ParseEpochReleased(log types.Log) (*ContractEpochReleased, error) {
	event := new(ContractEpochReleased)
	if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProjectTypeUpdatedIterator is returned from FilterProjectTypeUpdated and is used to iterate over the raw logs and unpacked data for ProjectTypeUpdated events raised by the Contract contract.
type ContractProjectTypeUpdatedIterator struct {
	Event *ContractProjectTypeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProjectTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProjectTypeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProjectTypeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProjectTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProjectTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProjectTypeUpdated represents a ProjectTypeUpdated event raised by the Contract contract.
type ContractProjectTypeUpdated struct {
	ProjectType   string
	Allowed       bool
	EnableEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProjectTypeUpdated is a free log retrieval operation binding the contract event 0xcd1b466e0de4e12ecf0bc286450d3e5b4aa88db70272ed5c11508b16acb35bfc.
//
// Solidity: event ProjectTypeUpdated(string projectType, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) FilterProjectTypeUpdated(opts *bind.FilterOpts) (*ContractProjectTypeUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProjectTypeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractProjectTypeUpdatedIterator{contract: _Contract.contract, event: "ProjectTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectTypeUpdated is a free log subscription operation binding the contract event 0xcd1b466e0de4e12ecf0bc286450d3e5b4aa88db70272ed5c11508b16acb35bfc.
//
// Solidity: event ProjectTypeUpdated(string projectType, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) WatchProjectTypeUpdated(opts *bind.WatchOpts, sink chan<- *ContractProjectTypeUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProjectTypeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProjectTypeUpdated)
				if err := _Contract.contract.UnpackLog(event, "ProjectTypeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectTypeUpdated is a log parse operation binding the contract event 0xcd1b466e0de4e12ecf0bc286450d3e5b4aa88db70272ed5c11508b16acb35bfc.
//
// Solidity: event ProjectTypeUpdated(string projectType, bool allowed, uint256 enableEpochId)
func (_Contract *ContractFilterer) ParseProjectTypeUpdated(log types.Log) (*ContractProjectTypeUpdated, error) {
	event := new(ContractProjectTypeUpdated)
	if err := _Contract.contract.UnpackLog(event, "ProjectTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSequencersUpdatedIterator is returned from FilterSequencersUpdated and is used to iterate over the raw logs and unpacked data for SequencersUpdated events raised by the Contract contract.
type ContractSequencersUpdatedIterator struct {
	Event *ContractSequencersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSequencersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSequencersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSequencersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSequencersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSequencersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSequencersUpdated represents a SequencersUpdated event raised by the Contract contract.
type ContractSequencersUpdated struct {
	SequencerAddress common.Address
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSequencersUpdated is a free log retrieval operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) FilterSequencersUpdated(opts *bind.FilterOpts) (*ContractSequencersUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SequencersUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSequencersUpdatedIterator{contract: _Contract.contract, event: "SequencersUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencersUpdated is a free log subscription operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) WatchSequencersUpdated(opts *bind.WatchOpts, sink chan<- *ContractSequencersUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SequencersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSequencersUpdated)
				if err := _Contract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencersUpdated is a log parse operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) ParseSequencersUpdated(log types.Log) (*ContractSequencersUpdated, error) {
	event := new(ContractSequencersUpdated)
	if err := _Contract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchAttestationSubmittedIterator is returned from FilterSnapshotBatchAttestationSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchAttestationSubmitted events raised by the Contract contract.
type ContractSnapshotBatchAttestationSubmittedIterator struct {
	Event *ContractSnapshotBatchAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchAttestationSubmitted represents a SnapshotBatchAttestationSubmitted event raised by the Contract contract.
type ContractSnapshotBatchAttestationSubmitted struct {
	BatchId       *big.Int
	EpochId       *big.Int
	Timestamp     *big.Int
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchAttestationSubmitted is a free log retrieval operation binding the contract event 0xae8d67af6c2cb7ba8591211569b509740c04331b67af5d4839cf25d3b0208604.
//
// Solidity: event SnapshotBatchAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) FilterSnapshotBatchAttestationSubmitted(opts *bind.FilterOpts, epochId []*big.Int, validatorAddr []common.Address) (*ContractSnapshotBatchAttestationSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchAttestationSubmittedIterator{contract: _Contract.contract, event: "SnapshotBatchAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchAttestationSubmitted is a free log subscription operation binding the contract event 0xae8d67af6c2cb7ba8591211569b509740c04331b67af5d4839cf25d3b0208604.
//
// Solidity: event SnapshotBatchAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) WatchSnapshotBatchAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchAttestationSubmitted, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchAttestationSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchAttestationSubmitted is a log parse operation binding the contract event 0xae8d67af6c2cb7ba8591211569b509740c04331b67af5d4839cf25d3b0208604.
//
// Solidity: event SnapshotBatchAttestationSubmitted(uint256 batchId, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) ParseSnapshotBatchAttestationSubmitted(log types.Log) (*ContractSnapshotBatchAttestationSubmitted, error) {
	event := new(ContractSnapshotBatchAttestationSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchFinalizedIterator is returned from FilterSnapshotBatchFinalized and is used to iterate over the raw logs and unpacked data for SnapshotBatchFinalized events raised by the Contract contract.
type ContractSnapshotBatchFinalizedIterator struct {
	Event *ContractSnapshotBatchFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchFinalized represents a SnapshotBatchFinalized event raised by the Contract contract.
type ContractSnapshotBatchFinalized struct {
	EpochId   *big.Int
	BatchId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchFinalized is a free log retrieval operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchFinalized(opts *bind.FilterOpts, epochId []*big.Int, batchId []*big.Int) (*ContractSnapshotBatchFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchFinalized", epochIdRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchFinalizedIterator{contract: _Contract.contract, event: "SnapshotBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchFinalized is a free log subscription operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchFinalized(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchFinalized, epochId []*big.Int, batchId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchFinalized", epochIdRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchFinalized)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchFinalized is a log parse operation binding the contract event 0x5234f54aa94bd8e44799076dbbe59404d669eb94833cfbff03360715b9820e40.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, uint256 indexed batchId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchFinalized(log types.Log) (*ContractSnapshotBatchFinalized, error) {
	event := new(ContractSnapshotBatchFinalized)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchSubmittedIterator is returned from FilterSnapshotBatchSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchSubmitted events raised by the Contract contract.
type ContractSnapshotBatchSubmittedIterator struct {
	Event *ContractSnapshotBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchSubmitted represents a SnapshotBatchSubmitted event raised by the Contract contract.
type ContractSnapshotBatchSubmitted struct {
	BatchId   *big.Int
	BatchCid  string
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchSubmitted is a free log retrieval operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchSubmitted(opts *bind.FilterOpts, epochId []*big.Int) (*ContractSnapshotBatchSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchSubmittedIterator{contract: _Contract.contract, event: "SnapshotBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchSubmitted is a free log subscription operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchSubmitted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchSubmitted is a log parse operation binding the contract event 0x71f7e37d775c89547d94f23aae240bec29afd96657dc34975abbdd326e8e6d5d.
//
// Solidity: event SnapshotBatchSubmitted(uint256 batchId, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchSubmitted(log types.Log) (*ContractSnapshotBatchSubmitted, error) {
	event := new(ContractSnapshotBatchSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotFinalizedIterator is returned from FilterSnapshotFinalized and is used to iterate over the raw logs and unpacked data for SnapshotFinalized events raised by the Contract contract.
type ContractSnapshotFinalizedIterator struct {
	Event *ContractSnapshotFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotFinalized represents a SnapshotFinalized event raised by the Contract contract.
type ContractSnapshotFinalized struct {
	EpochId     *big.Int
	EpochEnd    *big.Int
	ProjectId   string
	SnapshotCid string
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSnapshotFinalized is a free log retrieval operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotFinalized(opts *bind.FilterOpts, epochId []*big.Int) (*ContractSnapshotFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotFinalizedIterator{contract: _Contract.contract, event: "SnapshotFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotFinalized is a free log subscription operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotFinalized(opts *bind.WatchOpts, sink chan<- *ContractSnapshotFinalized, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotFinalized)
				if err := _Contract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotFinalized is a log parse operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotFinalized(log types.Log) (*ContractSnapshotFinalized, error) {
	event := new(ContractSnapshotFinalized)
	if err := _Contract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotSubmittedIterator is returned from FilterSnapshotSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotSubmitted events raised by the Contract contract.
type ContractSnapshotSubmittedIterator struct {
	Event *ContractSnapshotSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotSubmitted represents a SnapshotSubmitted event raised by the Contract contract.
type ContractSnapshotSubmitted struct {
	SnapshotterAddr common.Address
	SlotId          *big.Int
	SnapshotCid     string
	EpochId         *big.Int
	ProjectId       string
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSnapshotSubmitted is a free log retrieval operation binding the contract event 0x8d9341fa1766ade9f55ddb87e37a11648afefce0f76a389675dd56a5d555b8d3.
//
// Solidity: event SnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotSubmitted(opts *bind.FilterOpts, snapshotterAddr []common.Address, epochId []*big.Int) (*ContractSnapshotSubmittedIterator, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotSubmittedIterator{contract: _Contract.contract, event: "SnapshotSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotSubmitted is a free log subscription operation binding the contract event 0x8d9341fa1766ade9f55ddb87e37a11648afefce0f76a389675dd56a5d555b8d3.
//
// Solidity: event SnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotSubmitted, snapshotterAddr []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotSubmitted", snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotSubmitted is a log parse operation binding the contract event 0x8d9341fa1766ade9f55ddb87e37a11648afefce0f76a389675dd56a5d555b8d3.
//
// Solidity: event SnapshotSubmitted(address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotSubmitted(log types.Log) (*ContractSnapshotSubmitted, error) {
	event := new(ContractSnapshotSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the Contract contract.
type ContractValidatorsUpdatedIterator struct {
	Event *ContractValidatorsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorsUpdated represents a ValidatorsUpdated event raised by the Contract contract.
type ContractValidatorsUpdated struct {
	ValidatorAddress common.Address
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*ContractValidatorsUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorsUpdatedIterator{contract: _Contract.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *ContractValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorsUpdated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsUpdated is a log parse operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) ParseValidatorsUpdated(log types.Log) (*ContractValidatorsUpdated, error) {
	event := new(ContractValidatorsUpdated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllSnapshottersUpdatedIterator is returned from FilterAllSnapshottersUpdated and is used to iterate over the raw logs and unpacked data for AllSnapshottersUpdated events raised by the Contract contract.
type ContractAllSnapshottersUpdatedIterator struct {
	Event *ContractAllSnapshottersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllSnapshottersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllSnapshottersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllSnapshottersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllSnapshottersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllSnapshottersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllSnapshottersUpdated represents a AllSnapshottersUpdated event raised by the Contract contract.
type ContractAllSnapshottersUpdated struct {
	SnapshotterAddress common.Address
	Allowed            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllSnapshottersUpdated is a free log retrieval operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) FilterAllSnapshottersUpdated(opts *bind.FilterOpts) (*ContractAllSnapshottersUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "allSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllSnapshottersUpdatedIterator{contract: _Contract.contract, event: "allSnapshottersUpdated", logs: logs, sub: sub}, nil
}

// WatchAllSnapshottersUpdated is a free log subscription operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) WatchAllSnapshottersUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllSnapshottersUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "allSnapshottersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllSnapshottersUpdated)
				if err := _Contract.contract.UnpackLog(event, "allSnapshottersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllSnapshottersUpdated is a log parse operation binding the contract event 0x743e47fbcd2e3a64a2ab8f5dcdeb4c17f892c17c9f6ab58d3a5c235953d60058.
//
// Solidity: event allSnapshottersUpdated(address snapshotterAddress, bool allowed)
func (_Contract *ContractFilterer) ParseAllSnapshottersUpdated(log types.Log) (*ContractAllSnapshottersUpdated, error) {
	event := new(ContractAllSnapshottersUpdated)
	if err := _Contract.contract.UnpackLog(event, "allSnapshottersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
