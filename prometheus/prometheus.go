package prometheus

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jphuc96/skymavis-section2/provider"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	echo *echo.Echo
}

func New() *Prometheus {
	e := echo.New()
	e.HideBanner = true

	p := &Prometheus{
		echo: e,
	}

	upCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "eth_up",
	})

	blockNumberGaugeInfura := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "eth_block_number",
		ConstLabels: prometheus.Labels{
			"provider": "infura",
		},
	})

	blockNumberGaugeAnkr := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "eth_block_number",
		ConstLabels: prometheus.Labels{
			"provider": "ankr",
		},
	})

	blockNumberDifference := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "eth_block_number_difference",
		ConstLabels: prometheus.Labels{
			"between": "infura_ankr",
		},
	})

	prometheus.MustRegister(
		upCounter,
		blockNumberGaugeInfura,
		blockNumberGaugeAnkr,
		blockNumberDifference,
	)

	// setup metrics
	p.echo.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		AfterNext: func(c echo.Context, err error) {
			upCounter.Inc()
		},

		BeforeNext: func(c echo.Context) {
			infuraBlockNumber := uint64(0)
			ankrBlockNumber := uint64(0)

			wg := sync.WaitGroup{}

			wg.Add(2)

			go func() {
				defer wg.Done()
				b, err := provider.NewProvider(fmt.Sprintf("https://mainnet.infura.io/v3/%s", os.Getenv("INFURA_API_KEY"))).BLockNumber()
				if err != nil {
					log.Println(err)
					return
				}

				infuraBlockNumber = b
			}()

			go func() {
				b, err := provider.NewProvider(fmt.Sprintf("https://rpc.ankr.com/eth")).BLockNumber()
				if err != nil {
					log.Println(err)
					return
				}

				ankrBlockNumber = b
				wg.Done()
			}()

			wg.Wait()

			blockNumberGaugeInfura.Set(float64(infuraBlockNumber))
			blockNumberGaugeAnkr.Set(float64(ankrBlockNumber))
			blockNumberDifference.Set(float64(infuraBlockNumber - ankrBlockNumber))
		},
	}))

	// setup handler
	p.echo.GET("/metrics", echoprometheus.NewHandler())

	return p
}

func (p *Prometheus) Start() error {
	log.Println("starting prometheus server on :9999")
	return p.echo.Start(":9999")
}
