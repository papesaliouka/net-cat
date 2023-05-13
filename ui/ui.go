package ui

import(
 "github.com/jroimartin/gocui"
 "log"
)
func SetupUI() {
	// Initialize gocui
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
			log.Panicln(err)
	}
	defer g.Close()

	// Set up the layout
	g.SetManagerFunc(func(g *gocui.Gui) error {
			maxX, maxY := g.Size()
			if v, err := g.SetView("input", 0, maxY-3, maxX-1, maxY-1); err != nil {
					if err != gocui.ErrUnknownView {
							return err
					}
					v.Editable = true
					v.Wrap = true
					if _, err := g.SetCurrentView("input"); err != nil {
							return err
					}
			}
			if v, err := g.SetView("log", 0, 0, maxX-1, maxY-4); err != nil {
					if err != gocui.ErrUnknownView {
							return err
					}
					v.Autoscroll = true
			}
			return nil
	})

	// Start the main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			log.Panicln(err)
	}
}
