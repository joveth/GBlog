package controllers

import (
	"github.com/revel/revel"
	"GBlog/app/models"
	"strings"
	"GBlog/app/myutils"
)
type WBlog struct {
	App
}
func (c WBlog) Putup(blog *models.Blog) revel.Result {
	blog.Title = strings.TrimSpace(blog.Title);
	blog.Email = strings.TrimSpace(blog.Email);
	blog.Subject = strings.TrimSpace(blog.Subject);
	blog.Validate(c.Validation)
	//c.Validation.Required(blog.Title).Message("Title should not be null.")
	//c.Validation.MaxSize(blog.Title,256).Message("The title is too long,please checkout it.")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Please correct the errors below.")
		return c.Redirect(App.WBlog)
	}
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	//dao := models.NewDao(c.MongoSession)
	if len(blog.Title)>35 {
		blog.ShortTitle = myutils.Substr(blog.Title,0,35)
	}
	if len(blog.Subject)>200 {
		blog.ShortSubject = myutils.Substr(blog.Subject,0,200)
	}
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
