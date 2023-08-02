TRUCKS_SOURCES := $(shell find trucks/ -name '*.fnl' -type f)
CONTRACTS_SOURCES := $(shell find contracts/ -name '*.fnl')

TRUCKS_OBJECTS := $(patsubst trucks/%.fnl, include/%.lua, $(TRUCKS_SOURCES))
CONTRACTS_OBJECTS := $(patsubst contracts/%.fnl, contracts/build/%.lua, $(CONTRACTS_SOURCES))

.PHONY: all clean

all: $(TRUCKS_OBJECTS) $(CONTRACTS_OBJECTS)

clean:
	rm -rf include
	rm -rf contracts/build

include/%.lua: trucks/%.fnl
	mkdir -p $(dir $@)
	fennel --compile $< > $@

contracts/build/%.lua: contracts/%.fnl
	mkdir -p $(dir $@)
	fennel --compile $< > $@

