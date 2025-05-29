package metrics

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"
	"io"
)

const Namespace = "manta_fp_aggregator"

type Metricer interface {
	StartBalanceMetrics(l log.Logger, client *ethclient.Client, account common.Address) io.Closer
	RecordBabylonInterval() (done func(err error))
	RecordCelestiaInterval() (done func(err error))
	RecordEthInterval() (done func(err error))
	RecordLatestBabylonBlock(number uint64)
	RecordLatestCelestiaBlock(number uint64)
	RecordLatestEthBlock(number uint64)
	RecordBabylonIndexedHeaders(size int)
	RecordCelestiaIndexedHeaders(size int)
	RecordEthIndexedHeaders(size int)
	RecordBabylonBatchLog(contractAddress string)
	RecordEthBatchLog(contractAddress string)
	RecordBatchId(id uint64)
	RecordWindowPeriodStartTime(timestamp uint64)
	RecordWindowPeriodEndTime(timestamp uint64)
}

type Metrics struct {
	ns       string
	registry *prometheus.Registry

	latestBabylonBlockNumber  prometheus.Gauge
	latestCelestiaBlockNumber prometheus.Gauge
	latestEthBlockNumber      prometheus.Gauge
	babylonIndexedHeaders     prometheus.Counter
	celestiaIndexedHeaders    prometheus.Counter
	ethIndexedHeaders         prometheus.Counter
	babylonBatchLogs          *prometheus.CounterVec
	ethBatchLogs              *prometheus.CounterVec
	babylonIntervalTick       prometheus.Counter
	celestiaIntervalTick      prometheus.Counter
	ethIntervalTick           prometheus.Counter
	babylonIntervalDuration   prometheus.Histogram
	celestiaIntervalDuration  prometheus.Histogram
	ethIntervalDuration       prometheus.Histogram
	babylonBatchFailures      prometheus.Counter
	celestiaBatchFailures     prometheus.Counter
	ethBatchFailures          prometheus.Counter

	batchId               prometheus.Gauge
	windowPeriodStartTime prometheus.Gauge
	windowPeriodEndTime   prometheus.Gauge
}

func NewMetrics(registry *prometheus.Registry) Metricer {
	factory := With(registry)
	return &Metrics{
		ns:       Namespace,
		registry: registry,
		latestBabylonBlockNumber: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "latest_babylon_block_number",
			Help:      "the latest babylon block height observed by a synchronizer interval",
		}),
		latestCelestiaBlockNumber: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "latest_celestia_block_number",
			Help:      "the latest celestia block height observed by a synchronizer interval",
		}),
		latestEthBlockNumber: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "latest_eth_block_number",
			Help:      "the latest eth block height observed by a synchronizer interval",
		}),
		batchId: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "batch_id",
			Help:      "batch id of the signed transaction submitted in the frm contract",
		}),
		babylonIntervalTick: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "babylon_intervals_total",
			Help:      "number of times the babylon synchronizer has run its extraction loop",
		}),
		celestiaIntervalTick: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "celestia_intervals_total",
			Help:      "number of times the celestia synchronizer has run its extraction loop",
		}),
		ethIntervalTick: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "eth_intervals_total",
			Help:      "number of times the eth synchronizer has run its extraction loop",
		}),
		babylonIntervalDuration: factory.NewHistogram(prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "babylon_interval_seconds",
			Help:      "duration elapsed for during the babylon processing loop",
		}),
		celestiaIntervalDuration: factory.NewHistogram(prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "celestia_interval_seconds",
			Help:      "duration elapsed for during the celestia processing loop",
		}),
		ethIntervalDuration: factory.NewHistogram(prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "eth_interval_seconds",
			Help:      "duration elapsed for during the eth processing loop",
		}),
		babylonBatchFailures: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "babylon_failures_total",
			Help:      "number of times the babylon synchronizer encountered a failure to extract a batch",
		}),
		celestiaBatchFailures: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "celestia_failures_total",
			Help:      "number of times the celestia synchronizer encountered a failure to extract a batch",
		}),
		ethBatchFailures: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "eth_failures_total",
			Help:      "number of times the eth synchronizer encountered a failure to extract a batch",
		}),
		babylonBatchLogs: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "babylon_logs_total",
			Help:      "number of logs observed by the babylon synchronizer",
		}, []string{
			"contract",
		}),
		ethBatchLogs: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "eth_logs_total",
			Help:      "number of logs observed by the eth synchronizer",
		}, []string{
			"contract",
		}),
		babylonIndexedHeaders: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "babylon_indexed_headers_total",
			Help:      "number of headers indexed by the babylon synchronizer",
		}),
		celestiaIndexedHeaders: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "celestia_indexed_headers_total",
			Help:      "number of headers indexed by the celestia synchronizer",
		}),
		ethIndexedHeaders: factory.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "eth_indexed_headers_total",
			Help:      "number of headers indexed by the eth synchronizer",
		}),
		windowPeriodStartTime: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "window_period_start_time",
			Help:      "the window period start time of verifying state root",
		}),
		windowPeriodEndTime: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      "window_period_end_time",
			Help:      "the window period end time of verifying state root",
		}),
	}
}

