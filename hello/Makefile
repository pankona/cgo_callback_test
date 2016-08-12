# makefile

LIB_OUTDIR=$(CURDIR)/../lib
HEADER_OUTDIR=$(CURDIR)/../include

PROJ_NAME=libhello

LIB_NAME=$(PROJ_NAME).a
TARGET_LIB=$(LIB_OUTDIR)/$(LIB_NAME)

HEADER_NAME=$(PROJ_NAME).h
TARGET_HEADER=$(HEADER_OUTDIR)/$(HEADER_NAME)

SRCS=\
	hello.go \
	callback.go

all: $(TARGET_LIB)

$(TARGET_LIB):
	go build -buildmode=c-archive -o $(LIB_NAME) $(SRCS)
	mv $(LIB_NAME) $(TARGET_LIB)
	mv $(HEADER_NAME) $(TARGET_HEADER)

clean:
	rm -f $(TARGET_LIB) $(TARGET_HEADER)


