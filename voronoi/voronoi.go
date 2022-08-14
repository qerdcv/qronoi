package voronoi

import (
	"fmt"
	img "image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
)

const (
	width  = 1920
	height = 1080

	coefX = width / 255
	coefY = height / 255

	seedMarkerRadius = 5
)

var (
	seedMarkerColor = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
)

type VDiagram struct {
	image *img.RGBA

	kw      string
	seeds   []img.Point
	palette []color.RGBA
}

func New(kw string, withDots bool) *VDiagram {
	vd := new(VDiagram).createImage().
		setKeyWord(kw).
		generateSeed().
		generatePalette().
		renderVoronoi()

	if withDots {
		return vd.renderSeeds()
	}

	return vd

}

func (vd *VDiagram) createImage() *VDiagram {
	vd.image = img.NewRGBA(img.Rectangle{Min: img.Point{}, Max: img.Point{X: width, Y: height}})
	return vd
}

func (vd *VDiagram) setKeyWord(kw string) *VDiagram {
	vd.kw = kw
	return vd
}

func (vd *VDiagram) generateSeed() *VDiagram {
	for i, b := range []byte(vd.kw) {
		ib := int(b)
		vd.seeds = append(vd.seeds, img.Point{
			X: ((ib + 100*i) * coefX) % width,
			Y: ((ib + 50*i) * coefY) % height,
		})
	}

	return vd
}

func (vd *VDiagram) generatePalette() *VDiagram {
	var seed int64 = 0
	for _, b := range []byte(vd.kw) {
		seed += int64(b)
	}

	rand.Seed(seed)
	for range vd.seeds {
		vd.palette = append(vd.palette, color.RGBA{
			R: uint8(rand.Int() % 255),
			G: uint8(rand.Int() % 255),
			B: uint8(rand.Int() % 255),
			A: 255,
		})
	}

	return vd
}

func (vd *VDiagram) renderVoronoi() *VDiagram {
	pLen := len(vd.palette)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			j := 0
			for i := range vd.seeds {
				if sqrDist(vd.seeds[i], img.Point{X: x, Y: y}) <
					sqrDist(vd.seeds[j], img.Point{X: x, Y: y}) {
					j = i
				}
			}

			vd.image.Set(x, y, vd.palette[j%pLen])
		}
	}
	return vd
}

func (vd *VDiagram) renderCircle(cp img.Point, r int, c color.RGBA) {
	p1 := img.Point{
		X: cp.X - r,
		Y: cp.Y - r,
	}

	p2 := img.Point{
		X: cp.X + r,
		Y: cp.Y + r,
	}

	sqR := r * r

	for x := p1.X; x <= p2.X; x++ {
		if x >= width || x < 0 {
			continue
		}

		for y := p1.Y; y <= p2.Y; y++ {
			if y >= height || y < 0 {
				continue
			}

			if sqrDist(img.Point{X: cp.X, Y: cp.Y}, img.Point{X: x, Y: y}) <= sqR {
				vd.image.Set(x, y, c)
			}
		}
	}
}

func (vd *VDiagram) renderSeeds() *VDiagram {
	for _, seed := range vd.seeds {
		vd.renderCircle(seed, seedMarkerRadius, seedMarkerColor)
	}

	return vd
}

func (vd *VDiagram) Export(w io.Writer) error {
	if err := png.Encode(w, vd.image); err != nil {
		return fmt.Errorf("png encode: %w", err)
	}

	return nil
}

func sqrDist(p1, p2 img.Point) int {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return dx*dx + dy*dy
}
