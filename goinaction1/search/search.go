package search

import (
	"log"
	"sync"
)

//用于搜索匹配器的映射，key是feed类型，value是feed匹配器
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//构造一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	//构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	//设置需要的goroutine的数量
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//Launch a goroutine to monitor when all the work is done
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	
	Display(results)
}

//Register is called to register a matcher for use by the program
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher aleady registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
