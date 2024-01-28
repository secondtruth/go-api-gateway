package gateway

import "strings"

// matchSegment matches a single segment of a domain against a segment of the pattern.
func matchSegment(valueSeg, patternSeg string) bool {
	if patternSeg == "**" {
		return true
	}
	if patternSeg == "*" {
		return true
	}
    if len(valueSeg) != len(patternSeg) {
        return false
    }
    for i := range valueSeg {
        if patternSeg[i] != '?' && valueSeg[i] != patternSeg[i] {
            return false
        }
    }
	return true
}

// matchSegments recursively matches domain and pattern segments.
func matchSegments(valueParts, patternParts []string) bool {
	if len(patternParts) == 0 {
		return len(valueParts) == 0
	}
	if len(valueParts) == 0 {
		for _, part := range patternParts {
			if part != "**" {
				return false
			}
		}
		return true
	}

	if patternParts[0] == "**" {
		// Try to match ** with the current segment and without it.
		return matchSegments(valueParts, patternParts[1:]) || matchSegments(valueParts[1:], patternParts)
	}
	if matchSegment(valueParts[0], patternParts[0]) {
		return matchSegments(valueParts[1:], patternParts[1:])
	}
	return false
}

// matchDomain matches the domain against the pattern.
func matchDomain(pattern, domain string) bool {
	domainParts := strings.Split(domain, ".")
	patternParts := strings.Split(pattern, ".")
	return matchSegments(domainParts, patternParts)
}
