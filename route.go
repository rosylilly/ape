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
	mounted     bool
}

func NewRoute(verbs []string, path string, handler Handler) *Route {
	route := &Route{
		Verbs:       verbs,
		Path:        path,
		Handler:     handler,
		Constraints: map[string]string{},
		mounted:     false,
	}

	route.compile()

	return route
}

func (r *Route) Match(verb string, path string) bool {
	verb = strings.ToUpper(verb)
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

func (r *Route) Constrain(key, val string) {
	r.Constraints[key] = val
	r.compile()
}

func (r *Route) Mounted() *Route {
	r.mounted = true
	r.compile()

	return r
}

func (r *Route) String() string {
	verbs := strings.Join(r.Verbs, "|")

	return "[" + verbs + "] " + r.Path
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

	regexpSource := "^" + strings.Join(compiledSegments, "/") + formantConstraint

	if !r.mounted {
		regexpSource += "$"
	}

	r.regexp = regexp.MustCompile(regexpSource)
}
