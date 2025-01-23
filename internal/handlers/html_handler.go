package handlers

import (
	"html/template"
	"net/http"
	"github.com/pinokiochan/social-network-render/internal/logger"
	"github.com/sirupsen/logrus"
)

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/templates/auth.html")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"path":  "./web/templates/auth.html",
		}).Error("Failed to parse template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	logger.Log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Debug("Serving index page")
	
	t.Execute(w, nil)
}
func ServeIndexHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"path":  "./web/templates/index.html",
		}).Error("Failed to parse auth template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	logger.Log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Debug("Serving auth page")
	
	t.Execute(w, nil)
}

func ServeUserProfileHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/templates/user-profile.html")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"path":  "./web/templates/user-profile.html",
		}).Error("Failed to parse user-profile template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	logger.Log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Debug("Serving user-profile page")
	
	t.Execute(w, nil)
}

func ServeAdminHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/templates/admin.html")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"path":  "./web/templates/admin.html",
		}).Error("Failed to parse admin template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	logger.Log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Debug("Serving admin page")
	
	t.Execute(w, nil)
}

func ServeEmailHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/templates/email.html")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"path":  "./web/templates/email.html",
		}).Error("Failed to parse email template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	logger.Log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Debug("Serving email page")
	
	t.Execute(w, nil)
}

