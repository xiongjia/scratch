; elips tests

(load-file "~/.emacs.d/init.el")

(defun publish ()
  (org-publish-project "abathur" 'true))

(publish)

