package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/store"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Registry struct {
	signService types.SignService
	db          *store.Storage
}

func NewRegistry(signService types.SignService, db *store.Storage) *Registry {
	return &Registry{
		signService: signService,
		db:          db,
	}
}

func (registry *Registry) SignMsgHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.SignMsgRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if len(request.TxHash) == 0 || request.BlockNumber == nil || request.TxType == "" {
			c.JSON(http.StatusBadRequest, errors.New("tx_hash, block_number and tx_type must not be nil"))
			return
		}
		var result *types.SignResult
		var err error

		result, err = registry.signService.SignMsgBatch(request)

		if err != nil {
			c.String(http.StatusInternalServerError, "failed to sign msg")
			log.Error("failed to sign msg", "error", err)
			return
		}
		if _, err = c.Writer.Write(result.Signature.Serialize()); err != nil {
			log.Error("failed to write signature to response writer", "error", err)
		}
	}
}

func (registry *Registry) StakerAmountHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		amount, err := registry.db.GetBTCDelegateAmount()
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to get staker delegation amount")
			log.Error("failed to get staker delegation amount", "error", err)
			return
		}
		c.String(http.StatusOK, strconv.FormatUint(amount, 10))
	}
}

func (registry *Registry) StakerDetailsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.StakerDetailsRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if request.BatchId < 0 {
			c.JSON(http.StatusBadRequest, errors.New("invalid request batch_id"))
			return
		}
		var result store.StakeDetails
		var bTotalAmount uint64
		var err error

		result, err = registry.db.GetBatchStakeDetails(request.BatchId)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to get staker details")
			log.Error("failed to get staker details", "error", err)
			return
		}
		if result.BatchID == request.BatchId {
			bTotalAmount, err = registry.db.GetBatchTotalBabylonStakeAmount(request.BatchId)
			if err != nil {
				c.String(http.StatusInternalServerError, "failed to get stake amount")
				log.Error("failed to get stake amount", "error", err)
				return
			}
			result.TotalBTCVote = bTotalAmount
		}

		data, err := json.Marshal(result)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to marshal staker details")
			log.Error("failed to marshal staker details", "error", err)
			return
		}
		if _, err = c.Writer.Write(data); err != nil {
			log.Error("failed to write staker details to response writer", "error", err)
		}
	}
}

func (registry *Registry) PrometheusHandler() gin.HandlerFunc {
	h := promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{MaxRequestsInFlight: 3},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
