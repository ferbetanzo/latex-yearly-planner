package note

import (
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

type Notes []*Note
type Note struct {
	Year   int
	Number int
	Page   int
}

func NewNote(year, page, number int) *Note {
	return &Note{Year: year, Page: page, Number: number}
}

func (p Notes) Breadcrumb(year, idx int) string {
	postfix := ""
	if idx > 0 {
		postfix = " " + strconv.Itoa(idx+1)
	}

	return header.Items{
		header.NewIntItem(year),
		header.NewTextItem("Notes Index" + postfix).Ref(true),
	}.Table(true)
}

func (p Notes) HeadingMOS(page, pages int) string {
	var out string

	anglesize := `\dimexpr\myLenHeaderResizeBox-0.86pt`
	if page > 1 {
		out += tex.Hyperlink(p.ref(page-1), tex.ResizeBoxW(anglesize, `$\langle$`)) + " "
	}

	out += tex.Hypertarget(p.ref(page), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Collections`)

	if page < pages {
		out += " " + tex.Hyperlink(p.ref(page+1), tex.ResizeBoxW(anglesize, `$\rangle$`))
	}

	return out
}

func (p Notes) ref(page int) string {
	var suffix string

	if page > 1 {
		suffix = " " + strconv.Itoa(page)
	}

	return "Notes Index" + suffix
}

func (n Note) HyperLink() string {
	return hyper.Link(n.ref(), fmt.Sprintf("C%02d", n.Number))
}

func (n Note) Breadcrumb() string {
	page := ""

	if n.Page > 1 {
		page = " " + strconv.Itoa(n.Page)
	}

	return header.Items{
		header.NewIntItem(n.Year),
		header.NewTextItem("Notes Index" + page),
		header.NewTextItem(n.ref()).Ref(true),
	}.Table(true)
}

func (n Note) PrevNext(notes int) header.Items {
	items := header.Items{}

	if n.Number > 1 {
		items = append(items, header.NewTextItem("Collection "+strconv.Itoa(n.Number-1)))
	}

	if n.Number < notes {
		items = append(items, header.NewTextItem("Collection "+strconv.Itoa(n.Number+1)))
	}

	return items
}

func (n Note) HeadingMOS(page int) string {
	//num := strconv.Itoa(n.Number)

	return tex.Hypertarget(n.ref(), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `C`+fmt.Sprintf("%02d", n.Number)+`\myDummyQ`)
}

func (n Note) ref() string {
	return "Note " + strconv.Itoa(n.Number)
}

func (n Note) promptRef(prompt string) string {
	return "Prompt " + prompt

}

func (n Note) HeadingMOSPrompts(prompt string) string {
	return tex.Hypertarget(n.promptRef(prompt), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, ``+prompt+` prompts`)
}
