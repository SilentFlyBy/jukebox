package cluster

import (
	"net/http"

	"github.com/SilentFlyBy/jukebox/base/id"
	"github.com/SilentFlyBy/jukebox/cluster"
	"github.com/go-chi/render"
)

func GetClusterList(w http.ResponseWriter, r *http.Request) {
	c := cluster.Cluster{
		ID:    id.NewID(),
		Label: "default",
	}
	list := []cluster.Cluster{c}

	render.JSON(w, r, list)
}
