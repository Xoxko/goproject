/*
Compile like this:

..\..\bin\lcc -Wa-l -Wl-m -Wl-j -DUSE_SFR_FOR_REG -c -o piano.o piano.c
..\..\bin\lcc -Wa-l -Wl-m -Wl-j -DUSE_SFR_FOR_REG -o piano.gb piano.o

(Same as in make.bat in examples)

*/

#include <gb/gb.h>
#include <stdio.h>
#include <gb/console.h>

void play_note(UBYTE notevalue);

enum notes {
  C0, Cd0, D0, Dd0, E0, F0, Fd0, G0, Gd0, A0, Ad0, B0,
  C1, Cd1, D1, Dd1, E1, F1, Fd1, G1, Gd1, A1, Ad1, B1,
  C2, Cd2, D2, Dd2, E2, F2, Fd2, G2, Gd2, A2, Ad2, B2,
  C3, Cd3, D3, Dd3, E3, F3, Fd3, G3, Gd3, A3, Ad3, B3,
  C4, Cd4, D4, Dd4, E4, F4, Fd4, G4, Gd4, A4, Ad4, B4,
  C5, Cd5, D5, Dd5, E5, F5, Fd5, G5, Gd5, A5, Ad5, B5
};

const UWORD frequencies[] = {
  44, 156, 262, 363, 457, 547, 631, 710, 786, 854, 923, 986,
  1046, 1102, 1155, 1205, 1253, 1297, 1339, 1379, 1417, 1452, 1486, 1517,
  1546, 1575, 1602, 1627, 1650, 1673, 1694, 1714, 1732, 1750, 1767, 1783,
  1798, 1812, 1825, 1837, 1849, 1860, 1871, 1881, 1890, 1899, 1907, 1915,
  1923, 1930, 1936, 1943, 1949, 1954, 1959, 1964, 1969, 1974, 1978, 1982,
  1985, 1988, 1992, 1995, 1998, 2001, 2004, 2006, 2009, 2011, 2013, 2015
};


void main()
{
	UBYTE input;		// Variable for joypad data
	NR50_REG = 0x77;	// This turns on sound (Off by default!)
	NR51_REG = 0xFF;	// This turns on sound (Off by default!)
	NR52_REG = 0x80;	// This turns on sound (Off by default!)

	puts("PIANO EXAMPLE CODE");

	while(1){	// Loop forever
		input = joypad();	// Read the joypad

		// Check for keys
		if(input & J_A){
			puts("A       C2");		// Print the name of the button
			play_note(C2);
			waitpadup();	// Wait until the button is released
		}else if(input & J_B){
			puts("B       B");		// Print the name of the button
			play_note(B1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_SELECT){
			puts("SELECT  G");		// Print the name of the button
			play_note(G1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_START){
			puts("START   A");		// Print the name of the button
			play_note(A1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_UP){
			puts("UP      F");		// Print the name of the button
			play_note(F1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_DOWN){
			puts("DOWN    E");		// Print the name of the button
			play_note(E1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_LEFT){
			puts("LEFT    C");		// Print the name of the button
			play_note(C1);
			waitpadup();	// Wait until the button is released
		}else if(input & J_RIGHT){
			puts("RIGHT   D");		// Print the name of the button
			play_note(D1);
			waitpadup();	// Wait until the button is released
		}


	}
}

void play_note(UBYTE notevalue){
	UBYTE flo, fhi;
	UWORD freq = frequencies[notevalue];

	// Volume envelope
	NR12_REG = 0xF3;	// F=maximum volume, 3=sound fade out
	// Waveform
	NR11_REG = 0x80;	// 50% square wave
	// Frequency
	flo = (UBYTE)freq & 0xFF;
	fhi = (UBYTE)((freq & 0x0700)>>8);
	NR13_REG = flo;	// Take lower 8 bits from the function argument
	NR14_REG = 0x80 | fhi;

	// Take 3 more bits from the function argument, and set the start bit


}