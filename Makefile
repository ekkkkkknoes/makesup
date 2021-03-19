BIN=makesup
SRC=cmd.go  editor.go  keybinds.go  layout.go  main.go

.PHONY: all clean

all: $(BIN)

$(BIN): $(SRC)
	go build -o $(BIN)

clean:
	go clean
	rm -f $(BIN)
