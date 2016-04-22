package mvc

type Configuration struct {
	SiteName string
	LeftNav  NavLinks
	RightNav NavLinks
}

func newConfiguration() (c *Configuration) {
	c = new(Configuration)
	c.LeftNav = make(map[string]string)
	c.RightNav = make(map[string]string)
	return
}

func (this *Configuration)clone() (c *Configuration) {
	c = newConfiguration()
	c.SiteName = this.SiteName

	for k, v := range this.LeftNav {
		c.LeftNav[k] = v
	}

	for k, v := range this.RightNav {
		c.RightNav[k] = v
	}
	return
}

type NavLinks map[string]string

func (this NavLinks)Add(display string, url string) {
	this[display] = url
}