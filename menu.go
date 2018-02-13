package tentsuyu

import "github.com/hajimehoshi/ebiten"

//Menu is a collection of MenuElements
type Menu struct {
	Elements                               [][]*MenuElement
	Active, background                     bool
	Name                                   string
	x, y, midStartX, midStartY, minx, miny float64
	paddingX, paddingY                     int
	maxWidth, maxHeight                    int
	backgroundImage                        *ebiten.Image
	backgroundImgParts                     *BasicImageParts
}

//NewMenu creates a new menu
func NewMenu() *Menu {
	m := &Menu{
		midStartX: Components.ScreenWidth / 2,
		midStartY: Components.ScreenHeight / 4,
		paddingX:  10,
		paddingY:  20,
	}
	return m
}

//SetPadding of the menus x and y values
func (m *Menu) SetPadding(x, y int) {
	m.paddingX = x
	m.paddingY = y
}

//SetBackground image for menu
func (m *Menu) SetBackground(src *ebiten.Image, imgParts *BasicImageParts) {
	m.background = true

	m.backgroundImage = src
	m.backgroundImgParts = imgParts
}

//Update the menu
func (m *Menu) Update() {
	//log.Printf("%v,%v\n", m.maxWidth, m.maxHeight)
	//m.x = Components.Camera.GetX()
	//m.y = Components.Camera.GetY()

	for x := range m.Elements {
		for y := range m.Elements[x] {
			mx := m.x + m.Elements[x][y].menuX
			my := m.y + m.Elements[x][y].menuY

			if x == 0 {
				if y == 0 {
					m.miny = my
				}
				m.minx = mx
			} else {
				if mx < m.minx {
					m.minx = mx
				}
			}

			//m.elements[x][y].UIElement.SetPosition(mx, my)
			m.Elements[x][y].Update()
		}
	}
}

//Draw window
func (m *Menu) Draw(screen *ebiten.Image) {
	if m.background {
		//w := 50.0 //float64(m.maxWidth) / 2
		//h := 50.0 //float64(m.maxHeight) / 2
		scalex := float64(m.maxWidth)/100 + .85
		scaley := float64(m.maxHeight)/40 + .2
		op := &ebiten.DrawImageOptions{}
		op.ImageParts = m.backgroundImgParts
		//op.GeoM.Translate(-w, -h)
		op.GeoM.Scale(float64(scalex), float64(scaley))
		//op.GeoM.Translate(w*float64(scalex), h*float64(scaley))
		op.GeoM.Translate(m.minx-50, m.miny-10)
		//log.Printf("%v,%v\n", scalex, scaley)
		//ApplyCameraTransform(op, false)

		screen.DrawImage(m.backgroundImage, op)

	}
	for x := range m.Elements {
		for y := range m.Elements[x] {
			m.Elements[x][y].UIElement.Draw(screen)
		}
	}
}

//AddElement adds a new Line of UIElements
func (m *Menu) AddElement(element []UIElement, action []func()) {

	menuY := m.midStartY + float64(m.maxHeight)

	/*if len(m.elements) > 0 {
		addition := 0
		for x := range m.elements {
			maxHeight := 0
			for y := range m.elements[x] {
				_, h := m.elements[x][y].Size()
				if h > maxHeight {
					maxHeight = h
				}

			}
			addition += m.paddingY + maxHeight
		}
		menuY += float64(addition)
		if addition > m.maxHeight {
			m.maxHeight += addition
		}
	}*/

	MenuElements := []*MenuElement{}
	maxWidth := 0
	maxHeight := 0
	for i := range element {
		width, height := element[i].Size()
		mE := &MenuElement{
			menuX:     m.midStartX - float64(width/2),
			UIElement: element[i],
			menuY:     menuY,
		}
		switch u := element[i].(type) {
		case *TextElement:
			u.Stationary = true
		}
		mE.SetAction(action[i])
		mE.SetPosition(mE.menuX, mE.menuY)
		MenuElements = append(MenuElements, mE)
		maxWidth += width + m.paddingX
		if height > maxHeight {
			maxHeight = height
		}
	}

	if maxWidth > m.maxWidth {
		m.maxWidth = maxWidth
	}
	maxWidth -= m.paddingX
	m.maxHeight += maxHeight + m.paddingY
	if len(element) > 1 {
		lineStartX := m.midStartX - float64(maxWidth/2)
		for i := range MenuElements {
			MenuElements[i].menuX = lineStartX
			MenuElements[i].SetPosition(MenuElements[i].menuX, MenuElements[i].menuY)
			width, _ := element[i].Size()
			lineStartX += float64(width + m.paddingX)
		}
		lineStartX -= float64(m.paddingX)
	}

	m.Elements = append(m.Elements, MenuElements)
}

type MenuElement struct {
	UIElement
	Action                  func()
	highlighted, Selectable bool
	menuX, menuY            float64
}

//SetAction of the MenuElement
func (m *MenuElement) SetAction(function func()) {
	m.Action = function
	if function == nil {
		m.Selectable = false
	} else {
		m.Selectable = true
	}
}

//Update the MenuElement
func (m *MenuElement) Update() {

	if m.UIElement.Contains(Input.Mouse.X, Input.Mouse.Y) {
		if m.Selectable {
			m.Highlighted()
			m.highlighted = true
			if Input.LeftClick().JustPressed() {

				if m.Action != nil {
					m.Action()
				}
			}
		}
	} else {
		if m.highlighted == true {
			m.highlighted = false
			m.UnHighlighted()
		}
	}

	m.UIElement.Update()
}