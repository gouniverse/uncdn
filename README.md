# uncdn

![tests](https://github.com/gouniverse/uncdn/workflows/tests/badge.svg)

CDNs are lovely, however there are certain drawbacks of using a CDN for your website may far outweigh the benefits.

This library is an attemp to take the control back from CDNs, while keeping all the good parts.

1. Security. Using CDN someone else is hosting these files. You must trust them 100%. No??? :)

2. Speed. Your Golang app will serve these faster than any CDN.

3. Reliability. CDNs do fail. When they do your website is useless.

4. Embeded CSS/JS. Even though you can either serve these locally, if you embed them in the webpage you will reduce the requests and speed the page.

5. Automated testing.

6. Better SEO

7. Costs. If you pay for CDN the costs can be substantial

## Installation

```
go get github.com/gouniverse/uncdn
```

## How to Use

```golang
uncdn.JQuery360()
```

## Notable Articles

- https://www.maxlaumeister.com/articles/3-reasons-not-to-use-a-cdn/

- https://ericsachsseo.com/to-use-a-cdn-or-not-to-use-a-cdn/

## Methods

### JQuery

- Jquery360() string

### JQuery UI

- JqueryUiCss1132() string
- JqueryUiJs1132() string

## Material Design Color Palette 1.1 
source (https://github.com/zavoloklom/material-design-color-palette)

- MaterialDesignColorPaletteCss11() string

### SweetAlert 2

- SweetAlert2_11432() string

### VueJs 3

- VuewJs3() string

### WebJS

- WebJs260() string

