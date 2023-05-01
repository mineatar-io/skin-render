# skin-render
A library for rendering Minecraft skins into 2D and 3D isometric images. Inspired by existing services like [Crafatar](https://crafatar.com/), but with an effort to make renders faster and more efficient.

## Documentation

https://pkg.go.dev/github.com/mineatar-io/skin-render

## Examples

### `RenderFace()`

![face_steve_test](https://user-images.githubusercontent.com/16949253/158674544-fb37ac0c-9e89-4b14-a000-c2195704b734.png)

### `RenderHead()`

![head_steve_test](https://user-images.githubusercontent.com/16949253/158674629-14037a48-0c6c-4454-81fc-57478356e417.png)

### `RenderBody()`

![fullbody_steve_test](https://user-images.githubusercontent.com/16949253/158674750-32c83eb6-454a-4ecd-bc2a-e28166b250b7.png)

### `RenderFrontBody()`

![frontbody_steve_test](https://user-images.githubusercontent.com/16949253/158674771-925daefd-93c6-4d09-8568-1989ca384bd5.png)

### `RenderBackBody()`

![backbody_steve_test](https://user-images.githubusercontent.com/16949253/158674806-aa7ba0c5-aa68-449f-ad21-e439f86b6556.png)

### `RenderLeftBody()`

![leftbody_steve_test](https://user-images.githubusercontent.com/16949253/158674841-180334b5-fec6-41db-beec-42cd30126736.png)

### `RenderRightBody()`

![rightbody_steve_test](https://user-images.githubusercontent.com/16949253/158674867-eb0ad8fb-b7f0-4dac-bbce-df410ce7ee75.png)

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
