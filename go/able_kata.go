package able_kata

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryNovel() string {
	return "novel"
}

func (b *Book) CategoryShortStory() string {
	return "short story"
}

func (b *Book) Category() string {
	if b.Pages > 300 {
		return b.CategoryNovel()
	}
	return b.CategoryShortStory()
}
