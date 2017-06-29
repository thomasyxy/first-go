package search

import (
  "log"
  "sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string)  {
  feeds, err := RetrieveFeeds()

  if err != nil {
    log.Fatal(err)
  }

  result := make(chan *Result)

  var waitGroup sync.WaitGroup

  waitGroup.Add(len(feeds))

  for _, feed := range feeds {
    matcher, exists := matchers[feed.Type]
    if !exists {
      matchers = matchers["default"]
    }

    go func(matcher Matcher, feed *Feed) {
      Match(matcher, feed, searchTerm, results)
      waitGroup.Done()
    }(matcher, feed)
  }

  go func() {
    waitGroup.Wait()

    close(results)
  }()

  Display(results)
}