func (m *Metrics) StartBalanceMetrics(l log.Logger, client *ethclient.Client, account common.Address) io.Closer {
	return LaunchBalanceMetrics(l, m.registry, m.ns, client, account)
}

func (m *Metrics) RecordLatestBabylonBlock(number uint64) {
	m.latestBabylonBlockNumber.Set(float64(number))
}

func (m *Metrics) RecordLatestCelestiaBlock(number uint64) {
	m.latestCelestiaBlockNumber.Set(float64(number))
}

func (m *Metrics) RecordLatestEthBlock(number uint64) {
	m.latestEthBlockNumber.Set(float64(number))
}

func (m *Metrics) RecordBabylonIndexedHeaders(size int) {
	m.babylonIndexedHeaders.Add(float64(size))
}

func (m *Metrics) RecordCelestiaIndexedHeaders(size int) {
	m.celestiaIndexedHeaders.Add(float64(size))
}

func (m *Metrics) RecordEthIndexedHeaders(size int) {
	m.ethIndexedHeaders.Add(float64(size))
}

func (m *Metrics) RecordBabylonBatchLog(contractAddress string) {
	m.babylonBatchLogs.WithLabelValues(contractAddress).Inc()
}

func (m *Metrics) RecordEthBatchLog(contractAddress string) {
	m.ethBatchLogs.WithLabelValues(contractAddress).Inc()
}

func (m *Metrics) RecordBatchId(number uint64) {
	m.batchId.Set(float64(number))
}

func (m *Metrics) RecordBabylonInterval() func(error) {
	m.babylonIntervalTick.Inc()
	timer := prometheus.NewTimer(m.babylonIntervalDuration)
	return func(err error) {
		if err != nil {
			m.babylonBatchFailures.Inc()
		}

		timer.ObserveDuration()
	}
}

func (m *Metrics) RecordCelestiaInterval() func(error) {
	m.celestiaIntervalTick.Inc()
	timer := prometheus.NewTimer(m.celestiaIntervalDuration)
	return func(err error) {
		if err != nil {
			m.celestiaBatchFailures.Inc()
		}

		timer.ObserveDuration()
	}
}

func (m *Metrics) RecordEthInterval() func(error) {
	m.ethIntervalTick.Inc()
	timer := prometheus.NewTimer(m.ethIntervalDuration)
	return func(err error) {
		if err != nil {
			m.ethBatchFailures.Inc()
		}

		timer.ObserveDuration()
	}
}

func (m *Metrics) RecordWindowPeriodStartTime(timestamp uint64) {
	m.windowPeriodStartTime.Set(float64(timestamp))
}

func (m *Metrics) RecordWindowPeriodEndTime(timestamp uint64) {
	m.windowPeriodEndTime.Set(float64(timestamp))
}
