package components

import (
	"github.com/cj123/vecty"
	"github.com/cj123/vecty/elem"
	"github.com/cj123/vecty/event"
	"github.com/cj123/vecty/example/todomvc/actions"
	"github.com/cj123/vecty/example/todomvc/dispatcher"
	"github.com/cj123/vecty/example/todomvc/store"
	"github.com/cj123/vecty/example/todomvc/store/model"
	"github.com/cj123/vecty/prop"
)

// FilterButton is a vecty.Component which allows the user to select a filter
// state.
type FilterButton struct {
	vecty.Core

	Label  string            `vecty:"prop"`
	Filter model.FilterState `vecty:"prop"`
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

// Render implements the vecty.Component interface.
func (b *FilterButton) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				vecty.MarkupIf(store.Filter == b.Filter, vecty.Class("selected")),
				prop.Href("#"),
				event.Click(b.onClick).PreventDefault(),
			),

			vecty.Text(b.Label),
		),
	)
}
