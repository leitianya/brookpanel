package controllers

type UserIndexController struct {
	BaseController
}

//　用户首页
func (c *UserIndexController) Index() {
	c.Data["title"] = "用户首页-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userindex/css.html"
	c.LayoutSections["footerjs"] = "userindex/js.html"
	c.setTpl("userindex/index.html", "shared/userpanel.html")

}
