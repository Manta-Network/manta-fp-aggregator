package metrics

import (
	"net"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Manta-Network/manta-fp-aggregator/common/httputil"
)

func StartServer(r *prometheus.Registry, hostname string, port int) (*httputil.HTTPServer, error) {
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))
	h := promhttp.InstrumentMetricHandler(
		r, promhttp.HandlerFor(r, promhttp.HandlerOpts{}),
	)
	return httputil.StartHTTPServer(addr, h)
}
