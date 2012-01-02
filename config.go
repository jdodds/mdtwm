package main

import (
	"code.google.com/p/x-go-binding/xgb"
	"os"
	"path/filepath"
)

type Config struct {
	Instance string
	Class    string

	NormalBorderColor  uint32
	FocusedBorderColor uint32
	BorderWidth        int16
	StatusLogger       StatusLogger `json:"-"`

	DefaultCursor    xgb.Id        `json:"-"`
	MoveCursor       xgb.Id        `json:"-"`
	MultiClickTime   xgb.Timestamp
	MovedClickRadius int

	ModMask uint16
	Keys    map[byte]Cmd `json:"-"`

	Ignore List
	Float  List
}

func configure() {
	// Default configuration
	cfg = &Config{
		Instance: filepath.Base(os.Args[0]),
		Class:    "Mdtwm",

		NormalBorderColor:  rgbColor(0x8888, 0x8888, 0x8888),
		FocusedBorderColor: rgbColor(0x4444, 0x0000, 0xffff),
		BorderWidth:        1,

		StatusLogger: &Dzen2Logger{
			Writer:     os.Stdout,
			FgColor:    "#ddddcc",
			BgColor:    "#555588",
			TimeFormat: "Mon, Jan _2 15:04:05",
			TimePos:    -154, // Negatife value means pixels from right border
		},

		DefaultCursor:    stdCursor(68),
		MoveCursor:       stdCursor(52),
		MultiClickTime:   300, // maximum interval for multiclick [ms]
		MovedClickRadius: 5,   // minimal radius for moved click [pixel]

		ModMask: xgb.ModMask4,
		Keys: map[byte]Cmd{
			Key1:     {chDesk, 1},
			Key2:     {chDesk, 2},
			Key3:     {chDesk, 3},
			KeyEnter: {spawn, "gnome-terminal"},
			KeyQ:     {exit, 0},
		},

		Ignore: List{},
		Float:  List{},
	}
	// Read configuration from file
	cfg.Load(filepath.Join(os.Getenv("HOME"), ".mdtwm"))

	// Layout
	root = NewRootPanel()
	// Setup all desks
	desk1 := NewPanel(Horizontal, 1.82)
	//desk1 := NewPanel(Horizontal, 1.97)
	desk2 := NewPanel(Horizontal, 1)
	desk3 := NewPanel(Horizontal, 1)
	root.Append(desk1)
	root.Append(desk2)
	root.Append(desk3)
	// Setup two main vertical panels on first desk
	left := NewPanel(Vertical, 1.03)
	right := NewPanel(Vertical, 0.3)
	//left := NewPanel(Vertical, 1.02)
	//right := NewPanel(Vertical, 0.29)
	desk1.Append(left)
	desk1.Append(right)
	// Divide right panel into two horizontal panels
	right.Append(NewPanel(Horizontal, 1))
	right.Append(NewPanel(Horizontal, 1))
	// Setup one main panel on second and thrid desk
	desk2.Append(NewPanel(Horizontal, 1))
	desk3.Append(NewPanel(Vertical, 1))
	// Set current desk and current box
	currentDesk = desk1
	currentDesk.Raise()
	// In this box all existing windows will be placed
	currentBox = currentDesk.Children().Front()

	// Some operation on configuration 
	cfg.MovedClickRadius *= cfg.MovedClickRadius // We need square of radius
	if cfg.StatusLogger != nil {
		cfg.StatusLogger.Start()
	}
}
