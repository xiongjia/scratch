; elips tests
(message "elisp tests: ")

; buffer filename
(message "buffFilename: %s "
  load-file-name (buffer-file-name))

(message "folder: %s" (file-name-directory "/abc/1.txt"))
(message "folder: %s" (file-name-directory "/abc/test/"))

; strings
(message "substr: %s" (substring "0123456789" 0 -1))
(message "and: %s" (and nil "abc"))

; project root
(defvar app-root-path
  (file-name-directory
    (substring (file-name-directory (or load-file-name (buffer-file-name))) 0 -1))
  "The root dir of this project")
(message "app root: %s " app-root-path)

