# gin Notes

## *Context* (*struct*)

### Os campos da estrutura

- `Params`: populado pelo próprio *gin*, são os valores de parâmetros que veem na requisição (`/some/endpoint/:value`).
- `Keys`: populado pelos *middlewares*, guarda informações para serem transitadas entre eles e que também pode ser usada no *handler* final.

### *Parse* da requisição na estrutura

todas as funções de *binding* do `gin.Context` se diferem apenas em:

- *binda* e "esvazia" a request, sem retorno.
- *binda* e "esvazia" a request, retorna erro se houver.
- *binda* e **NÃO** "esvazia" a request, retorna erro se houver.

OBS: that's ok?
