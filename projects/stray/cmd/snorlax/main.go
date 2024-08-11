package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"stray/cmd/snorlax/api"
	"stray/internal/log"
	"stray/pkg/dugtrio"

	"github.com/ghodss/yaml"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samber/lo"
	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func split(block *ipaddr.IPAddress, index int, allSplits []*ipaddr.IPAddress) []*ipaddr.IPAddress {
	fmt.Println("\tSplitting", block)
	iterator := block.BlockIterator(index + 1)
	addrs := make([]*ipaddr.IPAddress, 0, block.GetMaxSegmentValue()+1)
	for next := iterator.Next(); next != nil; next = iterator.Next() {
		addrs = append(addrs, next)
	}
	isLastSegment := index == block.GetSegmentCount()-1
	label := "subnets"
	if isLastSegment {
		label = "addresses"
	}
	fmt.Printf("\t\t%d %s: %v\n", len(addrs), label, addrs)
	allSplits = append(allSplits, addrs...)
	return allSplits
}

func splitBlocks(blocks []*ipaddr.IPAddress) []*ipaddr.IPAddress {
	// A sequential block can have at most one segment
	// that is not single nor full size,
	// so for each block we iterate over that one segment
	//ArrayList<IPAddress> allSplits = new ArrayList<IPAddress>();
	var allSplits []*ipaddr.IPAddress
	for _, block := range blocks {
		isSplit := false
		if block.IsMultiple() {
			for i := 0; i < block.GetSegmentCount(); i++ {
				segment := block.GetSegment(i)
				if segment.IsMultiple() {
					if segment.IsFullRange() {
						break
					}
					// segment is neither single nor full-range
					allSplits = split(block, i, allSplits)
					isSplit = true
					break
				}
			}
		}
		if !isSplit {
			label := "address"
			if block.IsMultiple() {
				label = "subnet"
			}
			fmt.Printf("\tNot splitting %s %v\n", label, block)
			allSplits = append(allSplits, block)
		}
	}
	return allSplits
}

func main() {
	err := errors.New("test")
	x := lo.If(err == nil, "").Else(err.Error())
	fmt.Println(x)

	logger := dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	})
	dugtrio.SetDefaultLogger(logger)

	lower, upper := ipaddr.NewIPAddressString("172.24.13.1"), ipaddr.NewIPAddressString("172.24.12.11")
	log.Errorf("range count %v - %v", lower, upper)
	rng := lower.GetAddress().SpanWithRange(upper.GetAddress())
	blocks := rng.SpanWithSequentialBlocks()
	log.Errorf(
		"Starting with %v, the minimal list of sequential blocks is: %v\n",
		rng, blocks)

	addr := splitBlocks(blocks)
	log.Errorf("addr: %v\n", addr)

	swagger, err := api.GetSwagger()
	if err != nil {
		logger.Error("Error loading swagger spec: %s", err)
		os.Exit(1)
	}

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	// server := api.NewServer()

	r := chi.NewMux()
	r.Use(middleware.Logger)
	// r.Use(middleware.OapiRequestValidator(swagger))

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		data, _ := yaml.Marshal(&swagger)
		_, _ = w.Write(data)
		// w.WriteHeader(http.StatusOK)
	})

	// get an `http.Handler` that we can use
	// h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	s.ListenAndServe()
}
