let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/src/budgetAutomation
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +32 ~/src/budgetAutomation/src/mappings.go
badd +104 src/main.go
badd +19 ~/go/src/goplayground/main.go
badd +7 ~/src/budgetAutomation/src/util.go
argglobal
%argdel
edit ~/src/budgetAutomation/src/mappings.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
2wincmd h
wincmd w
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 25 + 95) / 190)
exe 'vert 2resize ' . ((&columns * 81 + 95) / 190)
exe '3resize ' . ((&lines * 39 + 26) / 53)
exe 'vert 3resize ' . ((&columns * 82 + 95) / 190)
exe '4resize ' . ((&lines * 10 + 26) / 53)
exe 'vert 4resize ' . ((&columns * 82 + 95) / 190)
argglobal
enew
file NvimTree_1
balt ~/src/budgetAutomation/src/mappings.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
balt ~/src/budgetAutomation/src/util.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 32 - ((31 * winheight(0) + 25) / 50)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 32
normal! 015|
wincmd w
argglobal
if bufexists(fnamemodify("src/main.go", ":p")) | buffer src/main.go | else | edit src/main.go | endif
if &buftype ==# 'terminal'
  silent file src/main.go
endif
balt ~/src/budgetAutomation/src/util.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 104 - ((15 * winheight(0) + 19) / 39)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 104
normal! 029|
wincmd w
argglobal
enew
balt src/main.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
wincmd w
3wincmd w
exe 'vert 1resize ' . ((&columns * 25 + 95) / 190)
exe 'vert 2resize ' . ((&columns * 81 + 95) / 190)
exe '3resize ' . ((&lines * 39 + 26) / 53)
exe 'vert 3resize ' . ((&columns * 82 + 95) / 190)
exe '4resize ' . ((&lines * 10 + 26) / 53)
exe 'vert 4resize ' . ((&columns * 82 + 95) / 190)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
