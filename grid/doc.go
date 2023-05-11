// Types and functions that support the grid itself.
//
// The grid consists of an nxn matrix of cells, which are of three types:
//   - Black cells: Blocks in the grid
//   - Letter cells: Ordinary cells where letters of words can be placed.
//     They also contain a pair of pointers: one to the numbered cell that
//     begins the across word that contains this cell, and another to the
//     numbered cell that begins the down word that contains this cell.
//   - Numbered cells: Letter cells that mark the beginning of an across
//     and/or down word.
//
// The grid also supports undo/redo for black cells in this grid.
package grid
