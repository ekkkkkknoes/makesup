BIN=makesup
SRC=main.go

.PHONY: all clean

all: $(BIN)

$(BIN): $(SRC)
	go build -o $(BIN)

clean:
	go clean
	rm -f $(BIN)
