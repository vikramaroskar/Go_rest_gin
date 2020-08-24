package newsfeed

//interface

type Getter interface {
	GetAll() []Newsitem
}

type Adder interface {
	Add(item Newsitem)
}



type Newsitem struct {

	Title string `json:"title"`
	Post string `json:"post"`
}

type Newsfeed struct {
	Feeditems []Newsitem

}

func New() *Newsfeed {
	return &Newsfeed {
		Feeditems: []Newsitem{},
	}
}

func (nf *Newsfeed) Add(item Newsitem) {
	
	nf.Feeditems = append(nf.Feeditems, item)
	
	
}

func (nf *Newsfeed) GetAll() []Newsitem {

	return nf.Feeditems
	
}