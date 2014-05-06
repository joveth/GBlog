package controllers

import (
	"github.com/revel/revel"
	"GBlog/app/models"
	"strings"
)
type WBlog struct {
	App
}
func (c WBlog) Putup(blog *models.Blog) revel.Result {
	blog.Title = strings.TrimSpace(blog.Title);
	blog.Email = strings.TrimSpace(blog.Email);
	blog.Subject = strings.TrimSpace(blog.Subject);
	blog.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.WBlog)
	}
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	err = dao.CreateBlog(blog)
	if(err!=nil){
		c.Response.Status = 500
		return c.RenderError(err)
	}
	newEmail := new(models.EmailObj);
	newEmail.Email = blog.Email;
	dao.InsertEmail(newEmail);
	return c.Redirect(App.Index)
}
