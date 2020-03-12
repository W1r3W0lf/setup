call plug#begin('~/.vim/plugged')

" Files
Plug 'scrooloose/nerdtree'

Plug 'majutsushi/tagbar'

" Git
Plug 'tpope/vim-fugitive'
Plug 'airblade/vim-gitgutter'

"Plug 'vim-scripts/nginx.vim'

" Spell Check
Plug 'reedes/vim-lexical'
set nocompatible
filetype plugin on

augroup lexical
	autocmd!
	autocmd FileType markdown,mkd call lexical#init()
	autocmd FileType textile call lexical#init()
	autocmd FileType text call lexical#init()
augroup END

"Plug 'vim-syntastic/syntastic'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'scrooloose/nerdcommenter'
Plug 'jiangmiao/auto-pairs'
Plug 'lambdalisue/vim-manpager'

" Java stuff
Plug 'artur-shaik/vim-javacomplete2'

"Language packs/servers
Plug 'neoclide/coc.nvim', {'tag': '*', 'branch': 'release'}
Plug 'sheerun/vim-polyglot'
Plug 'autozimu/LanguageClient-neovim', {
    \ 'branch': 'next',
    \ 'do': 'bash install.sh',
    \ }

Plug 'dense-analysis/ale'

"Go stuff
Plug 'fatih/vim-go'

"Rust Stuff
Plug 'rust-lang/rust.vim'
Plug 'timonv/vim-cargo'
Plug 'racer-rust/vim-racer'

"Clojure Stuff
Plug 'guns/vim-clojure-static'
Plug 'tpope/vim-fireplace'
Plug 'luochen1990/rainbow'
Plug 'guns/vim-sexp'
Plug 'tpope/vim-sexp-mappings-for-regular-people'
"Plug 'gberenfield/cljfold.vim'
" clojure rainbow parens
let g:rainbow_active = 1
let g:rainbow_conf = {
      \  'guifgs': ['royalblue3', 'darkorange3', 'seagreen3', 'firebrick'],
      \  'ctermfgs': ['lightblue', 'lightyellow', 'lightcyan', 'lightmagenta'],
      \  'parentheses': ['start=/(/ end=/)/ fold', 'start=/\[/ end=/\]/ fold', 'start=/{/ end=/}/ fold'],
      \  'separately': {
      \      '*': 0,
      \      'clojure': {},
      \  }
      \}


" (Optional) Multi-entry selection UI.
Plug 'junegunn/fzf', { 'do': './install --bin' }
Plug 'junegunn/fzf.vim'

"Build
Plug 'vhdirk/vim-cmake'

" Looks
Plug 'vim-airline/vim-airline'
Plug 'challenger-deep-theme/vim', { 'as': 'challenger-deep' }
Plug 'editorconfig/editorconfig-vim'
Plug 'ryanoasis/vim-devicons'
call plug#end()

" Normal Config Stuff
set number
set tabstop=4 shiftwidth=4 smarttab autoindent
set encoding=utf-8
set fileencoding=utf-8
set cursorline
"set mouse=a
set exrc
set secure
syntax on

" NERDTree


"challenger-deep
colorscheme challenger_deep
set termguicolors


" vim-airline
let g:airline_theme='darkscene'
let g:airline_powerline_fonts = 1
let g:airline#extensions#tabline#enabled = 1
let g:airline#extensions#bufferline#enabled = 1
"let g:airline_left_sep = ''
"let g:airline_left_alt_sep = ''
"let g:airline_right_sep = ''

" key mappings
nmap z : NERDTreeToggle<cr>

map <F1> <nop>
imap <F1> <nop>

map <F2> :TagbarToggle<CR>

nnoremap <C-Tab> :bn<CR>
nnoremap <C-S-Tab> :bp<CR>
