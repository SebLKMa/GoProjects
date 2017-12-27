package search

// implement interface Matcher from match.go
type defaultMatcher struct{}

// registers the default matcher with the program
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// implements Matcher.Search as no-op
// where (m defaultMatcher) is receiver type, 
// Search() is now bound to values or pointers of defaultMatcher type
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}