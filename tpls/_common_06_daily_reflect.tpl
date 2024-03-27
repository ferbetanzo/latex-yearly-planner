{{- $today := .Body.Weekday -}}

{{ if eq $today 0 }}
\myUnderline{What did I do that moved me closer to reaching my \textbf{goals}?}
\myMash{\myNumWeeklyReflectGoals}{\myNumDotWidthFull}
\vskip\dimexpr5.4mm

\myUnderline{What am I most \textbf{grateful} for this week and why?}
\myMash{\myNumWeeklyReflectGood}{\myNumDotWidthFull}
\vskip\dimexpr5.4mm

\myUnderline{What would I have \textbf{done differently} this week and why?}
\myMash{\myNumWeeklyReflectBad}{\myNumDotWidthFull}
\vskip\dimexpr5.4mm

\myUnderline{}
\vskip\dimexpr2mm
\renewcommand{\arraystretch}{\myNumArrayStretch}
\begin{tabularx}{\linewidth}{*{2}{X}}
The funniest thing was: & \\
I most enjoyed: & \\
I learned: & \\
I procrastinated on: & \\
\end{tabularx}%
{{ else }}
\vskip\dimexpr9mm
\myUnderline{\hyperlink{Prompt Gratitude}{Gratitude}}
\myMash{\myNumDailyDiaryGrateful}{\myNumDotWidthFull}

\vskip\dimexpr9mm
\myUnderline{\hyperlink{Prompt Positivity}{Positivity}}
\myMash{\myNumDailyDiaryBest}{\myNumDotWidthFull}

\vskip\dimexpr9mm
\myUnderline{\hyperlink{Prompt Improve}{Improve}}
\myMash{\myNumDailyDiaryImprove}{\myNumDotWidthFull}

\iffalse
\vskip\dimexpr5.4mm
\myUnderline{Daily log}
\myMash{\myNumDailyDiaryLog}{\myNumDotWidthFull}
\fi
{{ end }}
