package prom

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
)

func labsSort(src labels.Labels) labels.Labels {
	slices.SortFunc(src, func(a labels.Label, b labels.Label) int {
		return strings.Compare(a.Name, b.Name)
	})
	return src
}

func labMetadataCopy(src metadata.Metadata) *metadata.Metadata {
	return &metadata.Metadata{
		Type: src.Type,
		Unit: strings.Clone(src.Unit),
		Help: strings.Clone(src.Help),
	}
}

func labsToJson(src labels.Labels) (string, error) {
	b, err := src.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func labsFromJson(src string) (labels.Labels, error) {
	var m map[string]string
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		return nil, err
	}
	return labels.FromMap(m), nil
}
