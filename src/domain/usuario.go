package domain

type Usuario struct{
	Correo string `json:"email"`
	Pass   string `json:"password"`
}

type UsuarioTextTweet struct{
	Usuario string `json:"user"`
	Texto   string `json:"tweet"`
}