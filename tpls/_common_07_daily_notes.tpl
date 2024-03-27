{{ if $.Cfg.Dotted -}} \vskip-.5\myLenLineHeightButLine\myMash{\myNumDotNotesHeightFull}{\myNumDotWidthFull} {{- else -}}
\vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
  \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
}
{{end}}
