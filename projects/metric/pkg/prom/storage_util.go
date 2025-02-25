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

func labMetadataToJson(src *metadata.Metadata) (string, error) {
	if src == nil {
		return "", nil
	}
	b, err := json.Marshal(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func labMetadataFromJsonDefault(src string, defaultVal *metadata.Metadata) *metadata.Metadata {
	m, err := labMetadataFromJson(src)
	if err != nil {
		return defaultVal
	}
	return m
}

func labMetadataFromJson(src string) (*metadata.Metadata, error) {
	if len(src) == 0 {
		return nil, nil
	}

	var m metadata.Metadata
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func labMetadataEqual(m1 *metadata.Metadata, m2 *metadata.Metadata) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil && m2 != nil {
		return false
	}
	if m1 != nil && m2 == nil {
		return false
	}
	if (m1.Help == m2.Help) && (m1.Unit == m2.Unit) && (m1.Type == m2.Type) {
		return true
	}
	return false
}

func labMetadataCopy(src *metadata.Metadata) *metadata.Metadata {
	if src == nil {
		return nil
	}
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
