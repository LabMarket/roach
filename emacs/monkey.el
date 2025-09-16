;;; roach.el --- mode for editing roach scripts

;; Copyright (C) 2018 Steve Kemp

;; Author: Steve Kemp <steve@steve.fi>
;; Keywords: languages
;; Version: 1.0

;;; Commentary:

;; Provides support for editing roach scripts with full support for
;; font-locking, but no special keybindings, or indentation handling.

;;;; Enabling:

;; Add the following to your .emacs file

;; (require 'roach)
;; (setq auto-mode-alist (append '(("\\.roach$" . roach-mode)) auto-mode-alist)))



;;; Code:

(defvar roach-constants
  '("true"
    "false"))

(defvar roach-keywords
  '(
    "case"
    "else"
    "fn"
    "for"
    "foreach"
    "function"
    "if"
    "in"
    "let"
    "return"
    "switch"
    ))

;; The language-core and functions from the standard-library.
(defvar roach-functions
  '(
    "args"
    "exit"
    "file.close"
    "file.lines"
    "file.open"
    "first"
    "int"
    "last"
    "len"
    "math.abs"
    "math.random"
    "math.sqrt"
    "push"
    "puts"
    "read"
    "rest"
    "set"
    "string"
    "string.interpolate"
    "string.reverse"
    "string.split"
    "string.tolower"
    "string.toupper"
    "string.trim"
    "type"
    "version"
    ))


(defvar roach-font-lock-defaults
  `((
     ("\"\\.\\*\\?" . font-lock-string-face)
     (";\\|,\\|=" . font-lock-keyword-face)
     ( ,(regexp-opt roach-keywords 'words) . font-lock-builtin-face)
     ( ,(regexp-opt roach-constants 'words) . font-lock-constant-face)
     ( ,(regexp-opt roach-functions 'words) . font-lock-function-name-face)
     )))

(define-derived-mode roach-mode fundamental-mode "roach script"
  "roach-mode is a major mode for editing roach scripts"
  (setq font-lock-defaults roach-font-lock-defaults)

  ;; Comment handler for single & multi-line modes
  (modify-syntax-entry ?\/ ". 124b" roach-mode-syntax-table)
  (modify-syntax-entry ?\* ". 23n" roach-mode-syntax-table)

  ;; Comment ender for single-line comments.
  (modify-syntax-entry ?\n "> b" roach-mode-syntax-table)
  (modify-syntax-entry ?\r "> b" roach-mode-syntax-table)
  )

(provide 'roach)
