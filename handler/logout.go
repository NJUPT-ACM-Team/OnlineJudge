package handler

func (this *Handler) Logout() {
	this.session.Logout()
}
