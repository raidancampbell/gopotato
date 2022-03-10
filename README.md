# gopotato
A CHIP-8 emulator written in golang

[Cowgod's Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
[wikipedia page](https://en.wikipedia.org/wiki/CHIP-8)
[timing reference](https://jackson-s.me/2019/07/13/Chip-8-Instruction-Scheduling-and-Frequency.html)

# TODO
 - [X] integrate with ~~SDL~~ Pixel to draw out the display
 - [X] add support for a keyboard
 - [X] implement opcodes Ex9E onwards
 - [ ] emit sound
 - [X] read a ROM into memory
 - [X] add hardcoded 5-byte hex font
 - [ ] (maybe): render output to terminal instead of a graphical display