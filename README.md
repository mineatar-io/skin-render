# skin-render
A fast and efficient library for rendering Minecraft skins into 2D and 3D isometric images. Inspired by existing services like [Crafatar](https://crafatar.com/) but with performance in mind. This library is used in production for [mineatar.io](https://mineatar.io).

## Documentation

https://pkg.go.dev/github.com/mineatar-io/skin-render

## Examples

Method              | Result
------------------- | ------
`RenderFace()`      | ![face_steve_test](https://api.mineatar.io/face/c06f89064c8a49119c29ea1dbd1aab82?scale=8)
`RenderHead()`      | ![head_steve_test](https://api.mineatar.io/head/c06f89064c8a49119c29ea1dbd1aab82?scale=6)
`RenderBody()`      | ![fullbody_steve_test](https://api.mineatar.io/body/full/c06f89064c8a49119c29ea1dbd1aab82?scale=6)
`RenderFrontBody()` | ![frontbody_steve_test](https://api.mineatar.io/body/front/c06f89064c8a49119c29ea1dbd1aab82?scale=6)
`RenderBackBody()`  | ![backbody_steve_test](https://api.mineatar.io/body/back/c06f89064c8a49119c29ea1dbd1aab82?scale=6)
`RenderLeftBody()`  | ![leftbody_steve_test](https://api.mineatar.io/body/left/c06f89064c8a49119c29ea1dbd1aab82?scale=6)
`RenderRightBody()` | ![rightbody_steve_test](https://api.mineatar.io/body/right/c06f89064c8a49119c29ea1dbd1aab82?scale=6)

## Credit

- [Isometric graphics in Inkscape &mdash; Nicolás Guarín-Zapata](https://web.archive.org/web/20220306062006/https://nicoguaro.github.io/posts/isometric_inkscape/)
- [go-gl/matgl](https://github.com/go-gl/mathgl)
- [LapisBlue/Lapitar](https://github.com/LapisBlue/Lapitar)
- [go/x/image `transform_Image_Image_Over` function](https://cs.opensource.google/go/x/image/+/refs/heads/master:draw/impl.go;drc=ed5dba0ea28f9438e4dac0320f7d9bb2fddd9737;l=965)
- [go/x/image matrix `invert` function](https://cs.opensource.google/go/x/image/+/refs/heads/master:draw/scale.go;l=332;drc=ed5dba0ea28f9438e4dac0320f7d9bb2fddd9737)
- And various other online matrix tutorials

A special thanks to `oakar258` in the [Minecraft Wiki Discord server](https://minecraft.fandom.com/wiki/Minecraft_Wiki:Discord) for support on how Minecraft scales and renders the overlay skin layer.

## License

[MIT License](https://github.com/mineatar-io/skin-render/blob/main/LICENSE)
