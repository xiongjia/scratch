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
(message "concat %s" (concat "123" "abc"))

; project root
(defvar app-root-path
  (file-name-directory
    (substring (file-name-directory (or load-file-name (buffer-file-name))) 0 -1))
  "The root dir of this project")
(message "app root: %s " app-root-path)

(load-file (concat app-root-path "scripts/misc.el"))
(test-export)

(message "Load data from file: ")
(message (file-contents
  (concat app-root-path "assets/test-data.txt")))


; env
(message "get env: %s" (getenv "HOME"))
(message "get env: %s" (not (string= (getenv "HOME") nil)))
(message "get env: %s" (not (string= (getenv "HOME100") nil)))

; misc
(setq testdata
  `(
    ("test"
     :base "abc"
     :base2 t
     )
    ("data" :components ("test"))))

(message "test data: %s" testdata)


