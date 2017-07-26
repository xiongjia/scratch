; elips tests

(defun test-export ()
  (message "test export"))

(defun file-contents (filename)
  "Return the contents of FILENAME."
  (with-temp-buffer
    (insert-file-contents filename)
    (buffer-string)))

