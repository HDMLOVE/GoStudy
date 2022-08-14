CXX=go build
CXXFLAGS=
SRCS:=$(wildcard *.go)
OBJECTS=$(SRCS:%.go=%)
TARGETS=$(SRCS:%.go=%)
all:
	@for target in $(TARGETS);\
	do \
	$(CXX) $(CXXFLAGS) -o $$target $$target.go;\
	done
clean:
	-rm -f $(TARGETS)
