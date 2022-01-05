package types

import (
	"gh-labeler/pkg/errors"

	"github.com/google/go-github/v41/github"
)

var (
	DEFAULT_LABEL_COLOR = "dddddd"
)

type Label struct {
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (l *Label) validate() error {
	if l.Name == nil || *l.Name == "" {
		return errors.MISSING_NAME
	}
	return nil
}

func (l *Label) sanitize() {
	if l.Color == nil || *l.Color == "" {
		l.Color = &DEFAULT_LABEL_COLOR
	}

	if l.Description == nil || *l.Description == "" {
		l.Description = l.Name
	}
}

func From(upstreamLabel *github.Label) *Label {
	if upstreamLabel == nil {
		return nil
	}
	return &Label{
		Name:        upstreamLabel.Name,
		Color:       upstreamLabel.Color,
		Description: upstreamLabel.Description,
	}
}
