package gopotato

type ram [0xFFF]byte

var mem ram

// 0x000 to 0x1FF reserved for interpreter
// 0x200 start of programs
