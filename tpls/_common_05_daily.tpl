{{- $today := .Body.Day -}}
{{- $weekday := .Body.Weekday -}}

\begin{minipage}[t]{\myLenTriCol}
{{template "schedule.tpl" dict "Cfg" .Cfg "Day" .Body.Day}}
  \vspace{\dimexpr4mm+.3pt}

{{- if .Cfg.CalAfterSchedule -}}
{{- template "monthTabularV2.tpl" dict "Month" .Body.Month "Today" $today -}}
{{- end -}}
\end{minipage}%
\hspace{\myLenTriColSep}%
\begin{minipage}[t]{\dimexpr2\myLenTriCol+\myLenTriColSep}
  \myUnderline{My intention for today is\myDummyQ}
  \myMash[\myDailySpring]{2}{\myNumDotWidthTwoThirds}
  \vskip\dimexpr5.4mm
  \myUnderline{Todo\myDummyQ}
  \iffalse
  \Repeat{\myNumDailyTodos}{\myTodoLineGray}
  \fi
  \myMash[\myDailySpring]{\myNumDailyTodos}{\myNumDotWidthTwoThirds}
  \vskip\dimexpr5.4mm
  \iffalse
  \myUnderline{Notes $\vert$ {{ $today.LinkLeaf "More" "More" }}\hfill{}{{ $today.LinkLeaf "Reflect" "Reflect" }}\hfill{}\hyperlink{Notes Index}{Collections}}
  \fi
  {{ if eq $weekday 0 }}
  \myUnderline{Notes $\vert$ {{ $today.LinkLeaf "More" "More" }}\hfill{}Reflect $\vert$ {{ $today.LinkLeaf "Reflect" "Week" }}}
  {{ else }}
  \myUnderline{Notes $\vert$ {{ $today.LinkLeaf "More" "More" }}\hfill{}Reflect $\vert$ {{ $today.LinkLeaf "Reflect" "Day" }}}
  {{ end }}
  \myMash[\myDailySpring]{\myNumDailyNotes}{\myNumDotWidthTwoThirds}
\end{minipage}
\par\pagebreak
