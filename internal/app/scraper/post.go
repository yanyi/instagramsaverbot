package scraper

import (
	"errors"
)

func (p *post) getOnlyImageURLs() error {
	scm := p.GraphQL.ShortCodeMedia
	edges := scm.EdgeSidecarToChildren.Edges
	multiPost := len(edges) > 0

	// Checks if the post is multi-post first
	if multiPost {
		for _, edge := range edges {
			scm := edge.Node
			if scm.IsVideo == false {
				err := p.appendImageURLs(scm)
				if err != nil {
					return err
				}
			}
		}
	} else if scm.IsVideo == false {
		err := p.appendImageURLs(scm)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *post) appendImageURLs(scm shortCodeMedia) error {
	imageURL, err := scm.getHighestDefImageURL()
	if err != nil {
		return err
	}

	p.ImageURLs = append(p.ImageURLs, imageURL)

	return nil
}

func (scm shortCodeMedia) getHighestDefImageURL() (string, error) {
	resources := scm.DisplayResources

	if len(resources) == 0 {
		return "", errors.New("Error getting highest definition image: might not have an image to get")
	}

	return resources[len(resources)-1].Src, nil
}
