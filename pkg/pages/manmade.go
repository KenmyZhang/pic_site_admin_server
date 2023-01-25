package pages

type Page struct {
	PageIndex  int `json:"pageIndex,omitempty" form:"pageIndex"` // 页码
	PageSize   int `json:"pageSize,omitempty" form:"pageSize"`   // 每页条数
	TotalCount int `json:"totalCount,omitempty"`                 // 总数
}

func (p *Page) SetPage(index, size, count int) {
	p.PageIndex = index
	p.PageSize = size
	p.TotalCount = count
}

func (p Page) GetStartAndEnd() (start, end int) {
	if p.PageIndex >= 1 {
		start = (p.PageIndex - 1) * p.PageSize
	}
	end = start + p.PageSize
	if end > p.TotalCount {
		end = p.TotalCount
	}
	if start > end {
		start = end
	}
	return start, end
}

func (p Page) Clone(totalCount int) Page {
	pageSize, pageIndex := p.PageSize, p.PageIndex
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	return Page{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalCount: totalCount,
	}
}

type Pager interface {
	Len() int
	Copy(start, end int)
}

// ManmadePage 手动分页，这里用泛型写更好，我是1.16的，就不用泛型改interface类型了
//
//	func manmadePage[T any](items []T, page Page) ([]T, Page) {
//		if len(items) == 0 {
//			return nil, page.Clone(0)
//		}
//		pg := page.Clone(len(items))
//		start, end := pg.GetStartAndEnd()
//		items = items[start:end]
//		return items, pg
//	}
func ManmadePage(pager Pager, page Page) (Pager, Page) {
	pg := page.Clone(pager.Len())
	pager.Copy(pg.GetStartAndEnd())
	return pager, pg
}

// PageOffsetAndLimit pageIndex transfer pageOffset
func (p Page) PageOffsetAndLimit() (int, int) {
	return (p.PageIndex - 1) * p.PageSize, p.PageSize
}
