package world

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type AURPackage struct {
	Description    string
	FirstSubmitted int
	ID             int
	LastModified   int
	Maintainer     string
	Name           string
	NumVotes       int
	OutOfDate      int
	PackageBase    string
	PackageBaseID  int
	Popularity     float32
	URL            string
	URLPath        string
	Version        string
}

type AURResponse struct {
	ResultCount int
	Results     []AURPackage
}

func Search(pkg string) []AURPackage {
	url := "https://aur.archlinux.org/rpc/v5/search/"

	resp, err := http.Get(fmt.Sprintf("%s%s", url, pkg))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var r AURResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	slices.SortStableFunc(r.Results, func(a, b AURPackage) int {
		return fuzzy.RankMatch(a.Name, pkg) - fuzzy.RankMatch(b.Name, pkg)
	})

	return r.Results
}
