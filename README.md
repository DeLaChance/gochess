# GoChess
A Chess game programmed in the [Go programming language](https://go.dev/). 

Currently simulates 2 AI players.

## To run
`go run .`

Example output:

```
INFO: 2022/07/12 18:16:17 chess_game.go:17: Player turn is W
INFO: 2022/07/12 18:16:17 chess_game.go:18: Board:
-----------------
|♙|♙|♝|♛|♙|♝|♞|♜|
-----------------
| | |♟| | | | | |
-----------------
| | | | | | | | |
-----------------
| | | | | | | | |
-----------------
| | | | | | | | |
-----------------
| | | | | | | | |
-----------------
| | | | | | | | |
-----------------
|♖|♟|♟|♟|♔|♟|♟|♖|
-----------------
```

## To build
`go install`

## TODO:

### Backend code
[x] Basic structure
[x] Pawn moves implemented
[x] Simulate AI vs AI game
[x] Implement DB layer
[x] Fix bug with only storing first entity in a list
[x] Implement service layer
[x] Implement REST api's

- Implement bishop moves
- Implement knight moves
- Implement rook moves
- Implement queen moves 
- Implement king moves 
- Implement check
- Implement check mate
- Implement pawn to queen logic 
- Implement rook and king swap
- Refactor + cleanup the code
- Implement human player
- Move code from domain to service layer

### CI/CD

- Containerize
- Gitlab pipeline