import somar from "./esm/soma.js"
// import media from "./esm/media.js"
// import { validar, XPTO, somar as somarMedia } from "./esm/media.js"
import media, { validar, XPTO, somar as somarMedia } from "./esm/media.js"

console.log(somar(1, 2, 3, 4, 5))
console.log(media(1, 2, 3, 4, 5))

validar()
console.log(XPTO)
somarMedia()