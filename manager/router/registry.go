package router

import (
	"encoding/json"
	"errors"
	"net/http"

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

func (registry *Registry) StakerDelegationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.StakerDelegationRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if request.Address == "" {
			c.JSON(http.StatusBadRequest, errors.New("address must not be nil"))
			return
		}
		var result *types.StakerDelegationResult
		var err error

		result.Amount, err = registry.db.GetBTCDelegateAmount([]byte(request.Address))
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to get staker delegation")
			log.Error("failed to get staker delegation", "error", err)
			return
		}
		data, err := json.Marshal(result)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to marshal staker delegation")
			log.Error("failed to marshal staker delegation", "error", err)
			return
		}
		if _, err = c.Writer.Write(data); err != nil {
			log.Error("failed to write staker delegation to response writer", "error", err)
		}
	}
}

func (registry *Registry) StakerDetailsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.StakerDetailsRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if request.BatchId <= 0 {
			c.JSON(http.StatusBadRequest, errors.New("invalid request batch_id"))
			return
		}
		var result store.StakeDetails
		var err error

		result, err = registry.db.GetStakeDetails(request.BatchId)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to get staker details")
			log.Error("failed to get staker details", "error", err)
			return
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
