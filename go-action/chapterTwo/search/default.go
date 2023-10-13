package search

import "log"

type defaultMatcher struct {
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string){
	return nil, nil
}

dm := new(defaultMatcher)
dm.Search(feed, "test")

func (m *defaultMatcher) Search(feed *Feed,searchTerm string)


var dm defaultMatcher

dm.Search(feed,"test")


func Register(feedType string,matcher Matcher){
	if _,exists := matchers[feedType];exists{
		log.Fatalln(feedType,"Matcher already registered")
	}
	log.Println("Register",feedType,"matcher")
	matchers[feedType] = matcher
}
