
NAME=cppcreate

all:
	go build -o $(NAME) main.go 

fclean:
	rm $(NAME)
