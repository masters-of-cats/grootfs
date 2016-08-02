package groot

import (
	"fmt"

	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter . Graph
//go:generate counterfeiter . Cloner

type Graph interface {
	MakeBundle(lager.Logger, string) (Bundle, error)
}

type CloneSpec struct {
	FromDir, ToDir           string
	UIDMappings, GIDMappings []IDMappingSpec
}

type IDMappingSpec struct {
	HostID, NamespaceID, Size int
}

type Cloner interface {
	Clone(lager.Logger, CloneSpec) error
}

type Groot struct {
	graph  Graph
	cloner Cloner
}

func IamGroot(graph Graph, cloner Cloner) *Groot {
	return &Groot{
		graph:  graph,
		cloner: cloner,
	}
}

type CreateSpec struct {
	ID          string
	ImagePath   string
	UIDMappings []IDMappingSpec
	GIDMappings []IDMappingSpec
}

func (g *Groot) Create(logger lager.Logger, spec CreateSpec) (Bundle, error) {
	bundle, err := g.graph.MakeBundle(logger, spec.ID)
	if err != nil {
		return nil, fmt.Errorf("making bundle: %s", err)
	}

	err = g.cloner.Clone(logger, CloneSpec{
		FromDir:     spec.ImagePath,
		ToDir:       bundle.RootFsPath(),
		UIDMappings: spec.UIDMappings,
		GIDMappings: spec.GIDMappings,
	})
	if err != nil {
		return nil, fmt.Errorf("cloning: %s", err)
	}

	return bundle, nil
}
