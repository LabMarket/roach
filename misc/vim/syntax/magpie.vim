if exists("b:current_syntax")
  finish
endif


syn case match

syn keyword     roachDirective         import
syn keyword     roachDeclaration       let const


hi def link     roachDirective         Statement
hi def link     roachDeclaration       Type

" Linq Keywords
syn keyword     roachLinq              from where select group into orderby join in on equals by ascending descending

syn keyword     roachStatement         return let const spawn defer struct enum using async await service
syn keyword     roachException         try catch finally throw
syn keyword     roachConditional       if else elif unless and or case is
syn keyword     roachRepeat            do while for break continue grep map
syn keyword     roachBranch            break continue
syn keyword     roachClass             class new property get set default this parent static public private protected interface

hi def link     roachStatement         Statement
hi def link     roachClass             Statement
hi def link     roachConditional       Conditional
hi def link     roachBranch            Conditional
hi def link     roachLabel             Label
hi def link     roachRepeat            Repeat
hi def link     roachLinq              Keyword

syn match       roachDeclaration       /\<fn\>/
syn match       roachDeclaration       /^fn\>/


syn keyword roachCommentTodo contained TODO FIXME XXX NOTE
hi def link roachCommentTodo Todo

syn match comment "#.*$"    contains=@Spell,roachCommentTodo
syn match comment "\/\/.*$" contains=@Spell,roachCommentTodo

syn keyword     roachCast              int str float array


hi def link     roachCast              Type


syn keyword     roachBuiltins          len
syn keyword     roachBuiltins          println print stdin stdout stderr
syn keyword     roachBoolean           true false
syn keyword     roachNull              nil

hi def link     roachBuiltins          Keyword
hi def link     roachNull              Keyword
hi def link     roachBoolean           Boolean


" Comments; their contents
syn keyword     roachTodo              contained TODO FIXME XXX BUG
syn cluster     roachCommentGroup      contains=roachTodo
syn region      roachComment           start="#" end="$" contains=@roachCommentGroup,@Spell,@roachTodo


hi def link     roachComment           Comment
hi def link     roachTodo              Todo


" roach escapes
syn match       roachEscapeOctal       display contained "\\[0-7]\{3}"
syn match       roachEscapeC           display contained +\\[abfnrtv\\'"]+
syn match       roachEscapeX           display contained "\\x\x\{2}"
syn match       roachEscapeU           display contained "\\u\x\{4}"
syn match       roachEscapeBigU        display contained "\\U\x\{8}"
syn match       roachEscapeError       display contained +\\[^0-7xuUabfnrtv\\'"]+


hi def link     roachEscapeOctal       roachSpecialString
hi def link     roachEscapeC           roachSpecialString
hi def link     roachEscapeX           roachSpecialString
hi def link     roachEscapeU           roachSpecialString
hi def link     roachEscapeBigU        roachSpecialString
hi def link     roachSpecialString     Special
hi def link     roachEscapeError       Error
hi def link     roachException		Exception

" Strings and their contents
syn cluster     roachStringGroup       contains=roachEscapeOctal,roachEscapeC,roachEscapeX,roachEscapeU,roachEscapeBigU,roachEscapeError
syn region      roachString            start=+"+ skip=+\\\\\|\\"+ end=+"+ contains=@roachStringGroup
syn region      roachRegExString       start=+/[^/*]+me=e-1 skip=+\\\\\|\\/+ end=+/\s*$+ end=+/\s*[;.,)\]}]+me=e-1 oneline
syn region      roachRawString         start=+`+ end=+`+

hi def link     roachString            String
hi def link     roachRawString         String
hi def link     roachRegExString       String

" Characters; their contents
syn cluster     roachCharacterGroup    contains=roachEscapeOctal,roachEscapeC,roachEscapeX,roachEscapeU,roachEscapeBigU
syn region      roachCharacter         start=+'+ skip=+\\\\\|\\'+ end=+'+ contains=@roachCharacterGroup


hi def link     roachCharacter         Character


" Regions
syn region      roachBlock             start="{" end="}" transparent fold
syn region      roachParen             start='(' end=')' transparent


" Integers
syn match       roachDecimalInt        "\<\d\+\([Ee]\d\+\)\?\>"
syn match       roachHexadecimalInt    "\<0x\x\+\>"
syn match       roachOctalInt          "\<0\o\+\>"
syn match       roachOctalError        "\<0\o*[89]\d*\>"


hi def link     roachDecimalInt        Integer
hi def link     roachHexadecimalInt    Integer
hi def link     roachOctalInt          Integer
hi def link     Integer                 Number

" Floating point
syn match       roachFloat             "\<\d\+\.\d*\([Ee][-+]\d\+\)\?\>"
syn match       roachFloat             "\<\.\d\+\([Ee][-+]\d\+\)\?\>"
syn match       roachFloat             "\<\d\+[Ee][-+]\d\+\>"


hi def link     roachFloat             Float
"hi def link     roachImaginary         Number


if exists("roach_fold")
    syn match	roachFunction	"\<fn\>"
    syn region	roachFunctionFold	start="\<fn\>.*[^};]$" end="^\z1}.*$" transparent fold keepend

    syn sync match roachSync	grouphere roachFunctionFold "\<fn\>"
    syn sync match roachSync	grouphere NONE "^}"

    setlocal foldmethod=syntax
    setlocal foldtext=getline(v:foldstart)
else
    syn keyword roachFunction	fn
    syn match	roachBraces	"[{}\[\]]"
    syn match	roachParens	"[()]"
endif

syn sync fromstart
syn sync maxlines=100

hi def link roachFunction		Function
hi def link roachBraces		Function

let b:current_syntax = "roach"
