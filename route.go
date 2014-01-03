package ape

import (
	"regexp"
	"strings"
)

const (
	defaultConstraint    = "[^/]+?"
	formantConstraint    = "(?:\\.(?P<_format>[a-z0-9_]+))?"
	NumberOnlyConstraint = "[0-9]+"
)

var (
	namedSegmentRegexp = regexp.MustCompile(":[a-z][a-z0-9_]*")
)

type Route struct {
	Verbs       []string
	Path        string
	Handler     Handler
	Constraints map[string]string
	regexp      *regexp.Regexp
}

func NewRoute(verbs []string, path string, handler Handler) *Route {
	route := &Route{
		Verbs:   verbs,
		Path:    path,
		Handler: handler,
	}

	route.compile()

	return route
}

func (r *Route) Match(verb string, path string) bool {
	verbHit := false

	for _, expectedVerb := range r.Verbs {
		if expectedVerb == verb {
			verbHit = true
			break
		}
	}

	return verbHit && r.regexp.MatchString(path)
}

func (r *Route) Params(path string) map[string]string {
	params := map[string]string{}
	segments := r.regexp.SubexpNames()

	segmentParams := r.regexp.FindAllStringSubmatch(path, -1)[0]

	for i, segment := range segments {
		if i == 0 {
			continue
		}

		params[segment] = segmentParams[i]
	}

	return params
}

func (r *Route) compile() {
	segments := strings.Split(r.Path, "/")
	compiledSegments := make([]string, len(segments))

	for i, segment := range segments {
		switch {
		case namedSegmentRegexp.MatchString(segment):
			segmentName := regexp.QuoteMeta(segment[1:len(segment)])
			constraint := defaultConstraint

			if _, hit := r.Constraints[segmentName]; hit {
				constraint = r.Constraints[segmentName]
			}

			segment = "(?P<" + segmentName + ">" + constraint + ")"
		default:
			segment = regexp.QuoteMeta(segment)
		}
		compiledSegments[i] = segment
	}

	r.regexp = regexp.MustCompile(
		"^" +
			strings.Join(compiledSegments, "/") +
			formantConstraint +
			"$",
	)
}

func (r *Route) String() string {
	verbs := strings.Join(r.Verbs, "|")

	return "[" + verbs + "] " + r.Path
}
