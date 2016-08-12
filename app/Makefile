# makefile

TARGET=main

OBJS=\
	 main.o

INCLUDES=\
	-I$(CURDIR)/../include

LIBS=\
	 $(CURDIR)/../lib/libhello.a

LDFLAGS=\
	-lpthread

all: $(LIBS) $(TARGET)

$(TARGET): $(OBJS) 
	gcc -o main $(OBJS) $(LIBS) $(LDFLAGS)


$(CURDIR)/../lib/libhello.a:
	make -C $(CURDIR)/../hello

%.o : %.c
	gcc -c $(INCLUDES) $<

clean:
	make clean -C $(CURDIR)/../hello
	rm -f $(TARGET) $(OBJS)


