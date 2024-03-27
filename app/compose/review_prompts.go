package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/note"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func ReviewPrompts(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	prompts := [3]string{"Gratitude","Positivity","Improve"}
	for _, prompt := range prompts {
		nt := note.NewNote(2000, 0, 0)
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Note":         nt,
				"Breadcrumb":   nt.Breadcrumb(),
				"HeadingMOS":   nt.HeadingMOSPrompts(prompt),
				"SideQuarters": year.SideQuarters(0),
				"SideMonths":   year.SideMonths(0),
				"Extra": nt.
					PrevNext(cfg.Layout.Numbers.NotesOnPage * cfg.Layout.Numbers.NotesIndexPages).
					WithTopRightCorner(cfg.ClearTopRightCorner),
				"Extra2": extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
			},
		})
	}

	return modules, nil
}
