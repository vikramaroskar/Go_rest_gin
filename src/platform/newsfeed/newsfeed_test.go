package newsfeed

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {


	feed := New()
	
	fmt.Println(feed)
	
	feed.Add(Newsitem{"One", "This is One"})
	feed.Add(Newsitem{"Two", "This is Two"})
	
	fmt.Println(feed)
	
	
	if len(feed.Feeditems) !=2 {
		t.Errorf("items not added properly")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	
	fmt.Println(feed)
	
	feed.Add(Newsitem{"One", "This is One"})
	feed.Add(Newsitem{"Two", "This is Two"})
	
	fmt.Println(feed)
	
	results := feed.GetAll()
	
	if len(results) !=2 {
		t.Errorf("items not added properly")
	}
}